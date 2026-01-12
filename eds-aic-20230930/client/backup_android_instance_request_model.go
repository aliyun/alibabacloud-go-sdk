// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackupAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *BackupAndroidInstanceRequest
	GetAndroidInstanceIdList() []*string
	SetBackupFileName(v string) *BackupAndroidInstanceRequest
	GetBackupFileName() *string
	SetBackupFilePath(v string) *BackupAndroidInstanceRequest
	GetBackupFilePath() *string
	SetDescription(v string) *BackupAndroidInstanceRequest
	GetDescription() *string
	SetUploadEndpoint(v string) *BackupAndroidInstanceRequest
	GetUploadEndpoint() *string
}

type BackupAndroidInstanceRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// abc
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// example:
	//
	// this is a backup android instance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
}

func (s BackupAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BackupAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *BackupAndroidInstanceRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *BackupAndroidInstanceRequest) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *BackupAndroidInstanceRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *BackupAndroidInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *BackupAndroidInstanceRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *BackupAndroidInstanceRequest) SetAndroidInstanceIdList(v []*string) *BackupAndroidInstanceRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *BackupAndroidInstanceRequest) SetBackupFileName(v string) *BackupAndroidInstanceRequest {
	s.BackupFileName = &v
	return s
}

func (s *BackupAndroidInstanceRequest) SetBackupFilePath(v string) *BackupAndroidInstanceRequest {
	s.BackupFilePath = &v
	return s
}

func (s *BackupAndroidInstanceRequest) SetDescription(v string) *BackupAndroidInstanceRequest {
	s.Description = &v
	return s
}

func (s *BackupAndroidInstanceRequest) SetUploadEndpoint(v string) *BackupAndroidInstanceRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *BackupAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
