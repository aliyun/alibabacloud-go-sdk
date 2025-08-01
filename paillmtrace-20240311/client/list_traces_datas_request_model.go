// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesDatasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *ListTracesDatasRequest
	GetEndUserId() *string
	SetFilters(v []*ListTracesDatasRequestFilters) *ListTracesDatasRequest
	GetFilters() []*ListTracesDatasRequestFilters
	SetHasEvents(v bool) *ListTracesDatasRequest
	GetHasEvents() *bool
	SetHasStatusMessage(v bool) *ListTracesDatasRequest
	GetHasStatusMessage() *bool
	SetLlmAppName(v string) *ListTracesDatasRequest
	GetLlmAppName() *string
	SetMaxDuration(v float32) *ListTracesDatasRequest
	GetMaxDuration() *float32
	SetMaxTime(v string) *ListTracesDatasRequest
	GetMaxTime() *string
	SetMinDuration(v float32) *ListTracesDatasRequest
	GetMinDuration() *float32
	SetMinTime(v string) *ListTracesDatasRequest
	GetMinTime() *string
	SetOpentelemetryCompatible(v bool) *ListTracesDatasRequest
	GetOpentelemetryCompatible() *bool
	SetOwnerId(v string) *ListTracesDatasRequest
	GetOwnerId() *string
	SetOwnerSubId(v string) *ListTracesDatasRequest
	GetOwnerSubId() *string
	SetPageNumber(v int32) *ListTracesDatasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTracesDatasRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListTracesDatasRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListTracesDatasRequest
	GetSortOrder() *string
	SetSpanIds(v []*string) *ListTracesDatasRequest
	GetSpanIds() []*string
	SetSpanName(v string) *ListTracesDatasRequest
	GetSpanName() *string
	SetTraceIds(v []*string) *ListTracesDatasRequest
	GetTraceIds() []*string
	SetTraceReduceMethod(v string) *ListTracesDatasRequest
	GetTraceReduceMethod() *string
}

type ListTracesDatasRequest struct {
	// The value of the attributes.service.app.user_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// end-user.12345
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Other filter parameters
	Filters []*ListTracesDatasRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Whether to return only trace records containing spans with a non-empty events. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty events. The default value is False. The events is not used for filtering.
	//
	// example:
	//
	// False
	HasEvents *bool `json:"HasEvents,omitempty" xml:"HasEvents,omitempty"`
	// Whether to return only trace records containing spans with a non-empty statusMessage. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty statusMessage. The default value is False. The statusMessage is not used for filtering.
	//
	// example:
	//
	// False
	HasStatusMessage *bool `json:"HasStatusMessage,omitempty" xml:"HasStatusMessage,omitempty"`
	// The value of the resources.service.app.name field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). Must be an exact match. It is empty by default.
	//
	// example:
	//
	// My.super_LLM-app2
	LlmAppName  *string  `json:"LlmAppName,omitempty" xml:"LlmAppName,omitempty"`
	MaxDuration *float32 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The upper limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time +10 minutes)
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MaxTime     *string  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinDuration *float32 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The lower limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time - 2 days).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// Whether the returned JSON data can be directly converted to OpenTelemetry TracesData protobuf object. Default value: False. JSON data that is compatible with OpenTelemetry is more complex. Such data is generally not required unless you want to generate a protobuf object of OpenTelemetry.
	//
	// example:
	//
	// False
	OpentelemetryCompatible *bool   `json:"OpentelemetryCompatible,omitempty" xml:"OpentelemetryCompatible,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The value of the resources.service.owner.sub_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// 123456789
	OwnerSubId *string `json:"OwnerSubId,omitempty" xml:"OwnerSubId,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used to sort the returned results. Valid values: StartTime and Duration.
	//
	// example:
	//
	// StartTime
	//
	// Duration
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- **ASC**
	//
	// 	- **DESC*	- (default)
	//
	// example:
	//
	// DESC
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The list of span IDs. Each trace record contains one or more spans.
	SpanIds  []*string `json:"SpanIds,omitempty" xml:"SpanIds,omitempty" type:"Repeated"`
	SpanName *string   `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The list of trace IDs.
	TraceIds []*string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty" type:"Repeated"`
	// The content simplification method for returned trace data to reduce the data volume.
	//
	// REMOVE_EMBEDDING: Removes all embedding array contents.
	//
	// ROOT_ONLY: Returns only the root span for each trace, with the root span content also having the REMOVE_EMBEDDING applied.
	//
	// Blank: Maintains the original data without simplification.
	//
	// example:
	//
	// REMOVE_EMBEDDING
	//
	// ROOT_ONLY
	TraceReduceMethod *string `json:"TraceReduceMethod,omitempty" xml:"TraceReduceMethod,omitempty"`
}

