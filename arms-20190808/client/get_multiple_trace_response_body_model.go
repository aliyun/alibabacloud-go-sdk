// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipleTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody
	GetMultiCallChainInfos() []*GetMultipleTraceResponseBodyMultiCallChainInfos
	SetRequestId(v string) *GetMultipleTraceResponseBody
	GetRequestId() *string
}

type GetMultipleTraceResponseBody struct {
	// The details of traces.
	MultiCallChainInfos []*GetMultipleTraceResponseBodyMultiCallChainInfos `json:"MultiCallChainInfos,omitempty" xml:"MultiCallChainInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2983BEF7-4A0D-47A2-94A2-8E9C5E63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultipleTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBody) GetMultiCallChainInfos() []*GetMultipleTraceResponseBodyMultiCallChainInfos {
	return s.MultiCallChainInfos
}

func (s *GetMultipleTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultipleTraceResponseBody) SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody {
	s.MultiCallChainInfos = v
	return s
}

func (s *GetMultipleTraceResponseBody) SetRequestId(v string) *GetMultipleTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultipleTraceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfos struct {
	// The details of the trace.
	Spans []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
	// The trace ID.
	//
	// example:
	//
	// ac1400a115951745017447033d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) GetSpans() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	return s.Spans
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) GetTraceID() *string {
	return s.TraceID
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetSpans(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.Spans = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.TraceID = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpans struct {
	// The amount of time consumed by the trace. Unit: milliseconds.
	//
	// example:
	//
	// 11
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether a method stack was provided.
	//
	// 	- `true`: A method stack was provided.
	//
	// 	- `false`: No method stack was provided.
	//
	// example:
	//
	// true
	HaveStack *bool `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	// The log events in the trace.
	LogEventList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	// The name of the traced span.
	//
	// example:
	//
	// /demo/queryNotExistDB/11
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The ID of the parent span.
	//
	// example:
	//
	// 18
	ParentSpanId *string `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	// The status code returned.
	//
	// example:
	//
	// 1
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// RPC ID
	//
	// example:
	//
	// 0.1
	RpcId *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	// The type of the remote procedure call (RPC) mode.
	//
	// 	- 0: HTTP entry
	//
	// 	- 25: HTTP call
	//
	// 	- 1: High-speed Service Framework (HSF) call
	//
	// 	- 2: HSF provision
	//
	// 	- 40: on-premises API call
	//
	// 	- 60: MySQL call
	//
	// 	- 62: Oracle call
	//
	// 	- 63: PostgreSQL call
	//
	// 	- 70: Redis call
	//
	// 	- 4: Taobao Distributed Data Layer (TDDL) call
	//
	// 	- 5: Tair call
	//
	// 	- 13: MetaQ message sending
	//
	// 	- 252: MetaQ message receiving
	//
	// 	- 3: notification sending
	//
	// 	- 254: notification receiving
	//
	// 	- 7: Apache Dubbo call
	//
	// 	- 8: Apache Dubbo provision
	//
	// 	- 19: SOFARPC call
	//
	// 	- 18: SOFARPC provision
	//
	// 	- 11: Distributed Service Framework (DSF) call
	//
	// 	- 12: DSF provision
	//
	// 	- \\-1: unknown call
	//
	// example:
	//
	// 0
	RpcType *int32 `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	// The IP address of the host where the application resides.
	//
	// example:
	//
	// 172.20.XX.XX
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// arms-k8s-demo-subcomponent
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The span ID.
	//
	// example:
	//
	// 1234
	SpanId *string `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	// The tags of the trace.
	TagEntryList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	// The timestamp.
	//
	// example:
	//
	// 1595174501747
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// ac1400a115951745017447033d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetDuration() *int64 {
	return s.Duration
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetHaveStack() *bool {
	return s.HaveStack
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetLogEventList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	return s.LogEventList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetOperationName() *string {
	return s.OperationName
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetRpcId() *string {
	return s.RpcId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetRpcType() *int32 {
	return s.RpcType
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetSpanId() *string {
	return s.SpanId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTagEntryList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	return s.TagEntryList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTraceID() *string {
	return s.TraceID
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetDuration(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Duration = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetHaveStack(v bool) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.HaveStack = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetLogEventList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.LogEventList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetOperationName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.OperationName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetParentSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetResultCode(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ResultCode = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcType(v int32) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcType = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceIp(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceIp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.SpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TraceID = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList struct {
	// The tags of the trace.
	TagEntryList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	// The time when the log was generated. The value is a timestamp.
	//
	// example:
	//
	// 1595174501747
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GetTagEntryList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	return s.TagEntryList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList struct {
	// The key of the tag.
	//
	// example:
	//
	// http.status.code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList struct {
	// The key of the tag.
	//
	// example:
	//
	// http.status.code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Value = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) Validate() error {
	return dara.Validate(s)
}
