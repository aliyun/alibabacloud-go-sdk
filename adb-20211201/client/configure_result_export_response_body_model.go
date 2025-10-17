// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureResultExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfigureResultExportResponseBody
	GetCode() *string
	SetData(v *ConfigureResultExportResponseBodyData) *ConfigureResultExportResponseBody
	GetData() *ConfigureResultExportResponseBodyData
	SetHttpStatusCode(v int32) *ConfigureResultExportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ConfigureResultExportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigureResultExportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigureResultExportResponseBody
	GetSuccess() *bool
}

type ConfigureResultExportResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ConfigureResultExportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureResultExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigureResultExportResponseBody) GetData() *ConfigureResultExportResponseBodyData {
	return s.Data
}

func (s *ConfigureResultExportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConfigureResultExportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigureResultExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureResultExportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigureResultExportResponseBody) SetCode(v string) *ConfigureResultExportResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigureResultExportResponseBody) SetData(v *ConfigureResultExportResponseBodyData) *ConfigureResultExportResponseBody {
	s.Data = v
	return s
}

func (s *ConfigureResultExportResponseBody) SetHttpStatusCode(v int32) *ConfigureResultExportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureResultExportResponseBody) SetMessage(v string) *ConfigureResultExportResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigureResultExportResponseBody) SetRequestId(v string) *ConfigureResultExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureResultExportResponseBody) SetSuccess(v bool) *ConfigureResultExportResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureResultExportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureResultExportResponseBodyData struct {
	// example:
	//
	// SLS
	ExportType *string                                       `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	OssInfo    *ConfigureResultExportResponseBodyDataOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	SlsInfo    *ConfigureResultExportResponseBodyDataSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s ConfigureResultExportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportResponseBodyData) GetExportType() *string {
	return s.ExportType
}

func (s *ConfigureResultExportResponseBodyData) GetOssInfo() *ConfigureResultExportResponseBodyDataOssInfo {
	return s.OssInfo
}

func (s *ConfigureResultExportResponseBodyData) GetSlsInfo() *ConfigureResultExportResponseBodyDataSlsInfo {
	return s.SlsInfo
}

func (s *ConfigureResultExportResponseBodyData) SetExportType(v string) *ConfigureResultExportResponseBodyData {
	s.ExportType = &v
	return s
}

func (s *ConfigureResultExportResponseBodyData) SetOssInfo(v *ConfigureResultExportResponseBodyDataOssInfo) *ConfigureResultExportResponseBodyData {
	s.OssInfo = v
	return s
}

func (s *ConfigureResultExportResponseBodyData) SetSlsInfo(v *ConfigureResultExportResponseBodyDataSlsInfo) *ConfigureResultExportResponseBodyData {
	s.SlsInfo = v
	return s
}

func (s *ConfigureResultExportResponseBodyData) Validate() error {
	if s.OssInfo != nil {
		if err := s.OssInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SlsInfo != nil {
		if err := s.SlsInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureResultExportResponseBodyDataOssInfo struct {
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

func (s ConfigureResultExportResponseBodyDataOssInfo) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportResponseBodyDataOssInfo) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) GetExportBasePath() *string {
	return s.ExportBasePath
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) GetResultFileTtl() *int32 {
	return s.ResultFileTtl
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) SetExportBasePath(v string) *ConfigureResultExportResponseBodyDataOssInfo {
	s.ExportBasePath = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) SetResourceGroup(v string) *ConfigureResultExportResponseBodyDataOssInfo {
	s.ResourceGroup = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) SetResultFileTtl(v int32) *ConfigureResultExportResponseBodyDataOssInfo {
	s.ResultFileTtl = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataOssInfo) Validate() error {
	return dara.Validate(s)
}

type ConfigureResultExportResponseBodyDataSlsInfo struct {
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

func (s ConfigureResultExportResponseBodyDataSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s ConfigureResultExportResponseBodyDataSlsInfo) GoString() string {
	return s.String()
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) GetLogstoreTtl() *int32 {
	return s.LogstoreTtl
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) SetLogstoreTtl(v int32) *ConfigureResultExportResponseBodyDataSlsInfo {
	s.LogstoreTtl = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) SetResourceGroup(v string) *ConfigureResultExportResponseBodyDataSlsInfo {
	s.ResourceGroup = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) SetSlsProject(v string) *ConfigureResultExportResponseBodyDataSlsInfo {
	s.SlsProject = &v
	return s
}

func (s *ConfigureResultExportResponseBodyDataSlsInfo) Validate() error {
	return dara.Validate(s)
}
