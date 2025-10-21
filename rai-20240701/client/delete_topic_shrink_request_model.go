// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteTopicShrinkRequest
	GetRegionId() *string
	SetTopicIdListShrink(v string) *DeleteTopicShrinkRequest
	GetTopicIdListShrink() *string
}

type DeleteTopicShrinkRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopicIdListShrink *string `json:"TopicIdList,omitempty" xml:"TopicIdList,omitempty"`
}

func (s DeleteTopicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTopicShrinkRequest) GetTopicIdListShrink() *string {
	return s.TopicIdListShrink
}

func (s *DeleteTopicShrinkRequest) SetRegionId(v string) *DeleteTopicShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTopicShrinkRequest) SetTopicIdListShrink(v string) *DeleteTopicShrinkRequest {
	s.TopicIdListShrink = &v
	return s
}

func (s *DeleteTopicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
