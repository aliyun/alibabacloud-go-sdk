// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureResultExportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ConfigureResultExportShrinkRequest
	GetDBClusterId() *string
	SetExportType(v string) *ConfigureResultExportShrinkRequest
	GetExportType() *string
	SetOssInfoShrink(v string) *ConfigureResultExportShrinkRequest
	GetOssInfoShrink() *string
	SetRegionId(v string) *ConfigureResultExportShrinkRequest
	GetRegionId() *string
	SetSlsInfoShrink(v string) *ConfigureResultExportShrinkRequest
	GetSlsInfoShrink() *string
}

type ConfigureResultExportShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// SLS
	ExportType    *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	OssInfoShrink *string `json:"OssInfo,omitempty" xml:"OssInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsInfoShrink *string `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty"`
}

func (s ConfigureResultExportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ConfigureResultExportShrinkRequest) GetExportType() *string {
	return s.ExportType
}

func (s *ConfigureResultExportShrinkRequest) GetOssInfoShrink() *string {
	return s.OssInfoShrink
}

func (s *ConfigureResultExportShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureResultExportShrinkRequest) GetSlsInfoShrink() *string {
	return s.SlsInfoShrink
}

func (s *ConfigureResultExportShrinkRequest) SetDBClusterId(v string) *ConfigureResultExportShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ConfigureResultExportShrinkRequest) SetExportType(v string) *ConfigureResultExportShrinkRequest {
	s.ExportType = &v
	return s
}

func (s *ConfigureResultExportShrinkRequest) SetOssInfoShrink(v string) *ConfigureResultExportShrinkRequest {
	s.OssInfoShrink = &v
	return s
}

func (s *ConfigureResultExportShrinkRequest) SetRegionId(v string) *ConfigureResultExportShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureResultExportShrinkRequest) SetSlsInfoShrink(v string) *ConfigureResultExportShrinkRequest {
	s.SlsInfoShrink = &v
	return s
}

func (s *ConfigureResultExportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
