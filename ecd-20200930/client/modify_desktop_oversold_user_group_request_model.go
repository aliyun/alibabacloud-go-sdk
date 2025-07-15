// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ModifyDesktopOversoldUserGroupRequest
	GetImageId() *string
	SetName(v string) *ModifyDesktopOversoldUserGroupRequest
	GetName() *string
	SetOversoldGroupId(v string) *ModifyDesktopOversoldUserGroupRequest
	GetOversoldGroupId() *string
	SetPolicyGroupId(v string) *ModifyDesktopOversoldUserGroupRequest
	GetPolicyGroupId() *string
	SetUserGroupId(v string) *ModifyDesktopOversoldUserGroupRequest
	GetUserGroupId() *string
}

type ModifyDesktopOversoldUserGroupRequest struct {
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	PolicyGroupId   *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ModifyDesktopOversoldUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldUserGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyDesktopOversoldUserGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDesktopOversoldUserGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *ModifyDesktopOversoldUserGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyDesktopOversoldUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ModifyDesktopOversoldUserGroupRequest) SetImageId(v string) *ModifyDesktopOversoldUserGroupRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupRequest) SetName(v string) *ModifyDesktopOversoldUserGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupRequest) SetOversoldGroupId(v string) *ModifyDesktopOversoldUserGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopOversoldUserGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupRequest) SetUserGroupId(v string) *ModifyDesktopOversoldUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
