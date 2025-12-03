// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateResourceMemberRequest
	GetAccountId() *string
	SetRoleName(v string) *CreateResourceMemberRequest
	GetRoleName() *string
}

type CreateResourceMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1111111111111
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s CreateResourceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateResourceMemberRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateResourceMemberRequest) SetAccountId(v string) *CreateResourceMemberRequest {
	s.AccountId = &v
	return s
}

func (s *CreateResourceMemberRequest) SetRoleName(v string) *CreateResourceMemberRequest {
	s.RoleName = &v
	return s
}

func (s *CreateResourceMemberRequest) Validate() error {
	return dara.Validate(s)
}
