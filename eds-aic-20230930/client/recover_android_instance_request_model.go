// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *RecoverAndroidInstanceRequest
	GetAndroidInstanceIdList() []*string
	SetBackupFileId(v string) *RecoverAndroidInstanceRequest
	GetBackupFileId() *string
	SetBackupFilePath(v string) *RecoverAndroidInstanceRequest
	GetBackupFilePath() *string
	SetExcludeAppList(v []*string) *RecoverAndroidInstanceRequest
	GetExcludeAppList() []*string
	SetUploadEndpoint(v string) *RecoverAndroidInstanceRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *RecoverAndroidInstanceRequest
	GetUploadType() *string
}

type RecoverAndroidInstanceRequest struct {
	// The list of instance IDs.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The OSS address to which the backup file is uploaded.
	//
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The list of app package names excluded from the full-device restore. Excluded apps are not uninstalled and their APKs are not installed, which is compatible with legacy backups that contain APKs. Only user data is replaced. Empty shells must be installed by the customer by using the pm install command. If this parameter is not specified, no apps are excluded.
	ExcludeAppList []*string `json:"ExcludeAppList,omitempty" xml:"ExcludeAppList,omitempty" type:"Repeated"`
	// The backup region.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The upload type.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s RecoverAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *RecoverAndroidInstanceRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *RecoverAndroidInstanceRequest) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *RecoverAndroidInstanceRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *RecoverAndroidInstanceRequest) GetExcludeAppList() []*string {
	return s.ExcludeAppList
}

func (s *RecoverAndroidInstanceRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *RecoverAndroidInstanceRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *RecoverAndroidInstanceRequest) SetAndroidInstanceIdList(v []*string) *RecoverAndroidInstanceRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *RecoverAndroidInstanceRequest) SetBackupFileId(v string) *RecoverAndroidInstanceRequest {
	s.BackupFileId = &v
	return s
}

func (s *RecoverAndroidInstanceRequest) SetBackupFilePath(v string) *RecoverAndroidInstanceRequest {
	s.BackupFilePath = &v
	return s
}

func (s *RecoverAndroidInstanceRequest) SetExcludeAppList(v []*string) *RecoverAndroidInstanceRequest {
	s.ExcludeAppList = v
	return s
}

func (s *RecoverAndroidInstanceRequest) SetUploadEndpoint(v string) *RecoverAndroidInstanceRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *RecoverAndroidInstanceRequest) SetUploadType(v string) *RecoverAndroidInstanceRequest {
	s.UploadType = &v
	return s
}

func (s *RecoverAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
