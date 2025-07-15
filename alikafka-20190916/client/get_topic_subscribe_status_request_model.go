// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSubscribeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTopicSubscribeStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetTopicSubscribeStatusRequest
	GetRegionId() *string
	SetTopic(v string) *GetTopicSubscribeStatusRequest
	GetTopic() *string
}

type GetTopicSubscribeStatusRequest struct {
	// The instance ID.
	//
	// You can call the [GetInstanceList](https://help.aliyun.com/document_detail/437663.html) operation to query the list of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// You can call the [GetTopicList](https://help.aliyun.com/document_detail/437677.html) operation to query the list of topics.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicSubscribeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSubscribeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTopicSubscribeStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicSubscribeStatusRequest) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicSubscribeStatusRequest) SetInstanceId(v string) *GetTopicSubscribeStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicSubscribeStatusRequest) SetRegionId(v string) *GetTopicSubscribeStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicSubscribeStatusRequest) SetTopic(v string) *GetTopicSubscribeStatusRequest {
	s.Topic = &v
	return s
}

func (s *GetTopicSubscribeStatusRequest) Validate() error {
	return dara.Validate(s)
}
