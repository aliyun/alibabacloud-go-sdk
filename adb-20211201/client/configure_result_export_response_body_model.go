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
	// The backend error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The path where the result sets are stored.
	Data *ConfigureResultExportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- OK is returned if the request is normal.
	//
	// 	- The specific error code is returned if the request is abnormal,
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: Succeeded.
	//
	// 	- **false**: Failed.
	//
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
	// The export type. Valid values:
	//
	// 	- SLS: Indicates that the export destination is SLS.
	//
	// 	- OSS: Indicates that the export destination is OSS.
	//
	// example:
	//
	// SLS
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The OSS configuration if the destination is of the OSS type.
	OssInfo *ConfigureResultExportResponseBodyDataOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	// The SLS configuration if the destination is of the SLS type.
	SlsInfo *ConfigureResultExportResponseBodyDataSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
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
	// The OSS path where the result sets are stored.
	//
	// example:
	//
	// oss://testBucketName/data_result
	ExportBasePath *string `json:"ExportBasePath,omitempty" xml:"ExportBasePath,omitempty"`
	// The name of the resource group that runs the job.
	//
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The expiration time of the OSS file. Unit: days. Minimum: 1 day, maximum: 30 days.
	//
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
	// The expiration time of the Logstore temporarily generated during the result set export. Unit: days. The Logstore is automatically deleted after it expires. Minimum value: 1 day. Maximum value: 30 days.
	//
	// example:
	//
	// 10
	LogstoreTtl *int32 `json:"LogstoreTtl,omitempty" xml:"LogstoreTtl,omitempty"`
	// The name of the resource group that runs the job.
	//
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The name of the SLS project.
	//
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
