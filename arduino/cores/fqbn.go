// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package cores

import (
	"fmt"
	"strings"

	properties "github.com/arduino/go-properties-orderedmap"
)

// FQBN represents a Board with a specific configuration
type FQBN struct {
	Package      string
	PlatformArch string
	BoardID      string
	Configs      *properties.Map
}

// ParseFQBN extract an FQBN object from the input string
func ParseFQBN(fqbnIn string) (*FQBN, error) {
	// Split fqbn
	fqbnParts := strings.Split(fqbnIn, ":")
	if len(fqbnParts) < 3 || len(fqbnParts) > 4 {
		return nil, fmt.Errorf("not an FQBN: %s", fqbnIn)
	}

	fqbn := &FQBN{
		Package:      fqbnParts[0],
		PlatformArch: fqbnParts[1],
		BoardID:      fqbnParts[2],
		Configs:      properties.NewMap(),
	}
	if fqbn.BoardID == "" {
		return nil, fmt.Errorf(tr("empty board identifier"))
	}
	if len(fqbnParts) > 3 {
		for _, pair := range strings.Split(fqbnParts[3], ",") {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf(tr("invalid config option: %s"), pair)
			}
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			if k == "" {
				return nil, fmt.Errorf(tr("invalid config option: %s"), pair)
			}
			fqbn.Configs.Set(k, v)
		}
	}
	return fqbn, nil
}

func (fqbn *FQBN) String() string {
	res := fqbn.StringWithoutConfig()
	if fqbn.Configs.Size() > 0 {
		sep := ":"
		for _, k := range fqbn.Configs.Keys() {
			res += sep + k + "=" + fqbn.Configs.Get(k)
			sep = ","
		}
	}
	return res
}

// Match check if the target FQBN corresponds to the receiver one.
// The core parts are checked for exact equality while board options are loosely
// matched: the set of boards options of the target must be fully contained within
// the one of the receiver and their values must be equal.
func (fqbn *FQBN) Match(target *FQBN) bool {
	if fqbn.StringWithoutConfig() != target.StringWithoutConfig() {
		return false
	}

	searchedProperties := target.Configs.Clone()
	actualConfigs := fqbn.Configs.AsMap()
	for neededKey, neededValue := range searchedProperties.AsMap() {
		targetValue, hasKey := actualConfigs[neededKey]
		if !hasKey || targetValue != neededValue {
			return false
		}
	}
	return true
}

// StringWithoutConfig returns the FQBN without the Config part
func (fqbn *FQBN) StringWithoutConfig() string {
	return fqbn.Package + ":" + fqbn.PlatformArch + ":" + fqbn.BoardID
}
