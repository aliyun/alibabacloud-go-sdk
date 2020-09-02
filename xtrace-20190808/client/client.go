// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateSamplingRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	Sampling    *string `json:"Sampling,omitempty" xml:"Sampling,omitempty" require:"true"`
}

func (s UpdateSamplingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSamplingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSamplingRequest) SetRegionId(v string) *UpdateSamplingRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSamplingRequest) SetProxyUserId(v string) *UpdateSamplingRequest {
	s.ProxyUserId = &v
	return s
}

func (s *UpdateSamplingRequest) SetSampling(v string) *UpdateSamplingRequest {
	s.Sampling = &v
	return s
}

type UpdateSamplingResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s UpdateSamplingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSamplingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSamplingResponse) SetRequestId(v string) *UpdateSamplingResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSamplingResponse) SetData(v string) *UpdateSamplingResponse {
	s.Data = &v
	return s
}

type GetSamplingRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s GetSamplingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSamplingRequest) GoString() string {
	return s.String()
}

func (s *GetSamplingRequest) SetRegionId(v string) *GetSamplingRequest {
	s.RegionId = &v
	return s
}

func (s *GetSamplingRequest) SetProxyUserId(v string) *GetSamplingRequest {
	s.ProxyUserId = &v
	return s
}

type GetSamplingResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s GetSamplingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSamplingResponse) GoString() string {
	return s.String()
}

func (s *GetSamplingResponse) SetRequestId(v string) *GetSamplingResponse {
	s.RequestId = &v
	return s
}

func (s *GetSamplingResponse) SetData(v string) *GetSamplingResponse {
	s.Data = &v
	return s
}

type QueryMetricRequest struct {
	IntervalInSec *int                         `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	StartTime     *int64                       `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	OrderBy       *string                      `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Limit         *int                         `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Filters       []*QueryMetricRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	Dimensions    []*string                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Metric        *string                      `json:"Metric,omitempty" xml:"Metric,omitempty" require:"true"`
	Measures      []*string                    `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Order         *string                      `json:"Order,omitempty" xml:"Order,omitempty"`
	ProxyUserId   *string                      `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s QueryMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricRequest) SetIntervalInSec(v int) *QueryMetricRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricRequest) SetStartTime(v int64) *QueryMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMetricRequest) SetEndTime(v int64) *QueryMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricRequest) SetOrderBy(v string) *QueryMetricRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricRequest) SetLimit(v int) *QueryMetricRequest {
	s.Limit = &v
	return s
}

func (s *QueryMetricRequest) SetFilters(v []*QueryMetricRequestFilters) *QueryMetricRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricRequest) SetDimensions(v []*string) *QueryMetricRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricRequest) SetMetric(v string) *QueryMetricRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricRequest) SetMeasures(v []*string) *QueryMetricRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricRequest) SetOrder(v string) *QueryMetricRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricRequest) SetProxyUserId(v string) *QueryMetricRequest {
	s.ProxyUserId = &v
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

type QueryMetricResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s QueryMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricResponse) SetRequestId(v string) *QueryMetricResponse {
	s.RequestId = &v
	return s
}

func (s *QueryMetricResponse) SetData(v string) *QueryMetricResponse {
	s.Data = &v
	return s
}

