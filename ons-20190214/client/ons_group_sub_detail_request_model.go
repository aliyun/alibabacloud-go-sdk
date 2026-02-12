// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupSubDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsGroupSubDetailRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsGroupSubDetailRequest
	GetInstanceId() *string
}

type OnsGroupSubDetailRequest struct {
	// The ID of the consumer group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsGroupSubDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupSubDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupSubDetailRequest) SetGroupId(v string) *OnsGroupSubDetailRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupSubDetailRequest) SetInstanceId(v string) *OnsGroupSubDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupSubDetailRequest) Validate() error {
	return dara.Validate(s)
}
