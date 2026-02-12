// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsGroupDeleteRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsGroupDeleteRequest
	GetInstanceId() *string
}

type OnsGroupDeleteRequest struct {
	// The ID of the consumer group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance to which the specified consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsGroupDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupDeleteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupDeleteRequest) SetGroupId(v string) *OnsGroupDeleteRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupDeleteRequest) SetInstanceId(v string) *OnsGroupDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupDeleteRequest) Validate() error {
	return dara.Validate(s)
}
