// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckCommercialStatusRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xtrace
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s CheckCommercialStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCommercialStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusRequest) SetRegionId(v string) *CheckCommercialStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCommercialStatusRequest) SetService(v string) *CheckCommercialStatusRequest {
	s.Service = &v
	return s
}

type CheckCommercialStatusResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1FC47DED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCommercialStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCommercialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusResponseBody) SetData(v string) *CheckCommercialStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckCommercialStatusResponseBody) SetRequestId(v string) *CheckCommercialStatusResponseBody {
	s.RequestId = &v
	return s
}

type CheckCommercialStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCommercialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCommercialStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCommercialStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusResponse) SetHeaders(v map[string]*string) *CheckCommercialStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckCommercialStatusResponse) SetStatusCode(v int32) *CheckCommercialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCommercialStatusResponse) SetBody(v *CheckCommercialStatusResponseBody) *CheckCommercialStatusResponse {
	s.Body = v
	return s
}

type GetTagKeyRequest struct {
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575622455686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// appTest
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The name of the span.
	//
	// example:
	//
	// createOrder
	SpanName *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575561600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetTagKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyRequest) GoString() string {
	return s.String()
}

func (s *GetTagKeyRequest) SetEndTime(v int64) *GetTagKeyRequest {
	s.EndTime = &v
	return s
}

func (s *GetTagKeyRequest) SetRegionId(v string) *GetTagKeyRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagKeyRequest) SetServiceName(v string) *GetTagKeyRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagKeyRequest) SetSpanName(v string) *GetTagKeyRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagKeyRequest) SetStartTime(v int64) *GetTagKeyRequest {
	s.StartTime = &v
	return s
}

type GetTagKeyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys *GetTagKeyResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s GetTagKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBody) SetRequestId(v string) *GetTagKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagKeyResponseBody) SetTagKeys(v *GetTagKeyResponseBodyTagKeys) *GetTagKeyResponseBody {
	s.TagKeys = v
	return s
}

type GetTagKeyResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s GetTagKeyResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseBodyTagKeys) SetTagKey(v []*string) *GetTagKeyResponseBodyTagKeys {
	s.TagKey = v
	return s
}

type GetTagKeyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponse) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponse) SetHeaders(v map[string]*string) *GetTagKeyResponse {
	s.Headers = v
	return s
}

func (s *GetTagKeyResponse) SetStatusCode(v int32) *GetTagKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagKeyResponse) SetBody(v *GetTagKeyResponseBody) *GetTagKeyResponse {
	s.Body = v
	return s
}

type GetTagValRequest struct {
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575622455686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// appTest
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The name of the span.
	//
	// example:
	//
	// createOrder
	SpanName *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575561600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// span.kind
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s GetTagValRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagValRequest) GoString() string {
	return s.String()
}

func (s *GetTagValRequest) SetEndTime(v int64) *GetTagValRequest {
	s.EndTime = &v
	return s
}

func (s *GetTagValRequest) SetRegionId(v string) *GetTagValRequest {
	s.RegionId = &v
	return s
}

func (s *GetTagValRequest) SetServiceName(v string) *GetTagValRequest {
	s.ServiceName = &v
	return s
}

func (s *GetTagValRequest) SetSpanName(v string) *GetTagValRequest {
	s.SpanName = &v
	return s
}

func (s *GetTagValRequest) SetStartTime(v int64) *GetTagValRequest {
	s.StartTime = &v
	return s
}

func (s *GetTagValRequest) SetTagKey(v string) *GetTagValRequest {
	s.TagKey = &v
	return s
}

type GetTagValResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues *GetTagValResponseBodyTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Struct"`
}

func (s GetTagValResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBody) SetRequestId(v string) *GetTagValResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagValResponseBody) SetTagValues(v *GetTagValResponseBodyTagValues) *GetTagValResponseBody {
	s.TagValues = v
	return s
}

type GetTagValResponseBodyTagValues struct {
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s GetTagValResponseBodyTagValues) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseBodyTagValues) GoString() string {
	return s.String()
}

func (s *GetTagValResponseBodyTagValues) SetTagValue(v []*string) *GetTagValResponseBodyTagValues {
	s.TagValue = v
	return s
}

type GetTagValResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagValResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagValResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponse) GoString() string {
	return s.String()
}

func (s *GetTagValResponse) SetHeaders(v map[string]*string) *GetTagValResponse {
	s.Headers = v
	return s
}

func (s *GetTagValResponse) SetStatusCode(v int32) *GetTagValResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagValResponse) SetBody(v *GetTagValResponseBody) *GetTagValResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	// The type of the application. You can set the value to **XTRACE*	- or leave this parameter unspecified.
	//
	// example:
	//
	// XTRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The unique ID of the trace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1c6881aab84191a4
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetAppType(v string) *GetTraceRequest {
	s.AppType = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

type GetTraceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the trace.
	Spans *GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Struct"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v *GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

type GetTraceResponseBodySpans struct {
	Span []*GetTraceResponseBodySpansSpan `json:"Span,omitempty" xml:"Span,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) SetSpan(v []*GetTraceResponseBodySpansSpan) *GetTraceResponseBodySpans {
	s.Span = v
	return s
}

type GetTraceResponseBodySpansSpan struct {
	// The time used to call the trace. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether the span has child spans. Valid values:
	//
	// - true: The span has child spans.
	//
	// - false: The span has no child spans.
	//
	// example:
	//
	// false
	HaveStack *bool `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	// The log events in the trace.
	LogEventList *GetTraceResponseBodySpansSpanLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Struct"`
	// The name of the span.
	//
	// example:
	//
	// /api
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The ID of the parent span.
	//
	// example:
	//
	// fec891bb8f8XXX
	ParentSpanId *string `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The parent-child and sibling relationship between spans. For example, span 1.1 is the parent of span 1.1.1, and span 1.1.2 and span 1.1.1 are siblings.
	//
	// example:
	//
	// 1.1
	RpcId *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	// The IP address of the server where the span resides.
	//
	// example:
	//
	// 192.168.XXX.XXX
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// server1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Span ID.
	//
	// example:
	//
	// fec891bb8f8XXX
	SpanId *string `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	// The tags in the span.
	TagEntryList *GetTraceResponseBodySpansSpanTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	// The timestamp when the span was generated. Unit: microseconds.
	//
	// example:
	//
	// 1689845513298000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The unique ID of the trace.
	//
	// example:
	//
	// 1c6881aab84191a4****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpansSpan) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpan) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpan) SetDuration(v int64) *GetTraceResponseBodySpansSpan {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetHaveStack(v bool) *GetTraceResponseBodySpansSpan {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetLogEventList(v *GetTraceResponseBodySpansSpanLogEventList) *GetTraceResponseBodySpansSpan {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetOperationName(v string) *GetTraceResponseBodySpansSpan {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetParentSpanId(v string) *GetTraceResponseBodySpansSpan {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetResultCode(v string) *GetTraceResponseBodySpansSpan {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetRpcId(v string) *GetTraceResponseBodySpansSpan {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetServiceIp(v string) *GetTraceResponseBodySpansSpan {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetServiceName(v string) *GetTraceResponseBodySpansSpan {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetSpanId(v string) *GetTraceResponseBodySpansSpan {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTagEntryList(v *GetTraceResponseBodySpansSpanTagEntryList) *GetTraceResponseBodySpansSpan {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTimestamp(v int64) *GetTraceResponseBodySpansSpan {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpansSpan) SetTraceID(v string) *GetTraceResponseBodySpansSpan {
	s.TraceID = &v
	return s
}

type GetTraceResponseBodySpansSpanLogEventList struct {
	LogEvent []*GetTraceResponseBodySpansSpanLogEventListLogEvent `json:"LogEvent,omitempty" xml:"LogEvent,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventList) SetLogEvent(v []*GetTraceResponseBodySpansSpanLogEventListLogEvent) *GetTraceResponseBodySpansSpanLogEventList {
	s.LogEvent = v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEvent struct {
	// The tags in the log event.
	TagEntryList *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	// The timestamp when the log event was generated.
	//
	// example:
	//
	// 1583683202047000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEvent) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEvent) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEvent) SetTagEntryList(v *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) *GetTraceResponseBodySpansSpanLogEventListLogEvent {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEvent) SetTimestamp(v int64) *GetTraceResponseBodySpansSpanLogEventListLogEvent {
	s.Timestamp = &v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList struct {
	TagEntry []*GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList) SetTagEntry(v []*GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry struct {
	// The tag key in the log event.
	//
	// example:
	//
	// logLevel
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value in the log event.
	//
	// example:
	//
	// Warning
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponseBodySpansSpanTagEntryList struct {
	TagEntry []*GetTraceResponseBodySpansSpanTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodySpansSpanTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanTagEntryList) SetTagEntry(v []*GetTraceResponseBodySpansSpanTagEntryListTagEntry) *GetTraceResponseBodySpansSpanTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseBodySpansSpanTagEntryListTagEntry struct {
	// The tag key in the span.
	//
	// example:
	//
	// logLevel
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value in the span.
	//
	// example:
	//
	// Warning
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansSpanTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansSpanTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansSpanTagEntryListTagEntry) SetKey(v string) *GetTraceResponseBodySpansSpanTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansSpanTagEntryListTagEntry) SetValue(v string) *GetTraceResponseBodySpansSpanTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type ListIpOrHostsRequest struct {
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1583723776974
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the application. If you do not set this parameter, the IP addresses of all applications in the specified region are returned.
	//
	// example:
	//
	// service1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1583683200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIpOrHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsRequest) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsRequest) SetEndTime(v int64) *ListIpOrHostsRequest {
	s.EndTime = &v
	return s
}

func (s *ListIpOrHostsRequest) SetRegionId(v string) *ListIpOrHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpOrHostsRequest) SetServiceName(v string) *ListIpOrHostsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListIpOrHostsRequest) SetStartTime(v int64) *ListIpOrHostsRequest {
	s.StartTime = &v
	return s
}

type ListIpOrHostsResponseBody struct {
	// The IP addresses.
	IpNames *ListIpOrHostsResponseBodyIpNames `json:"IpNames,omitempty" xml:"IpNames,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1FC4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIpOrHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponseBody) SetIpNames(v *ListIpOrHostsResponseBodyIpNames) *ListIpOrHostsResponseBody {
	s.IpNames = v
	return s
}

func (s *ListIpOrHostsResponseBody) SetRequestId(v string) *ListIpOrHostsResponseBody {
	s.RequestId = &v
	return s
}

type ListIpOrHostsResponseBodyIpNames struct {
	IpName []*string `json:"IpName,omitempty" xml:"IpName,omitempty" type:"Repeated"`
}

func (s ListIpOrHostsResponseBodyIpNames) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponseBodyIpNames) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponseBodyIpNames) SetIpName(v []*string) *ListIpOrHostsResponseBodyIpNames {
	s.IpName = v
	return s
}

type ListIpOrHostsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpOrHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpOrHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponse) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponse) SetHeaders(v map[string]*string) *ListIpOrHostsResponse {
	s.Headers = v
	return s
}

func (s *ListIpOrHostsResponse) SetStatusCode(v int32) *ListIpOrHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpOrHostsResponse) SetBody(v *ListIpOrHostsResponseBody) *ListIpOrHostsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// The type of the application. You can set the value to **XTRACE*	- or leave this parameter unspecified.
	//
	// example:
	//
	// XTRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetAppType(v string) *ListServicesRequest {
	s.AppType = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

type ListServicesResponseBody struct {
	// The ID of the region.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1FC47DED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The applications.
	Services *ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Struct"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v *ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

type ListServicesResponseBodyServices struct {
	Service []*ListServicesResponseBodyServicesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetService(v []*ListServicesResponseBodyServicesService) *ListServicesResponseBodyServices {
	s.Service = v
	return s
}

type ListServicesResponseBodyServicesService struct {
	// The ID of the application.
	//
	// example:
	//
	// XXXqn3ly@741623b4e915df8
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// a3
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListServicesResponseBodyServicesService) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesService) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesService) SetPid(v string) *ListServicesResponseBodyServicesService {
	s.Pid = &v
	return s
}

func (s *ListServicesResponseBodyServicesService) SetRegionId(v string) *ListServicesResponseBodyServicesService {
	s.RegionId = &v
	return s
}

