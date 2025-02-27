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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: cc/arduino/cli/commands/v1/monitor.proto

package commands

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MonitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Arduino Core Service instance from the `Init` response.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Port to open, must be filled only on the first request
	Port *Port `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// The board FQBN we are trying to connect to. This is optional, and  it's
	// needed to disambiguate if more than one platform provides the pluggable
	// monitor for a given port protocol.
	Fqbn string `protobuf:"bytes,3,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	// Data to send to the port
	TxData []byte `protobuf:"bytes,4,opt,name=tx_data,json=txData,proto3" json:"tx_data,omitempty"`
	// Port configuration, optional, contains settings of the port to be applied
	PortConfiguration *MonitorPortConfiguration `protobuf:"bytes,5,opt,name=port_configuration,json=portConfiguration,proto3" json:"port_configuration,omitempty"`
}

func (x *MonitorRequest) Reset() {
	*x = MonitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorRequest) ProtoMessage() {}

func (x *MonitorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorRequest.ProtoReflect.Descriptor instead.
func (*MonitorRequest) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *MonitorRequest) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *MonitorRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *MonitorRequest) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

func (x *MonitorRequest) GetTxData() []byte {
	if x != nil {
		return x.TxData
	}
	return nil
}

func (x *MonitorRequest) GetPortConfiguration() *MonitorPortConfiguration {
	if x != nil {
		return x.PortConfiguration
	}
	return nil
}

type MonitorPortConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The port configuration parameters to configure
	Settings []*MonitorPortSetting `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty"`
}

func (x *MonitorPortConfiguration) Reset() {
	*x = MonitorPortConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorPortConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorPortConfiguration) ProtoMessage() {}

func (x *MonitorPortConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorPortConfiguration.ProtoReflect.Descriptor instead.
func (*MonitorPortConfiguration) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *MonitorPortConfiguration) GetSettings() []*MonitorPortSetting {
	if x != nil {
		return x.Settings
	}
	return nil
}

type MonitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Eventual errors dealing with monitor port
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Data received from the port
	RxData []byte `protobuf:"bytes,2,opt,name=rx_data,json=rxData,proto3" json:"rx_data,omitempty"`
	// Settings applied to the port, may be returned after a port is opened (to
	// report the default settings) or after a new port_configuration is sent
	// (to report the new settings applied)
	AppliedSettings []*MonitorPortSetting `protobuf:"bytes,3,rep,name=applied_settings,json=appliedSettings,proto3" json:"applied_settings,omitempty"`
	// A message with this field set to true is sent as soon as the port is
	// succesfully opened
	Success bool `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MonitorResponse) Reset() {
	*x = MonitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorResponse) ProtoMessage() {}

func (x *MonitorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorResponse.ProtoReflect.Descriptor instead.
func (*MonitorResponse) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *MonitorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MonitorResponse) GetRxData() []byte {
	if x != nil {
		return x.RxData
	}
	return nil
}

func (x *MonitorResponse) GetAppliedSettings() []*MonitorPortSetting {
	if x != nil {
		return x.AppliedSettings
	}
	return nil
}

func (x *MonitorResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MonitorPortSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SettingId string `protobuf:"bytes,1,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MonitorPortSetting) Reset() {
	*x = MonitorPortSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorPortSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorPortSetting) ProtoMessage() {}

func (x *MonitorPortSetting) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorPortSetting.ProtoReflect.Descriptor instead.
func (*MonitorPortSetting) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *MonitorPortSetting) GetSettingId() string {
	if x != nil {
		return x.SettingId
	}
	return ""
}

