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
	// Principal Type. Valid values:**user**or**custom role**.
	//
	// Valid values:
	//
	// 	- USER
	//
	// 	- ROLE
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the role.
	//
	// > If IdentityType is set to ROLE, this parameter is required.
	//
	// example:
	//
	// 31****
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The ID of the user. You can call the [GetUser](https://help.aliyun.com/document_detail/465816.html) operation to query the user ID.
	//
	// > If IdentityType is set to USER, this parameter is required.
	//
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
