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
	SetUploadEndpoint(v string) *RecoverAndroidInstanceRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *RecoverAndroidInstanceRequest
	GetUploadType() *string
}

type RecoverAndroidInstanceRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
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
