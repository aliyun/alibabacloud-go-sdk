// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerGetConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsConsumerGetConnectionRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsConsumerGetConnectionRequest
	GetInstanceId() *string
}

type OnsConsumerGetConnectionRequest struct {
	// The ID of the consumer group whose client connection status you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerGetConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerGetConnectionRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsConsumerGetConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerGetConnectionRequest) SetGroupId(v string) *OnsConsumerGetConnectionRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerGetConnectionRequest) SetInstanceId(v string) *OnsConsumerGetConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerGetConnectionRequest) Validate() error {
	return dara.Validate(s)
}
