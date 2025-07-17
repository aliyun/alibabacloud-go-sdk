// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateResourceFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateResourceFileAdvanceRequest
	GetContent() *string
	SetFileDescription(v string) *CreateResourceFileAdvanceRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *CreateResourceFileAdvanceRequest
	GetFileFolderPath() *string
	SetFileName(v string) *CreateResourceFileAdvanceRequest
	GetFileName() *string
	SetFileType(v int32) *CreateResourceFileAdvanceRequest
	GetFileType() *int32
	SetOriginResourceName(v string) *CreateResourceFileAdvanceRequest
	GetOriginResourceName() *string
	SetOwner(v string) *CreateResourceFileAdvanceRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateResourceFileAdvanceRequest
	GetProjectId() *int64
	SetRegisterToCalcEngine(v bool) *CreateResourceFileAdvanceRequest
	GetRegisterToCalcEngine() *bool
	SetResourceFileObject(v io.Reader) *CreateResourceFileAdvanceRequest
	GetResourceFileObject() io.Reader
	SetStorageURL(v string) *CreateResourceFileAdvanceRequest
	GetStorageURL() *string
	SetUploadMode(v bool) *CreateResourceFileAdvanceRequest
	GetUploadMode() *bool
}

type CreateResourceFileAdvanceRequest struct {
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
	ResourceFileObject io.Reader `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
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

func (s CreateResourceFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceFileAdvanceRequest) GetContent() *string {
	return s.Content
}

func (s *CreateResourceFileAdvanceRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *CreateResourceFileAdvanceRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *CreateResourceFileAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateResourceFileAdvanceRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *CreateResourceFileAdvanceRequest) GetOriginResourceName() *string {
	return s.OriginResourceName
}

func (s *CreateResourceFileAdvanceRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateResourceFileAdvanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateResourceFileAdvanceRequest) GetRegisterToCalcEngine() *bool {
	return s.RegisterToCalcEngine
}

func (s *CreateResourceFileAdvanceRequest) GetResourceFileObject() io.Reader {
	return s.ResourceFileObject
}

func (s *CreateResourceFileAdvanceRequest) GetStorageURL() *string {
	return s.StorageURL
}

func (s *CreateResourceFileAdvanceRequest) GetUploadMode() *bool {
	return s.UploadMode
}

func (s *CreateResourceFileAdvanceRequest) SetContent(v string) *CreateResourceFileAdvanceRequest {
	s.Content = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetFileDescription(v string) *CreateResourceFileAdvanceRequest {
	s.FileDescription = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetFileFolderPath(v string) *CreateResourceFileAdvanceRequest {
	s.FileFolderPath = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetFileName(v string) *CreateResourceFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetFileType(v int32) *CreateResourceFileAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetOriginResourceName(v string) *CreateResourceFileAdvanceRequest {
	s.OriginResourceName = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetOwner(v string) *CreateResourceFileAdvanceRequest {
	s.Owner = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetProjectId(v int64) *CreateResourceFileAdvanceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetRegisterToCalcEngine(v bool) *CreateResourceFileAdvanceRequest {
	s.RegisterToCalcEngine = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetResourceFileObject(v io.Reader) *CreateResourceFileAdvanceRequest {
	s.ResourceFileObject = v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetStorageURL(v string) *CreateResourceFileAdvanceRequest {
	s.StorageURL = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) SetUploadMode(v bool) *CreateResourceFileAdvanceRequest {
	s.UploadMode = &v
	return s
}

func (s *CreateResourceFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
