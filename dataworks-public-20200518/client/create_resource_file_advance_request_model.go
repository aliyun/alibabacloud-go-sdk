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
	// The code for the file. The code format varies based on the file type. To view the code format for a specific file type, go to Operation Center, right-click a node of the file type, and then select View Code.
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
	// File name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The type of the code for the file. The code for files varies based on the file type. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// You can call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) operation to query the type of the code for the file.
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
	// The Alibaba Cloud User ID of the file owner. If this parameter is empty, the caller\\"s Alibaba cloud user ID is used by default.
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
	// Whether to synchronously Upload resources to the computing engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RegisterToCalcEngine *bool `json:"RegisterToCalcEngine,omitempty" xml:"RegisterToCalcEngine,omitempty"`
	// The OSS URL for uploading files provided by POP.
	//
	// example:
	//
	// http://bucketname1.oss-cn-shanghai.aliyuncs.com/example
	ResourceFileObject io.Reader `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
	// The storage path of the resource file on the computing engine. Currently, only EMR and CDH use this field. EMR format is [oss | hdfs]:// path/to/object. CDH must be/user/admin/lib by default.
	//
	// example:
	//
	// oss://oss-cn-shanghai.aliyuncs.com/emr-test
	StorageURL *string `json:"StorageURL,omitempty" xml:"StorageURL,omitempty"`
	// File resource Upload mode. Currently, only files of the File type of MaxCompute are valid. The values are as follows:
	//
	// - true: The Resource mode that can be downloaded.
	//
	// - false: The text mode for online editing.
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
