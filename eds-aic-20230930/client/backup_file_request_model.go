// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *BackupFileRequest
	GetAndroidInstanceIdList() []*string
	SetBackupAll(v bool) *BackupFileRequest
	GetBackupAll() *bool
	SetBackupFileName(v string) *BackupFileRequest
	GetBackupFileName() *string
	SetBackupFilePath(v string) *BackupFileRequest
	GetBackupFilePath() *string
	SetDescription(v string) *BackupFileRequest
	GetDescription() *string
	SetExcludeSourceFilePathList(v []*string) *BackupFileRequest
	GetExcludeSourceFilePathList() []*string
	SetSourceAppList(v []*string) *BackupFileRequest
	GetSourceAppList() []*string
	SetSourceFilePathList(v []*string) *BackupFileRequest
	GetSourceFilePathList() []*string
	SetUploadEndpoint(v string) *BackupFileRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *BackupFileRequest
	GetUploadType() *string
}

type BackupFileRequest struct {
	// A list of instance IDs.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to back up the entire instance.
	//
	// example:
	//
	// false
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The name of the backup file.
	//
	// example:
	//
	// MyBackup
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The upload URL for the backup file.
	//
	// > If you upload the file to an OSS bucket, call the DescribeBuckets operation to get the bucketName. Then, select a key from ossObjectList. The key represents the folder path in the OSS bucket. Combine these values into the format `oss://${bucketName}/${key}`.
	//
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// This is a backup/data request.
	Description               *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeSourceFilePathList []*string `json:"ExcludeSourceFilePathList,omitempty" xml:"ExcludeSourceFilePathList,omitempty" type:"Repeated"`
	// A list of application package names to back up.
	SourceAppList []*string `json:"SourceAppList,omitempty" xml:"SourceAppList,omitempty" type:"Repeated"`
	// A list of file paths to back up.
	SourceFilePathList []*string `json:"SourceFilePathList,omitempty" xml:"SourceFilePathList,omitempty" type:"Repeated"`
	// The domain name of the upload URL.
	//
	// > If you upload the file to an OSS bucket, call the DescribeBuckets operation to obtain the bucket information. If the cloud phone and the bucket are in the same region, use the value of the intranetEndpoint field. If they are in different regions, use the value of the extranetEndpoint field.
	//
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The backup type.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s BackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s BackupFileRequest) GoString() string {
	return s.String()
}

func (s *BackupFileRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *BackupFileRequest) GetBackupAll() *bool {
	return s.BackupAll
}

func (s *BackupFileRequest) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupFileRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *BackupFileRequest) GetDescription() *string {
	return s.Description
}

func (s *BackupFileRequest) GetExcludeSourceFilePathList() []*string {
	return s.ExcludeSourceFilePathList
}

func (s *BackupFileRequest) GetSourceAppList() []*string {
	return s.SourceAppList
}

func (s *BackupFileRequest) GetSourceFilePathList() []*string {
	return s.SourceFilePathList
}

func (s *BackupFileRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *BackupFileRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *BackupFileRequest) SetAndroidInstanceIdList(v []*string) *BackupFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *BackupFileRequest) SetBackupAll(v bool) *BackupFileRequest {
	s.BackupAll = &v
	return s
}

func (s *BackupFileRequest) SetBackupFileName(v string) *BackupFileRequest {
	s.BackupFileName = &v
	return s
}

func (s *BackupFileRequest) SetBackupFilePath(v string) *BackupFileRequest {
	s.BackupFilePath = &v
	return s
}

func (s *BackupFileRequest) SetDescription(v string) *BackupFileRequest {
	s.Description = &v
	return s
}

func (s *BackupFileRequest) SetExcludeSourceFilePathList(v []*string) *BackupFileRequest {
	s.ExcludeSourceFilePathList = v
	return s
}

func (s *BackupFileRequest) SetSourceAppList(v []*string) *BackupFileRequest {
	s.SourceAppList = v
	return s
}

func (s *BackupFileRequest) SetSourceFilePathList(v []*string) *BackupFileRequest {
	s.SourceFilePathList = v
	return s
}

func (s *BackupFileRequest) SetUploadEndpoint(v string) *BackupFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *BackupFileRequest) SetUploadType(v string) *BackupFileRequest {
	s.UploadType = &v
	return s
}

func (s *BackupFileRequest) Validate() error {
	return dara.Validate(s)
}
