// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *GetGroupRequest
	GetInstanceId() *string
}

type GetGroupRequest struct {
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

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupRequest) SetInstanceId(v string) *GetGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
