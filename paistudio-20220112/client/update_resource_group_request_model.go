// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateResourceGroupRequest
	GetDescription() *string
	SetName(v string) *UpdateResourceGroupRequest
	GetName() *string
	SetUnbind(v bool) *UpdateResourceGroupRequest
	GetUnbind() *bool
	SetUserVpc(v *UserVpc) *UpdateResourceGroupRequest
	GetUserVpc() *UserVpc
}

type UpdateResourceGroupRequest struct {
	// example:
	//
	// test_new_havpn_tf
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// prophet
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	Unbind  *bool    `json:"Unbind,omitempty" xml:"Unbind,omitempty"`
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateResourceGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResourceGroupRequest) GetUnbind() *bool {
	return s.Unbind
}

func (s *UpdateResourceGroupRequest) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *UpdateResourceGroupRequest) SetDescription(v string) *UpdateResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetName(v string) *UpdateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetUnbind(v bool) *UpdateResourceGroupRequest {
	s.Unbind = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetUserVpc(v *UserVpc) *UpdateResourceGroupRequest {
	s.UserVpc = v
	return s
}

func (s *UpdateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
