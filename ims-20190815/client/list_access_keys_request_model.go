// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *ListAccessKeysRequest
	GetUserPrincipalName() *string
}

type ListAccessKeysRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListAccessKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListAccessKeysRequest) SetUserPrincipalName(v string) *ListAccessKeysRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *ListAccessKeysRequest) Validate() error {
	return dara.Validate(s)
}
