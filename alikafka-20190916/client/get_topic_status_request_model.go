// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTopicStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetTopicStatusRequest
	GetRegionId() *string
	SetTopic(v string) *GetTopicStatusRequest
	GetTopic() *string
}

type GetTopicStatusRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h15tjm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
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
	// normal_topic_9d034262835916103455551be06cc****
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTopicStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTopicStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicStatusRequest) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicStatusRequest) SetInstanceId(v string) *GetTopicStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicStatusRequest) SetRegionId(v string) *GetTopicStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicStatusRequest) SetTopic(v string) *GetTopicStatusRequest {
	s.Topic = &v
	return s
}

func (s *GetTopicStatusRequest) Validate() error {
	return dara.Validate(s)
}
