// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsGroupListResponseBodyData) *OnsGroupListResponseBody
	GetData() *OnsGroupListResponseBodyData
	SetRequestId(v string) *OnsGroupListResponseBody
	GetRequestId() *string
}

type OnsGroupListResponseBody struct {
	Data *OnsGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 16996623-AC4A-43AF-9248-FD9D2D75****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBody) GetData() *OnsGroupListResponseBodyData {
	return s.Data
}

func (s *OnsGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsGroupListResponseBody) SetData(v *OnsGroupListResponseBodyData) *OnsGroupListResponseBody {
	s.Data = v
	return s
}

func (s *OnsGroupListResponseBody) SetRequestId(v string) *OnsGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsGroupListResponseBodyData struct {
	SubscribeInfoDo []*OnsGroupListResponseBodyDataSubscribeInfoDo `json:"SubscribeInfoDo,omitempty" xml:"SubscribeInfoDo,omitempty" type:"Repeated"`
}

func (s OnsGroupListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyData) GetSubscribeInfoDo() []*OnsGroupListResponseBodyDataSubscribeInfoDo {
	return s.SubscribeInfoDo
}

func (s *OnsGroupListResponseBodyData) SetSubscribeInfoDo(v []*OnsGroupListResponseBodyDataSubscribeInfoDo) *OnsGroupListResponseBodyData {
	s.SubscribeInfoDo = v
	return s
}

func (s *OnsGroupListResponseBodyData) Validate() error {
	if s.SubscribeInfoDo != nil {
		for _, item := range s.SubscribeInfoDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsGroupListResponseBodyDataSubscribeInfoDo struct {
	CreateTime        *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType         *string                                          `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IndependentNaming *bool                                            `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Owner             *string                                          `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Remark            *string                                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Tags              *OnsGroupListResponseBodyDataSubscribeInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateTime        *int64                                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDo) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDo) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetGroupType() *string {
	return s.GroupType
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetIndependentNaming() *bool {
	return s.IndependentNaming
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetOwner() *string {
	return s.Owner
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetRemark() *string {
	return s.Remark
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetTags() *OnsGroupListResponseBodyDataSubscribeInfoDoTags {
	return s.Tags
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetCreateTime(v int64) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.CreateTime = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetGroupId(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetGroupType(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.GroupType = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetIndependentNaming(v bool) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetInstanceId(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetOwner(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Owner = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetRemark(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Remark = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetTags(v *OnsGroupListResponseBodyDataSubscribeInfoDoTags) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Tags = v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetUpdateTime(v int64) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.UpdateTime = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsGroupListResponseBodyDataSubscribeInfoDoTags struct {
	Tag []*OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTags) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTags) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTags) GetTag() []*OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag {
	return s.Tag
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTags) SetTag(v []*OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) *OnsGroupListResponseBodyDataSubscribeInfoDoTags {
	s.Tag = v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTags) Validate() error {
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

type OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) GetKey() *string {
	return s.Key
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) GetValue() *string {
	return s.Value
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) SetKey(v string) *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag {
	s.Key = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) SetValue(v string) *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag {
	s.Value = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) Validate() error {
	return dara.Validate(s)
}
