// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsGroupCreateRequest
	GetGroupId() *string
	SetGroupType(v string) *OnsGroupCreateRequest
	GetGroupType() *string
	SetInstanceId(v string) *OnsGroupCreateRequest
	GetInstanceId() *string
	SetRemark(v string) *OnsGroupCreateRequest
	GetRemark() *string
}

type OnsGroupCreateRequest struct {
	// The ID of the consumer group that you want to create. The group ID must meet the following rules:
	//
	// 	- The group ID must be 2 to 64 characters in length and can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- If the ApsaraMQ for RocketMQ instance in which you want to create the consumer group uses a namespace, the group ID must be unique in the instance. The group ID cannot be the same as an existing group ID or a topic name in the instance. The group ID can be the same as a group ID or a topic name in another instance that uses a different namespace. For example, if Instance A and Instance B use different namespaces, a group ID in Instance A can be the same as a group ID or a topic name in Instance B.
	//
	// 	- If the instance does not use a namespace, the group ID must be globally unique across instances and regions. The group ID cannot be the same as an existing group ID or topic name in ApsaraMQ for RocketMQ instances that belong to your Alibaba Cloud account.
	//
	// >
	//
	// 	- After the consumer group is created, the group ID cannot be changed.
	//
	// 	- To check whether an instance uses a namespace, log on to the ApsaraMQ for RocketMQ console, go to the **Instance Details*	- page, and then view the value of the Namespace field in the **Basic Information*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The protocol over which clients in the consumer group communicate with the ApsaraMQ for RocketMQ broker. All clients in a consumer group communicate with the ApsaraMQ for RocketMQ broker over the same protocol. You must create different groups for TCP clients and HTTP clients. Valid values:
	//
	// 	- **tcp**: Clients in the consumer group consume messages over TCP. This is the default value.
	//
	// 	- **http**: Clients in the consumer group consume messages over HTTP.
	//
	// example:
	//
	// tcp
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the instance in which you want to create the consumer group.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsGroupCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupCreateRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *OnsGroupCreateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupCreateRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnsGroupCreateRequest) SetGroupId(v string) *OnsGroupCreateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupCreateRequest) SetGroupType(v string) *OnsGroupCreateRequest {
	s.GroupType = &v
	return s
}

func (s *OnsGroupCreateRequest) SetInstanceId(v string) *OnsGroupCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupCreateRequest) SetRemark(v string) *OnsGroupCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsGroupCreateRequest) Validate() error {
	return dara.Validate(s)
}
