// The MIT License
//
// Copyright (c) 2024 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/persistence/v1/hsm.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A node in a hierarchical state machine tree.
type StateMachineNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of transitions on this state machine object.
	// Used to verify that a task is not stale if the state machine does not allow concurrent task execution.
	TransitionCount int64 `protobuf:"varint,1,opt,name=transition_count,json=transitionCount,proto3" json:"transition_count,omitempty"`
	// Seralized data of the underlying state machine.
	Data                            []byte                     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Children                        map[int32]*StateMachineMap `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InitialNamespaceFailoverVersion int64                      `protobuf:"varint,4,opt,name=initial_namespace_failover_version,json=initialNamespaceFailoverVersion,proto3" json:"initial_namespace_failover_version,omitempty"`
	CurrentNamespaceFailoverVersion int64                      `protobuf:"varint,5,opt,name=current_namespace_failover_version,json=currentNamespaceFailoverVersion,proto3" json:"current_namespace_failover_version,omitempty"`
}

func (x *StateMachineNode) Reset() {
	*x = StateMachineNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineNode) ProtoMessage() {}

func (x *StateMachineNode) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineNode.ProtoReflect.Descriptor instead.
func (*StateMachineNode) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{0}
}

func (x *StateMachineNode) GetTransitionCount() int64 {
	if x != nil {
		return x.TransitionCount
	}
	return 0
}

func (x *StateMachineNode) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StateMachineNode) GetChildren() map[int32]*StateMachineMap {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *StateMachineNode) GetInitialNamespaceFailoverVersion() int64 {
	if x != nil {
		return x.InitialNamespaceFailoverVersion
	}
	return 0
}

func (x *StateMachineNode) GetCurrentNamespaceFailoverVersion() int64 {
	if x != nil {
		return x.CurrentNamespaceFailoverVersion
	}
	return 0
}

// Map of state machine ID to StateMachineNode.
type StateMachineMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (-- api-linter: core::0140::prepositions=disabled
	//
	//	aip.dev/not-precedent: "by" is used to clarify the keys and values. --)
	MachinesById map[string]*StateMachineNode `protobuf:"bytes,1,rep,name=machines_by_id,json=machinesById,proto3" json:"machines_by_id,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StateMachineMap) Reset() {
	*x = StateMachineMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineMap) ProtoMessage() {}

func (x *StateMachineMap) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineMap.ProtoReflect.Descriptor instead.
func (*StateMachineMap) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{1}
}

func (x *StateMachineMap) GetMachinesById() map[string]*StateMachineNode {
	if x != nil {
		return x.MachinesById
	}
	return nil
}

type StateMachineKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Addressable type of the corresponding state machine in a single tree level.
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Addressable ID of the corresponding state machine in a single tree level.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StateMachineKey) Reset() {
	*x = StateMachineKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineKey) ProtoMessage() {}

func (x *StateMachineKey) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineKey.ProtoReflect.Descriptor instead.
func (*StateMachineKey) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{2}
}

