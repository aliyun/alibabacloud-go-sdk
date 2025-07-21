// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseVersion(v string) *GetDeviceOtaInfoRequest
	GetBaseVersion() *string
	SetChannel(v string) *GetDeviceOtaInfoRequest
	GetChannel() *string
	SetDeviceId(v string) *GetDeviceOtaInfoRequest
	GetDeviceId() *string
	SetModel(v string) *GetDeviceOtaInfoRequest
	GetModel() *string
	SetNetworkType(v string) *GetDeviceOtaInfoRequest
	GetNetworkType() *string
	SetOsVersion(v string) *GetDeviceOtaInfoRequest
	GetOsVersion() *string
	SetRegion(v string) *GetDeviceOtaInfoRequest
	GetRegion() *string
	SetRegionId(v string) *GetDeviceOtaInfoRequest
	GetRegionId() *string
	SetTargetVersionType(v string) *GetDeviceOtaInfoRequest
	GetTargetVersionType() *string
	SetTenantId(v string) *GetDeviceOtaInfoRequest
	GetTenantId() *string
}

type GetDeviceOtaInfoRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	Model             *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OsVersion         *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetVersionType *string `json:"TargetVersionType,omitempty" xml:"TargetVersionType,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceOtaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoRequest) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *GetDeviceOtaInfoRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetDeviceOtaInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceOtaInfoRequest) GetModel() *string {
	return s.Model
}

func (s *GetDeviceOtaInfoRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetDeviceOtaInfoRequest) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetDeviceOtaInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *GetDeviceOtaInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDeviceOtaInfoRequest) GetTargetVersionType() *string {
	return s.TargetVersionType
}

func (s *GetDeviceOtaInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDeviceOtaInfoRequest) SetBaseVersion(v string) *GetDeviceOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetChannel(v string) *GetDeviceOtaInfoRequest {
	s.Channel = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetDeviceId(v string) *GetDeviceOtaInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetModel(v string) *GetDeviceOtaInfoRequest {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetNetworkType(v string) *GetDeviceOtaInfoRequest {
	s.NetworkType = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetOsVersion(v string) *GetDeviceOtaInfoRequest {
	s.OsVersion = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetRegion(v string) *GetDeviceOtaInfoRequest {
	s.Region = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetRegionId(v string) *GetDeviceOtaInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetTargetVersionType(v string) *GetDeviceOtaInfoRequest {
	s.TargetVersionType = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) SetTenantId(v string) *GetDeviceOtaInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetDeviceOtaInfoRequest) Validate() error {
	return dara.Validate(s)
}
