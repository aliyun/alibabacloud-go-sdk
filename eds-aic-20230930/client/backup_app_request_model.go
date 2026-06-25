// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *BackupAppRequest
	GetAndroidInstanceIdList() []*string
	SetBackupFileName(v string) *BackupAppRequest
	GetBackupFileName() *string
	SetBackupFilePath(v string) *BackupAppRequest
	GetBackupFilePath() *string
	SetDescription(v string) *BackupAppRequest
	GetDescription() *string
	SetSourceAppList(v []*string) *BackupAppRequest
	GetSourceAppList() []*string
	SetUploadEndpoint(v string) *BackupAppRequest
	GetUploadEndpoint() *string
}

type BackupAppRequest struct {
	// A list of instance IDs.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The name of the backup file.
	//
	// example:
	//
	// MyBackup
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The URL of the backup file.
	//
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// The description of the application backup.
	//
	// example:
	//
	// this is a backup app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A list of package names for the applications to back up.
	//
	// This parameter is required.
	SourceAppList []*string `json:"SourceAppList,omitempty" xml:"SourceAppList,omitempty" type:"Repeated"`
	// Specifies the region where the backup is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
}

func (s BackupAppRequest) String() string {
	return dara.Prettify(s)
}

func (s BackupAppRequest) GoString() string {
	return s.String()
}

func (s *BackupAppRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *BackupAppRequest) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupAppRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *BackupAppRequest) GetDescription() *string {
	return s.Description
}

func (s *BackupAppRequest) GetSourceAppList() []*string {
	return s.SourceAppList
}

func (s *BackupAppRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *BackupAppRequest) SetAndroidInstanceIdList(v []*string) *BackupAppRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *BackupAppRequest) SetBackupFileName(v string) *BackupAppRequest {
	s.BackupFileName = &v
	return s
}

func (s *BackupAppRequest) SetBackupFilePath(v string) *BackupAppRequest {
	s.BackupFilePath = &v
	return s
}

func (s *BackupAppRequest) SetDescription(v string) *BackupAppRequest {
	s.Description = &v
	return s
}

func (s *BackupAppRequest) SetSourceAppList(v []*string) *BackupAppRequest {
	s.SourceAppList = v
	return s
}

func (s *BackupAppRequest) SetUploadEndpoint(v string) *BackupAppRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *BackupAppRequest) Validate() error {
	return dara.Validate(s)
}
