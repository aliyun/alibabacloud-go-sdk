// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v bool) *OnsConsumerStatusRequest
	GetDetail() *bool
	SetGroupId(v string) *OnsConsumerStatusRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsConsumerStatusRequest
	GetInstanceId() *string
	SetNeedJstack(v bool) *OnsConsumerStatusRequest
	GetNeedJstack() *bool
}

type OnsConsumerStatusRequest struct {
	// Specifies whether to query the details of the consumer group. Valid values:
	//
	// 	- **true**: The details of the consumer group are queried. You can obtain details from the **ConsumerConnectionInfoList*	- and **DetailInTopicList*	- response parameters.
	//
	// 	- **false**: The details of the consumer group are not queried. The values of the **ConsumerConnectionInfoList*	- and **DetailInTopicList*	- response parameters are empty. This value is the default value of the Detail parameter.
	//
	// example:
	//
	// true
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the consumer group whose details you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query the information about thread stack traces. Valid values:
	//
	// 	- **true**: The information about thread stack traces is queried. You can obtain the information from the **Jstack*	- response parameter.
	//
	// > If you want to obtain the information about thread stack traces, make sure that the **Detail*	- parameter in the request is set to **true**.
	//
	// 	- **false**: The information about thread stack traces is not queried. The value of the **Jstack*	- response parameter is empty. This value is the default value of the NeedJstack parameter.
	//
	// example:
	//
	// true
	NeedJstack *bool `json:"NeedJstack,omitempty" xml:"NeedJstack,omitempty"`
}

func (s OnsConsumerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusRequest) GetDetail() *bool {
	return s.Detail
}

func (s *OnsConsumerStatusRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsConsumerStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerStatusRequest) GetNeedJstack() *bool {
	return s.NeedJstack
}

func (s *OnsConsumerStatusRequest) SetDetail(v bool) *OnsConsumerStatusRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetGroupId(v string) *OnsConsumerStatusRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetInstanceId(v string) *OnsConsumerStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetNeedJstack(v bool) *OnsConsumerStatusRequest {
	s.NeedJstack = &v
	return s
}

func (s *OnsConsumerStatusRequest) Validate() error {
	return dara.Validate(s)
}
