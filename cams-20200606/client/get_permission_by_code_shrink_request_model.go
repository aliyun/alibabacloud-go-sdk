// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionByCodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPermissionByCodeShrinkRequest
	GetCode() *string
	SetCustSpaceId(v string) *GetPermissionByCodeShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetPermissionByCodeShrinkRequest
	GetOwnerId() *int64
	SetPermissionsShrink(v string) *GetPermissionByCodeShrinkRequest
	GetPermissionsShrink() *string
	SetResourceOwnerAccount(v string) *GetPermissionByCodeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPermissionByCodeShrinkRequest
	GetResourceOwnerId() *int64
}

type GetPermissionByCodeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PermissionsShrink    *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPermissionByCodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionByCodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *GetPermissionByCodeShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetPermissionByCodeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPermissionByCodeShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *GetPermissionByCodeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPermissionByCodeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPermissionByCodeShrinkRequest) SetCode(v string) *GetPermissionByCodeShrinkRequest {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetCustSpaceId(v string) *GetPermissionByCodeShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetOwnerId(v int64) *GetPermissionByCodeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetPermissionsShrink(v string) *GetPermissionByCodeShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetResourceOwnerAccount(v string) *GetPermissionByCodeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) SetResourceOwnerId(v int64) *GetPermissionByCodeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPermissionByCodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
