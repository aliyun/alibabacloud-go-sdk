// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *GetUserPoolRequest
	GetUserPoolName() *string
}

type GetUserPoolRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetUserPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolRequest) GoString() string {
	return s.String()
}

func (s *GetUserPoolRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolRequest) SetUserPoolName(v string) *GetUserPoolRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolRequest) Validate() error {
	return dara.Validate(s)
}
