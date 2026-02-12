// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerAccumulateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v bool) *OnsConsumerAccumulateRequest
	GetDetail() *bool
	SetGroupId(v string) *OnsConsumerAccumulateRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsConsumerAccumulateRequest
	GetInstanceId() *string
}

type OnsConsumerAccumulateRequest struct {
	// Specifies whether to query the details of each topic to which the consumer group subscribes. Valid values:
	//
	// 	- **true**: The details of each topic are queried. You can obtain the details from the **DetailInTopicList*	- response parameter.
	//
	// 	- **false**: The details of each topic are not queried. This is the default value. If you use this value, the value of the **DetailInTopicList*	- response parameter is empty.
	//
	// example:
	//
	// true
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerAccumulateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateRequest) GetDetail() *bool {
	return s.Detail
}

func (s *OnsConsumerAccumulateRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsConsumerAccumulateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerAccumulateRequest) SetDetail(v bool) *OnsConsumerAccumulateRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetGroupId(v string) *OnsConsumerAccumulateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetInstanceId(v string) *OnsConsumerAccumulateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) Validate() error {
	return dara.Validate(s)
}
