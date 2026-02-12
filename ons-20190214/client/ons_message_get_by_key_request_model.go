// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsMessageGetByKeyRequest
	GetInstanceId() *string
	SetKey(v string) *OnsMessageGetByKeyRequest
	GetKey() *string
	SetTopic(v string) *OnsMessageGetByKeyRequest
	GetTopic() *string
}

type OnsMessageGetByKeyRequest struct {
	// The ID of the instance to which the messages that you want to query belong.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The key of the messages that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// messageKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The topic that contains the messages that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageGetByKeyRequest) GetKey() *string {
	return s.Key
}

func (s *OnsMessageGetByKeyRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageGetByKeyRequest) SetInstanceId(v string) *OnsMessageGetByKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetKey(v string) *OnsMessageGetByKeyRequest {
	s.Key = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetTopic(v string) *OnsMessageGetByKeyRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) Validate() error {
	return dara.Validate(s)
}
