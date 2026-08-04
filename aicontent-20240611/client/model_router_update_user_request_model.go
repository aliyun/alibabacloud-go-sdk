// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ModelRouterUpdateUserRequest
	GetName() *string
	SetPhone(v string) *ModelRouterUpdateUserRequest
	GetPhone() *string
}

type ModelRouterUpdateUserRequest struct {
	// The name of the user.
	//
	// example:
	//
	// John
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The phone number of the user.
	//
	// example:
	//
	// 13800000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ModelRouterUpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateUserRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateUserRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterUpdateUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *ModelRouterUpdateUserRequest) SetName(v string) *ModelRouterUpdateUserRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterUpdateUserRequest) SetPhone(v string) *ModelRouterUpdateUserRequest {
	s.Phone = &v
	return s
}

func (s *ModelRouterUpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
