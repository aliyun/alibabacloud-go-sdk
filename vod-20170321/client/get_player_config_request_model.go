// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayerConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *GetPlayerConfigRequest
	GetApiVersion() *string
	SetAuthInfo(v string) *GetPlayerConfigRequest
	GetAuthInfo() *string
	SetAuthTimestamp(v int64) *GetPlayerConfigRequest
	GetAuthTimestamp() *int64
	SetDeviceBrand(v string) *GetPlayerConfigRequest
	GetDeviceBrand() *string
	SetDeviceModel(v string) *GetPlayerConfigRequest
	GetDeviceModel() *string
	SetOsName(v string) *GetPlayerConfigRequest
	GetOsName() *string
	SetReserved(v string) *GetPlayerConfigRequest
	GetReserved() *string
	SetRule(v string) *GetPlayerConfigRequest
	GetRule() *string
}

type GetPlayerConfigRequest struct {
	ApiVersion    *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	AuthInfo      *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	AuthTimestamp *int64  `json:"AuthTimestamp,omitempty" xml:"AuthTimestamp,omitempty"`
	DeviceBrand   *string `json:"DeviceBrand,omitempty" xml:"DeviceBrand,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	OsName        *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	Reserved      *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	Rule          *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s GetPlayerConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlayerConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPlayerConfigRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetPlayerConfigRequest) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *GetPlayerConfigRequest) GetAuthTimestamp() *int64 {
	return s.AuthTimestamp
}

func (s *GetPlayerConfigRequest) GetDeviceBrand() *string {
	return s.DeviceBrand
}

func (s *GetPlayerConfigRequest) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetPlayerConfigRequest) GetOsName() *string {
	return s.OsName
}

func (s *GetPlayerConfigRequest) GetReserved() *string {
	return s.Reserved
}

func (s *GetPlayerConfigRequest) GetRule() *string {
	return s.Rule
}

func (s *GetPlayerConfigRequest) SetApiVersion(v string) *GetPlayerConfigRequest {
	s.ApiVersion = &v
	return s
}

func (s *GetPlayerConfigRequest) SetAuthInfo(v string) *GetPlayerConfigRequest {
	s.AuthInfo = &v
	return s
}

func (s *GetPlayerConfigRequest) SetAuthTimestamp(v int64) *GetPlayerConfigRequest {
	s.AuthTimestamp = &v
	return s
}

func (s *GetPlayerConfigRequest) SetDeviceBrand(v string) *GetPlayerConfigRequest {
	s.DeviceBrand = &v
	return s
}

func (s *GetPlayerConfigRequest) SetDeviceModel(v string) *GetPlayerConfigRequest {
	s.DeviceModel = &v
	return s
}

func (s *GetPlayerConfigRequest) SetOsName(v string) *GetPlayerConfigRequest {
	s.OsName = &v
	return s
}

func (s *GetPlayerConfigRequest) SetReserved(v string) *GetPlayerConfigRequest {
	s.Reserved = &v
	return s
}

func (s *GetPlayerConfigRequest) SetRule(v string) *GetPlayerConfigRequest {
	s.Rule = &v
	return s
}

func (s *GetPlayerConfigRequest) Validate() error {
	return dara.Validate(s)
}
