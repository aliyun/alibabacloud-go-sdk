// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeOssRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *AssumeOssRoleRequest
	GetAppVersion() *string
	SetAuthInfo(v string) *AssumeOssRoleRequest
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *AssumeOssRoleRequest
	GetAuthTimestamp() *int64
	SetBusinessType(v string) *AssumeOssRoleRequest
	GetBusinessType() *string
	SetClientId(v string) *AssumeOssRoleRequest
	GetClientId() *string
	SetDeviceModel(v string) *AssumeOssRoleRequest
	GetDeviceModel() *string
	SetOwnerId(v int64) *AssumeOssRoleRequest
	GetOwnerId() *int64
	SetPrefix(v string) *AssumeOssRoleRequest
	GetPrefix() *string
	SetResourceOwnerAccount(v string) *AssumeOssRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssumeOssRoleRequest
	GetResourceOwnerId() *int64
	SetSource(v string) *AssumeOssRoleRequest
	GetSource() *string
	SetTerminalType(v string) *AssumeOssRoleRequest
	GetTerminalType() *string
}

type AssumeOssRoleRequest struct {
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
	Prefix               *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s AssumeOssRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeOssRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeOssRoleRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *AssumeOssRoleRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *AssumeOssRoleRequest) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *AssumeOssRoleRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AssumeOssRoleRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AssumeOssRoleRequest) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *AssumeOssRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssumeOssRoleRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *AssumeOssRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssumeOssRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssumeOssRoleRequest) GetSource() *string {
	return s.Source
}

func (s *AssumeOssRoleRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *AssumeOssRoleRequest) SetAppVersion(v string) *AssumeOssRoleRequest {
	s.AppVersion = &v
	return s
}

func (s *AssumeOssRoleRequest) SetAuthInfo(v string) *AssumeOssRoleRequest {
	s.AuthInfo = &v
	return s
}

func (s *AssumeOssRoleRequest) SetAuthTimestamp(v int64) *AssumeOssRoleRequest {
	s.AuthTimestamp = &v
	return s
}

func (s *AssumeOssRoleRequest) SetBusinessType(v string) *AssumeOssRoleRequest {
	s.BusinessType = &v
	return s
}

func (s *AssumeOssRoleRequest) SetClientId(v string) *AssumeOssRoleRequest {
	s.ClientId = &v
	return s
}

func (s *AssumeOssRoleRequest) SetDeviceModel(v string) *AssumeOssRoleRequest {
	s.DeviceModel = &v
	return s
}

func (s *AssumeOssRoleRequest) SetOwnerId(v int64) *AssumeOssRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AssumeOssRoleRequest) SetPrefix(v string) *AssumeOssRoleRequest {
	s.Prefix = &v
	return s
}

func (s *AssumeOssRoleRequest) SetResourceOwnerAccount(v string) *AssumeOssRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssumeOssRoleRequest) SetResourceOwnerId(v int64) *AssumeOssRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssumeOssRoleRequest) SetSource(v string) *AssumeOssRoleRequest {
	s.Source = &v
	return s
}

func (s *AssumeOssRoleRequest) SetTerminalType(v string) *AssumeOssRoleRequest {
	s.TerminalType = &v
	return s
}

func (s *AssumeOssRoleRequest) Validate() error {
	return dara.Validate(s)
}
