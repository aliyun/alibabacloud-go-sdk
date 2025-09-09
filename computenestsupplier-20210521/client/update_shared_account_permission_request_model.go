// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSharedAccountPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSharedAccountPermissionRequest
	GetClientToken() *string
	SetPermission(v string) *UpdateSharedAccountPermissionRequest
	GetPermission() *string
	SetRegionId(v string) *UpdateSharedAccountPermissionRequest
	GetRegionId() *string
	SetServiceId(v string) *UpdateSharedAccountPermissionRequest
	GetServiceId() *string
	SetType(v string) *UpdateSharedAccountPermissionRequest
	GetType() *string
	SetUserAliUid(v int64) *UpdateSharedAccountPermissionRequest
	GetUserAliUid() *int64
}

type UpdateSharedAccountPermissionRequest struct {
	// Client token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Permission type. Possible values:
	//
	// - Deployable: Can be deployed.
	//
	// - Accessible: Can be accessed.
	//
	// - AccessibleIncludeBeta: Can access all versions, including Beta versions.
	//
	// - DeployableIncludeBeta: Can deploy all versions, including Beta versions.
	//
	// - Authorized: Authorized (for reselling scenarios)
	//
	// - Unauthorized: Unauthorized (for reselling scenarios)
	//
	// This parameter is required.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0d6e1d846e4c4axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service sharing type, with a default value of SharedAccount. Available options:
	//
	// - SharedAccount: Regular sharing type.
	//
	// - Reseller: Reselling sharing type.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Whitelist account for service sharing.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s UpdateSharedAccountPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSharedAccountPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSharedAccountPermissionRequest) GetPermission() *string {
	return s.Permission
}

func (s *UpdateSharedAccountPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSharedAccountPermissionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateSharedAccountPermissionRequest) GetType() *string {
	return s.Type
}

func (s *UpdateSharedAccountPermissionRequest) GetUserAliUid() *int64 {
	return s.UserAliUid
}

func (s *UpdateSharedAccountPermissionRequest) SetClientToken(v string) *UpdateSharedAccountPermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetPermission(v string) *UpdateSharedAccountPermissionRequest {
	s.Permission = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetRegionId(v string) *UpdateSharedAccountPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetServiceId(v string) *UpdateSharedAccountPermissionRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetType(v string) *UpdateSharedAccountPermissionRequest {
	s.Type = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetUserAliUid(v int64) *UpdateSharedAccountPermissionRequest {
	s.UserAliUid = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) Validate() error {
	return dara.Validate(s)
}
