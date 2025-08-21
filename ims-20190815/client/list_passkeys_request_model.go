// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPasskeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *ListPasskeysRequest
	GetUserPrincipalName() *string
}

type ListPasskeysRequest struct {
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListPasskeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPasskeysRequest) GoString() string {
	return s.String()
}

func (s *ListPasskeysRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListPasskeysRequest) SetUserPrincipalName(v string) *ListPasskeysRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *ListPasskeysRequest) Validate() error {
	return dara.Validate(s)
}
