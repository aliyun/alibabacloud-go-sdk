// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseVersion(v string) *GetDeviceOtaInfoTestRequest
	GetBaseVersion() *string
	SetDeviceId(v string) *GetDeviceOtaInfoTestRequest
	GetDeviceId() *string
	SetModel(v string) *GetDeviceOtaInfoTestRequest
	GetModel() *string
	SetTenantId(v string) *GetDeviceOtaInfoTestRequest
	GetTenantId() *string
}

type GetDeviceOtaInfoTestRequest struct {
	// This parameter is required.
	BaseVersion *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDeviceOtaInfoTestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoTestRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestRequest) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *GetDeviceOtaInfoTestRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceOtaInfoTestRequest) GetModel() *string {
	return s.Model
}

func (s *GetDeviceOtaInfoTestRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDeviceOtaInfoTestRequest) SetBaseVersion(v string) *GetDeviceOtaInfoTestRequest {
	s.BaseVersion = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetDeviceId(v string) *GetDeviceOtaInfoTestRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetModel(v string) *GetDeviceOtaInfoTestRequest {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) SetTenantId(v string) *GetDeviceOtaInfoTestRequest {
	s.TenantId = &v
	return s
}

func (s *GetDeviceOtaInfoTestRequest) Validate() error {
	return dara.Validate(s)
}
