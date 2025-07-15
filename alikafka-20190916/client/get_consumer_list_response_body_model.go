// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetConsumerListResponseBody
	GetCode() *int32
	SetConsumerList(v *GetConsumerListResponseBodyConsumerList) *GetConsumerListResponseBody
	GetConsumerList() *GetConsumerListResponseBodyConsumerList
	SetCurrentPage(v int32) *GetConsumerListResponseBody
	GetCurrentPage() *int32
	SetMessage(v string) *GetConsumerListResponseBody
	GetMessage() *string
	SetPageSize(v int32) *GetConsumerListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetConsumerListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerListResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *GetConsumerListResponseBody
	GetTotal() *int64
}

type GetConsumerListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumer groups.
	ConsumerList *GetConsumerListResponseBodyConsumerList `json:"ConsumerList,omitempty" xml:"ConsumerList,omitempty" type:"Struct"`
	// The number of the page to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 808F042B-CB9A-4FBC-9009-00E7DDB6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetConsumerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetConsumerListResponseBody) GetConsumerList() *GetConsumerListResponseBodyConsumerList {
	return s.ConsumerList
}

func (s *GetConsumerListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetConsumerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetConsumerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetConsumerListResponseBody) SetCode(v int32) *GetConsumerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerListResponseBody) SetConsumerList(v *GetConsumerListResponseBodyConsumerList) *GetConsumerListResponseBody {
	s.ConsumerList = v
	return s
}

func (s *GetConsumerListResponseBody) SetCurrentPage(v int32) *GetConsumerListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetConsumerListResponseBody) SetMessage(v string) *GetConsumerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerListResponseBody) SetPageSize(v int32) *GetConsumerListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetConsumerListResponseBody) SetRequestId(v string) *GetConsumerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerListResponseBody) SetSuccess(v bool) *GetConsumerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerListResponseBody) SetTotal(v int64) *GetConsumerListResponseBody {
	s.Total = &v
	return s
}

func (s *GetConsumerListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConsumerListResponseBodyConsumerList struct {
	ConsumerVO []*GetConsumerListResponseBodyConsumerListConsumerVO `json:"ConsumerVO,omitempty" xml:"ConsumerVO,omitempty" type:"Repeated"`
}

func (s GetConsumerListResponseBodyConsumerList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerList) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerList) GetConsumerVO() []*GetConsumerListResponseBodyConsumerListConsumerVO {
	return s.ConsumerVO
}

func (s *GetConsumerListResponseBodyConsumerList) SetConsumerVO(v []*GetConsumerListResponseBodyConsumerListConsumerVO) *GetConsumerListResponseBodyConsumerList {
	s.ConsumerVO = v
	return s
}

func (s *GetConsumerListResponseBodyConsumerList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerListResponseBodyConsumerListConsumerVO struct {
	// Indicates that the consumer group was automatically created by the system.
	//
	// example:
	//
	// false
	AutomaticallyCreatedGroup *bool `json:"AutomaticallyCreatedGroup,omitempty" xml:"AutomaticallyCreatedGroup,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The timestamp that indicates when the consumer group was created. Unit: milliseconds.
	//
	// example:
	//
	// 1729736584002
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_post-cn-v0h18sav****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags.
	Tags *GetConsumerListResponseBodyConsumerListConsumerVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetAutomaticallyCreatedGroup() *bool {
	return s.AutomaticallyCreatedGroup
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetRemark() *string {
	return s.Remark
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) GetTags() *GetConsumerListResponseBodyConsumerListConsumerVOTags {
	return s.Tags
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetAutomaticallyCreatedGroup(v bool) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.AutomaticallyCreatedGroup = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetConsumerId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetCreateTime(v int64) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.CreateTime = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetInstanceId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetRegionId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.RegionId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetRemark(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.Remark = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetTags(v *GetConsumerListResponseBodyConsumerListConsumerVOTags) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.Tags = v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) Validate() error {
	return dara.Validate(s)
}

type GetConsumerListResponseBodyConsumerListConsumerVOTags struct {
	TagVO []*GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTags) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTags) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTags) GetTagVO() []*GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO {
	return s.TagVO
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTags) SetTagVO(v []*GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) *GetConsumerListResponseBodyConsumerListConsumerVOTags {
	s.TagVO = v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTags) Validate() error {
	return dara.Validate(s)
}

type GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) GetKey() *string {
	return s.Key
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) GetValue() *string {
	return s.Value
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) SetKey(v string) *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) SetValue(v string) *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO {
	s.Value = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) Validate() error {
	return dara.Validate(s)
}