func (s ListTracesDatasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTracesDatasRequest) GoString() string {
	return s.String()
}

func (s *ListTracesDatasRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListTracesDatasRequest) GetFilters() []*ListTracesDatasRequestFilters {
	return s.Filters
}

func (s *ListTracesDatasRequest) GetHasEvents() *bool {
	return s.HasEvents
}

func (s *ListTracesDatasRequest) GetHasStatusMessage() *bool {
	return s.HasStatusMessage
}

func (s *ListTracesDatasRequest) GetLlmAppName() *string {
	return s.LlmAppName
}

func (s *ListTracesDatasRequest) GetMaxDuration() *float32 {
	return s.MaxDuration
}

func (s *ListTracesDatasRequest) GetMaxTime() *string {
	return s.MaxTime
}

func (s *ListTracesDatasRequest) GetMinDuration() *float32 {
	return s.MinDuration
}

func (s *ListTracesDatasRequest) GetMinTime() *string {
	return s.MinTime
}

func (s *ListTracesDatasRequest) GetOpentelemetryCompatible() *bool {
	return s.OpentelemetryCompatible
}

func (s *ListTracesDatasRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTracesDatasRequest) GetOwnerSubId() *string {
	return s.OwnerSubId
}

func (s *ListTracesDatasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTracesDatasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTracesDatasRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTracesDatasRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTracesDatasRequest) GetSpanIds() []*string {
	return s.SpanIds
}

func (s *ListTracesDatasRequest) GetSpanName() *string {
	return s.SpanName
}

func (s *ListTracesDatasRequest) GetTraceIds() []*string {
	return s.TraceIds
}

func (s *ListTracesDatasRequest) GetTraceReduceMethod() *string {
	return s.TraceReduceMethod
}

func (s *ListTracesDatasRequest) SetEndUserId(v string) *ListTracesDatasRequest {
	s.EndUserId = &v
	return s
}

func (s *ListTracesDatasRequest) SetFilters(v []*ListTracesDatasRequestFilters) *ListTracesDatasRequest {
	s.Filters = v
	return s
}

func (s *ListTracesDatasRequest) SetHasEvents(v bool) *ListTracesDatasRequest {
	s.HasEvents = &v
	return s
}

func (s *ListTracesDatasRequest) SetHasStatusMessage(v bool) *ListTracesDatasRequest {
	s.HasStatusMessage = &v
	return s
}

func (s *ListTracesDatasRequest) SetLlmAppName(v string) *ListTracesDatasRequest {
	s.LlmAppName = &v
	return s
}

func (s *ListTracesDatasRequest) SetMaxDuration(v float32) *ListTracesDatasRequest {
	s.MaxDuration = &v
	return s
}

func (s *ListTracesDatasRequest) SetMaxTime(v string) *ListTracesDatasRequest {
	s.MaxTime = &v
	return s
}

func (s *ListTracesDatasRequest) SetMinDuration(v float32) *ListTracesDatasRequest {
	s.MinDuration = &v
	return s
}

