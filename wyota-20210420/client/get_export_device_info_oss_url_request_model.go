// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportDeviceInfoOssUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantId(v string) *GetExportDeviceInfoOssUrlRequest
	GetTenantId() *string
	SetZoneId(v string) *GetExportDeviceInfoOssUrlRequest
	GetZoneId() *string
	SetZoneName(v string) *GetExportDeviceInfoOssUrlRequest
	GetZoneName() *string
}

type GetExportDeviceInfoOssUrlRequest struct {
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s GetExportDeviceInfoOssUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlRequest) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetExportDeviceInfoOssUrlRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetExportDeviceInfoOssUrlRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *GetExportDeviceInfoOssUrlRequest) SetTenantId(v string) *GetExportDeviceInfoOssUrlRequest {
	s.TenantId = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlRequest) SetZoneId(v string) *GetExportDeviceInfoOssUrlRequest {
	s.ZoneId = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlRequest) SetZoneName(v string) *GetExportDeviceInfoOssUrlRequest {
	s.ZoneName = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlRequest) Validate() error {
	return dara.Validate(s)
}
