// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionByCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPermissionByCodeRequest
	GetCode() *string
	SetCustSpaceId(v string) *GetPermissionByCodeRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetPermissionByCodeRequest
	GetOwnerId() *int64
	SetPermissions(v []*string) *GetPermissionByCodeRequest
	GetPermissions() []*string
	SetResourceOwnerAccount(v string) *GetPermissionByCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPermissionByCodeRequest
	GetResourceOwnerId() *int64
}

type GetPermissionByCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	CustSpaceId          *string   `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Permissions          []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPermissionByCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeRequest) GetCode() *string {
	return s.Code
}

func (s *GetPermissionByCodeRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetPermissionByCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPermissionByCodeRequest) GetPermissions() []*string {
	return s.Permissions
}

func (s *GetPermissionByCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPermissionByCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPermissionByCodeRequest) SetCode(v string) *GetPermissionByCodeRequest {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetCustSpaceId(v string) *GetPermissionByCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetOwnerId(v int64) *GetPermissionByCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetPermissions(v []*string) *GetPermissionByCodeRequest {
	s.Permissions = v
	return s
}

func (s *GetPermissionByCodeRequest) SetResourceOwnerAccount(v string) *GetPermissionByCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPermissionByCodeRequest) SetResourceOwnerId(v int64) *GetPermissionByCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPermissionByCodeRequest) Validate() error {
	return dara.Validate(s)
}
