// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListBaselinesResponseBodyData) *ListBaselinesResponseBody
	GetData() *ListBaselinesResponseBodyData
	SetErrorCode(v string) *ListBaselinesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListBaselinesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListBaselinesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListBaselinesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaselinesResponseBody
	GetSuccess() *bool
}

type ListBaselinesResponseBody struct {
	// The data returned.
	Data *ListBaselinesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 103630001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The user does not exist. 276571706358178756
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 952795279527ab****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaselinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaselinesResponseBody) GetData() *ListBaselinesResponseBodyData {
	return s.Data
}

func (s *ListBaselinesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListBaselinesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBaselinesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaselinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaselinesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaselinesResponseBody) SetData(v *ListBaselinesResponseBodyData) *ListBaselinesResponseBody {
	s.Data = v
	return s
}

func (s *ListBaselinesResponseBody) SetErrorCode(v string) *ListBaselinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListBaselinesResponseBody) SetErrorMessage(v string) *ListBaselinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListBaselinesResponseBody) SetHttpStatusCode(v int32) *ListBaselinesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaselinesResponseBody) SetRequestId(v string) *ListBaselinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaselinesResponseBody) SetSuccess(v bool) *ListBaselinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaselinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBaselinesResponseBodyData struct {
	// The baselines.
	Baselines []*ListBaselinesResponseBodyDataBaselines `json:"Baselines,omitempty" xml:"Baselines,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of baselines returned.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBaselinesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaselinesResponseBodyData) GetBaselines() []*ListBaselinesResponseBodyDataBaselines {
	return s.Baselines
}

func (s *ListBaselinesResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListBaselinesResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *ListBaselinesResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListBaselinesResponseBodyData) SetBaselines(v []*ListBaselinesResponseBodyDataBaselines) *ListBaselinesResponseBodyData {
	s.Baselines = v
	return s
}

func (s *ListBaselinesResponseBodyData) SetPageNumber(v string) *ListBaselinesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBaselinesResponseBodyData) SetPageSize(v string) *ListBaselinesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBaselinesResponseBodyData) SetTotalCount(v string) *ListBaselinesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBaselinesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBaselinesResponseBodyDataBaselines struct {
	// Indicates whether the alerting feature is enabled. Valid values: true and false.
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The alert margin threshold for the baseline instance. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values: DAILY and HOURLY.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Indicates whether the baseline is enabled. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The settings of the committed completion time of the baseline.
	OverTimeSettings []*ListBaselinesResponseBodyDataBaselinesOverTimeSettings `json:"OverTimeSettings,omitempty" xml:"OverTimeSettings,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs can be specified. The IDs are separated by commas (,).
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,2,5,7,8}.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workspace to which the baseline belongs.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListBaselinesResponseBodyDataBaselines) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesResponseBodyDataBaselines) GoString() string {
	return s.String()
}

func (s *ListBaselinesResponseBodyDataBaselines) GetAlertEnabled() *bool {
	return s.AlertEnabled
}

func (s *ListBaselinesResponseBodyDataBaselines) GetAlertMarginThreshold() *int32 {
	return s.AlertMarginThreshold
}

func (s *ListBaselinesResponseBodyDataBaselines) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListBaselinesResponseBodyDataBaselines) GetBaselineName() *string {
	return s.BaselineName
}

func (s *ListBaselinesResponseBodyDataBaselines) GetBaselineType() *string {
	return s.BaselineType
}

func (s *ListBaselinesResponseBodyDataBaselines) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListBaselinesResponseBodyDataBaselines) GetOverTimeSettings() []*ListBaselinesResponseBodyDataBaselinesOverTimeSettings {
	return s.OverTimeSettings
}

func (s *ListBaselinesResponseBodyDataBaselines) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselinesResponseBodyDataBaselines) GetPriority() *int32 {
	return s.Priority
}

func (s *ListBaselinesResponseBodyDataBaselines) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBaselinesResponseBodyDataBaselines) SetAlertEnabled(v bool) *ListBaselinesResponseBodyDataBaselines {
	s.AlertEnabled = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetAlertMarginThreshold(v int32) *ListBaselinesResponseBodyDataBaselines {
	s.AlertMarginThreshold = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetBaselineId(v int64) *ListBaselinesResponseBodyDataBaselines {
	s.BaselineId = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetBaselineName(v string) *ListBaselinesResponseBodyDataBaselines {
	s.BaselineName = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetBaselineType(v string) *ListBaselinesResponseBodyDataBaselines {
	s.BaselineType = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetEnabled(v bool) *ListBaselinesResponseBodyDataBaselines {
	s.Enabled = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetOverTimeSettings(v []*ListBaselinesResponseBodyDataBaselinesOverTimeSettings) *ListBaselinesResponseBodyDataBaselines {
	s.OverTimeSettings = v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetOwner(v string) *ListBaselinesResponseBodyDataBaselines {
	s.Owner = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetPriority(v int32) *ListBaselinesResponseBodyDataBaselines {
	s.Priority = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) SetProjectId(v int64) *ListBaselinesResponseBodyDataBaselines {
	s.ProjectId = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselines) Validate() error {
	return dara.Validate(s)
}

type ListBaselinesResponseBodyDataBaselinesOverTimeSettings struct {
	// The cycle that corresponds to the committed completion time. For a day-level baseline, the value of this parameter is 1. For an hour-level baseline, the value of this parameter cannot exceed 24.
	//
	// example:
	//
	// 1
	Cycle *int32 `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The committed completion time in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListBaselinesResponseBodyDataBaselinesOverTimeSettings) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesResponseBodyDataBaselinesOverTimeSettings) GoString() string {
	return s.String()
}

func (s *ListBaselinesResponseBodyDataBaselinesOverTimeSettings) GetCycle() *int32 {
	return s.Cycle
}

func (s *ListBaselinesResponseBodyDataBaselinesOverTimeSettings) GetTime() *string {
	return s.Time
}

func (s *ListBaselinesResponseBodyDataBaselinesOverTimeSettings) SetCycle(v int32) *ListBaselinesResponseBodyDataBaselinesOverTimeSettings {
	s.Cycle = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselinesOverTimeSettings) SetTime(v string) *ListBaselinesResponseBodyDataBaselinesOverTimeSettings {
	s.Time = &v
	return s
}

func (s *ListBaselinesResponseBodyDataBaselinesOverTimeSettings) Validate() error {
	return dara.Validate(s)
}
