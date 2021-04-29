// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: opentelemetry/proto/metrics/experimental/configservice.proto

package experimental

import (
	v1 "github.com/influxdata/influxdb-observability/otlp/resource/v1"
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

type MetricConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource for which configuration should be returned.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// Optional. The value of MetricConfigResponse.fingerprint for the last
	// configuration that the caller received and successfully applied.
	LastKnownFingerprint []byte `protobuf:"bytes,2,opt,name=last_known_fingerprint,json=lastKnownFingerprint,proto3" json:"last_known_fingerprint,omitempty"`
}

func (x *MetricConfigRequest) Reset() {
	*x = MetricConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricConfigRequest) ProtoMessage() {}

func (x *MetricConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricConfigRequest.ProtoReflect.Descriptor instead.
func (*MetricConfigRequest) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescGZIP(), []int{0}
}

func (x *MetricConfigRequest) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *MetricConfigRequest) GetLastKnownFingerprint() []byte {
	if x != nil {
		return x.LastKnownFingerprint
	}
	return nil
}

type MetricConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The fingerprint associated with this MetricConfigResponse. Each
	// change in configs yields a different fingerprint. The resource SHOULD copy
	// this value to MetricConfigRequest.last_known_fingerprint for the next
	// configuration request. If there are no changes between fingerprint and
	// MetricConfigRequest.last_known_fingerprint, then all other fields besides
	// fingerprint in the response are optional, or the same as the last update if
	// present.
	//
	// The exact mechanics of generating the fingerprint is up to the
	// implementation. However, a fingerprint must be deterministically determined
	// by the configurations -- the same configuration will generate the same
	// fingerprint on any instance of an implementation. Hence using a timestamp is
	// unacceptable, but a deterministic hash is fine.
	Fingerprint []byte `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// A single metric may match multiple schedules. In such cases, the schedule
	// that specifies the smallest period is applied.
	//
	// Note, for optimization purposes, it is recommended to use as few schedules
	// as possible to capture all required metric updates. Where you can be
	// conservative, do take full advantage of the inclusion/exclusion patterns to
	// capture as much of your targeted metrics.
	Schedules []*MetricConfigResponse_Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty"`
	// Optional. The client is suggested to wait this long (in seconds) before
	// pinging the configuration service again.
	SuggestedWaitTimeSec int32 `protobuf:"varint,3,opt,name=suggested_wait_time_sec,json=suggestedWaitTimeSec,proto3" json:"suggested_wait_time_sec,omitempty"`
}

func (x *MetricConfigResponse) Reset() {
	*x = MetricConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricConfigResponse) ProtoMessage() {}

func (x *MetricConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricConfigResponse.ProtoReflect.Descriptor instead.
func (*MetricConfigResponse) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescGZIP(), []int{1}
}

func (x *MetricConfigResponse) GetFingerprint() []byte {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

func (x *MetricConfigResponse) GetSchedules() []*MetricConfigResponse_Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *MetricConfigResponse) GetSuggestedWaitTimeSec() int32 {
	if x != nil {
		return x.SuggestedWaitTimeSec
	}
	return 0
}

// A Schedule is used to apply a particular scheduling configuration to
// a metric. If a metric name matches a schedule's patterns, then the metric
// adopts the configuration specified by the schedule.
type MetricConfigResponse_Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metrics with names that match a rule in the inclusion_patterns are
	// targeted by this schedule. Metrics that match the exclusion_patterns
	// are not targeted for this schedule, even if they match an inclusion
	// pattern.
	ExclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,1,rep,name=exclusion_patterns,json=exclusionPatterns,proto3" json:"exclusion_patterns,omitempty"`
	InclusionPatterns []*MetricConfigResponse_Schedule_Pattern `protobuf:"bytes,2,rep,name=inclusion_patterns,json=inclusionPatterns,proto3" json:"inclusion_patterns,omitempty"`
	// Describes the collection period for each metric in seconds.
	// A period of 0 means to not export.
	PeriodSec int32 `protobuf:"varint,3,opt,name=period_sec,json=periodSec,proto3" json:"period_sec,omitempty"`
}

func (x *MetricConfigResponse_Schedule) Reset() {
	*x = MetricConfigResponse_Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricConfigResponse_Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricConfigResponse_Schedule) ProtoMessage() {}

func (x *MetricConfigResponse_Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricConfigResponse_Schedule.ProtoReflect.Descriptor instead.
func (*MetricConfigResponse_Schedule) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MetricConfigResponse_Schedule) GetExclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if x != nil {
		return x.ExclusionPatterns
	}
	return nil
}

func (x *MetricConfigResponse_Schedule) GetInclusionPatterns() []*MetricConfigResponse_Schedule_Pattern {
	if x != nil {
		return x.InclusionPatterns
	}
	return nil
}

func (x *MetricConfigResponse_Schedule) GetPeriodSec() int32 {
	if x != nil {
		return x.PeriodSec
	}
	return 0
}

// A light-weight pattern that can match 1 or more
// metrics, for which this schedule will apply. The string is used to
// match against metric names. It should not exceed 100k characters.
type MetricConfigResponse_Schedule_Pattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Match:
	//	*MetricConfigResponse_Schedule_Pattern_Equals
	//	*MetricConfigResponse_Schedule_Pattern_StartsWith
	Match isMetricConfigResponse_Schedule_Pattern_Match `protobuf_oneof:"match"`
}

