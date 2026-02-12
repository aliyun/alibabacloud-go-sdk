// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceInServiceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsInstanceInServiceListResponseBodyData) *OnsInstanceInServiceListResponseBody
	GetData() *OnsInstanceInServiceListResponseBodyData
	SetRequestId(v string) *OnsInstanceInServiceListResponseBody
	GetRequestId() *string
}

type OnsInstanceInServiceListResponseBody struct {
	Data *OnsInstanceInServiceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0598E46F-DB06-40E2-AD7B-C45923EE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceInServiceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBody) GetData() *OnsInstanceInServiceListResponseBodyData {
	return s.Data
}

func (s *OnsInstanceInServiceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsInstanceInServiceListResponseBody) SetData(v *OnsInstanceInServiceListResponseBodyData) *OnsInstanceInServiceListResponseBody {
	s.Data = v
	return s
}

func (s *OnsInstanceInServiceListResponseBody) SetRequestId(v string) *OnsInstanceInServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsInstanceInServiceListResponseBodyData struct {
	InstanceVO []*OnsInstanceInServiceListResponseBodyDataInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyData) GetInstanceVO() []*OnsInstanceInServiceListResponseBodyDataInstanceVO {
	return s.InstanceVO
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceVO(v []*OnsInstanceInServiceListResponseBodyDataInstanceVO) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceVO = v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) Validate() error {
	if s.InstanceVO != nil {
		for _, item := range s.InstanceVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsInstanceInServiceListResponseBodyDataInstanceVO struct {
	CreateTime        *int64                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupCount        *int32                                                  `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	IndependentNaming *bool                                                   `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus    *int32                                                  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType      *int32                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ReleaseTime       *int64                                                  `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	Tags              *OnsInstanceInServiceListResponseBodyDataInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TopicCount        *int32                                                  `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVO) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVO) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetIndependentNaming() *bool {
	return s.IndependentNaming
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetInstanceStatus() *int32 {
	return s.InstanceStatus
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetInstanceType() *int32 {
	return s.InstanceType
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetTags() *OnsInstanceInServiceListResponseBodyDataInstanceVOTags {
	return s.Tags
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetCreateTime(v int64) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetGroupCount(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.GroupCount = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetIndependentNaming(v bool) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceId(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceName(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceStatus(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceType(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceType = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetReleaseTime(v int64) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetTags(v *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.Tags = v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetTopicCount(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.TopicCount = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsInstanceInServiceListResponseBodyDataInstanceVOTags struct {
	Tag []*OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTags) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTags) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) GetTag() []*OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag {
	return s.Tag
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) SetTag(v []*OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) *OnsInstanceInServiceListResponseBodyDataInstanceVOTags {
	s.Tag = v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) Validate() error {
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

type OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) GetKey() *string {
	return s.Key
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) GetValue() *string {
	return s.Value
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) SetKey(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag {
	s.Key = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) SetValue(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag {
	s.Value = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) Validate() error {
	return dara.Validate(s)
}
