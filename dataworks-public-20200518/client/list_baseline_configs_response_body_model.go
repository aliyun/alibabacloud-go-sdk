// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListBaselineConfigsResponseBodyData) *ListBaselineConfigsResponseBody
	GetData() *ListBaselineConfigsResponseBodyData
	SetErrorCode(v string) *ListBaselineConfigsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListBaselineConfigsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListBaselineConfigsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListBaselineConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaselineConfigsResponseBody
	GetSuccess() *bool
}

type ListBaselineConfigsResponseBody struct {
	// The returned data.
	Data *ListBaselineConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
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
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaselineConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaselineConfigsResponseBody) GetData() *ListBaselineConfigsResponseBodyData {
	return s.Data
}

func (s *ListBaselineConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListBaselineConfigsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBaselineConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaselineConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaselineConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaselineConfigsResponseBody) SetData(v *ListBaselineConfigsResponseBodyData) *ListBaselineConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListBaselineConfigsResponseBody) SetErrorCode(v string) *ListBaselineConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListBaselineConfigsResponseBody) SetErrorMessage(v string) *ListBaselineConfigsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListBaselineConfigsResponseBody) SetHttpStatusCode(v int32) *ListBaselineConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaselineConfigsResponseBody) SetRequestId(v string) *ListBaselineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaselineConfigsResponseBody) SetSuccess(v bool) *ListBaselineConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaselineConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBaselineConfigsResponseBodyData struct {
	// The baselines.
	Baselines []*ListBaselineConfigsResponseBodyDataBaselines `json:"Baselines,omitempty" xml:"Baselines,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of baselines returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBaselineConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaselineConfigsResponseBodyData) GetBaselines() []*ListBaselineConfigsResponseBodyDataBaselines {
	return s.Baselines
}

func (s *ListBaselineConfigsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBaselineConfigsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineConfigsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBaselineConfigsResponseBodyData) SetBaselines(v []*ListBaselineConfigsResponseBodyDataBaselines) *ListBaselineConfigsResponseBodyData {
	s.Baselines = v
	return s
}

func (s *ListBaselineConfigsResponseBodyData) SetPageNumber(v int32) *ListBaselineConfigsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyData) SetPageSize(v int32) *ListBaselineConfigsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyData) SetTotalCount(v int32) *ListBaselineConfigsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBaselineConfigsResponseBodyDataBaselines struct {
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
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values: DAILY and HOURLY.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// The hour in the alerting time of the day-level baseline. Valid values: [0, 47].
	//
	// example:
	//
	// 7
	ExpHour *int32 `json:"ExpHour,omitempty" xml:"ExpHour,omitempty"`
	// The minute in the alerting time of the day-level baseline. Valid values: [0, 59].
	//
	// example:
	//
	// 30
	ExpMinu *int32 `json:"ExpMinu,omitempty" xml:"ExpMinu,omitempty"`
	// The alerting time of the hour-level baseline. This parameter is presented as key-value pairs in the JSON format. The key indicates the ID of the cycle, and the value is presented in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// {"1":"03:28","2":"04:28","3":"05:28","4":"06:28","5":"07:28","6":"08:28","7":"09:28","8":"10:28","9":"11:28","10":"12:28","11":"13:28","12":"14:28","13":"15:28","14":"16:28","15":"17:28","16":"18:28","17":"19:28","18":"20:28","19":"21:28","20":"22:28","21":"23:28","22":"24:28","23":"25:28","24":"26:28"}
	HourExpDetail *string `json:"HourExpDetail,omitempty" xml:"HourExpDetail,omitempty"`
	// The committed completion time of the hour-level baseline. This parameter is presented as key-value pairs in the JSON format. The key indicates the ID of the cycle, and the value is presented in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// {"1":"03:58","2":"04:58","3":"05:58","4":"06:58","5":"07:58","6":"08:58","7":"09:58","8":"10:58","9":"11:58","10":"12:58","11":"13:58","12":"14:58","13":"15:58","14":"16:58","15":"17:58","16":"18:58","17":"19:58","18":"20:58","19":"21:58","20":"22:58","21":"23:58","22":"24:58","23":"25:58","24":"26:58"}
	HourSlaDetail *string `json:"HourSlaDetail,omitempty" xml:"HourSlaDetail,omitempty"`
	// Indicates whether the baseline is a default baseline of the workspace. Valid values: true and false.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs can be specified. The IDs are separated by commas (,).
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
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
	// The hour in the committed completion time of the day-level baseline. Valid values: [0, 47].
	//
	// example:
	//
	// 9
	SlaHour *int32 `json:"SlaHour,omitempty" xml:"SlaHour,omitempty"`
	// The minute in the alerting time of the day-level baseline. Valid values: [0, 59].
	//
	// example:
	//
	// 30
	SlaMinu *int32 `json:"SlaMinu,omitempty" xml:"SlaMinu,omitempty"`
	// Indicates whether the baseline is enabled. Valid values: true and false.
	//
	// example:
	//
	// true
	UseFlag *bool `json:"UseFlag,omitempty" xml:"UseFlag,omitempty"`
}

func (s ListBaselineConfigsResponseBodyDataBaselines) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineConfigsResponseBodyDataBaselines) GoString() string {
	return s.String()
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetBaselineName() *string {
	return s.BaselineName
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetBaselineType() *string {
	return s.BaselineType
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetExpHour() *int32 {
	return s.ExpHour
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetExpMinu() *int32 {
	return s.ExpMinu
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetHourExpDetail() *string {
	return s.HourExpDetail
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetHourSlaDetail() *string {
	return s.HourSlaDetail
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetPriority() *int32 {
	return s.Priority
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetSlaHour() *int32 {
	return s.SlaHour
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetSlaMinu() *int32 {
	return s.SlaMinu
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) GetUseFlag() *bool {
	return s.UseFlag
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetBaselineId(v int64) *ListBaselineConfigsResponseBodyDataBaselines {
	s.BaselineId = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetBaselineName(v string) *ListBaselineConfigsResponseBodyDataBaselines {
	s.BaselineName = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetBaselineType(v string) *ListBaselineConfigsResponseBodyDataBaselines {
	s.BaselineType = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetExpHour(v int32) *ListBaselineConfigsResponseBodyDataBaselines {
	s.ExpHour = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetExpMinu(v int32) *ListBaselineConfigsResponseBodyDataBaselines {
	s.ExpMinu = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetHourExpDetail(v string) *ListBaselineConfigsResponseBodyDataBaselines {
	s.HourExpDetail = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetHourSlaDetail(v string) *ListBaselineConfigsResponseBodyDataBaselines {
	s.HourSlaDetail = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetIsDefault(v bool) *ListBaselineConfigsResponseBodyDataBaselines {
	s.IsDefault = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetOwner(v string) *ListBaselineConfigsResponseBodyDataBaselines {
	s.Owner = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetPriority(v int32) *ListBaselineConfigsResponseBodyDataBaselines {
	s.Priority = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetProjectId(v int64) *ListBaselineConfigsResponseBodyDataBaselines {
	s.ProjectId = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetSlaHour(v int32) *ListBaselineConfigsResponseBodyDataBaselines {
	s.SlaHour = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetSlaMinu(v int32) *ListBaselineConfigsResponseBodyDataBaselines {
	s.SlaMinu = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) SetUseFlag(v bool) *ListBaselineConfigsResponseBodyDataBaselines {
	s.UseFlag = &v
	return s
}

func (s *ListBaselineConfigsResponseBodyDataBaselines) Validate() error {
	return dara.Validate(s)
}
