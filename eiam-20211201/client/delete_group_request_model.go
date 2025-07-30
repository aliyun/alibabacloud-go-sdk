// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *DeleteGroupRequest
	GetInstanceId() *string
}

type DeleteGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupRequest) SetInstanceId(v string) *DeleteGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
