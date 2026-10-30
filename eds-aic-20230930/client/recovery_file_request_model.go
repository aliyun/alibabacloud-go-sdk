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
	SetExcludeAppList(v []*string) *RecoveryFileRequest
	GetExcludeAppList() []*string
	SetUploadEndpoint(v string) *RecoveryFileRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *RecoveryFileRequest
	GetUploadType() *string
}

type RecoveryFileRequest struct {
	// The list of instance IDs.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to back up the entire device.
	//
	// example:
	//
	// false
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The OSS URL to which the backup file is uploaded.
	//
	// <notice>Call the describeBuckets operation to obtain the bucketName, and select ossObjectList to obtain the key. The key is the path of the backup file. Concatenate them into "oss://${bucketName}/${key}".</notice>
	BackupFilePath *string   `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	ExcludeAppList []*string `json:"ExcludeAppList,omitempty" xml:"ExcludeAppList,omitempty" type:"Repeated"`
	// The domain name of the backup address.
	//
	// <notice>Call the DescribeBuckets operation to obtain the bucket. If the cloud phone and the bucket are in the same region, obtain the IntranetEndpoint field. If they are in different regions, obtain the ExtranetEndpoint field.</notice>
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The backup type.
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

func (s *RecoveryFileRequest) GetExcludeAppList() []*string {
	return s.ExcludeAppList
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

func (s *RecoveryFileRequest) SetExcludeAppList(v []*string) *RecoveryFileRequest {
	s.ExcludeAppList = v
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
