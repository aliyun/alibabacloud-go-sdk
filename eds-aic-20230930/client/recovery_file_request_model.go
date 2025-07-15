// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *RecoveryFileRequest
	GetAndroidInstanceIdList() []*string
	SetBackupAll(v bool) *RecoveryFileRequest
	GetBackupAll() *bool
	SetBackupFileId(v string) *RecoveryFileRequest
	GetBackupFileId() *string
	SetBackupFilePath(v string) *RecoveryFileRequest
	GetBackupFilePath() *string
	SetUploadEndpoint(v string) *RecoveryFileRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *RecoveryFileRequest
	GetUploadType() *string
}

type RecoveryFileRequest struct {
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
	// The ID of the backup file.
	//
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The OSS path to which the backup file is uploaded.
	//
	// >  When calling the describeBuckets operation to retrieve a bucket name, you must also call the ossObjectList operation to obtain the object key. Combine these to form the full path: oss://${bucketName}/${key}.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The endpoint of the OSS bucket that stores the backup file.
	//
	// > : When calling the DescribeBuckets operation to query buckets, retrieve the IntranetEndpoint value if the cloud phone and the OSS bucket are in the same region. If they are in different regions, retrieve the ExtranetEndpoint value instead.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The type of the backup.
	//
	// Valid values:
	//
	// 	- OSS: backup files are stored in OSS buckets.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s RecoveryFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoveryFileRequest) GoString() string {
	return s.String()
}

func (s *RecoveryFileRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *RecoveryFileRequest) GetBackupAll() *bool {
	return s.BackupAll
}

func (s *RecoveryFileRequest) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *RecoveryFileRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *RecoveryFileRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *RecoveryFileRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *RecoveryFileRequest) SetAndroidInstanceIdList(v []*string) *RecoveryFileRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *RecoveryFileRequest) SetBackupAll(v bool) *RecoveryFileRequest {
	s.BackupAll = &v
	return s
}

func (s *RecoveryFileRequest) SetBackupFileId(v string) *RecoveryFileRequest {
	s.BackupFileId = &v
	return s
}

func (s *RecoveryFileRequest) SetBackupFilePath(v string) *RecoveryFileRequest {
	s.BackupFilePath = &v
	return s
}

func (s *RecoveryFileRequest) SetUploadEndpoint(v string) *RecoveryFileRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *RecoveryFileRequest) SetUploadType(v string) *RecoveryFileRequest {
	s.UploadType = &v
	return s
}

func (s *RecoveryFileRequest) Validate() error {
	return dara.Validate(s)
}