func (s *ListTracesDatasRequest) SetMinTime(v string) *ListTracesDatasRequest {
	s.MinTime = &v
	return s
}

func (s *ListTracesDatasRequest) SetOpentelemetryCompatible(v bool) *ListTracesDatasRequest {
	s.OpentelemetryCompatible = &v
	return s
}

func (s *ListTracesDatasRequest) SetOwnerId(v string) *ListTracesDatasRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTracesDatasRequest) SetOwnerSubId(v string) *ListTracesDatasRequest {
	s.OwnerSubId = &v
	return s
}

func (s *ListTracesDatasRequest) SetPageNumber(v int32) *ListTracesDatasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesDatasRequest) SetPageSize(v int32) *ListTracesDatasRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesDatasRequest) SetSortBy(v string) *ListTracesDatasRequest {
	s.SortBy = &v
	return s
}

func (s *ListTracesDatasRequest) SetSortOrder(v string) *ListTracesDatasRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTracesDatasRequest) SetSpanIds(v []*string) *ListTracesDatasRequest {
	s.SpanIds = v
	return s
}

func (s *ListTracesDatasRequest) SetSpanName(v string) *ListTracesDatasRequest {
	s.SpanName = &v
	return s
}

func (s *ListTracesDatasRequest) SetTraceIds(v []*string) *ListTracesDatasRequest {
	s.TraceIds = v
	return s
}

func (s *ListTracesDatasRequest) SetTraceReduceMethod(v string) *ListTracesDatasRequest {
	s.TraceReduceMethod = &v
	return s
}

func (s *ListTracesDatasRequest) Validate() error {
	return dara.Validate(s)
}

type ListTracesDatasRequestFilters struct {
	// The name of the filter parameter, case-insensitive. Supported parameters: \\"serviceid\\", \\"servicename\\", \\"input\\", \\"output\\", \\"status\\", \\"tracetype\\", and \\"tracename\\".
	//
	// The otel span attributes corresponding to the parameters:
	//
	// serviceid: resources.service.id
	//
	// servicename: resources.service.name
	//
	// input: attributes.input.value
	//
	// output: attributes.output.value
	//
	// status: statusCode
	//
	// tracetype: the attributes.gen_ai.span.kind of span whose parentSpanId is 0
	//
	// tracename: the spanName of span whose parentSpanId is 0
	//
	// Valid values:
	//
	// 	- Status
	//
	// 	- SpanName
	//
	// 	- Input
	//
	// 	- TraceType
	//
	// 	- SpanType
	//
	// 	- ServiceName
	//
	// 	- Output
	//
	// 	- TraceName
	//
	// 	- ServiceId
	//
	// example:
	//
	// output
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter operator. Case-insensitive. Supported operators: \\"=\\", \\"contains\\", and \\"startswith\\".
	//
	// Valid values:
	//
	// 	- contains
	//
	// 	- \\=
	//
	// 	- startsWith
	//
	// example:
	//
	// contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter parameter. For the contains operation, it is case-sensitive. For other operations, it is case-insensitive.
	//
	// example:
	//
	// cretain filter string
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTracesDatasRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListTracesDatasRequestFilters) GoString() string {
	return s.String()
}

func (s *ListTracesDatasRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListTracesDatasRequestFilters) GetOperator() *string {
	return s.Operator
}

func (s *ListTracesDatasRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListTracesDatasRequestFilters) SetKey(v string) *ListTracesDatasRequestFilters {
	s.Key = &v
	return s
}

func (s *ListTracesDatasRequestFilters) SetOperator(v string) *ListTracesDatasRequestFilters {
	s.Operator = &v
	return s
}

func (s *ListTracesDatasRequestFilters) SetValue(v string) *ListTracesDatasRequestFilters {
	s.Value = &v
	return s
}

func (s *ListTracesDatasRequestFilters) Validate() error {
	return dara.Validate(s)
}
