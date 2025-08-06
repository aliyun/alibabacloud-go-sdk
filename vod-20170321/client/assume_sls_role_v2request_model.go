// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *AssumeSlsRoleV2Request
	GetAppVersion() *string
	SetAuthInfo(v string) *AssumeSlsRoleV2Request
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *AssumeSlsRoleV2Request
	GetAuthTimestamp() *int64
	SetBusinessType(v string) *AssumeSlsRoleV2Request
	GetBusinessType() *string
	SetClientId(v string) *AssumeSlsRoleV2Request
	GetClientId() *string
	SetDeviceModel(v string) *AssumeSlsRoleV2Request
	GetDeviceModel() *string
	SetOwnerId(v int64) *AssumeSlsRoleV2Request
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssumeSlsRoleV2Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssumeSlsRoleV2Request
	GetResourceOwnerId() *int64
	SetTerminalType(v string) *AssumeSlsRoleV2Request
	GetTerminalType() *string
}

type AssumeSlsRoleV2Request struct {
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

func (s AssumeSlsRoleV2Request) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleV2Request) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleV2Request) GetAppVersion() *string {
	return s.AppVersion
}

func (s *AssumeSlsRoleV2Request) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *AssumeSlsRoleV2Request) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *AssumeSlsRoleV2Request) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AssumeSlsRoleV2Request) GetClientId() *string {
	return s.ClientId
}

func (s *AssumeSlsRoleV2Request) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *AssumeSlsRoleV2Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssumeSlsRoleV2Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssumeSlsRoleV2Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssumeSlsRoleV2Request) GetTerminalType() *string {
	return s.TerminalType
}

func (s *AssumeSlsRoleV2Request) SetAppVersion(v string) *AssumeSlsRoleV2Request {
	s.AppVersion = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetAuthInfo(v string) *AssumeSlsRoleV2Request {
	s.AuthInfo = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetAuthTimestamp(v int64) *AssumeSlsRoleV2Request {
	s.AuthTimestamp = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetBusinessType(v string) *AssumeSlsRoleV2Request {
	s.BusinessType = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetClientId(v string) *AssumeSlsRoleV2Request {
	s.ClientId = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetDeviceModel(v string) *AssumeSlsRoleV2Request {
	s.DeviceModel = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetOwnerId(v int64) *AssumeSlsRoleV2Request {
	s.OwnerId = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetResourceOwnerAccount(v string) *AssumeSlsRoleV2Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetResourceOwnerId(v int64) *AssumeSlsRoleV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssumeSlsRoleV2Request) SetTerminalType(v string) *AssumeSlsRoleV2Request {
	s.TerminalType = &v
	return s
}

func (s *AssumeSlsRoleV2Request) Validate() error {
	return dara.Validate(s)
}
