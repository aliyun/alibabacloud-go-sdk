// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTopicRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteTopicRequest
	GetRegionId() *string
	SetTopic(v string) *DeleteTopicRequest
	GetTopic() *string
}

type DeleteTopicRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTopicRequest) GetTopic() *string {
	return s.Topic
}

func (s *DeleteTopicRequest) SetInstanceId(v string) *DeleteTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTopicRequest) SetRegionId(v string) *DeleteTopicRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTopicRequest) SetTopic(v string) *DeleteTopicRequest {
	s.Topic = &v
	return s
}

func (s *DeleteTopicRequest) Validate() error {
	return dara.Validate(s)
}
