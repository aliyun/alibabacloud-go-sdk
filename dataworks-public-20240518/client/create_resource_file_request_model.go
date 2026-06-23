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
	// The code content of the file. The code format varies depending on the code type (fileType). You can locate a job of the corresponding type in the Operation Center, right-click it, and select View Code to check the specific code format.
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
	// The code type of the file.
	//
	// Different file types correspond to different code types. For more information, see [DataWorks Node Types](https://help.aliyun.com/document_detail/600169.html).<br>
	//
	// You can also invoke the [ListFileType](https://help.aliyun.com/document_detail/212428.html) API to query the code types of files.
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
	// The Alibaba Cloud User ID of the file owner. If this parameter is empty, the Alibaba Cloud User ID of the caller is used by default.
	//
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console, go to the workspace configuration page, and obtain the workspace ID. This parameter is required to identify the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Indicates whether to synchronize and upload the resource to the compute engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RegisterToCalcEngine *bool `json:"RegisterToCalcEngine,omitempty" xml:"RegisterToCalcEngine,omitempty"`
	// The OSS URL for file upload provided by POP.
	//
	// example:
	//
	// http://bucketname1.oss-cn-shanghai.aliyuncs.com/example
	ResourceFile *string `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
	// The Storage Path of the resource file on the compute engine. This field is currently used only by EMR and CDH. The EMR format is [osshdfs]://path/to/object, and for CDH, the default value must be /user/admin/lib.
	//
	// example:
	//
	// oss://oss-cn-shanghai.aliyuncs.com/emr-test
	StorageURL *string `json:"StorageURL,omitempty" xml:"StorageURL,omitempty"`
	// The upload mode for the resource file. This parameter currently applies only to File-type files in MaxCompute. Valid values:
	//
	// - true: Downloadable resource mode.
	//
	// - false: Online-editable text mode.
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
