// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupConsumerUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsGroupConsumerUpdateRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsGroupConsumerUpdateRequest
	GetInstanceId() *string
	SetReadEnable(v bool) *OnsGroupConsumerUpdateRequest
	GetReadEnable() *bool
}

type OnsGroupConsumerUpdateRequest struct {
	// The ID of the consumer group for which you want to configure read permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to configure belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to authorize the consumer group to read messages. Valid values:
	//
	// 	- **true**: The consumer group can read messages.
	//
	// 	- **false**: The consumer group cannot read messages.
	//
	// Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReadEnable *bool `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
}

func (s OnsGroupConsumerUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupConsumerUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupConsumerUpdateRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupConsumerUpdateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupConsumerUpdateRequest) GetReadEnable() *bool {
	return s.ReadEnable
}

func (s *OnsGroupConsumerUpdateRequest) SetGroupId(v string) *OnsGroupConsumerUpdateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) SetInstanceId(v string) *OnsGroupConsumerUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) SetReadEnable(v bool) *OnsGroupConsumerUpdateRequest {
	s.ReadEnable = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) Validate() error {
	return dara.Validate(s)
}
