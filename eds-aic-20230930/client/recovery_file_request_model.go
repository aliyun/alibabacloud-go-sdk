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
	SetCustomScripts(v []*RecoveryFileRequestCustomScripts) *RecoveryFileRequest
	GetCustomScripts() []*RecoveryFileRequestCustomScripts
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
	// Specifies whether it is a full-device backup.
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
	// The OSS address to which the backup file is uploaded.
	//
	// <notice>Call the describeBuckets operation to obtain the bucketName, and select ossObjectList to obtain the key. The key is the path of the backup file. Concatenate them into "oss://${bucketName}/${key}".</notice>
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// example:
	//
	// [{"OssBucket":"my-bucket","OssObjectKey":"script/recover.sh","OssEndpoint":"oss-cn-hangzhou-internal.aliyuncs.com","OssRegion":"cn-hangzhou"}]
	CustomScripts []*RecoveryFileRequestCustomScripts `json:"CustomScripts,omitempty" xml:"CustomScripts,omitempty" type:"Repeated"`
	// The list of excluded app package names for full-device restore. Matched apps are not uninstalled, and their APKs are not installed (compatible with legacy backups that contain APKs). Only user data is replaced. The empty shell must be installed by the customer using pm install. If this parameter is not specified, no apps are excluded.
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

func (s *RecoveryFileRequest) GetCustomScripts() []*RecoveryFileRequestCustomScripts {
	return s.CustomScripts
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

func (s *RecoveryFileRequest) SetCustomScripts(v []*RecoveryFileRequestCustomScripts) *RecoveryFileRequest {
	s.CustomScripts = v
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
	if s.CustomScripts != nil {
		for _, item := range s.CustomScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecoveryFileRequestCustomScripts struct {
	// example:
	//
	// my-bucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// example:
	//
	// script/recover.sh
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
	// example:
	//
	// cn-hangzhou
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// example:
	//
	// ["acp-0n4k52uiyon472rq0"]
	TestInstanceIds []*string `json:"TestInstanceIds,omitempty" xml:"TestInstanceIds,omitempty" type:"Repeated"`
}

func (s RecoveryFileRequestCustomScripts) String() string {
	return dara.Prettify(s)
}

func (s RecoveryFileRequestCustomScripts) GoString() string {
	return s.String()
}

func (s *RecoveryFileRequestCustomScripts) GetOssBucket() *string {
	return s.OssBucket
}

func (s *RecoveryFileRequestCustomScripts) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *RecoveryFileRequestCustomScripts) GetOssObjectKey() *string {
	return s.OssObjectKey
}

func (s *RecoveryFileRequestCustomScripts) GetOssRegion() *string {
	return s.OssRegion
}

func (s *RecoveryFileRequestCustomScripts) GetTestInstanceIds() []*string {
	return s.TestInstanceIds
}

func (s *RecoveryFileRequestCustomScripts) SetOssBucket(v string) *RecoveryFileRequestCustomScripts {
	s.OssBucket = &v
	return s
}

func (s *RecoveryFileRequestCustomScripts) SetOssEndpoint(v string) *RecoveryFileRequestCustomScripts {
	s.OssEndpoint = &v
	return s
}

func (s *RecoveryFileRequestCustomScripts) SetOssObjectKey(v string) *RecoveryFileRequestCustomScripts {
	s.OssObjectKey = &v
	return s
}

func (s *RecoveryFileRequestCustomScripts) SetOssRegion(v string) *RecoveryFileRequestCustomScripts {
	s.OssRegion = &v
	return s
}

func (s *RecoveryFileRequestCustomScripts) SetTestInstanceIds(v []*string) *RecoveryFileRequestCustomScripts {
	s.TestInstanceIds = v
	return s
}

func (s *RecoveryFileRequestCustomScripts) Validate() error {
	return dara.Validate(s)
}
