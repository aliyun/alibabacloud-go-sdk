// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnvsResponseBody
	GetCode() *string
	SetEnvs(v []*ListEnvsResponseBodyEnvs) *ListEnvsResponseBody
	GetEnvs() []*ListEnvsResponseBodyEnvs
	SetHttpStatusCode(v int32) *ListEnvsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEnvsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListEnvsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnvsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEnvsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListEnvsResponseBody
	GetTotalCount() *int64
}

type ListEnvsResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The environments.
	Envs []*ListEnvsResponseBodyEnvs `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of environments per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of environments.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEnvsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnvsResponseBody) GetEnvs() []*ListEnvsResponseBodyEnvs {
	return s.Envs
}

func (s *ListEnvsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEnvsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnvsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnvsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEnvsResponseBody) SetCode(v string) *ListEnvsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvsResponseBody) SetEnvs(v []*ListEnvsResponseBodyEnvs) *ListEnvsResponseBody {
	s.Envs = v
	return s
}

func (s *ListEnvsResponseBody) SetHttpStatusCode(v int32) *ListEnvsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEnvsResponseBody) SetMessage(v string) *ListEnvsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvsResponseBody) SetPageNumber(v int32) *ListEnvsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEnvsResponseBody) SetPageSize(v int32) *ListEnvsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEnvsResponseBody) SetRequestId(v string) *ListEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvsResponseBody) SetSuccess(v bool) *ListEnvsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvsResponseBody) SetTotalCount(v int64) *ListEnvsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEnvsResponseBody) Validate() error {
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvsResponseBodyEnvs struct {
	// The time when the environment was created.
	//
	// example:
	//
	// 1637053715165
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// 86S1LH
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// The name of the environment.
	//
	// example:
	//
	// test-create
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// The JMeter version of the environment.
	//
	// example:
	//
	// 5.0
	EnvVersion *string `json:"EnvVersion,omitempty" xml:"EnvVersion,omitempty"`
	// The JAR files.
	Files []*ListEnvsResponseBodyEnvsFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The time when the environment was last modified.
	//
	// example:
	//
	// 1637053719165
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The JMeter attributes.
	Properties []*ListEnvsResponseBodyEnvsProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The scenarios related to the environment.
	RelatedScenes []*string `json:"RelatedScenes,omitempty" xml:"RelatedScenes,omitempty" type:"Repeated"`
	// The IDs of the scenarios that run in the environment.
	RunningScenes []*string `json:"RunningScenes,omitempty" xml:"RunningScenes,omitempty" type:"Repeated"`
	// The total size of the environment file.
	//
	// example:
	//
	// 26668
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s ListEnvsResponseBodyEnvs) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsResponseBodyEnvs) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListEnvsResponseBodyEnvs) GetEnvId() *string {
	return s.EnvId
}

func (s *ListEnvsResponseBodyEnvs) GetEnvName() *string {
	return s.EnvName
}

func (s *ListEnvsResponseBodyEnvs) GetEnvVersion() *string {
	return s.EnvVersion
}

func (s *ListEnvsResponseBodyEnvs) GetFiles() []*ListEnvsResponseBodyEnvsFiles {
	return s.Files
}

func (s *ListEnvsResponseBodyEnvs) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListEnvsResponseBodyEnvs) GetProperties() []*ListEnvsResponseBodyEnvsProperties {
	return s.Properties
}

func (s *ListEnvsResponseBodyEnvs) GetRelatedScenes() []*string {
	return s.RelatedScenes
}

func (s *ListEnvsResponseBodyEnvs) GetRunningScenes() []*string {
	return s.RunningScenes
}

func (s *ListEnvsResponseBodyEnvs) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *ListEnvsResponseBodyEnvs) SetCreateTime(v int64) *ListEnvsResponseBodyEnvs {
	s.CreateTime = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvId(v string) *ListEnvsResponseBodyEnvs {
	s.EnvId = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvName(v string) *ListEnvsResponseBodyEnvs {
	s.EnvName = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvVersion(v string) *ListEnvsResponseBodyEnvs {
	s.EnvVersion = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetFiles(v []*ListEnvsResponseBodyEnvsFiles) *ListEnvsResponseBodyEnvs {
	s.Files = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetModifiedTime(v int64) *ListEnvsResponseBodyEnvs {
	s.ModifiedTime = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetProperties(v []*ListEnvsResponseBodyEnvsProperties) *ListEnvsResponseBodyEnvs {
	s.Properties = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetRelatedScenes(v []*string) *ListEnvsResponseBodyEnvs {
	s.RelatedScenes = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetRunningScenes(v []*string) *ListEnvsResponseBodyEnvs {
	s.RunningScenes = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetUsedCapacity(v int64) *ListEnvsResponseBodyEnvs {
	s.UsedCapacity = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvsResponseBodyEnvsFiles struct {
	// The ID of the file.
	//
	// example:
	//
	// 61660
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// json.jar
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS address of the file.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/json.jar
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 788
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The MD5 checksum of the file.
	//
	// example:
	//
	// 43B584026CE5E570F3DE638FA7EEF9E0
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s ListEnvsResponseBodyEnvsFiles) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsResponseBodyEnvsFiles) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvsFiles) GetFileId() *int64 {
	return s.FileId
}

func (s *ListEnvsResponseBodyEnvsFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListEnvsResponseBodyEnvsFiles) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *ListEnvsResponseBodyEnvsFiles) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListEnvsResponseBodyEnvsFiles) GetMd5() *string {
	return s.Md5
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileId(v int64) *ListEnvsResponseBodyEnvsFiles {
	s.FileId = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileName(v string) *ListEnvsResponseBodyEnvsFiles {
	s.FileName = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileOssAddress(v string) *ListEnvsResponseBodyEnvsFiles {
	s.FileOssAddress = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileSize(v int64) *ListEnvsResponseBodyEnvsFiles {
	s.FileSize = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetMd5(v string) *ListEnvsResponseBodyEnvsFiles {
	s.Md5 = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) Validate() error {
	return dara.Validate(s)
}

type ListEnvsResponseBodyEnvsProperties struct {
	// The description of the attribute.
	//
	// example:
	//
	// 远程主机
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the attribute.
	//
	// example:
	//
	// remote_hosts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// 127.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEnvsResponseBodyEnvsProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsResponseBodyEnvsProperties) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvsProperties) GetDescription() *string {
	return s.Description
}

func (s *ListEnvsResponseBodyEnvsProperties) GetName() *string {
	return s.Name
}

func (s *ListEnvsResponseBodyEnvsProperties) GetValue() *string {
	return s.Value
}

func (s *ListEnvsResponseBodyEnvsProperties) SetDescription(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Description = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsProperties) SetName(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Name = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsProperties) SetValue(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Value = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsProperties) Validate() error {
	return dara.Validate(s)
}
