// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTagKeyRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   *GetTagKeyResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTagKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagKey      *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTagValResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
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
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spans     *GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Struct"`
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
	Duration      *int64                                     `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                      `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  *GetTraceResponseBodySpansSpanLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Struct"`
	OperationName *string                                    `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                    `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                    `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceIp     *string                                    `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                    `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  *GetTraceResponseBodySpansSpanTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	Timestamp     *int64                                     `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                    `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
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
	TagEntryList *GetTraceResponseBodySpansSpanLogEventListLogEventTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Struct"`
	Timestamp    *int64                                                         `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	IpNames   *ListIpOrHostsResponseBodyIpNames `json:"IpNames,omitempty" xml:"IpNames,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIpOrHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  *ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Struct"`
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
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSpanNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type QueryMetricRequest struct {
	Dimensions    []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EndTime       *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters       []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec *int32                       `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	Limit         *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Measures      []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Metric        *string                      `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order         *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy       *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	ProxyUserId   *string                      `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	StartTime     *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppType       *string                   `json:"AppType,omitempty" xml:"AppType,omitempty"`
	EndTime       *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MinDuration   *int64                    `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName *string                   `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	PageNumber    *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse       *bool                     `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp     *string                   `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime     *int64                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag           []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	PageBean  *SearchTracesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Duration      *int64                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string                `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TagMap        map[string]interface{} `json:"TagMap,omitempty" xml:"TagMap,omitempty"`
	Timestamp     *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
