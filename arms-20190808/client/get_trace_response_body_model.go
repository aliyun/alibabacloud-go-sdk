// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTraceResponseBody
	GetRequestId() *string
	SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody
	GetSpans() []*GetTraceResponseBodySpans
}

type GetTraceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the trace.
	Spans []*GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceResponseBody) GetSpans() []*GetTraceResponseBodySpans {
	return s.Spans
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

func (s *GetTraceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodySpans struct {
	// The child spans of the current span.
	Children []map[string]interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The amount of time consumed by the trace. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether a method stack was provided.
	//
	// 	- `true`: A method stack was provided.
	//
	// 	- `false`: No method stack was provided.
	//
	// example:
	//
	// false
	HaveStack *bool `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	// The log events in the trace.
	LogEventList []*GetTraceResponseBodySpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	// The name of the traced span.
	//
	// example:
	//
	// /api/demo
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The ID of the parent span.
	//
	// example:
	//
	// 18
	ParentSpanId *string `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 222
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The ID of the RPC mode.
	//
	// example:
	//
	// 0
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
	// 1
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
	// arms-demo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The span ID.
	//
	// example:
	//
	// 1234
	SpanId *string `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	// The tags of the trace.
	TagEntryList []*GetTraceResponseBodySpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	// The timestamp generated when the span was generated.
	//
	// example:
	//
	// 1590388651
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// ac14001a15954493811405707d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpans) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) GetChildren() []map[string]interface{} {
	return s.Children
}

func (s *GetTraceResponseBodySpans) GetDuration() *int64 {
	return s.Duration
}

func (s *GetTraceResponseBodySpans) GetHaveStack() *bool {
	return s.HaveStack
}

func (s *GetTraceResponseBodySpans) GetLogEventList() []*GetTraceResponseBodySpansLogEventList {
	return s.LogEventList
}

func (s *GetTraceResponseBodySpans) GetOperationName() *string {
	return s.OperationName
}

func (s *GetTraceResponseBodySpans) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *GetTraceResponseBodySpans) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetTraceResponseBodySpans) GetRpcId() *string {
	return s.RpcId
}

func (s *GetTraceResponseBodySpans) GetRpcType() *int32 {
	return s.RpcType
}

func (s *GetTraceResponseBodySpans) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *GetTraceResponseBodySpans) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetTraceResponseBodySpans) GetSpanId() *string {
	return s.SpanId
}

func (s *GetTraceResponseBodySpans) GetTagEntryList() []*GetTraceResponseBodySpansTagEntryList {
	return s.TagEntryList
}

func (s *GetTraceResponseBodySpans) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetTraceResponseBodySpans) GetTraceID() *string {
	return s.TraceID
}

func (s *GetTraceResponseBodySpans) SetChildren(v []map[string]interface{}) *GetTraceResponseBodySpans {
	s.Children = v
	return s
}

func (s *GetTraceResponseBodySpans) SetDuration(v int64) *GetTraceResponseBodySpans {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetHaveStack(v bool) *GetTraceResponseBodySpans {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetLogEventList(v []*GetTraceResponseBodySpansLogEventList) *GetTraceResponseBodySpans {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetOperationName(v string) *GetTraceResponseBodySpans {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetParentSpanId(v string) *GetTraceResponseBodySpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetResultCode(v string) *GetTraceResponseBodySpans {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcId(v string) *GetTraceResponseBodySpans {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcType(v int32) *GetTraceResponseBodySpans {
	s.RpcType = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceIp(v string) *GetTraceResponseBodySpans {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceName(v string) *GetTraceResponseBodySpans {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetSpanId(v string) *GetTraceResponseBodySpans {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTagEntryList(v []*GetTraceResponseBodySpansTagEntryList) *GetTraceResponseBodySpans {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetTimestamp(v int64) *GetTraceResponseBodySpans {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTraceID(v string) *GetTraceResponseBodySpans {
	s.TraceID = &v
	return s
}

func (s *GetTraceResponseBodySpans) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodySpansLogEventList struct {
	// The tags of the trace.
	TagEntryList []*GetTraceResponseBodySpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	// The timestamp when the log event was generated.
	//
	// example:
	//
	// 1590388651
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventList) GetTagEntryList() []*GetTraceResponseBodySpansLogEventListTagEntryList {
	return s.TagEntryList
}

func (s *GetTraceResponseBodySpansLogEventList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetTraceResponseBodySpansLogEventList) SetTagEntryList(v []*GetTraceResponseBodySpansLogEventListTagEntryList) *GetTraceResponseBodySpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) SetTimestamp(v int64) *GetTraceResponseBodySpansLogEventList {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodySpansLogEventListTagEntryList struct {
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

func (s GetTraceResponseBodySpansLogEventListTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetKey(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetValue(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodySpansTagEntryList struct {
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

func (s GetTraceResponseBodySpansTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetTraceResponseBodySpansTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetTraceResponseBodySpansTagEntryList) SetKey(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) SetValue(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Value = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) Validate() error {
	return dara.Validate(s)
}