type GetTraceRequest struct {
	TraceID  *string `json:"TraceID,omitempty" xml:"TraceID,omitempty" require:"true"`
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

func (s *GetTraceRequest) SetAppType(v string) *GetTraceRequest {
	s.AppType = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

type GetTraceResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Spans     *GetTraceResponseSpans `json:"Spans,omitempty" xml:"Spans,omitempty" require:"true" type:"Struct"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetRequestId(v string) *GetTraceResponse {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponse) SetSpans(v *GetTraceResponseSpans) *GetTraceResponse {
	s.Spans = v
	return s
}

type GetTraceResponseSpans struct {
	Span []*GetTraceResponseSpansSpan `json:"Span,omitempty" xml:"Span,omitempty" require:"true" type:"Repeated"`
}

func (s GetTraceResponseSpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpans) SetSpan(v []*GetTraceResponseSpansSpan) *GetTraceResponseSpans {
	s.Span = v
	return s
}

type GetTraceResponseSpansSpan struct {
	TraceID       *string                                `json:"TraceID,omitempty" xml:"TraceID,omitempty" require:"true"`
	OperationName *string                                `json:"OperationName,omitempty" xml:"OperationName,omitempty" require:"true"`
	Duration      *int64                                 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	ServiceName   *string                                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
	ServiceIp     *string                                `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty" require:"true"`
	Timestamp     *int64                                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty" require:"true"`
	RpcId         *string                                `json:"RpcId,omitempty" xml:"RpcId,omitempty" require:"true"`
	ResultCode    *string                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty" require:"true"`
	HaveStack     *bool                                  `json:"HaveStack,omitempty" xml:"HaveStack,omitempty" require:"true"`
	TagEntryList  *GetTraceResponseSpansSpanTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" require:"true" type:"Struct"`
	LogEventList  *GetTraceResponseSpansSpanLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" require:"true" type:"Struct"`
}

func (s GetTraceResponseSpansSpan) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpan) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpan) SetTraceID(v string) *GetTraceResponseSpansSpan {
	s.TraceID = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetOperationName(v string) *GetTraceResponseSpansSpan {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetDuration(v int64) *GetTraceResponseSpansSpan {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetServiceName(v string) *GetTraceResponseSpansSpan {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetServiceIp(v string) *GetTraceResponseSpansSpan {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetTimestamp(v int64) *GetTraceResponseSpansSpan {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetRpcId(v string) *GetTraceResponseSpansSpan {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetResultCode(v string) *GetTraceResponseSpansSpan {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetHaveStack(v bool) *GetTraceResponseSpansSpan {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseSpansSpan) SetTagEntryList(v *GetTraceResponseSpansSpanTagEntryList) *GetTraceResponseSpansSpan {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseSpansSpan) SetLogEventList(v *GetTraceResponseSpansSpanLogEventList) *GetTraceResponseSpansSpan {
	s.LogEventList = v
	return s
}

type GetTraceResponseSpansSpanTagEntryList struct {
	TagEntry []*GetTraceResponseSpansSpanTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" require:"true" type:"Repeated"`
}

func (s GetTraceResponseSpansSpanTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanTagEntryList) SetTagEntry(v []*GetTraceResponseSpansSpanTagEntryListTagEntry) *GetTraceResponseSpansSpanTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseSpansSpanTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetTraceResponseSpansSpanTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanTagEntryListTagEntry) SetKey(v string) *GetTraceResponseSpansSpanTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseSpansSpanTagEntryListTagEntry) SetValue(v string) *GetTraceResponseSpansSpanTagEntryListTagEntry {
	s.Value = &v
	return s
}

type GetTraceResponseSpansSpanLogEventList struct {
	LogEvent []*GetTraceResponseSpansSpanLogEventListLogEvent `json:"LogEvent,omitempty" xml:"LogEvent,omitempty" require:"true" type:"Repeated"`
}

func (s GetTraceResponseSpansSpanLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanLogEventList) SetLogEvent(v []*GetTraceResponseSpansSpanLogEventListLogEvent) *GetTraceResponseSpansSpanLogEventList {
	s.LogEvent = v
	return s
}

type GetTraceResponseSpansSpanLogEventListLogEvent struct {
	Timestamp    *int64                                                     `json:"Timestamp,omitempty" xml:"Timestamp,omitempty" require:"true"`
	TagEntryList *GetTraceResponseSpansSpanLogEventListLogEventTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" require:"true" type:"Struct"`
}

func (s GetTraceResponseSpansSpanLogEventListLogEvent) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanLogEventListLogEvent) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanLogEventListLogEvent) SetTimestamp(v int64) *GetTraceResponseSpansSpanLogEventListLogEvent {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseSpansSpanLogEventListLogEvent) SetTagEntryList(v *GetTraceResponseSpansSpanLogEventListLogEventTagEntryList) *GetTraceResponseSpansSpanLogEventListLogEvent {
	s.TagEntryList = v
	return s
}

type GetTraceResponseSpansSpanLogEventListLogEventTagEntryList struct {
	TagEntry []*GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry `json:"TagEntry,omitempty" xml:"TagEntry,omitempty" require:"true" type:"Repeated"`
}

func (s GetTraceResponseSpansSpanLogEventListLogEventTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanLogEventListLogEventTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanLogEventListLogEventTagEntryList) SetTagEntry(v []*GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry) *GetTraceResponseSpansSpanLogEventListLogEventTagEntryList {
	s.TagEntry = v
	return s
}

type GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry) GoString() string {
	return s.String()
}

func (s *GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry) SetKey(v string) *GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Key = &v
	return s
}

func (s *GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry) SetValue(v string) *GetTraceResponseSpansSpanLogEventListLogEventTagEntryListTagEntry {
	s.Value = &v
	return s
}

type SearchTracesRequest struct {
	StartTime     *int64                    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *int64                    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	RegionId      *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ServiceName   *string                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	OperationName *string                   `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	MinDuration   *int64                    `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	AppType       *string                   `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Tag           []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	PageNumber    *int                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse       *bool                     `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp     *string                   `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetAppType(v string) *SearchTracesRequest {
	s.AppType = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

func (s *SearchTracesRequest) SetPageNumber(v int) *SearchTracesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesRequest) SetPageSize(v int) *SearchTracesRequest {
	s.PageSize = &v
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

type SearchTracesResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageBean  *SearchTracesResponsePageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" require:"true" type:"Struct"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetRequestId(v string) *SearchTracesResponse {
	s.RequestId = &v
	return s
}

func (s *SearchTracesResponse) SetPageBean(v *SearchTracesResponsePageBean) *SearchTracesResponse {
	s.PageBean = v
	return s
}

type SearchTracesResponsePageBean struct {
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TraceInfos *SearchTracesResponsePageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" require:"true" type:"Struct"`
}

func (s SearchTracesResponsePageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponsePageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesResponsePageBean) SetTotalCount(v int64) *SearchTracesResponsePageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTracesResponsePageBean) SetPageSize(v int) *SearchTracesResponsePageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesResponsePageBean) SetPageNumber(v int) *SearchTracesResponsePageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesResponsePageBean) SetTraceInfos(v *SearchTracesResponsePageBeanTraceInfos) *SearchTracesResponsePageBean {
	s.TraceInfos = v
	return s
}