func (s *ListServicesResponseBodyServicesService) SetServiceName(v string) *ListServicesResponseBodyServicesService {
	s.ServiceName = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSpanNamesRequest struct {
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575622455686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// service 1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// example:
	//
	// 1575561600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSpanNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesRequest) GoString() string {
	return s.String()
}

func (s *ListSpanNamesRequest) SetEndTime(v int64) *ListSpanNamesRequest {
	s.EndTime = &v
	return s
}

func (s *ListSpanNamesRequest) SetRegionId(v string) *ListSpanNamesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSpanNamesRequest) SetServiceName(v string) *ListSpanNamesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListSpanNamesRequest) SetStartTime(v int64) *ListSpanNamesRequest {
	s.StartTime = &v
	return s
}

type ListSpanNamesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The span names.
	SpanNames *ListSpanNamesResponseBodySpanNames `json:"SpanNames,omitempty" xml:"SpanNames,omitempty" type:"Struct"`
}

func (s ListSpanNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponseBody) SetRequestId(v string) *ListSpanNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpanNamesResponseBody) SetSpanNames(v *ListSpanNamesResponseBodySpanNames) *ListSpanNamesResponseBody {
	s.SpanNames = v
	return s
}

type ListSpanNamesResponseBodySpanNames struct {
	SpanName []*string `json:"SpanName,omitempty" xml:"SpanName,omitempty" type:"Repeated"`
}

func (s ListSpanNamesResponseBodySpanNames) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponseBodySpanNames) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponseBodySpanNames) SetSpanName(v []*string) *ListSpanNamesResponseBodySpanNames {
	s.SpanName = v
	return s
}

type ListSpanNamesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpanNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSpanNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponse) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponse) SetHeaders(v map[string]*string) *ListSpanNamesResponse {
	s.Headers = v
	return s
}

func (s *ListSpanNamesResponse) SetStatusCode(v int32) *ListSpanNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpanNamesResponse) SetBody(v *ListSpanNamesResponseBody) *ListSpanNamesResponse {
	s.Body = v
	return s
}

type OpenXtraceServiceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenXtraceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenXtraceServiceRequest) SetRegionId(v string) *OpenXtraceServiceRequest {
	s.RegionId = &v
	return s
}

type OpenXtraceServiceResponseBody struct {
	// example:
	//
	// 155709986
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1FC4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s OpenXtraceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenXtraceServiceResponseBody) SetOrderId(v string) *OpenXtraceServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenXtraceServiceResponseBody) SetRequestId(v string) *OpenXtraceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenXtraceServiceResponseBody) SetResult(v string) *OpenXtraceServiceResponseBody {
	s.Result = &v
	return s
}

type OpenXtraceServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenXtraceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenXtraceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenXtraceServiceResponse) SetHeaders(v map[string]*string) *OpenXtraceServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenXtraceServiceResponse) SetStatusCode(v int32) *OpenXtraceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenXtraceServiceResponse) SetBody(v *OpenXtraceServiceResponseBody) *OpenXtraceServiceResponse {
	s.Body = v
	return s
}

type QueryMetricRequest struct {
	// The dimensions of the metric that you want to query.
	//
	// example:
	//
	// RT
	Dimensions []*string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1575622455686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter conditions.
	Filters []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The time interval at which you want to query metric data. Unit: milliseconds. Minimum value: 60000.
	//
	// > If you set this parameter to 2147483647, all data in the specified time interval is returned.
	//
	// example:
	//
	// 100000
	IntervalInSec *int32 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	// The maximum number of entries that you want to return.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The measures of the metric that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// count
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// The name of the metric. Valid values:
	//
	// - `appstat.incall`: trace statistics
	//
	// - `appstat.sql`: SQL statistics
	//
	// This parameter is required.
	//
	// example:
	//
	// appstat.incall
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The order in which you want to sort the returned entries. Valid values:
	//
	// - ASC: ascending order
	//
	// - DESC: descending order
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field based on which you want to sort the returned entries.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The ID of the proxy user.
	//
	// example:
	//
	// testefgag12
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1575561600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetIntervalInSec(v int32) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int32) *QueryMetricRequest {
	s.Limit = &v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricRequest) SetMetric(v string) *QueryMetricRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricRequest) SetOrder(v string) *QueryMetricRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

type QueryMetricRequestFilters struct {
	// The key of the field that you want to use to filter the returned entries.
	//
	// example:
	//
	// http.status_code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the field that you want to use to filter the returned entries.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricRequestFilters) SetKey(v string) *QueryMetricRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricRequestFilters) SetValue(v string) *QueryMetricRequestFilters {
	s.Value = &v
	return s
}

