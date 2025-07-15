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
	// The IDs of the instances.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to back up the whole instance.
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The name of the backup file.
	//
	// example:
	//
	// defaultBackupFile
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The OSS path of the backup file.
	//
	// >  To upload a backup file to an OSS bucket, you must obtain the name of the bucket. When calling the describeBuckets operation to retrieve a bucket name, you must also call the ossObjectList operation to obtain the object key. Combine these to form the full path: oss://${bucketName}/${key}.
	//
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// This is a backup file description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The names of the application packages that you want to back up.
	SourceAppList []*string `json:"SourceAppList,omitempty" xml:"SourceAppList,omitempty" type:"Repeated"`
	// The paths to the source files.
	SourceFilePathList []*string `json:"SourceFilePathList,omitempty" xml:"SourceFilePathList,omitempty" type:"Repeated"`
	// The endpoint of the OSS bucket to which you want to upload the backup file.
	//
	// > : When calling the DescribeBuckets operation to query buckets, retrieve the IntranetEndpoint value if the cloud phone and the OSS bucket are in the same region. If they are in different regions, retrieve the ExtranetEndpoint value instead.
	//
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the backup.
	//
	// Valid values:
	//
	// 	- OSS: uploads the backup file to an OSS bucket.
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