func (x *MetricConfigResponse_Schedule_Pattern) Reset() {
	*x = MetricConfigResponse_Schedule_Pattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricConfigResponse_Schedule_Pattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricConfigResponse_Schedule_Pattern) ProtoMessage() {}

func (x *MetricConfigResponse_Schedule_Pattern) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricConfigResponse_Schedule_Pattern.ProtoReflect.Descriptor instead.
func (*MetricConfigResponse_Schedule_Pattern) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (m *MetricConfigResponse_Schedule_Pattern) GetMatch() isMetricConfigResponse_Schedule_Pattern_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *MetricConfigResponse_Schedule_Pattern) GetEquals() string {
	if x, ok := x.GetMatch().(*MetricConfigResponse_Schedule_Pattern_Equals); ok {
		return x.Equals
	}
	return ""
}

func (x *MetricConfigResponse_Schedule_Pattern) GetStartsWith() string {
	if x, ok := x.GetMatch().(*MetricConfigResponse_Schedule_Pattern_StartsWith); ok {
		return x.StartsWith
	}
	return ""
}

type isMetricConfigResponse_Schedule_Pattern_Match interface {
	isMetricConfigResponse_Schedule_Pattern_Match()
}

type MetricConfigResponse_Schedule_Pattern_Equals struct {
	Equals string `protobuf:"bytes,1,opt,name=equals,proto3,oneof"` // matches the metric name exactly
}

type MetricConfigResponse_Schedule_Pattern_StartsWith struct {
	StartsWith string `protobuf:"bytes,2,opt,name=starts_with,json=startsWith,proto3,oneof"` // prefix-matches the metric name
}

func (*MetricConfigResponse_Schedule_Pattern_Equals) isMetricConfigResponse_Schedule_Pattern_Match() {
}

func (*MetricConfigResponse_Schedule_Pattern_StartsWith) isMetricConfigResponse_Schedule_Pattern_Match() {
}

var File_opentelemetry_proto_metrics_experimental_configservice_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0xd3, 0x04,
	0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x65, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x35, 0x0a, 0x17, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x69,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x57, 0x61, 0x69, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x1a, 0xfa, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x7e, 0x0a, 0x12, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x12, 0x7e, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53,
	0x65, 0x63, 0x1a, 0x4f, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x18, 0x0a,
	0x06, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x32, 0xa1, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x90, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x91, 0x01, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x42, 0x18, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x64, 0x62, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2f, 0x6f, 0x74, 0x6c, 0x70, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescData = file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDesc
)

func file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescData)
	})
	return file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDescData
}

var file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_opentelemetry_proto_metrics_experimental_configservice_proto_goTypes = []interface{}{
	(*MetricConfigRequest)(nil),                   // 0: opentelemetry.proto.metrics.experimental.MetricConfigRequest
	(*MetricConfigResponse)(nil),                  // 1: opentelemetry.proto.metrics.experimental.MetricConfigResponse
	(*MetricConfigResponse_Schedule)(nil),         // 2: opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule
	(*MetricConfigResponse_Schedule_Pattern)(nil), // 3: opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.Pattern
	(*v1.Resource)(nil),                           // 4: opentelemetry.proto.resource.v1.Resource
}
var file_opentelemetry_proto_metrics_experimental_configservice_proto_depIdxs = []int32{
	4, // 0: opentelemetry.proto.metrics.experimental.MetricConfigRequest.resource:type_name -> opentelemetry.proto.resource.v1.Resource
	2, // 1: opentelemetry.proto.metrics.experimental.MetricConfigResponse.schedules:type_name -> opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule
	3, // 2: opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.exclusion_patterns:type_name -> opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.Pattern
	3, // 3: opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.inclusion_patterns:type_name -> opentelemetry.proto.metrics.experimental.MetricConfigResponse.Schedule.Pattern
	0, // 4: opentelemetry.proto.metrics.experimental.MetricConfig.GetMetricConfig:input_type -> opentelemetry.proto.metrics.experimental.MetricConfigRequest
	1, // 5: opentelemetry.proto.metrics.experimental.MetricConfig.GetMetricConfig:output_type -> opentelemetry.proto.metrics.experimental.MetricConfigResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_metrics_experimental_configservice_proto_init() }
func file_opentelemetry_proto_metrics_experimental_configservice_proto_init() {
	if File_opentelemetry_proto_metrics_experimental_configservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricConfigRequest); i {
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
		file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricConfigResponse); i {
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
		file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricConfigResponse_Schedule); i {
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
		file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricConfigResponse_Schedule_Pattern); i {
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
	file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MetricConfigResponse_Schedule_Pattern_Equals)(nil),
		(*MetricConfigResponse_Schedule_Pattern_StartsWith)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opentelemetry_proto_metrics_experimental_configservice_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_metrics_experimental_configservice_proto_depIdxs,
		MessageInfos:      file_opentelemetry_proto_metrics_experimental_configservice_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_metrics_experimental_configservice_proto = out.File
	file_opentelemetry_proto_metrics_experimental_configservice_proto_rawDesc = nil
	file_opentelemetry_proto_metrics_experimental_configservice_proto_goTypes = nil
	file_opentelemetry_proto_metrics_experimental_configservice_proto_depIdxs = nil
}
