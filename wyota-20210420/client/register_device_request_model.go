// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBluetooth(v string) *RegisterDeviceRequest
	GetBluetooth() *string
	SetBuildId(v string) *RegisterDeviceRequest
	GetBuildId() *string
	SetChipId(v string) *RegisterDeviceRequest
	GetChipId() *string
	SetClientId(v string) *RegisterDeviceRequest
	GetClientId() *string
	SetClientType(v int32) *RegisterDeviceRequest
	GetClientType() *int32
	SetCpu(v string) *RegisterDeviceRequest
	GetCpu() *string
	SetCustomId(v string) *RegisterDeviceRequest
	GetCustomId() *string
	SetEtherMac(v string) *RegisterDeviceRequest
	GetEtherMac() *string
	SetMemory(v string) *RegisterDeviceRequest
	GetMemory() *string
	SetModel(v string) *RegisterDeviceRequest
	GetModel() *string
	SetSerialNo(v string) *RegisterDeviceRequest
	GetSerialNo() *string
	SetStorage(v string) *RegisterDeviceRequest
	GetStorage() *string
	SetToken(v string) *RegisterDeviceRequest
	GetToken() *string
	SetWlan(v string) *RegisterDeviceRequest
	GetWlan() *string
}

type RegisterDeviceRequest struct {
	Bluetooth  *string `json:"Bluetooth,omitempty" xml:"Bluetooth,omitempty"`
	BuildId    *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ChipId     *string `json:"ChipId,omitempty" xml:"ChipId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Cpu        *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CustomId   *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	EtherMac   *string `json:"EtherMac,omitempty" xml:"EtherMac,omitempty"`
	Memory     *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SerialNo   *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Storage    *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Wlan       *string `json:"Wlan,omitempty" xml:"Wlan,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) GetBluetooth() *string {
	return s.Bluetooth
}

func (s *RegisterDeviceRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *RegisterDeviceRequest) GetChipId() *string {
	return s.ChipId
}

func (s *RegisterDeviceRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RegisterDeviceRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *RegisterDeviceRequest) GetCpu() *string {
	return s.Cpu
}

func (s *RegisterDeviceRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *RegisterDeviceRequest) GetEtherMac() *string {
	return s.EtherMac
}

func (s *RegisterDeviceRequest) GetMemory() *string {
	return s.Memory
}

func (s *RegisterDeviceRequest) GetModel() *string {
	return s.Model
}

func (s *RegisterDeviceRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *RegisterDeviceRequest) GetStorage() *string {
	return s.Storage
}

func (s *RegisterDeviceRequest) GetToken() *string {
	return s.Token
}

func (s *RegisterDeviceRequest) GetWlan() *string {
	return s.Wlan
}

func (s *RegisterDeviceRequest) SetBluetooth(v string) *RegisterDeviceRequest {
	s.Bluetooth = &v
	return s
}

func (s *RegisterDeviceRequest) SetBuildId(v string) *RegisterDeviceRequest {
	s.BuildId = &v
	return s
}

func (s *RegisterDeviceRequest) SetChipId(v string) *RegisterDeviceRequest {
	s.ChipId = &v
	return s
}

func (s *RegisterDeviceRequest) SetClientId(v string) *RegisterDeviceRequest {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceRequest) SetClientType(v int32) *RegisterDeviceRequest {
	s.ClientType = &v
	return s
}

func (s *RegisterDeviceRequest) SetCpu(v string) *RegisterDeviceRequest {
	s.Cpu = &v
	return s
}

func (s *RegisterDeviceRequest) SetCustomId(v string) *RegisterDeviceRequest {
	s.CustomId = &v
	return s
}

func (s *RegisterDeviceRequest) SetEtherMac(v string) *RegisterDeviceRequest {
	s.EtherMac = &v
	return s
}

func (s *RegisterDeviceRequest) SetMemory(v string) *RegisterDeviceRequest {
	s.Memory = &v
	return s
}

func (s *RegisterDeviceRequest) SetModel(v string) *RegisterDeviceRequest {
	s.Model = &v
	return s
}

func (s *RegisterDeviceRequest) SetSerialNo(v string) *RegisterDeviceRequest {
	s.SerialNo = &v
	return s
}

func (s *RegisterDeviceRequest) SetStorage(v string) *RegisterDeviceRequest {
	s.Storage = &v
	return s
}

func (s *RegisterDeviceRequest) SetToken(v string) *RegisterDeviceRequest {
	s.Token = &v
	return s
}

func (s *RegisterDeviceRequest) SetWlan(v string) *RegisterDeviceRequest {
	s.Wlan = &v
	return s
}

func (s *RegisterDeviceRequest) Validate() error {
	return dara.Validate(s)
}
