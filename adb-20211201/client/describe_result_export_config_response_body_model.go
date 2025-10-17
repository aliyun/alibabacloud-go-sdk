// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultExportConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResultExportConfigResponseBody
	GetCode() *string
	SetData(v *DescribeResultExportConfigResponseBodyData) *DescribeResultExportConfigResponseBody
	GetData() *DescribeResultExportConfigResponseBodyData
	SetHttpStatusCode(v int32) *DescribeResultExportConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeResultExportConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResultExportConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResultExportConfigResponseBody
	GetSuccess() *bool
}

type DescribeResultExportConfigResponseBody struct {
	// example:
	//
	// InvalidInput
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeResultExportConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeResultExportConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResultExportConfigResponseBody) GetData() *DescribeResultExportConfigResponseBodyData {
	return s.Data
}

func (s *DescribeResultExportConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeResultExportConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResultExportConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResultExportConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResultExportConfigResponseBody) SetCode(v string) *DescribeResultExportConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResultExportConfigResponseBody) SetData(v *DescribeResultExportConfigResponseBodyData) *DescribeResultExportConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResultExportConfigResponseBody) SetHttpStatusCode(v int32) *DescribeResultExportConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeResultExportConfigResponseBody) SetMessage(v string) *DescribeResultExportConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResultExportConfigResponseBody) SetRequestId(v string) *DescribeResultExportConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResultExportConfigResponseBody) SetSuccess(v bool) *DescribeResultExportConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResultExportConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResultExportConfigResponseBodyData struct {
	// example:
	//
	// OSS
	ExportType *string                                            `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	OssInfo    *DescribeResultExportConfigResponseBodyDataOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	SlsInfo    *DescribeResultExportConfigResponseBodyDataSlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s DescribeResultExportConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigResponseBodyData) GetExportType() *string {
	return s.ExportType
}

func (s *DescribeResultExportConfigResponseBodyData) GetOssInfo() *DescribeResultExportConfigResponseBodyDataOssInfo {
	return s.OssInfo
}

func (s *DescribeResultExportConfigResponseBodyData) GetSlsInfo() *DescribeResultExportConfigResponseBodyDataSlsInfo {
	return s.SlsInfo
}

func (s *DescribeResultExportConfigResponseBodyData) SetExportType(v string) *DescribeResultExportConfigResponseBodyData {
	s.ExportType = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyData) SetOssInfo(v *DescribeResultExportConfigResponseBodyDataOssInfo) *DescribeResultExportConfigResponseBodyData {
	s.OssInfo = v
	return s
}

func (s *DescribeResultExportConfigResponseBodyData) SetSlsInfo(v *DescribeResultExportConfigResponseBodyDataSlsInfo) *DescribeResultExportConfigResponseBodyData {
	s.SlsInfo = v
	return s
}

func (s *DescribeResultExportConfigResponseBodyData) Validate() error {
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

type DescribeResultExportConfigResponseBodyDataOssInfo struct {
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

func (s DescribeResultExportConfigResponseBodyDataOssInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigResponseBodyDataOssInfo) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) GetExportBasePath() *string {
	return s.ExportBasePath
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) GetResultFileTtl() *int32 {
	return s.ResultFileTtl
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) SetExportBasePath(v string) *DescribeResultExportConfigResponseBodyDataOssInfo {
	s.ExportBasePath = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) SetResourceGroup(v string) *DescribeResultExportConfigResponseBodyDataOssInfo {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) SetResultFileTtl(v int32) *DescribeResultExportConfigResponseBodyDataOssInfo {
	s.ResultFileTtl = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataOssInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeResultExportConfigResponseBodyDataSlsInfo struct {
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

func (s DescribeResultExportConfigResponseBodyDataSlsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigResponseBodyDataSlsInfo) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) GetLogstoreTtl() *int32 {
	return s.LogstoreTtl
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) SetLogstoreTtl(v int32) *DescribeResultExportConfigResponseBodyDataSlsInfo {
	s.LogstoreTtl = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) SetResourceGroup(v string) *DescribeResultExportConfigResponseBodyDataSlsInfo {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) SetSlsProject(v string) *DescribeResultExportConfigResponseBodyDataSlsInfo {
	s.SlsProject = &v
	return s
}

func (s *DescribeResultExportConfigResponseBodyDataSlsInfo) Validate() error {
	return dara.Validate(s)
}
