// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *UpdateTopicRequestBodyData) *UpdateTopicRequest
	GetBodyData() *UpdateTopicRequestBodyData
	SetRegionId(v string) *UpdateTopicRequest
	GetRegionId() *string
	SetTopicDefinition(v string) *UpdateTopicRequest
	GetTopicDefinition() *string
	SetTopicId(v int64) *UpdateTopicRequest
	GetTopicId() *int64
	SetTopicName(v string) *UpdateTopicRequest
	GetTopicName() *string
}

type UpdateTopicRequest struct {
	BodyData *UpdateTopicRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// cn-shanghai
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopicDefinition *string `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	// example:
	//
	// 216
	TopicId   *int64  `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UpdateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequest) GetBodyData() *UpdateTopicRequestBodyData {
	return s.BodyData
}

func (s *UpdateTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTopicRequest) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *UpdateTopicRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *UpdateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdateTopicRequest) SetBodyData(v *UpdateTopicRequestBodyData) *UpdateTopicRequest {
	s.BodyData = v
	return s
}

func (s *UpdateTopicRequest) SetRegionId(v string) *UpdateTopicRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTopicRequest) SetTopicDefinition(v string) *UpdateTopicRequest {
	s.TopicDefinition = &v
	return s
}

func (s *UpdateTopicRequest) SetTopicId(v int64) *UpdateTopicRequest {
	s.TopicId = &v
	return s
}

func (s *UpdateTopicRequest) SetTopicName(v string) *UpdateTopicRequest {
	s.TopicName = &v
	return s
}

func (s *UpdateTopicRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTopicRequestBodyData struct {
	TopicExampleInfoList []*UpdateTopicRequestBodyDataTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
}

func (s UpdateTopicRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicRequestBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequestBodyData) GetTopicExampleInfoList() []*UpdateTopicRequestBodyDataTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *UpdateTopicRequestBodyData) SetTopicExampleInfoList(v []*UpdateTopicRequestBodyDataTopicExampleInfoList) *UpdateTopicRequestBodyData {
	s.TopicExampleInfoList = v
	return s
}

func (s *UpdateTopicRequestBodyData) Validate() error {
	if s.TopicExampleInfoList != nil {
		for _, item := range s.TopicExampleInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTopicRequestBodyDataTopicExampleInfoList struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 0
	ExampleType *int32 `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s UpdateTopicRequestBodyDataTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicRequestBodyDataTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequestBodyDataTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *UpdateTopicRequestBodyDataTopicExampleInfoList) GetExampleType() *int32 {
	return s.ExampleType
}

func (s *UpdateTopicRequestBodyDataTopicExampleInfoList) SetContent(v string) *UpdateTopicRequestBodyDataTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *UpdateTopicRequestBodyDataTopicExampleInfoList) SetExampleType(v int32) *UpdateTopicRequestBodyDataTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *UpdateTopicRequestBodyDataTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}
