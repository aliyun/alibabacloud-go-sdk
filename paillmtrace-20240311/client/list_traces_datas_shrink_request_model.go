// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesDatasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *ListTracesDatasShrinkRequest
	GetEndUserId() *string
	SetFiltersShrink(v string) *ListTracesDatasShrinkRequest
	GetFiltersShrink() *string
	SetHasEvents(v bool) *ListTracesDatasShrinkRequest
	GetHasEvents() *bool
	SetHasStatusMessage(v bool) *ListTracesDatasShrinkRequest
	GetHasStatusMessage() *bool
	SetLlmAppName(v string) *ListTracesDatasShrinkRequest
	GetLlmAppName() *string
	SetMaxDuration(v float32) *ListTracesDatasShrinkRequest
	GetMaxDuration() *float32
	SetMaxTime(v string) *ListTracesDatasShrinkRequest
	GetMaxTime() *string
	SetMinDuration(v float32) *ListTracesDatasShrinkRequest
	GetMinDuration() *float32
	SetMinTime(v string) *ListTracesDatasShrinkRequest
	GetMinTime() *string
	SetOpentelemetryCompatible(v bool) *ListTracesDatasShrinkRequest
	GetOpentelemetryCompatible() *bool
	SetOwnerId(v string) *ListTracesDatasShrinkRequest
	GetOwnerId() *string
	SetOwnerSubId(v string) *ListTracesDatasShrinkRequest
	GetOwnerSubId() *string
	SetPageNumber(v int32) *ListTracesDatasShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTracesDatasShrinkRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListTracesDatasShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListTracesDatasShrinkRequest
	GetSortOrder() *string
	SetSpanIdsShrink(v string) *ListTracesDatasShrinkRequest
	GetSpanIdsShrink() *string
	SetSpanName(v string) *ListTracesDatasShrinkRequest
	GetSpanName() *string
	SetTraceIdsShrink(v string) *ListTracesDatasShrinkRequest
	GetTraceIdsShrink() *string
	SetTraceReduceMethod(v string) *ListTracesDatasShrinkRequest
	GetTraceReduceMethod() *string
}

type ListTracesDatasShrinkRequest struct {
	// The value of the attributes.service.app.user_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// end-user.12345
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Other filter parameters
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
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
	SpanIdsShrink *string `json:"SpanIds,omitempty" xml:"SpanIds,omitempty"`
	SpanName      *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The list of trace IDs.
	TraceIdsShrink *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
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

func (s ListTracesDatasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTracesDatasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTracesDatasShrinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListTracesDatasShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *ListTracesDatasShrinkRequest) GetHasEvents() *bool {
	return s.HasEvents
}

func (s *ListTracesDatasShrinkRequest) GetHasStatusMessage() *bool {
	return s.HasStatusMessage
}

func (s *ListTracesDatasShrinkRequest) GetLlmAppName() *string {
	return s.LlmAppName
}

func (s *ListTracesDatasShrinkRequest) GetMaxDuration() *float32 {
	return s.MaxDuration
}

func (s *ListTracesDatasShrinkRequest) GetMaxTime() *string {
	return s.MaxTime
}

func (s *ListTracesDatasShrinkRequest) GetMinDuration() *float32 {
	return s.MinDuration
}

func (s *ListTracesDatasShrinkRequest) GetMinTime() *string {
	return s.MinTime
}

func (s *ListTracesDatasShrinkRequest) GetOpentelemetryCompatible() *bool {
	return s.OpentelemetryCompatible
}

func (s *ListTracesDatasShrinkRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTracesDatasShrinkRequest) GetOwnerSubId() *string {
	return s.OwnerSubId
}

func (s *ListTracesDatasShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTracesDatasShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTracesDatasShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTracesDatasShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTracesDatasShrinkRequest) GetSpanIdsShrink() *string {
	return s.SpanIdsShrink
}

func (s *ListTracesDatasShrinkRequest) GetSpanName() *string {
	return s.SpanName
}

func (s *ListTracesDatasShrinkRequest) GetTraceIdsShrink() *string {
	return s.TraceIdsShrink
}

func (s *ListTracesDatasShrinkRequest) GetTraceReduceMethod() *string {
	return s.TraceReduceMethod
}

func (s *ListTracesDatasShrinkRequest) SetEndUserId(v string) *ListTracesDatasShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetFiltersShrink(v string) *ListTracesDatasShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetHasEvents(v bool) *ListTracesDatasShrinkRequest {
	s.HasEvents = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetHasStatusMessage(v bool) *ListTracesDatasShrinkRequest {
	s.HasStatusMessage = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetLlmAppName(v string) *ListTracesDatasShrinkRequest {
	s.LlmAppName = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMaxDuration(v float32) *ListTracesDatasShrinkRequest {
	s.MaxDuration = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMaxTime(v string) *ListTracesDatasShrinkRequest {
	s.MaxTime = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMinDuration(v float32) *ListTracesDatasShrinkRequest {
	s.MinDuration = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMinTime(v string) *ListTracesDatasShrinkRequest {
	s.MinTime = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOpentelemetryCompatible(v bool) *ListTracesDatasShrinkRequest {
	s.OpentelemetryCompatible = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOwnerId(v string) *ListTracesDatasShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOwnerSubId(v string) *ListTracesDatasShrinkRequest {
	s.OwnerSubId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetPageNumber(v int32) *ListTracesDatasShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetPageSize(v int32) *ListTracesDatasShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSortBy(v string) *ListTracesDatasShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSortOrder(v string) *ListTracesDatasShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSpanIdsShrink(v string) *ListTracesDatasShrinkRequest {
	s.SpanIdsShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSpanName(v string) *ListTracesDatasShrinkRequest {
	s.SpanName = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetTraceIdsShrink(v string) *ListTracesDatasShrinkRequest {
	s.TraceIdsShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetTraceReduceMethod(v string) *ListTracesDatasShrinkRequest {
	s.TraceReduceMethod = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
