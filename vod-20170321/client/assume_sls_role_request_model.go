// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *AssumeSlsRoleRequest
	GetAppVersion() *string
	SetAuthInfo(v string) *AssumeSlsRoleRequest
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *AssumeSlsRoleRequest
	GetAuthTimestamp() *int64
	SetBusinessType(v string) *AssumeSlsRoleRequest
	GetBusinessType() *string
	SetClientId(v string) *AssumeSlsRoleRequest
	GetClientId() *string
	SetDeviceModel(v string) *AssumeSlsRoleRequest
	GetDeviceModel() *string
	SetOwnerId(v int64) *AssumeSlsRoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssumeSlsRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssumeSlsRoleRequest
	GetResourceOwnerId() *int64
	SetTerminalType(v string) *AssumeSlsRoleRequest
	GetTerminalType() *string
}

type AssumeSlsRoleRequest struct {
	// This parameter is required.
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	AuthInfo *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	// This parameter is required.
	AuthTimestamp *int64 `json:"AuthTimestamp,omitempty" xml:"AuthTimestamp,omitempty"`
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	DeviceModel          *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s AssumeSlsRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *AssumeSlsRoleRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *AssumeSlsRoleRequest) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *AssumeSlsRoleRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AssumeSlsRoleRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AssumeSlsRoleRequest) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *AssumeSlsRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssumeSlsRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssumeSlsRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssumeSlsRoleRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *AssumeSlsRoleRequest) SetAppVersion(v string) *AssumeSlsRoleRequest {
	s.AppVersion = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetAuthInfo(v string) *AssumeSlsRoleRequest {
	s.AuthInfo = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetAuthTimestamp(v int64) *AssumeSlsRoleRequest {
	s.AuthTimestamp = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetBusinessType(v string) *AssumeSlsRoleRequest {
	s.BusinessType = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetClientId(v string) *AssumeSlsRoleRequest {
	s.ClientId = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetDeviceModel(v string) *AssumeSlsRoleRequest {
	s.DeviceModel = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetOwnerId(v int64) *AssumeSlsRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetResourceOwnerAccount(v string) *AssumeSlsRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetResourceOwnerId(v int64) *AssumeSlsRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssumeSlsRoleRequest) SetTerminalType(v string) *AssumeSlsRoleRequest {
	s.TerminalType = &v
	return s
}

func (s *AssumeSlsRoleRequest) Validate() error {
	return dara.Validate(s)
}
