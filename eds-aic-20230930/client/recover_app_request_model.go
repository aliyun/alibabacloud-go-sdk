// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *RecoverAppRequest
	GetAndroidInstanceIdList() []*string
	SetBackupFileId(v string) *RecoverAppRequest
	GetBackupFileId() *string
	SetBackupFilePath(v string) *RecoverAppRequest
	GetBackupFilePath() *string
	SetUploadEndpoint(v string) *RecoverAppRequest
	GetUploadEndpoint() *string
	SetUploadType(v string) *RecoverAppRequest
	GetUploadType() *string
}

type RecoverAppRequest struct {
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// bf-azhps4rdyi2th****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// This parameter is required.
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	// example:
	//
	// oss-cn-shanghai-internal.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s RecoverAppRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppRequest) GoString() string {
	return s.String()
}

func (s *RecoverAppRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *RecoverAppRequest) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *RecoverAppRequest) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *RecoverAppRequest) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *RecoverAppRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *RecoverAppRequest) SetAndroidInstanceIdList(v []*string) *RecoverAppRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *RecoverAppRequest) SetBackupFileId(v string) *RecoverAppRequest {
	s.BackupFileId = &v
	return s
}

func (s *RecoverAppRequest) SetBackupFilePath(v string) *RecoverAppRequest {
	s.BackupFilePath = &v
	return s
}

func (s *RecoverAppRequest) SetUploadEndpoint(v string) *RecoverAppRequest {
	s.UploadEndpoint = &v
	return s
}

func (s *RecoverAppRequest) SetUploadType(v string) *RecoverAppRequest {
	s.UploadType = &v
	return s
}

func (s *RecoverAppRequest) Validate() error {
	return dara.Validate(s)
}
