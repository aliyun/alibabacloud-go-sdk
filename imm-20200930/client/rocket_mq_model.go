// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRocketMQ interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RocketMQ
	GetInstanceId() *string
	SetTopicName(v string) *RocketMQ
	GetTopicName() *string
}

type RocketMQ struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TopicName  *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s RocketMQ) String() string {
	return dara.Prettify(s)
}

func (s RocketMQ) GoString() string {
	return s.String()
}

func (s *RocketMQ) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RocketMQ) GetTopicName() *string {
	return s.TopicName
}

func (s *RocketMQ) SetInstanceId(v string) *RocketMQ {
	s.InstanceId = &v
	return s
}

func (s *RocketMQ) SetTopicName(v string) *RocketMQ {
	s.TopicName = &v
	return s
}

func (s *RocketMQ) Validate() error {
	return dara.Validate(s)
}