type SearchTracesResponsePageBeanTraceInfos struct {
	TraceInfo []*SearchTracesResponsePageBeanTraceInfosTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" require:"true" type:"Repeated"`
}

func (s SearchTracesResponsePageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponsePageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponsePageBeanTraceInfos) SetTraceInfo(v []*SearchTracesResponsePageBeanTraceInfosTraceInfo) *SearchTracesResponsePageBeanTraceInfos {
	s.TraceInfo = v
	return s
}

type SearchTracesResponsePageBeanTraceInfosTraceInfo struct {
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty" require:"true"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty" require:"true"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty" require:"true"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty" require:"true"`
}

func (s SearchTracesResponsePageBeanTraceInfosTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponsePageBeanTraceInfosTraceInfo) GoString() string {
	return s.String()
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetTraceID(v string) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.TraceID = &v
	return s
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetOperationName(v string) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetServiceName(v string) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetServiceIp(v string) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetDuration(v int64) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponsePageBeanTraceInfosTraceInfo) SetTimestamp(v int64) *SearchTracesResponsePageBeanTraceInfosTraceInfo {
	s.Timestamp = &v
	return s
}

type GetTagValRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	TagKey      *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetTagValRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagValRequest) GoString() string {
	return s.String()
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

func (s *GetTagValRequest) SetTagKey(v string) *GetTagValRequest {
	s.TagKey = &v
	return s
}

func (s *GetTagValRequest) SetStartTime(v int64) *GetTagValRequest {
	s.StartTime = &v
	return s
}

func (s *GetTagValRequest) SetEndTime(v int64) *GetTagValRequest {
	s.EndTime = &v
	return s
}

type GetTagValResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TagValues *GetTagValResponseTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" require:"true" type:"Struct"`
}

func (s GetTagValResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponse) GoString() string {
	return s.String()
}

func (s *GetTagValResponse) SetRequestId(v string) *GetTagValResponse {
	s.RequestId = &v
	return s
}

func (s *GetTagValResponse) SetTagValues(v *GetTagValResponseTagValues) *GetTagValResponse {
	s.TagValues = v
	return s
}

type GetTagValResponseTagValues struct {
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true" type:"Repeated"`
}

func (s GetTagValResponseTagValues) String() string {
	return tea.Prettify(s)
}

func (s GetTagValResponseTagValues) GoString() string {
	return s.String()
}

