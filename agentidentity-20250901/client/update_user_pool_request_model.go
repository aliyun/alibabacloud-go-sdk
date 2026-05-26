// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserPoolRequest
	GetDescription() *string
	SetUserPoolName(v string) *UpdateUserPoolRequest
	GetUserPoolName() *string
}

type UpdateUserPoolRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s UpdateUserPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserPoolRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateUserPoolRequest) SetDescription(v string) *UpdateUserPoolRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserPoolRequest) SetUserPoolName(v string) *UpdateUserPoolRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateUserPoolRequest) Validate() error {
	return dara.Validate(s)
}
