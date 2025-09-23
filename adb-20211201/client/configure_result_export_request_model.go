// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureResultExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ConfigureResultExportRequest
	GetDBClusterId() *string
	SetExportType(v string) *ConfigureResultExportRequest
	GetExportType() *string
	SetOssInfo(v *ConfigureResultExportRequestOssInfo) *ConfigureResultExportRequest
	GetOssInfo() *ConfigureResultExportRequestOssInfo
	SetRegionId(v string) *ConfigureResultExportRequest
	GetRegionId() *string
	SetSlsInfo(v *ConfigureResultExportRequestSlsInfo) *ConfigureResultExportRequest
	GetSlsInfo() *ConfigureResultExportRequestSlsInfo
}

type ConfigureResultExportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// SLS
	ExportType *string                              `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	OssInfo    *ConfigureResultExportRequestOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsInfo  *ConfigureResultExportRequestSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s ConfigureResultExportRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportRequest) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ConfigureResultExportRequest) GetExportType() *string {
	return s.ExportType
}

func (s *ConfigureResultExportRequest) GetOssInfo() *ConfigureResultExportRequestOssInfo {
	return s.OssInfo
}

func (s *ConfigureResultExportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureResultExportRequest) GetSlsInfo() *ConfigureResultExportRequestSlsInfo {
	return s.SlsInfo
}

func (s *ConfigureResultExportRequest) SetDBClusterId(v string) *ConfigureResultExportRequest {
	s.DBClusterId = &v
	return s
}

func (s *ConfigureResultExportRequest) SetExportType(v string) *ConfigureResultExportRequest {
	s.ExportType = &v
	return s
}

func (s *ConfigureResultExportRequest) SetOssInfo(v *ConfigureResultExportRequestOssInfo) *ConfigureResultExportRequest {
	s.OssInfo = v
	return s
}

func (s *ConfigureResultExportRequest) SetRegionId(v string) *ConfigureResultExportRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureResultExportRequest) SetSlsInfo(v *ConfigureResultExportRequestSlsInfo) *ConfigureResultExportRequest {
	s.SlsInfo = v
	return s
}

func (s *ConfigureResultExportRequest) Validate() error {
	return dara.Validate(s)
}

type ConfigureResultExportRequestOssInfo struct {
	// example:
	//
	// oss://testBucketName/data_result
	ExportBasePath *string `json:"ExportBasePath,omitempty" xml:"ExportBasePath,omitempty"`
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// 10
	ResultFileTtl *int32 `json:"ResultFileTtl,omitempty" xml:"ResultFileTtl,omitempty"`
}

func (s ConfigureResultExportRequestOssInfo) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportRequestOssInfo) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportRequestOssInfo) GetExportBasePath() *string {
	return s.ExportBasePath
}

func (s *ConfigureResultExportRequestOssInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ConfigureResultExportRequestOssInfo) GetResultFileTtl() *int32 {
	return s.ResultFileTtl
}

func (s *ConfigureResultExportRequestOssInfo) SetExportBasePath(v string) *ConfigureResultExportRequestOssInfo {
	s.ExportBasePath = &v
	return s
}

func (s *ConfigureResultExportRequestOssInfo) SetResourceGroup(v string) *ConfigureResultExportRequestOssInfo {
	s.ResourceGroup = &v
	return s
}

func (s *ConfigureResultExportRequestOssInfo) SetResultFileTtl(v int32) *ConfigureResultExportRequestOssInfo {
	s.ResultFileTtl = &v
	return s
}

func (s *ConfigureResultExportRequestOssInfo) Validate() error {
	return dara.Validate(s)
}

type ConfigureResultExportRequestSlsInfo struct {
	// example:
	//
	// 10
	LogstoreTtl *int32 `json:"LogstoreTtl,omitempty" xml:"LogstoreTtl,omitempty"`
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// test-project
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s ConfigureResultExportRequestSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportRequestSlsInfo) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportRequestSlsInfo) GetLogstoreTtl() *int32 {
	return s.LogstoreTtl
}

func (s *ConfigureResultExportRequestSlsInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ConfigureResultExportRequestSlsInfo) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ConfigureResultExportRequestSlsInfo) SetLogstoreTtl(v int32) *ConfigureResultExportRequestSlsInfo {
	s.LogstoreTtl = &v
	return s
}

func (s *ConfigureResultExportRequestSlsInfo) SetResourceGroup(v string) *ConfigureResultExportRequestSlsInfo {
	s.ResourceGroup = &v
	return s
}

func (s *ConfigureResultExportRequestSlsInfo) SetSlsProject(v string) *ConfigureResultExportRequestSlsInfo {
	s.SlsProject = &v
	return s
}

func (s *ConfigureResultExportRequestSlsInfo) Validate() error {
	return dara.Validate(s)
}
