// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *CreateTopicRequestBodyData) *CreateTopicRequest
	GetBodyData() *CreateTopicRequestBodyData
	SetRegionId(v string) *CreateTopicRequest
	GetRegionId() *string
	SetTopicDefinition(v string) *CreateTopicRequest
	GetTopicDefinition() *string
	SetTopicName(v string) *CreateTopicRequest
	GetTopicName() *string
	SetWorkspaceId(v int64) *CreateTopicRequest
	GetWorkspaceId() *int64
}

type CreateTopicRequest struct {
	BodyData *CreateTopicRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// cn-shanghai
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopicDefinition *string `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	TopicName       *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 643168
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) GetBodyData() *CreateTopicRequestBodyData {
	return s.BodyData
}

func (s *CreateTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTopicRequest) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *CreateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateTopicRequest) SetBodyData(v *CreateTopicRequestBodyData) *CreateTopicRequest {
	s.BodyData = v
	return s
}

func (s *CreateTopicRequest) SetRegionId(v string) *CreateTopicRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTopicRequest) SetTopicDefinition(v string) *CreateTopicRequest {
	s.TopicDefinition = &v
	return s
}

func (s *CreateTopicRequest) SetTopicName(v string) *CreateTopicRequest {
	s.TopicName = &v
	return s
}

func (s *CreateTopicRequest) SetWorkspaceId(v int64) *CreateTopicRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTopicRequestBodyData struct {
	TopicExampleInfoList []*CreateTopicRequestBodyDataTopicExampleInfoList `json:"TopicExampleInfoList,omitempty" xml:"TopicExampleInfoList,omitempty" type:"Repeated"`
}

func (s CreateTopicRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequestBodyData) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestBodyData) GetTopicExampleInfoList() []*CreateTopicRequestBodyDataTopicExampleInfoList {
	return s.TopicExampleInfoList
}

func (s *CreateTopicRequestBodyData) SetTopicExampleInfoList(v []*CreateTopicRequestBodyDataTopicExampleInfoList) *CreateTopicRequestBodyData {
	s.TopicExampleInfoList = v
	return s
}

func (s *CreateTopicRequestBodyData) Validate() error {
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

type CreateTopicRequestBodyDataTopicExampleInfoList struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ExampleType *int32 `json:"ExampleType,omitempty" xml:"ExampleType,omitempty"`
}

func (s CreateTopicRequestBodyDataTopicExampleInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequestBodyDataTopicExampleInfoList) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestBodyDataTopicExampleInfoList) GetContent() *string {
	return s.Content
}

func (s *CreateTopicRequestBodyDataTopicExampleInfoList) GetExampleType() *int32 {
	return s.ExampleType
}

func (s *CreateTopicRequestBodyDataTopicExampleInfoList) SetContent(v string) *CreateTopicRequestBodyDataTopicExampleInfoList {
	s.Content = &v
	return s
}

func (s *CreateTopicRequestBodyDataTopicExampleInfoList) SetExampleType(v int32) *CreateTopicRequestBodyDataTopicExampleInfoList {
	s.ExampleType = &v
	return s
}

func (s *CreateTopicRequestBodyDataTopicExampleInfoList) Validate() error {
	return dara.Validate(s)
}
