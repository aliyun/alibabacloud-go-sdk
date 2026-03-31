// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *ListPoliciesForUserRequest
	GetUserName() *string
}

type ListPoliciesForUserRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListPoliciesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListPoliciesForUserRequest) SetUserName(v string) *ListPoliciesForUserRequest {
	s.UserName = &v
	return s
}

func (s *ListPoliciesForUserRequest) Validate() error {
	return dara.Validate(s)
}