type QueryMetricResponseBody struct {
	// The returned statistics.
	//
	// example:
	//
	// {   "RequestId": "E2373982-D8CD-413D-B991-8EB678******",   "Data": "{\\"data\\":[{\\"date\\":1583686800000,\\"count\\":0,\\"rt\\":0,\\"rpc\\":\\"childSpan3\\"}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricResponseBody) SetData(v string) *QueryMetricResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMetricResponseBody) SetRequestId(v string) *QueryMetricResponseBody {
	s.RequestId = &v
	return s
}

type QueryMetricResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricResponse) SetHeaders(v map[string]*string) *QueryMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricResponse) SetStatusCode(v int32) *QueryMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricResponse) SetBody(v *QueryMetricResponseBody) *QueryMetricResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	// The type of the application. You can set the value to **XTRACE*	- or leave this parameter unspecified.
	//
	// example:
	//
	// XTRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The timestamp of the end time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1575622455686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time more than which is used to call the trace. Unit: milliseconds. For example, a value of 100 specifies to return the traces that more than 100 milliseconds are used to call.
	//
	// example:
	//
	// 1000
	MinDuration *int64 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The name of the span.
	//
	// example:
	//
	// /api
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The number of the page to return. For example, a value of 5 indicates page 5.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to sort the query results in chronological order or reverse chronological order. Default value: false. Valid values:
	//
	// - true: reverse chronological order
	//
	// - false: chronological order
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The IP address that corresponds to the span.
	//
	// example:
	//
	// 10.0.0.0
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// service 1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The timestamp of the start time of the time range to query. The timestamp is accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1575561600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of the tags.
	Tag []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetAppType(v string) *SearchTracesRequest {
	s.AppType = &v
	return s
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetPageNumber(v int32) *SearchTracesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesRequest) SetPageSize(v int32) *SearchTracesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

type SearchTracesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// http.status_cod
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

type SearchTracesResponseBody struct {
	// The information about the returned page.
	PageBean *SearchTracesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1E2B6A4C-6B83-4062-8B6F-AEEC1F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) SetPageBean(v *SearchTracesResponseBodyPageBean) *SearchTracesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

type SearchTracesResponseBodyPageBean struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the trace.
	TraceInfos *SearchTracesResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Struct"`
}

func (s SearchTracesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetPageSize(v int32) *SearchTracesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetTotalCount(v int64) *SearchTracesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTracesResponseBodyPageBean) SetTraceInfos(v *SearchTracesResponseBodyPageBeanTraceInfos) *SearchTracesResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

type SearchTracesResponseBodyPageBeanTraceInfos struct {
	TraceInfo []*SearchTracesResponseBodyPageBeanTraceInfosTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBodyPageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBeanTraceInfos) SetTraceInfo(v []*SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) *SearchTracesResponseBodyPageBeanTraceInfos {
	s.TraceInfo = v
	return s
}

