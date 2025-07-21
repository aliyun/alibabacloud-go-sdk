// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUuidValidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBluetooth(v string) *CheckUuidValidRequest
	GetBluetooth() *string
	SetBuildId(v string) *CheckUuidValidRequest
	GetBuildId() *string
	SetChipId(v string) *CheckUuidValidRequest
	GetChipId() *string
	SetClientId(v string) *CheckUuidValidRequest
	GetClientId() *string
	SetClientVersion(v string) *CheckUuidValidRequest
	GetClientVersion() *string
	SetCustomId(v string) *CheckUuidValidRequest
	GetCustomId() *string
	SetEtherMac(v string) *CheckUuidValidRequest
	GetEtherMac() *string
	SetHostOsInfo(v string) *CheckUuidValidRequest
	GetHostOsInfo() *string
	SetLoginRegionId(v string) *CheckUuidValidRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *CheckUuidValidRequest
	GetLoginToken() *string
	SetSerialNo(v string) *CheckUuidValidRequest
	GetSerialNo() *string
	SetSessionId(v string) *CheckUuidValidRequest
	GetSessionId() *string
	SetUuid(v string) *CheckUuidValidRequest
	GetUuid() *string
	SetWlan(v string) *CheckUuidValidRequest
	GetWlan() *string
	SetWosAppVersion(v string) *CheckUuidValidRequest
	GetWosAppVersion() *string
}

type CheckUuidValidRequest struct {
	Bluetooth *string `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId   *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// This parameter is required.
	ChipId        *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	CustomId      *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	EtherMac      *string `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	HostOsInfo    *string `json:"HostOsInfo,omitempty" xml:"HostOsInfo,omitempty"`
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	SerialNo  *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Wlan          *string `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
	WosAppVersion *string `json:"WosAppVersion,omitempty" xml:"WosAppVersion,omitempty"`
}

func (s CheckUuidValidRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUuidValidRequest) GoString() string {
	return s.String()
}

func (s *CheckUuidValidRequest) GetBluetooth() *string {
	return s.Bluetooth
}

func (s *CheckUuidValidRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *CheckUuidValidRequest) GetChipId() *string {
	return s.ChipId
}

func (s *CheckUuidValidRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CheckUuidValidRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *CheckUuidValidRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *CheckUuidValidRequest) GetEtherMac() *string {
	return s.EtherMac
}

func (s *CheckUuidValidRequest) GetHostOsInfo() *string {
	return s.HostOsInfo
}

func (s *CheckUuidValidRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *CheckUuidValidRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *CheckUuidValidRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *CheckUuidValidRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CheckUuidValidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CheckUuidValidRequest) GetWlan() *string {
	return s.Wlan
}

func (s *CheckUuidValidRequest) GetWosAppVersion() *string {
	return s.WosAppVersion
}

func (s *CheckUuidValidRequest) SetBluetooth(v string) *CheckUuidValidRequest {
	s.Bluetooth = &v
	return s
}

func (s *CheckUuidValidRequest) SetBuildId(v string) *CheckUuidValidRequest {
	s.BuildId = &v
	return s
}

func (s *CheckUuidValidRequest) SetChipId(v string) *CheckUuidValidRequest {
	s.ChipId = &v
	return s
}

func (s *CheckUuidValidRequest) SetClientId(v string) *CheckUuidValidRequest {
	s.ClientId = &v
	return s
}

func (s *CheckUuidValidRequest) SetClientVersion(v string) *CheckUuidValidRequest {
	s.ClientVersion = &v
	return s
}

func (s *CheckUuidValidRequest) SetCustomId(v string) *CheckUuidValidRequest {
	s.CustomId = &v
	return s
}

func (s *CheckUuidValidRequest) SetEtherMac(v string) *CheckUuidValidRequest {
	s.EtherMac = &v
	return s
}

func (s *CheckUuidValidRequest) SetHostOsInfo(v string) *CheckUuidValidRequest {
	s.HostOsInfo = &v
	return s
}

func (s *CheckUuidValidRequest) SetLoginRegionId(v string) *CheckUuidValidRequest {
	s.LoginRegionId = &v
	return s
}

func (s *CheckUuidValidRequest) SetLoginToken(v string) *CheckUuidValidRequest {
	s.LoginToken = &v
	return s
}

func (s *CheckUuidValidRequest) SetSerialNo(v string) *CheckUuidValidRequest {
	s.SerialNo = &v
	return s
}

func (s *CheckUuidValidRequest) SetSessionId(v string) *CheckUuidValidRequest {
	s.SessionId = &v
	return s
}

func (s *CheckUuidValidRequest) SetUuid(v string) *CheckUuidValidRequest {
	s.Uuid = &v
	return s
}

func (s *CheckUuidValidRequest) SetWlan(v string) *CheckUuidValidRequest {
	s.Wlan = &v
	return s
}

func (s *CheckUuidValidRequest) SetWosAppVersion(v string) *CheckUuidValidRequest {
	s.WosAppVersion = &v
	return s
}

func (s *CheckUuidValidRequest) Validate() error {
	return dara.Validate(s)
}
