// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *UpdateTopicShrinkRequest
	GetBodyDataShrink() *string
	SetRegionId(v string) *UpdateTopicShrinkRequest
	GetRegionId() *string
	SetTopicDefinition(v string) *UpdateTopicShrinkRequest
	GetTopicDefinition() *string
	SetTopicId(v int64) *UpdateTopicShrinkRequest
	GetTopicId() *int64
	SetTopicName(v string) *UpdateTopicShrinkRequest
	GetTopicName() *string
}

type UpdateTopicShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
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

func (s UpdateTopicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *UpdateTopicShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTopicShrinkRequest) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *UpdateTopicShrinkRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *UpdateTopicShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdateTopicShrinkRequest) SetBodyDataShrink(v string) *UpdateTopicShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *UpdateTopicShrinkRequest) SetRegionId(v string) *UpdateTopicShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTopicShrinkRequest) SetTopicDefinition(v string) *UpdateTopicShrinkRequest {
	s.TopicDefinition = &v
	return s
}

func (s *UpdateTopicShrinkRequest) SetTopicId(v int64) *UpdateTopicShrinkRequest {
	s.TopicId = &v
	return s
}

func (s *UpdateTopicShrinkRequest) SetTopicName(v string) *UpdateTopicShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *UpdateTopicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
