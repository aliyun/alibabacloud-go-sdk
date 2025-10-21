// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *CreateTopicShrinkRequest
	GetBodyDataShrink() *string
	SetRegionId(v string) *CreateTopicShrinkRequest
	GetRegionId() *string
	SetTopicDefinition(v string) *CreateTopicShrinkRequest
	GetTopicDefinition() *string
	SetTopicName(v string) *CreateTopicShrinkRequest
	GetTopicName() *string
	SetWorkspaceId(v int64) *CreateTopicShrinkRequest
	GetWorkspaceId() *int64
}

type CreateTopicShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
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

func (s CreateTopicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *CreateTopicShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTopicShrinkRequest) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *CreateTopicShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateTopicShrinkRequest) SetBodyDataShrink(v string) *CreateTopicShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetRegionId(v string) *CreateTopicShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetTopicDefinition(v string) *CreateTopicShrinkRequest {
	s.TopicDefinition = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetTopicName(v string) *CreateTopicShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetWorkspaceId(v int64) *CreateTopicShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTopicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
