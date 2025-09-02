// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBaselineConfigResponseBodyData) *GetBaselineConfigResponseBody
	GetData() *GetBaselineConfigResponseBodyData
	SetErrorCode(v string) *GetBaselineConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBaselineConfigResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBaselineConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBaselineConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaselineConfigResponseBody
	GetSuccess() *bool
}

type GetBaselineConfigResponseBody struct {
	// The details of the baseline.
	Data *GetBaselineConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 401
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ecb967ec-c137-48a5-860****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineConfigResponseBody) GetData() *GetBaselineConfigResponseBodyData {
	return s.Data
}

func (s *GetBaselineConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBaselineConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBaselineConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaselineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaselineConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaselineConfigResponseBody) SetData(v *GetBaselineConfigResponseBodyData) *GetBaselineConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineConfigResponseBody) SetErrorCode(v string) *GetBaselineConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBaselineConfigResponseBody) SetErrorMessage(v string) *GetBaselineConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBaselineConfigResponseBody) SetHttpStatusCode(v int32) *GetBaselineConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineConfigResponseBody) SetRequestId(v string) *GetBaselineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineConfigResponseBody) SetSuccess(v bool) *GetBaselineConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaselineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBaselineConfigResponseBodyData struct {
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
	// 9527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The hour in the committed completion time of the day-level baseline. Valid values: [0, 47].
	//
	// example:
	//
	// 9
	SlaHour *int32 `json:"SlaHour,omitempty" xml:"SlaHour,omitempty"`
	// The minute in the committed completion time of the day-level baseline. Valid values: [0, 59].
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

func (s GetBaselineConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineConfigResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineConfigResponseBodyData) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetBaselineConfigResponseBodyData) GetBaselineType() *string {
	return s.BaselineType
}

func (s *GetBaselineConfigResponseBodyData) GetExpHour() *int32 {
	return s.ExpHour
}

func (s *GetBaselineConfigResponseBodyData) GetExpMinu() *int32 {
	return s.ExpMinu
}

func (s *GetBaselineConfigResponseBodyData) GetHourExpDetail() *string {
	return s.HourExpDetail
}

func (s *GetBaselineConfigResponseBodyData) GetHourSlaDetail() *string {
	return s.HourSlaDetail
}

func (s *GetBaselineConfigResponseBodyData) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetBaselineConfigResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineConfigResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetBaselineConfigResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineConfigResponseBodyData) GetSlaHour() *int32 {
	return s.SlaHour
}

func (s *GetBaselineConfigResponseBodyData) GetSlaMinu() *int32 {
	return s.SlaMinu
}

func (s *GetBaselineConfigResponseBodyData) GetUseFlag() *bool {
	return s.UseFlag
}

func (s *GetBaselineConfigResponseBodyData) SetBaselineId(v int64) *GetBaselineConfigResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetBaselineName(v string) *GetBaselineConfigResponseBodyData {
	s.BaselineName = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetBaselineType(v string) *GetBaselineConfigResponseBodyData {
	s.BaselineType = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetExpHour(v int32) *GetBaselineConfigResponseBodyData {
	s.ExpHour = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetExpMinu(v int32) *GetBaselineConfigResponseBodyData {
	s.ExpMinu = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetHourExpDetail(v string) *GetBaselineConfigResponseBodyData {
	s.HourExpDetail = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetHourSlaDetail(v string) *GetBaselineConfigResponseBodyData {
	s.HourSlaDetail = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetIsDefault(v bool) *GetBaselineConfigResponseBodyData {
	s.IsDefault = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetOwner(v string) *GetBaselineConfigResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetPriority(v int32) *GetBaselineConfigResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetProjectId(v int64) *GetBaselineConfigResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetSlaHour(v int32) *GetBaselineConfigResponseBodyData {
	s.SlaHour = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetSlaMinu(v int32) *GetBaselineConfigResponseBodyData {
	s.SlaMinu = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) SetUseFlag(v bool) *GetBaselineConfigResponseBodyData {
	s.UseFlag = &v
	return s
}

func (s *GetBaselineConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
