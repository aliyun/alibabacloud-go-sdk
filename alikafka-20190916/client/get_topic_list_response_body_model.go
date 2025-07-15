// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTopicListResponseBody
	GetCode() *int32
	SetCurrentPage(v int32) *GetTopicListResponseBody
	GetCurrentPage() *int32
	SetMessage(v string) *GetTopicListResponseBody
	GetMessage() *string
	SetPageSize(v int32) *GetTopicListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTopicListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicListResponseBody
	GetSuccess() *bool
	SetTopicList(v *GetTopicListResponseBodyTopicList) *GetTopicListResponseBody
	GetTopicList() *GetTopicListResponseBodyTopicList
	SetTotal(v int32) *GetTopicListResponseBody
	GetTotal() *int32
}

type GetTopicListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C0D3DC5B-5C37-47AD-9F22-1F559880****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topics.
	TopicList *GetTopicListResponseBodyTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The number of topics.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTopicListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTopicListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTopicListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTopicListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicListResponseBody) GetTopicList() *GetTopicListResponseBodyTopicList {
	return s.TopicList
}

func (s *GetTopicListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetTopicListResponseBody) SetCode(v int32) *GetTopicListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicListResponseBody) SetCurrentPage(v int32) *GetTopicListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListResponseBody) SetMessage(v string) *GetTopicListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicListResponseBody) SetPageSize(v int32) *GetTopicListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTopicListResponseBody) SetRequestId(v string) *GetTopicListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicListResponseBody) SetSuccess(v bool) *GetTopicListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicListResponseBody) SetTopicList(v *GetTopicListResponseBodyTopicList) *GetTopicListResponseBody {
	s.TopicList = v
	return s
}

func (s *GetTopicListResponseBody) SetTotal(v int32) *GetTopicListResponseBody {
	s.Total = &v
	return s
}

func (s *GetTopicListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTopicListResponseBodyTopicList struct {
	TopicVO []*GetTopicListResponseBodyTopicListTopicVO `json:"TopicVO,omitempty" xml:"TopicVO,omitempty" type:"Repeated"`
}

func (s GetTopicListResponseBodyTopicList) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponseBodyTopicList) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicList) GetTopicVO() []*GetTopicListResponseBodyTopicListTopicVO {
	return s.TopicVO
}

func (s *GetTopicListResponseBodyTopicList) SetTopicVO(v []*GetTopicListResponseBodyTopicListTopicVO) *GetTopicListResponseBodyTopicList {
	s.TopicVO = v
	return s
}

func (s *GetTopicListResponseBodyTopicList) Validate() error {
	return dara.Validate(s)
}

type GetTopicListResponseBodyTopicListTopicVO struct {
	// Indicates whether the topic was automatically created.
	//
	// example:
	//
	// false
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The log cleanup policy for the topic. This parameter is returned only if **LocalTopic*	- is set to **true**. Valid values:
	//
	// 	- false: the default log cleanup policy.
	//
	// 	- true: the Apache Kafka log compaction policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The timestamp that indicates when the topic was created. Unit: milliseconds.
	//
	// example:
	//
	// 1576563109000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1954n****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The storage type that is used for the topic. Valid values:
	//
	// 	- false: cloud storage
	//
	// 	- true: local storage
	//
	// example:
	//
	// false
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The number of partitions in the topic.
	//
	// example:
	//
	// 6
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic description. Valid values:
	//
	// 	- The description can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The description must be 3 to 64 characters in length.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The topic status. Valid value:
	//
	// **0**: running.
	//
	// If the topic is deleted, this parameter is not returned.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic status. Valid value:
	//
	// **Running**.
	//
	// If the topic is deleted, this parameter is not returned.
	//
	// example:
	//
	// Running
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The tags.
	Tags *GetTopicListResponseBodyTopicListTopicVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic name. Valid values:
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. If the name contains more than 64 characters, the system automatically truncates the name.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The topic configuration.
	//
	// example:
	//
	// {"replication-factor":3}
	TopicConfig *string `json:"TopicConfig,omitempty" xml:"TopicConfig,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVO) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVO) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetCompactTopic() *bool {
	return s.CompactTopic
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetLocalTopic() *bool {
	return s.LocalTopic
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetRemark() *string {
	return s.Remark
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetStatus() *int32 {
	return s.Status
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetStatusName() *string {
	return s.StatusName
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetTags() *GetTopicListResponseBodyTopicListTopicVOTags {
	return s.Tags
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicListResponseBodyTopicListTopicVO) GetTopicConfig() *string {
	return s.TopicConfig
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetAutoCreate(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.AutoCreate = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetCompactTopic(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.CompactTopic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetCreateTime(v int64) *GetTopicListResponseBodyTopicListTopicVO {
	s.CreateTime = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetInstanceId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetLocalTopic(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.LocalTopic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetPartitionNum(v int32) *GetTopicListResponseBodyTopicListTopicVO {
	s.PartitionNum = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRegionId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.RegionId = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRemark(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Remark = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatus(v int32) *GetTopicListResponseBodyTopicListTopicVO {
	s.Status = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatusName(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.StatusName = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTags(v *GetTopicListResponseBodyTopicListTopicVOTags) *GetTopicListResponseBodyTopicListTopicVO {
	s.Tags = v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTopic(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Topic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTopicConfig(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.TopicConfig = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) Validate() error {
	return dara.Validate(s)
}

type GetTopicListResponseBodyTopicListTopicVOTags struct {
	TagVO []*GetTopicListResponseBodyTopicListTopicVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetTopicListResponseBodyTopicListTopicVOTags) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVOTags) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVOTags) GetTagVO() []*GetTopicListResponseBodyTopicListTopicVOTagsTagVO {
	return s.TagVO
}

func (s *GetTopicListResponseBodyTopicListTopicVOTags) SetTagVO(v []*GetTopicListResponseBodyTopicListTopicVOTagsTagVO) *GetTopicListResponseBodyTopicListTopicVOTags {
	s.TagVO = v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVOTags) Validate() error {
	return dara.Validate(s)
}

type GetTopicListResponseBodyTopicListTopicVOTagsTagVO struct {
	// The tag key.
	//
	// example:
	//
	// Test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVOTagsTagVO) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) GetKey() *string {
	return s.Key
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) GetValue() *string {
	return s.Value
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) SetKey(v string) *GetTopicListResponseBodyTopicListTopicVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) SetValue(v string) *GetTopicListResponseBodyTopicListTopicVOTagsTagVO {
	s.Value = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) Validate() error {
	return dara.Validate(s)
}
