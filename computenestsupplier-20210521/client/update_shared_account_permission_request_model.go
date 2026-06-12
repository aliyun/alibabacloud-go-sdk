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
	// A client token to ensure the idempotence of the request. Generate a unique value for this parameter on your client. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The permission type. Valid values:
	//
	// - Deployable: The service can be deployed.
	//
	// - Accessible: The service can be accessed.
	//
	// - AccessibleIncludeBeta: All versions of the service, including beta versions, can be accessed.
	//
	// - DeployableIncludeBeta: All versions of the service, including beta versions, can be deployed.
	//
	// - Authorized: The service is authorized. This value is used for distribution scenarios.
	//
	// - Unauthorized: The service is not authorized. This value is used for distribution scenarios.
	//
	// This parameter is required.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0d6e1d846e4c4axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The type of service sharing. The default value is SharedAccount. Valid values:
	//
	// - SharedAccount: Regular sharing.
	//
	// - Reseller: Distribution sharing.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The whitelisted account with which the service is shared.
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
