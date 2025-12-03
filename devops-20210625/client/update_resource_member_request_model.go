// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *UpdateResourceMemberRequest
	GetRoleName() *string
}

type UpdateResourceMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// admin
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s UpdateResourceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateResourceMemberRequest) SetRoleName(v string) *UpdateResourceMemberRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateResourceMemberRequest) Validate() error {
	return dara.Validate(s)
}