func (x *MonitorPortSetting) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type EnumerateMonitorPortSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Arduino Core Service instance from the `Init` response.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// The port protocol to enumerate settings.
	PortProtocol string `protobuf:"bytes,2,opt,name=port_protocol,json=portProtocol,proto3" json:"port_protocol,omitempty"`
	// The board FQBN we are trying to connect to. This is optional, and it's
	// needed to disambiguate if more than one platform provides the pluggable
	// monitor for a given port protocol.
	Fqbn string `protobuf:"bytes,3,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
}

func (x *EnumerateMonitorPortSettingsRequest) Reset() {
	*x = EnumerateMonitorPortSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumerateMonitorPortSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumerateMonitorPortSettingsRequest) ProtoMessage() {}

func (x *EnumerateMonitorPortSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumerateMonitorPortSettingsRequest.ProtoReflect.Descriptor instead.
func (*EnumerateMonitorPortSettingsRequest) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{4}
}

func (x *EnumerateMonitorPortSettingsRequest) GetInstance() *Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *EnumerateMonitorPortSettingsRequest) GetPortProtocol() string {
	if x != nil {
		return x.PortProtocol
	}
	return ""
}

func (x *EnumerateMonitorPortSettingsRequest) GetFqbn() string {
	if x != nil {
		return x.Fqbn
	}
	return ""
}

type EnumerateMonitorPortSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of descriptors of the settings that may be changed for the monitor
	// port.
	Settings []*MonitorPortSettingDescriptor `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty"`
}

func (x *EnumerateMonitorPortSettingsResponse) Reset() {
	*x = EnumerateMonitorPortSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumerateMonitorPortSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumerateMonitorPortSettingsResponse) ProtoMessage() {}

func (x *EnumerateMonitorPortSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumerateMonitorPortSettingsResponse.ProtoReflect.Descriptor instead.
func (*EnumerateMonitorPortSettingsResponse) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{5}
}

func (x *EnumerateMonitorPortSettingsResponse) GetSettings() []*MonitorPortSettingDescriptor {
	if x != nil {
		return x.Settings
	}
	return nil
}

type MonitorPortSettingDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The setting identifier
	SettingId string `protobuf:"bytes,1,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"`
	// A human-readable label of the setting (to be displayed on the GUI)
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// The setting type (at the moment only "enum" is avaiable)
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The values allowed on "enum" types
	EnumValues []string `protobuf:"bytes,4,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// The selected or default value
	Value string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MonitorPortSettingDescriptor) Reset() {
	*x = MonitorPortSettingDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorPortSettingDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorPortSettingDescriptor) ProtoMessage() {}

func (x *MonitorPortSettingDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorPortSettingDescriptor.ProtoReflect.Descriptor instead.
func (*MonitorPortSettingDescriptor) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP(), []int{6}
}

func (x *MonitorPortSettingDescriptor) GetSettingId() string {
	if x != nil {
		return x.SettingId
	}
	return ""
}

func (x *MonitorPortSettingDescriptor) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MonitorPortSettingDescriptor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MonitorPortSettingDescriptor) GetEnumValues() []string {
	if x != nil {
		return x.EnumValues
	}
	return nil
}

func (x *MonitorPortSettingDescriptor) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cc_arduino_cli_commands_v1_monitor_proto protoreflect.FileDescriptor

var file_cc_arduino_cli_commands_v1_monitor_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x63, 0x2e, 0x61,
	0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69,
	0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x63,
	0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x63, 0x2e, 0x61,
	0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x62, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x71, 0x62, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x63,
	0x0a, 0x12, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x63, 0x2e,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x18, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x0f,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x59,
	0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72,
	0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x12, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa0,
	0x01, 0x0a, 0x23, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72,
	0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x71, 0x62, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x62,
	0x6e, 0x22, 0x7c, 0x0a, 0x24, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x63,
	0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0x9e, 0x01, 0x0a, 0x1c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d, 0x63,
	0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e,
	0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cc_arduino_cli_commands_v1_monitor_proto_rawDescOnce sync.Once
	file_cc_arduino_cli_commands_v1_monitor_proto_rawDescData = file_cc_arduino_cli_commands_v1_monitor_proto_rawDesc
)

func file_cc_arduino_cli_commands_v1_monitor_proto_rawDescGZIP() []byte {
	file_cc_arduino_cli_commands_v1_monitor_proto_rawDescOnce.Do(func() {
		file_cc_arduino_cli_commands_v1_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_arduino_cli_commands_v1_monitor_proto_rawDescData)
	})
	return file_cc_arduino_cli_commands_v1_monitor_proto_rawDescData
}

var file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cc_arduino_cli_commands_v1_monitor_proto_goTypes = []interface{}{
	(*MonitorRequest)(nil),                       // 0: cc.arduino.cli.commands.v1.MonitorRequest
	(*MonitorPortConfiguration)(nil),             // 1: cc.arduino.cli.commands.v1.MonitorPortConfiguration
	(*MonitorResponse)(nil),                      // 2: cc.arduino.cli.commands.v1.MonitorResponse
	(*MonitorPortSetting)(nil),                   // 3: cc.arduino.cli.commands.v1.MonitorPortSetting
	(*EnumerateMonitorPortSettingsRequest)(nil),  // 4: cc.arduino.cli.commands.v1.EnumerateMonitorPortSettingsRequest
	(*EnumerateMonitorPortSettingsResponse)(nil), // 5: cc.arduino.cli.commands.v1.EnumerateMonitorPortSettingsResponse
	(*MonitorPortSettingDescriptor)(nil),         // 6: cc.arduino.cli.commands.v1.MonitorPortSettingDescriptor
	(*Instance)(nil),                             // 7: cc.arduino.cli.commands.v1.Instance
	(*Port)(nil),                                 // 8: cc.arduino.cli.commands.v1.Port
}
var file_cc_arduino_cli_commands_v1_monitor_proto_depIdxs = []int32{
	7, // 0: cc.arduino.cli.commands.v1.MonitorRequest.instance:type_name -> cc.arduino.cli.commands.v1.Instance
	8, // 1: cc.arduino.cli.commands.v1.MonitorRequest.port:type_name -> cc.arduino.cli.commands.v1.Port
	1, // 2: cc.arduino.cli.commands.v1.MonitorRequest.port_configuration:type_name -> cc.arduino.cli.commands.v1.MonitorPortConfiguration
	3, // 3: cc.arduino.cli.commands.v1.MonitorPortConfiguration.settings:type_name -> cc.arduino.cli.commands.v1.MonitorPortSetting
	3, // 4: cc.arduino.cli.commands.v1.MonitorResponse.applied_settings:type_name -> cc.arduino.cli.commands.v1.MonitorPortSetting
	7, // 5: cc.arduino.cli.commands.v1.EnumerateMonitorPortSettingsRequest.instance:type_name -> cc.arduino.cli.commands.v1.Instance
	6, // 6: cc.arduino.cli.commands.v1.EnumerateMonitorPortSettingsResponse.settings:type_name -> cc.arduino.cli.commands.v1.MonitorPortSettingDescriptor
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cc_arduino_cli_commands_v1_monitor_proto_init() }
func file_cc_arduino_cli_commands_v1_monitor_proto_init() {
	if File_cc_arduino_cli_commands_v1_monitor_proto != nil {
		return
	}
	file_cc_arduino_cli_commands_v1_common_proto_init()
	file_cc_arduino_cli_commands_v1_port_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorPortConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorPortSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumerateMonitorPortSettingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumerateMonitorPortSettingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorPortSettingDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cc_arduino_cli_commands_v1_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cc_arduino_cli_commands_v1_monitor_proto_goTypes,
		DependencyIndexes: file_cc_arduino_cli_commands_v1_monitor_proto_depIdxs,
		MessageInfos:      file_cc_arduino_cli_commands_v1_monitor_proto_msgTypes,
	}.Build()
	File_cc_arduino_cli_commands_v1_monitor_proto = out.File
	file_cc_arduino_cli_commands_v1_monitor_proto_rawDesc = nil
	file_cc_arduino_cli_commands_v1_monitor_proto_goTypes = nil
	file_cc_arduino_cli_commands_v1_monitor_proto_depIdxs = nil
}
