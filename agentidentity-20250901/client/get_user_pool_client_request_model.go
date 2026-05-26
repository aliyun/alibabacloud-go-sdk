// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientName(v string) *GetUserPoolClientRequest
	GetClientName() *string
	SetUserPoolName(v string) *GetUserPoolClientRequest
	GetUserPoolName() *string
}

type GetUserPoolClientRequest struct {
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetUserPoolClientRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolClientRequest) GoString() string {
	return s.String()
}

func (s *GetUserPoolClientRequest) GetClientName() *string {
	return s.ClientName
}

func (s *GetUserPoolClientRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolClientRequest) SetClientName(v string) *GetUserPoolClientRequest {
	s.ClientName = &v
	return s
}

func (s *GetUserPoolClientRequest) SetUserPoolName(v string) *GetUserPoolClientRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolClientRequest) Validate() error {
	return dara.Validate(s)
}