func (x *StateMachineKey) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *StateMachineKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// A reference to a state machine at a point in time.
type StateMachineRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nested path to a state machine.
	Path []*StateMachineKey `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Namespace failover version on the corresponding mutable state object, used for staleness detection when global
	// namespaces are enabled.
	// this is also the current namespace failover version when this ref is generated.
	MutableStateNamespaceFailoverVersion int64 `protobuf:"varint,2,opt,name=mutable_state_namespace_failover_version,json=mutableStateNamespaceFailoverVersion,proto3" json:"mutable_state_namespace_failover_version,omitempty"`
	// Number of transitions on the corresponding mutable state object. Used to verify that a task is not referencing a
	// stale state or, in some situations, that the task itself is not stale.
	MutableStateTransitionCount int64 `protobuf:"varint,3,opt,name=mutable_state_transition_count,json=mutableStateTransitionCount,proto3" json:"mutable_state_transition_count,omitempty"`
	// Number of transitions executed on the referenced state machine node at the time this Ref is instantiated.
	// If non-zero, the state machine is assumed to support only non-concurrent tasks, and this number should match the
	// number of state transitions on the corresponding state machine object.
	MachineTransitionCount          int64 `protobuf:"varint,4,opt,name=machine_transition_count,json=machineTransitionCount,proto3" json:"machine_transition_count,omitempty"`
	InitialNamespaceFailoverVersion int64 `protobuf:"varint,5,opt,name=initial_namespace_failover_version,json=initialNamespaceFailoverVersion,proto3" json:"initial_namespace_failover_version,omitempty"`
}

func (x *StateMachineRef) Reset() {
	*x = StateMachineRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineRef) ProtoMessage() {}

func (x *StateMachineRef) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineRef.ProtoReflect.Descriptor instead.
func (*StateMachineRef) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{3}
}

func (x *StateMachineRef) GetPath() []*StateMachineKey {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StateMachineRef) GetMutableStateNamespaceFailoverVersion() int64 {
	if x != nil {
		return x.MutableStateNamespaceFailoverVersion
	}
	return 0
}

func (x *StateMachineRef) GetMutableStateTransitionCount() int64 {
	if x != nil {
		return x.MutableStateTransitionCount
	}
	return 0
}

func (x *StateMachineRef) GetMachineTransitionCount() int64 {
	if x != nil {
		return x.MachineTransitionCount
	}
	return 0
}

func (x *StateMachineRef) GetInitialNamespaceFailoverVersion() int64 {
	if x != nil {
		return x.InitialNamespaceFailoverVersion
	}
	return 0
}

type StateMachineTaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to a state machine.
	Ref *StateMachineRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// Task type. Not to be confused with the state machine's type in the `ref` field.
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// Opaque data attached to this task. May be nil. Deserialized by a registered TaskSerializer for this type.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StateMachineTaskInfo) Reset() {
	*x = StateMachineTaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineTaskInfo) ProtoMessage() {}

func (x *StateMachineTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineTaskInfo.ProtoReflect.Descriptor instead.
func (*StateMachineTaskInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{4}
}

func (x *StateMachineTaskInfo) GetRef() *StateMachineRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *StateMachineTaskInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *StateMachineTaskInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Keeps track of the ranges of transition counts per namespace failover version.
// Each task generated by the HSM framework is imprinted with the current transaction’s `NamespaceFailoverVersion` and
// `MaxTransitionCount` at the end of the transaction.
// When a task is being processed, the `StateTransitionHistory` is compared with the imprinted task information to
// verify that a task is not referencing a stale state or that the task itself is not stale.
// For example, if the state has a history of `[{v: 1, t: 3}, {v: 2, t: 5}]`, task A `{v: 2, t: 4}` **is not**
// referencing stale state because for version `2` transitions `4-5` are valid, while task B `{v: 2, t: 6}` **is**
// referencing stale state because the transition count is out of range for version `2`.
// Furthermore, task C `{v: 1, t: 4}` itself is stale because it is referencing an impossible state, likely due to post
// split-brain reconciliation.
type VersionedTransition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace failover version at transition time.
	NamespaceFailoverVersion int64 `protobuf:"varint,1,opt,name=namespace_failover_version,json=namespaceFailoverVersion,proto3" json:"namespace_failover_version,omitempty"`
	// Maximum state transition count perceived during the specified namespace_failover_version.
	MaxTransitionCount int64 `protobuf:"varint,2,opt,name=max_transition_count,json=maxTransitionCount,proto3" json:"max_transition_count,omitempty"`
}

func (x *VersionedTransition) Reset() {
	*x = VersionedTransition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionedTransition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionedTransition) ProtoMessage() {}

func (x *VersionedTransition) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionedTransition.ProtoReflect.Descriptor instead.
func (*VersionedTransition) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP(), []int{5}
}

func (x *VersionedTransition) GetNamespaceFailoverVersion() int64 {
	if x != nil {
		return x.NamespaceFailoverVersion
	}
	return 0
}

func (x *VersionedTransition) GetMaxTransitionCount() int64 {
	if x != nil {
		return x.MaxTransitionCount
	}
	return 0
}

var File_temporal_server_api_persistence_v1_hsm_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_hsm_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xd9, 0x03,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x2d, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x16, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x62, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42,
	0x02, 0x68, 0x00, 0x12, 0x4f, 0x0a, 0x22, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1f, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x4f, 0x0a, 0x22, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x78, 0x0a, 0x0d,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x4d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x61,
	0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x81, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x6f, 0x0a, 0x0e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x7d, 0x0a,
	0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02,
	0x68, 0x00, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x12,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x02, 0x68,
	0x00, 0x22, 0x92, 0x03, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x66, 0x12, 0x4b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x02, 0x68, 0x00, 0x12, 0x5a, 0x0a,
	0x28, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x24, 0x6d, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x47, 0x0a, 0x1e, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3c, 0x0a, 0x18, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x4f, 0x0a, 0x22, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1f, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68,
	0x00, 0x22, 0x91, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x66, 0x52, 0x03, 0x72, 0x65, 0x66, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x16, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x16, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x02, 0x68, 0x00, 0x22, 0x8d, 0x01, 0x0a,
	0x13, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x1a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x34, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x6d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x02, 0x68, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_hsm_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_hsm_proto_rawDescData = file_temporal_server_api_persistence_v1_hsm_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_hsm_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_hsm_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_hsm_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_hsm_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_hsm_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_hsm_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_temporal_server_api_persistence_v1_hsm_proto_goTypes = []interface{}{
	(*StateMachineNode)(nil),     // 0: temporal.server.api.persistence.v1.StateMachineNode
	(*StateMachineMap)(nil),      // 1: temporal.server.api.persistence.v1.StateMachineMap
	(*StateMachineKey)(nil),      // 2: temporal.server.api.persistence.v1.StateMachineKey
	(*StateMachineRef)(nil),      // 3: temporal.server.api.persistence.v1.StateMachineRef
	(*StateMachineTaskInfo)(nil), // 4: temporal.server.api.persistence.v1.StateMachineTaskInfo
	(*VersionedTransition)(nil),  // 5: temporal.server.api.persistence.v1.VersionedTransition
	nil,                          // 6: temporal.server.api.persistence.v1.StateMachineNode.ChildrenEntry
	nil,                          // 7: temporal.server.api.persistence.v1.StateMachineMap.MachinesByIdEntry
}
var file_temporal_server_api_persistence_v1_hsm_proto_depIdxs = []int32{
	6, // 0: temporal.server.api.persistence.v1.StateMachineNode.children:type_name -> temporal.server.api.persistence.v1.StateMachineNode.ChildrenEntry
	7, // 1: temporal.server.api.persistence.v1.StateMachineMap.machines_by_id:type_name -> temporal.server.api.persistence.v1.StateMachineMap.MachinesByIdEntry
	2, // 2: temporal.server.api.persistence.v1.StateMachineRef.path:type_name -> temporal.server.api.persistence.v1.StateMachineKey
	3, // 3: temporal.server.api.persistence.v1.StateMachineTaskInfo.ref:type_name -> temporal.server.api.persistence.v1.StateMachineRef
	1, // 4: temporal.server.api.persistence.v1.StateMachineNode.ChildrenEntry.value:type_name -> temporal.server.api.persistence.v1.StateMachineMap
	0, // 5: temporal.server.api.persistence.v1.StateMachineMap.MachinesByIdEntry.value:type_name -> temporal.server.api.persistence.v1.StateMachineNode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_hsm_proto_init() }
func file_temporal_server_api_persistence_v1_hsm_proto_init() {
	if File_temporal_server_api_persistence_v1_hsm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineNode); i {
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
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineMap); i {
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
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineKey); i {
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
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineRef); i {
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
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineTaskInfo); i {
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
		file_temporal_server_api_persistence_v1_hsm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionedTransition); i {
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
			RawDescriptor: file_temporal_server_api_persistence_v1_hsm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_hsm_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_hsm_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_hsm_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_hsm_proto = out.File
	file_temporal_server_api_persistence_v1_hsm_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_hsm_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_hsm_proto_depIdxs = nil
}
