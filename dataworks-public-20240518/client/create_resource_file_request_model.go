// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateResourceFileRequest
	GetContent() *string
	SetFileDescription(v string) *CreateResourceFileRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *CreateResourceFileRequest
	GetFileFolderPath() *string
	SetFileName(v string) *CreateResourceFileRequest
	GetFileName() *string
	SetFileType(v int32) *CreateResourceFileRequest
	GetFileType() *int32
	SetOriginResourceName(v string) *CreateResourceFileRequest
	GetOriginResourceName() *string
	SetOwner(v string) *CreateResourceFileRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateResourceFileRequest
	GetProjectId() *int64
	SetRegisterToCalcEngine(v bool) *CreateResourceFileRequest
	GetRegisterToCalcEngine() *bool
	SetResourceFile(v string) *CreateResourceFileRequest
	GetResourceFile() *string
	SetStorageURL(v string) *CreateResourceFileRequest
	GetStorageURL() *string
	SetUploadMode(v bool) *CreateResourceFileRequest
	GetUploadMode() *bool
}

type CreateResourceFileRequest struct {
	// The code for the file. The code format varies based on the file type. To view the code format for a specific file type, go to Operation Center, open the directed acyclic graph (DAG) of a node of the file type, right-click the node, and then select View Code.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// This is a description
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The path of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business_process/First_Business_Process/MaxCompute/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// Filename
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the code for the file.
	//
	// The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html). You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the type of the code for the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The name of the original resource file.
	//
	// This parameter is required.
	//
	// example:
	//
	// origin_file_name.sql
	OriginResourceName *string `json:"OriginResourceName,omitempty" xml:"OriginResourceName,omitempty"`
	// The ID of the Alibaba Cloud account used by the file owner. If this parameter is not configured, the ID of the Alibaba Cloud account of the user who calls the operation is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to upload the resource file to a desired compute engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RegisterToCalcEngine *bool `json:"RegisterToCalcEngine,omitempty" xml:"RegisterToCalcEngine,omitempty"`
	// The URL of the Object Storage Service (OSS) bucket to which you upload the file. The URL is provided by the POP platform.
	//
	// example:
	//
	// http://bucketname1.oss-cn-shanghai.aliyuncs.com/example
	ResourceFile *string `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
	// The storage path of the resource file in a desired compute engine. This parameter takes effect only for E-MapReduce (EMR) and Cloudera\\"s Distribution including Apache Hadoop (CDH) compute engines. In an EMR compute engine, this parameter is configured in the [osshdfs]://path/to/object format. In a CDH compute engine, this parameter is set to /user/admin/lib by default.
	//
	// example:
	//
	// oss://oss-cn-shanghai.aliyuncs.com/emr-test
	StorageURL *string `json:"StorageURL,omitempty" xml:"StorageURL,omitempty"`
	// The upload mode of MaxCompute file resources. This parameter takes effect only for MaxCompute file resources. Valid values:
	//
	// 	- true: indicates the resource upload and download mode.
	//
	// 	- false: indicates the online editing mode.
	//
	// example:
	//
	// false
	UploadMode *bool `json:"UploadMode,omitempty" xml:"UploadMode,omitempty"`
}

func (s CreateResourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceFileRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceFileRequest) GetContent() *string {
	return s.Content
}

func (s *CreateResourceFileRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *CreateResourceFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *CreateResourceFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateResourceFileRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *CreateResourceFileRequest) GetOriginResourceName() *string {
	return s.OriginResourceName
}

func (s *CreateResourceFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateResourceFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateResourceFileRequest) GetRegisterToCalcEngine() *bool {
	return s.RegisterToCalcEngine
}

func (s *CreateResourceFileRequest) GetResourceFile() *string {
	return s.ResourceFile
}

func (s *CreateResourceFileRequest) GetStorageURL() *string {
	return s.StorageURL
}

func (s *CreateResourceFileRequest) GetUploadMode() *bool {
	return s.UploadMode
}

func (s *CreateResourceFileRequest) SetContent(v string) *CreateResourceFileRequest {
	s.Content = &v
	return s
}

func (s *CreateResourceFileRequest) SetFileDescription(v string) *CreateResourceFileRequest {
	s.FileDescription = &v
	return s
}

func (s *CreateResourceFileRequest) SetFileFolderPath(v string) *CreateResourceFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *CreateResourceFileRequest) SetFileName(v string) *CreateResourceFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateResourceFileRequest) SetFileType(v int32) *CreateResourceFileRequest {
	s.FileType = &v
	return s
}

func (s *CreateResourceFileRequest) SetOriginResourceName(v string) *CreateResourceFileRequest {
	s.OriginResourceName = &v
	return s
}

func (s *CreateResourceFileRequest) SetOwner(v string) *CreateResourceFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateResourceFileRequest) SetProjectId(v int64) *CreateResourceFileRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceFileRequest) SetRegisterToCalcEngine(v bool) *CreateResourceFileRequest {
	s.RegisterToCalcEngine = &v
	return s
}

func (s *CreateResourceFileRequest) SetResourceFile(v string) *CreateResourceFileRequest {
	s.ResourceFile = &v
	return s
}

func (s *CreateResourceFileRequest) SetStorageURL(v string) *CreateResourceFileRequest {
	s.StorageURL = &v
	return s
}

func (s *CreateResourceFileRequest) SetUploadMode(v bool) *CreateResourceFileRequest {
	s.UploadMode = &v
	return s
}

func (s *CreateResourceFileRequest) Validate() error {
	return dara.Validate(s)
}
