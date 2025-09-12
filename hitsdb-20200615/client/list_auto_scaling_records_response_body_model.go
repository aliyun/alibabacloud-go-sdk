// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAutoScalingRecordsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListAutoScalingRecordsResponseBody
	GetCode() *string
	SetData(v *ListAutoScalingRecordsResponseBodyData) *ListAutoScalingRecordsResponseBody
	GetData() *ListAutoScalingRecordsResponseBodyData
	SetDynamicCode(v string) *ListAutoScalingRecordsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAutoScalingRecordsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListAutoScalingRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAutoScalingRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAutoScalingRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAutoScalingRecordsResponseBody
	GetSuccess() *bool
}

type ListAutoScalingRecordsResponseBody struct {
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAutoScalingRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAutoScalingRecordsResponseBody) GetData() *ListAutoScalingRecordsResponseBodyData {
	return s.Data
}

func (s *ListAutoScalingRecordsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAutoScalingRecordsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAutoScalingRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAutoScalingRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAutoScalingRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoScalingRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutoScalingRecordsResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetCode(v string) *ListAutoScalingRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetData(v *ListAutoScalingRecordsResponseBodyData) *ListAutoScalingRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetDynamicCode(v string) *ListAutoScalingRecordsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetDynamicMessage(v string) *ListAutoScalingRecordsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetMessage(v string) *ListAutoScalingRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetRequestId(v string) *ListAutoScalingRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetSuccess(v bool) *ListAutoScalingRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAutoScalingRecordsResponseBodyData struct {
	PageNum      *int32                                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScaleRecords []*ListAutoScalingRecordsResponseBodyDataScaleRecords `json:"ScaleRecords,omitempty" xml:"ScaleRecords,omitempty" type:"Repeated"`
	TotalNum     *int32                                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage    *int32                                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListAutoScalingRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAutoScalingRecordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoScalingRecordsResponseBodyData) GetScaleRecords() []*ListAutoScalingRecordsResponseBodyDataScaleRecords {
	return s.ScaleRecords
}

func (s *ListAutoScalingRecordsResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListAutoScalingRecordsResponseBodyData) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListAutoScalingRecordsResponseBodyData) SetPageNum(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetPageSize(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetScaleRecords(v []*ListAutoScalingRecordsResponseBodyDataScaleRecords) *ListAutoScalingRecordsResponseBodyData {
	s.ScaleRecords = v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetTotalNum(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetTotalPage(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAutoScalingRecordsResponseBodyDataScaleRecords struct {
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OldValue     *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpecGroupId  *string `json:"SpecGroupId,omitempty" xml:"SpecGroupId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Strategy     *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	TargetValue  *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ListAutoScalingRecordsResponseBodyDataScaleRecords) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBodyDataScaleRecords) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetDetail() *string {
	return s.Detail
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetId() *string {
	return s.Id
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetOldValue() *string {
	return s.OldValue
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetSpecGroupId() *string {
	return s.SpecGroupId
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetStatus() *string {
	return s.Status
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetStrategy() *string {
	return s.Strategy
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) GetTargetValue() *string {
	return s.TargetValue
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetDetail(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Detail = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetEndTime(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Id = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetInstanceId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetOldValue(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.OldValue = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetResourceType(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.ResourceType = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetSpecGroupId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.SpecGroupId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStartTime(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStatus(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Status = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStrategy(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Strategy = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetTargetValue(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.TargetValue = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) Validate() error {
	return dara.Validate(s)
}