type SearchTracesResponseBodyPageBeanTraceInfosTraceInfo struct {
	// The time used to call the trace. Unit: milliseconds.
	//
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the span.
	//
	// example:
	//
	// /api
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The IP address of the server where the span resides.
	//
	// example:
	//
	// 192.163.XXX.XXX
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// service1
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The map of tags.
	//
	// example:
	//
	// {"env":"dev"}
	TagMap map[string]interface{} `json:"TagMap,omitempty" xml:"TagMap,omitempty"`
	// The time when the span was generated. Unit: microseconds.
	//
	// example:
	//
	// 1575561600000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 1c6881aab84191a4
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetDuration(v int64) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetOperationName(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetServiceIp(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetServiceName(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetTagMap(v map[string]interface{}) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.TagMap = v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetTimestamp(v int64) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo) SetTraceID(v string) *SearchTracesResponseBodyPageBeanTraceInfosTraceInfo {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetStatusCode(v int32) *SearchTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("xtrace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查商业化状态
//
// @param request - CheckCommercialStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCommercialStatusResponse
func (client *Client) CheckCommercialStatusWithOptions(request *CheckCommercialStatusRequest, runtime *util.RuntimeOptions) (_result *CheckCommercialStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCommercialStatus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckCommercialStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查商业化状态
//
// @param request - CheckCommercialStatusRequest
//
// @return CheckCommercialStatusResponse
func (client *Client) CheckCommercialStatus(request *CheckCommercialStatusRequest) (_result *CheckCommercialStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCommercialStatusResponse{}
	_body, _err := client.CheckCommercialStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - GetTagKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagKeyResponse
func (client *Client) GetTagKeyWithOptions(request *GetTagKeyRequest, runtime *util.RuntimeOptions) (_result *GetTagKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpanName)) {
		query["SpanName"] = request.SpanName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagKey"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - GetTagKeyRequest
//
// @return GetTagKeyResponse
func (client *Client) GetTagKey(request *GetTagKeyRequest) (_result *GetTagKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagKeyResponse{}
	_body, _err := client.GetTagKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values that correspond to a tag key.
//
// @param request - GetTagValRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagValResponse
func (client *Client) GetTagValWithOptions(request *GetTagValRequest, runtime *util.RuntimeOptions) (_result *GetTagValResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.SpanName)) {
		query["SpanName"] = request.SpanName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagVal"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagValResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values that correspond to a tag key.
//
// @param request - GetTagValRequest
//
// @return GetTagValResponse
func (client *Client) GetTagVal(request *GetTagValRequest) (_result *GetTagValResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagValResponse{}
	_body, _err := client.GetTagValWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a trace.
//
// @param request - GetTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceResponse
func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrace"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a trace.
//
// @param request - GetTraceRequest
//
// @return GetTraceResponse
func (client *Client) GetTrace(request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of an application.
//
// @param request - ListIpOrHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpOrHostsResponse
func (client *Client) ListIpOrHostsWithOptions(request *ListIpOrHostsRequest, runtime *util.RuntimeOptions) (_result *ListIpOrHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpOrHosts"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpOrHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of an application.
//
// @param request - ListIpOrHostsRequest
//
// @return ListIpOrHostsResponse
func (client *Client) ListIpOrHosts(request *ListIpOrHostsRequest) (_result *ListIpOrHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpOrHostsResponse{}
	_body, _err := client.ListIpOrHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries applications.
//
// @param request - ListServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries applications.
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all span names in a specified region or all span names of a microservice.
//
// @param request - ListSpanNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpanNamesResponse
func (client *Client) ListSpanNamesWithOptions(request *ListSpanNamesRequest, runtime *util.RuntimeOptions) (_result *ListSpanNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSpanNames"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSpanNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all span names in a specified region or all span names of a microservice.
//
// @param request - ListSpanNamesRequest
//
// @return ListSpanNamesResponse
func (client *Client) ListSpanNames(request *ListSpanNamesRequest) (_result *ListSpanNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpanNamesResponse{}
	_body, _err := client.ListSpanNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通xtrace和对应的sls
//
// @param request - OpenXtraceServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenXtraceServiceResponse
func (client *Client) OpenXtraceServiceWithOptions(request *OpenXtraceServiceRequest, runtime *util.RuntimeOptions) (_result *OpenXtraceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenXtraceService"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenXtraceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通xtrace和对应的sls
//
// @param request - OpenXtraceServiceRequest
//
// @return OpenXtraceServiceResponse
func (client *Client) OpenXtraceService(request *OpenXtraceServiceRequest) (_result *OpenXtraceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenXtraceServiceResponse{}
	_body, _err := client.OpenXtraceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a metric.
//
// @param request - QueryMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricResponse
func (client *Client) QueryMetricWithOptions(request *QueryMetricRequest, runtime *util.RuntimeOptions) (_result *QueryMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetric"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a metric.
//
// @param request - QueryMetricRequest
//
// @return QueryMetricResponse
func (client *Client) QueryMetric(request *QueryMetricRequest) (_result *QueryMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricResponse{}
	_body, _err := client.QueryMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries traces by time, application name, IP address, span name, and tag.
//
// @param request - SearchTracesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTracesResponse
func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraces"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traces by time, application name, IP address, span name, and tag.
//
// @param request - SearchTracesRequest
//
// @return SearchTracesResponse
func (client *Client) SearchTraces(request *SearchTracesRequest) (_result *SearchTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesResponse{}
	_body, _err := client.SearchTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