func (s *GetTagValResponseTagValues) SetTagValue(v []*string) *GetTagValResponseTagValues {
	s.TagValue = v
	return s
}

type GetTagKeyRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanName    *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetTagKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyRequest) GoString() string {
	return s.String()
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

func (s *GetTagKeyRequest) SetEndTime(v int64) *GetTagKeyRequest {
	s.EndTime = &v
	return s
}

type GetTagKeyResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TagKeys   *GetTagKeyResponseTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" require:"true" type:"Struct"`
}

func (s GetTagKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponse) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponse) SetRequestId(v string) *GetTagKeyResponse {
	s.RequestId = &v
	return s
}

func (s *GetTagKeyResponse) SetTagKeys(v *GetTagKeyResponseTagKeys) *GetTagKeyResponse {
	s.TagKeys = v
	return s
}

type GetTagKeyResponseTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true" type:"Repeated"`
}

func (s GetTagKeyResponseTagKeys) String() string {
	return tea.Prettify(s)
}

func (s GetTagKeyResponseTagKeys) GoString() string {
	return s.String()
}

func (s *GetTagKeyResponseTagKeys) SetTagKey(v []*string) *GetTagKeyResponseTagKeys {
	s.TagKey = v
	return s
}

type ListIpOrHostsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListIpOrHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsRequest) GoString() string {
	return s.String()
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

func (s *ListIpOrHostsRequest) SetEndTime(v int64) *ListIpOrHostsRequest {
	s.EndTime = &v
	return s
}

type ListIpOrHostsResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IpNames   *ListIpOrHostsResponseIpNames `json:"IpNames,omitempty" xml:"IpNames,omitempty" require:"true" type:"Struct"`
}

func (s ListIpOrHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponse) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponse) SetRequestId(v string) *ListIpOrHostsResponse {
	s.RequestId = &v
	return s
}

func (s *ListIpOrHostsResponse) SetIpNames(v *ListIpOrHostsResponseIpNames) *ListIpOrHostsResponse {
	s.IpNames = v
	return s
}

type ListIpOrHostsResponseIpNames struct {
	IpName []*string `json:"IpName,omitempty" xml:"IpName,omitempty" require:"true" type:"Repeated"`
}

func (s ListIpOrHostsResponseIpNames) String() string {
	return tea.Prettify(s)
}

func (s ListIpOrHostsResponseIpNames) GoString() string {
	return s.String()
}

func (s *ListIpOrHostsResponseIpNames) SetIpName(v []*string) *ListIpOrHostsResponseIpNames {
	s.IpName = v
	return s
}

type ListServicesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AppType  *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetAppType(v string) *ListServicesRequest {
	s.AppType = &v
	return s
}

type ListServicesResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Services  *ListServicesResponseServices `json:"Services,omitempty" xml:"Services,omitempty" require:"true" type:"Struct"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetRequestId(v string) *ListServicesResponse {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponse) SetServices(v *ListServicesResponseServices) *ListServicesResponse {
	s.Services = v
	return s
}

type ListServicesResponseServices struct {
	Service []*ListServicesResponseServicesService `json:"Service,omitempty" xml:"Service,omitempty" require:"true" type:"Repeated"`
}

func (s ListServicesResponseServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseServices) SetService(v []*ListServicesResponseServicesService) *ListServicesResponseServices {
	s.Service = v
	return s
}

type ListServicesResponseServicesService struct {
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty" require:"true"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s ListServicesResponseServicesService) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseServicesService) GoString() string {
	return s.String()
}

func (s *ListServicesResponseServicesService) SetServiceName(v string) *ListServicesResponseServicesService {
	s.ServiceName = &v
	return s
}

func (s *ListServicesResponseServicesService) SetPid(v string) *ListServicesResponseServicesService {
	s.Pid = &v
	return s
}

func (s *ListServicesResponseServicesService) SetRegionId(v string) *ListServicesResponseServicesService {
	s.RegionId = &v
	return s
}

type ListSpanNamesRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListSpanNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesRequest) GoString() string {
	return s.String()
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

func (s *ListSpanNamesRequest) SetEndTime(v int64) *ListSpanNamesRequest {
	s.EndTime = &v
	return s
}

type ListSpanNamesResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SpanNames *ListSpanNamesResponseSpanNames `json:"SpanNames,omitempty" xml:"SpanNames,omitempty" require:"true" type:"Struct"`
}

func (s ListSpanNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponse) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponse) SetRequestId(v string) *ListSpanNamesResponse {
	s.RequestId = &v
	return s
}

func (s *ListSpanNamesResponse) SetSpanNames(v *ListSpanNamesResponseSpanNames) *ListSpanNamesResponse {
	s.SpanNames = v
	return s
}

type ListSpanNamesResponseSpanNames struct {
	SpanName []*string `json:"SpanName,omitempty" xml:"SpanName,omitempty" require:"true" type:"Repeated"`
}

func (s ListSpanNamesResponseSpanNames) String() string {
	return tea.Prettify(s)
}

func (s ListSpanNamesResponseSpanNames) GoString() string {
	return s.String()
}

func (s *ListSpanNamesResponseSpanNames) SetSpanName(v []*string) *ListSpanNamesResponseSpanNames {
	s.SpanName = v
	return s
}

type GetTokenRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetRegionId(v string) *GetTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetTokenRequest) SetAppType(v string) *GetTokenRequest {
	s.AppType = &v
	return s
}

func (s *GetTokenRequest) SetProxyUserId(v string) *GetTokenRequest {
	s.ProxyUserId = &v
	return s
}

type GetTokenResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Token     *GetTokenResponseToken `json:"Token,omitempty" xml:"Token,omitempty" require:"true" type:"Struct"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetRequestId(v string) *GetTokenResponse {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponse) SetToken(v *GetTokenResponseToken) *GetTokenResponse {
	s.Token = v
	return s
}

