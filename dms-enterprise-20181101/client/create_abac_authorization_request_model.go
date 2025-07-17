// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityType(v string) *CreateAbacAuthorizationRequest
	GetIdentityType() *string
	SetPolicyId(v int64) *CreateAbacAuthorizationRequest
	GetPolicyId() *int64
	SetRoleId(v int64) *CreateAbacAuthorizationRequest
	GetRoleId() *int64
	SetTid(v int64) *CreateAbacAuthorizationRequest
	GetTid() *int64
	SetUserId(v int64) *CreateAbacAuthorizationRequest
	GetUserId() *int64
}

type CreateAbacAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 31****
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 51****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAbacAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *CreateAbacAuthorizationRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *CreateAbacAuthorizationRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *CreateAbacAuthorizationRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *CreateAbacAuthorizationRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateAbacAuthorizationRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateAbacAuthorizationRequest) SetIdentityType(v string) *CreateAbacAuthorizationRequest {
	s.IdentityType = &v
	return s
}

func (s *CreateAbacAuthorizationRequest) SetPolicyId(v int64) *CreateAbacAuthorizationRequest {
	s.PolicyId = &v
	return s
}

func (s *CreateAbacAuthorizationRequest) SetRoleId(v int64) *CreateAbacAuthorizationRequest {
	s.RoleId = &v
	return s
}

func (s *CreateAbacAuthorizationRequest) SetTid(v int64) *CreateAbacAuthorizationRequest {
	s.Tid = &v
	return s
}

func (s *CreateAbacAuthorizationRequest) SetUserId(v int64) *CreateAbacAuthorizationRequest {
	s.UserId = &v
	return s
}

func (s *CreateAbacAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
