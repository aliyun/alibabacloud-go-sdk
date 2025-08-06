// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *GetUploadProgressRequest
	GetAppVersion() *string
	SetAuthInfo(v string) *GetUploadProgressRequest
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *GetUploadProgressRequest
	GetAuthTimestamp() *int64
	SetBusinessType(v string) *GetUploadProgressRequest
	GetBusinessType() *string
	SetClientId(v string) *GetUploadProgressRequest
	GetClientId() *string
	SetDeviceModel(v string) *GetUploadProgressRequest
	GetDeviceModel() *string
	SetOwnerId(v int64) *GetUploadProgressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetUploadProgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetUploadProgressRequest
	GetResourceOwnerId() *int64
	SetSource(v string) *GetUploadProgressRequest
	GetSource() *string
	SetTerminalType(v string) *GetUploadProgressRequest
	GetTerminalType() *string
	SetUploadAddress(v string) *GetUploadProgressRequest
	GetUploadAddress() *string
	SetUploadInfoList(v string) *GetUploadProgressRequest
	GetUploadInfoList() *string
	SetUserId(v int64) *GetUploadProgressRequest
	GetUserId() *int64
}

type GetUploadProgressRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	AuthInfo *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	// This parameter is required.
	AuthTimestamp        *int64  `json:"AuthTimestamp,omitempty" xml:"AuthTimestamp,omitempty"`
	BusinessType         *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ClientId             *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DeviceModel          *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Source               *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TerminalType         *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	UploadAddress        *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// This parameter is required.
	UploadInfoList *string `json:"UploadInfoList,omitempty" xml:"UploadInfoList,omitempty"`
	UserId         *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUploadProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadProgressRequest) GoString() string {
	return s.String()
}

func (s *GetUploadProgressRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetUploadProgressRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *GetUploadProgressRequest) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *GetUploadProgressRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUploadProgressRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetUploadProgressRequest) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetUploadProgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetUploadProgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetUploadProgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetUploadProgressRequest) GetSource() *string {
	return s.Source
}

func (s *GetUploadProgressRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *GetUploadProgressRequest) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *GetUploadProgressRequest) GetUploadInfoList() *string {
	return s.UploadInfoList
}

func (s *GetUploadProgressRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *GetUploadProgressRequest) SetAppVersion(v string) *GetUploadProgressRequest {
	s.AppVersion = &v
	return s
}

func (s *GetUploadProgressRequest) SetAuthInfo(v string) *GetUploadProgressRequest {
	s.AuthInfo = &v
	return s
}

func (s *GetUploadProgressRequest) SetAuthTimestamp(v int64) *GetUploadProgressRequest {
	s.AuthTimestamp = &v
	return s
}

func (s *GetUploadProgressRequest) SetBusinessType(v string) *GetUploadProgressRequest {
	s.BusinessType = &v
	return s
}

func (s *GetUploadProgressRequest) SetClientId(v string) *GetUploadProgressRequest {
	s.ClientId = &v
	return s
}

func (s *GetUploadProgressRequest) SetDeviceModel(v string) *GetUploadProgressRequest {
	s.DeviceModel = &v
	return s
}

func (s *GetUploadProgressRequest) SetOwnerId(v int64) *GetUploadProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUploadProgressRequest) SetResourceOwnerAccount(v string) *GetUploadProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUploadProgressRequest) SetResourceOwnerId(v int64) *GetUploadProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUploadProgressRequest) SetSource(v string) *GetUploadProgressRequest {
	s.Source = &v
	return s
}

func (s *GetUploadProgressRequest) SetTerminalType(v string) *GetUploadProgressRequest {
	s.TerminalType = &v
	return s
}

func (s *GetUploadProgressRequest) SetUploadAddress(v string) *GetUploadProgressRequest {
	s.UploadAddress = &v
	return s
}

func (s *GetUploadProgressRequest) SetUploadInfoList(v string) *GetUploadProgressRequest {
	s.UploadInfoList = &v
	return s
}

func (s *GetUploadProgressRequest) SetUserId(v int64) *GetUploadProgressRequest {
	s.UserId = &v
	return s
}

func (s *GetUploadProgressRequest) Validate() error {
	return dara.Validate(s)
}