type GetTokenResponseToken struct {
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	LicenseKey     *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty" require:"true"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty" require:"true"`
	InternalDomain *string `json:"InternalDomain,omitempty" xml:"InternalDomain,omitempty" require:"true"`
}

func (s GetTokenResponseToken) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseToken) GoString() string {
	return s.String()
}

func (s *GetTokenResponseToken) SetDomain(v string) *GetTokenResponseToken {
	s.Domain = &v
	return s
}

func (s *GetTokenResponseToken) SetLicenseKey(v string) *GetTokenResponseToken {
	s.LicenseKey = &v
	return s
}

func (s *GetTokenResponseToken) SetPid(v string) *GetTokenResponseToken {
	s.Pid = &v
	return s
}

func (s *GetTokenResponseToken) SetInternalDomain(v string) *GetTokenResponseToken {
	s.InternalDomain = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) UpdateSamplingWithOptions(request *UpdateSamplingRequest, runtime *util.RuntimeOptions) (_result *UpdateSamplingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateSamplingResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateSampling"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSampling(request *UpdateSamplingRequest) (_result *UpdateSamplingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSamplingResponse{}
	_body, _err := client.UpdateSamplingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSamplingWithOptions(request *GetSamplingRequest, runtime *util.RuntimeOptions) (_result *GetSamplingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSamplingResponse{}
	_body, _err := client.DoRequest(tea.String("GetSampling"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSampling(request *GetSamplingRequest) (_result *GetSamplingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSamplingResponse{}
	_body, _err := client.GetSamplingWithOptions(request, runtime)
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
	_result = &QueryMetricResponse{}
	_body, _err := client.DoRequest(tea.String("QueryMetric"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTraceResponse{}
	_body, _err := client.DoRequest(tea.String("GetTrace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.DoRequest(tea.String("SearchTraces"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetTagValWithOptions(request *GetTagValRequest, runtime *util.RuntimeOptions) (_result *GetTagValResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTagValResponse{}
	_body, _err := client.DoRequest(tea.String("GetTagVal"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetTagKeyWithOptions(request *GetTagKeyRequest, runtime *util.RuntimeOptions) (_result *GetTagKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTagKeyResponse{}
	_body, _err := client.DoRequest(tea.String("GetTagKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListIpOrHostsWithOptions(request *ListIpOrHostsRequest, runtime *util.RuntimeOptions) (_result *ListIpOrHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListIpOrHostsResponse{}
	_body, _err := client.DoRequest(tea.String("ListIpOrHosts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListServicesResponse{}
	_body, _err := client.DoRequest(tea.String("ListServices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListSpanNamesResponse{}
	_body, _err := client.DoRequest(tea.String("ListSpanNames"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTokenResponse{}
	_body, _err := client.DoRequest(tea.String("GetToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
