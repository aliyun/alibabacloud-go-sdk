// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PushReportRequest
	GetAppId() *string
	SetAppVersion(v string) *PushReportRequest
	GetAppVersion() *string
	SetChannel(v string) *PushReportRequest
	GetChannel() *string
	SetConnectType(v string) *PushReportRequest
	GetConnectType() *string
	SetDeliveryToken(v string) *PushReportRequest
	GetDeliveryToken() *string
	SetImei(v string) *PushReportRequest
	GetImei() *string
	SetImsi(v string) *PushReportRequest
	GetImsi() *string
	SetModel(v string) *PushReportRequest
	GetModel() *string
	SetOsType(v int32) *PushReportRequest
	GetOsType() *int32
	SetPushVersion(v string) *PushReportRequest
	GetPushVersion() *string
	SetTenantId(v string) *PushReportRequest
	GetTenantId() *string
	SetThirdChannel(v int32) *PushReportRequest
	GetThirdChannel() *int32
	SetThirdChannelDeviceToken(v string) *PushReportRequest
	GetThirdChannelDeviceToken() *string
	SetWorkspaceId(v string) *PushReportRequest
	GetWorkspaceId() *string
}

type PushReportRequest struct {
	// This parameter is required.
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion  *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// This parameter is required.
	DeliveryToken *string `json:"DeliveryToken,omitempty" xml:"DeliveryToken,omitempty"`
	Imei          *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	Imsi          *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	OsType                  *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PushVersion             *string `json:"PushVersion,omitempty" xml:"PushVersion,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannel            *int32  `json:"ThirdChannel,omitempty" xml:"ThirdChannel,omitempty"`
	ThirdChannelDeviceToken *string `json:"ThirdChannelDeviceToken,omitempty" xml:"ThirdChannelDeviceToken,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushReportRequest) String() string {
	return dara.Prettify(s)
}

func (s PushReportRequest) GoString() string {
	return s.String()
}

func (s *PushReportRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushReportRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *PushReportRequest) GetChannel() *string {
	return s.Channel
}

func (s *PushReportRequest) GetConnectType() *string {
	return s.ConnectType
}

func (s *PushReportRequest) GetDeliveryToken() *string {
	return s.DeliveryToken
}

func (s *PushReportRequest) GetImei() *string {
	return s.Imei
}

func (s *PushReportRequest) GetImsi() *string {
	return s.Imsi
}

func (s *PushReportRequest) GetModel() *string {
	return s.Model
}

func (s *PushReportRequest) GetOsType() *int32 {
	return s.OsType
}

func (s *PushReportRequest) GetPushVersion() *string {
	return s.PushVersion
}

func (s *PushReportRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushReportRequest) GetThirdChannel() *int32 {
	return s.ThirdChannel
}

func (s *PushReportRequest) GetThirdChannelDeviceToken() *string {
	return s.ThirdChannelDeviceToken
}

func (s *PushReportRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushReportRequest) SetAppId(v string) *PushReportRequest {
	s.AppId = &v
	return s
}

func (s *PushReportRequest) SetAppVersion(v string) *PushReportRequest {
	s.AppVersion = &v
	return s
}

func (s *PushReportRequest) SetChannel(v string) *PushReportRequest {
	s.Channel = &v
	return s
}

func (s *PushReportRequest) SetConnectType(v string) *PushReportRequest {
	s.ConnectType = &v
	return s
}

func (s *PushReportRequest) SetDeliveryToken(v string) *PushReportRequest {
	s.DeliveryToken = &v
	return s
}

func (s *PushReportRequest) SetImei(v string) *PushReportRequest {
	s.Imei = &v
	return s
}

func (s *PushReportRequest) SetImsi(v string) *PushReportRequest {
	s.Imsi = &v
	return s
}

func (s *PushReportRequest) SetModel(v string) *PushReportRequest {
	s.Model = &v
	return s
}

func (s *PushReportRequest) SetOsType(v int32) *PushReportRequest {
	s.OsType = &v
	return s
}

func (s *PushReportRequest) SetPushVersion(v string) *PushReportRequest {
	s.PushVersion = &v
	return s
}

func (s *PushReportRequest) SetTenantId(v string) *PushReportRequest {
	s.TenantId = &v
	return s
}

func (s *PushReportRequest) SetThirdChannel(v int32) *PushReportRequest {
	s.ThirdChannel = &v
	return s
}

func (s *PushReportRequest) SetThirdChannelDeviceToken(v string) *PushReportRequest {
	s.ThirdChannelDeviceToken = &v
	return s
}

func (s *PushReportRequest) SetWorkspaceId(v string) *PushReportRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushReportRequest) Validate() error {
	return dara.Validate(s)
}
