// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserPoolRequest
	GetDescription() *string
	SetUserPoolName(v string) *CreateUserPoolRequest
	GetUserPoolName() *string
}

type CreateUserPoolRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateUserPoolRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserPoolRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolRequest) SetDescription(v string) *CreateUserPoolRequest {
	s.Description = &v
	return s
}

func (s *CreateUserPoolRequest) SetUserPoolName(v string) *CreateUserPoolRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolRequest) Validate() error {
	return dara.Validate(s)
}
