// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsTopicListResponseBodyData) *OnsTopicListResponseBody
	GetData() *OnsTopicListResponseBodyData
	SetRequestId(v string) *OnsTopicListResponseBody
	GetRequestId() *string
}

type OnsTopicListResponseBody struct {
	Data *OnsTopicListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4A978869-7681-4529-B470-107E1379****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBody) GetData() *OnsTopicListResponseBodyData {
	return s.Data
}

func (s *OnsTopicListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicListResponseBody) SetData(v *OnsTopicListResponseBodyData) *OnsTopicListResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicListResponseBody) SetRequestId(v string) *OnsTopicListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTopicListResponseBodyData struct {
	PublishInfoDo []*OnsTopicListResponseBodyDataPublishInfoDo `json:"PublishInfoDo,omitempty" xml:"PublishInfoDo,omitempty" type:"Repeated"`
}

func (s OnsTopicListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyData) GetPublishInfoDo() []*OnsTopicListResponseBodyDataPublishInfoDo {
	return s.PublishInfoDo
}

func (s *OnsTopicListResponseBodyData) SetPublishInfoDo(v []*OnsTopicListResponseBodyDataPublishInfoDo) *OnsTopicListResponseBodyData {
	s.PublishInfoDo = v
	return s
}

func (s *OnsTopicListResponseBodyData) Validate() error {
	if s.PublishInfoDo != nil {
		for _, item := range s.PublishInfoDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTopicListResponseBodyDataPublishInfoDo struct {
	CreateTime        *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IndependentNaming *bool                                          `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageType       *int32                                         `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Owner             *string                                        `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Relation          *int32                                         `json:"Relation,omitempty" xml:"Relation,omitempty"`
	RelationName      *string                                        `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
	Remark            *string                                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ServiceStatus     *int32                                         `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	Tags              *OnsTopicListResponseBodyDataPublishInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic             *string                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDo) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetIndependentNaming() *bool {
	return s.IndependentNaming
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetMessageType() *int32 {
	return s.MessageType
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetOwner() *string {
	return s.Owner
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetRelation() *int32 {
	return s.Relation
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetRelationName() *string {
	return s.RelationName
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetRemark() *string {
	return s.Remark
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetTags() *OnsTopicListResponseBodyDataPublishInfoDoTags {
	return s.Tags
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetCreateTime(v int64) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.CreateTime = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetIndependentNaming(v bool) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetInstanceId(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetMessageType(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.MessageType = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetOwner(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Owner = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRelation(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Relation = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRelationName(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.RelationName = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRemark(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Remark = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetServiceStatus(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.ServiceStatus = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetTags(v *OnsTopicListResponseBodyDataPublishInfoDoTags) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Tags = v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetTopic(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Topic = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTopicListResponseBodyDataPublishInfoDoTags struct {
	Tag []*OnsTopicListResponseBodyDataPublishInfoDoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTags) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTags) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTags) GetTag() []*OnsTopicListResponseBodyDataPublishInfoDoTagsTag {
	return s.Tag
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTags) SetTag(v []*OnsTopicListResponseBodyDataPublishInfoDoTagsTag) *OnsTopicListResponseBodyDataPublishInfoDoTags {
	s.Tag = v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTopicListResponseBodyDataPublishInfoDoTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTagsTag) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTagsTag) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) GetKey() *string {
	return s.Key
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) GetValue() *string {
	return s.Value
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) SetKey(v string) *OnsTopicListResponseBodyDataPublishInfoDoTagsTag {
	s.Key = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) SetValue(v string) *OnsTopicListResponseBodyDataPublishInfoDoTagsTag {
	s.Value = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) Validate() error {
	return dara.Validate(s)
}
