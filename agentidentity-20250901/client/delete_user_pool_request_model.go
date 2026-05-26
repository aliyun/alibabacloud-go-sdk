// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *DeleteUserPoolRequest
	GetUserPoolName() *string
}

type DeleteUserPoolRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteUserPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPoolRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteUserPoolRequest) SetUserPoolName(v string) *DeleteUserPoolRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteUserPoolRequest) Validate() error {
	return dara.Validate(s)
}
