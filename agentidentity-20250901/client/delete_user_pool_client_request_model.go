// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPoolClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientName(v string) *DeleteUserPoolClientRequest
	GetClientName() *string
	SetUserPoolName(v string) *DeleteUserPoolClientRequest
	GetUserPoolName() *string
}

type DeleteUserPoolClientRequest struct {
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteUserPoolClientRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPoolClientRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPoolClientRequest) GetClientName() *string {
	return s.ClientName
}

func (s *DeleteUserPoolClientRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteUserPoolClientRequest) SetClientName(v string) *DeleteUserPoolClientRequest {
	s.ClientName = &v
	return s
}

func (s *DeleteUserPoolClientRequest) SetUserPoolName(v string) *DeleteUserPoolClientRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteUserPoolClientRequest) Validate() error {
	return dara.Validate(s)
}
