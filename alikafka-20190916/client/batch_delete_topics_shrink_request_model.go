// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteTopicsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BatchDeleteTopicsShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *BatchDeleteTopicsShrinkRequest
	GetRegionId() *string
	SetTopicsShrink(v string) *BatchDeleteTopicsShrinkRequest
	GetTopicsShrink() *string
}

type BatchDeleteTopicsShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
}

func (s BatchDeleteTopicsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteTopicsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchDeleteTopicsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchDeleteTopicsShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *BatchDeleteTopicsShrinkRequest) SetInstanceId(v string) *BatchDeleteTopicsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchDeleteTopicsShrinkRequest) SetRegionId(v string) *BatchDeleteTopicsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *BatchDeleteTopicsShrinkRequest) SetTopicsShrink(v string) *BatchDeleteTopicsShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *BatchDeleteTopicsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
