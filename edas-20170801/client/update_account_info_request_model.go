// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *UpdateAccountInfoRequest
	GetEmail() *string
	SetName(v string) *UpdateAccountInfoRequest
	GetName() *string
	SetTelephone(v string) *UpdateAccountInfoRequest
	GetTelephone() *string
}

type UpdateAccountInfoRequest struct {
	// The email address of the account.
	//
	// example:
	//
	// 1321234****@alibaba-inc.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// edas-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The contact information of the account.
	//
	// example:
	//
	// 1321234****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
}

func (s UpdateAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateAccountInfoRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAccountInfoRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *UpdateAccountInfoRequest) SetEmail(v string) *UpdateAccountInfoRequest {
	s.Email = &v
	return s
}

func (s *UpdateAccountInfoRequest) SetName(v string) *UpdateAccountInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateAccountInfoRequest) SetTelephone(v string) *UpdateAccountInfoRequest {
	s.Telephone = &v
	return s
}

func (s *UpdateAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
