// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIncrementBackupSetDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetDataFormat(v string) *CreateIncrementBackupSetDownloadRequest
	GetBackupSetDataFormat() *string
	SetBackupSetId(v string) *CreateIncrementBackupSetDownloadRequest
	GetBackupSetId() *string
	SetBackupSetName(v string) *CreateIncrementBackupSetDownloadRequest
	GetBackupSetName() *string
	SetClientToken(v string) *CreateIncrementBackupSetDownloadRequest
	GetClientToken() *string
	SetOwnerId(v string) *CreateIncrementBackupSetDownloadRequest
	GetOwnerId() *string
}

type CreateIncrementBackupSetDownloadRequest struct {
	// The format in which the incremental backup set is downloaded. Valid values:
	//
	// 	- **Native**
	//
	// 	- **SQL**
	//
	// 	- **CSV**
	//
	// 	- **JSON**
	//
	// > Default value: Native.
	//
	// example:
	//
	// Native
	BackupSetDataFormat *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	// The ID of the incremental backup task. To obtain the task ID, you can call the [DescribeIncrementBackupList](https://help.aliyun.com/document_detail/2869833.html) operation and view the value of the **BackupSetJobId*	- parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01e****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The ID of the incremental backup set. To obtain the backup set ID, you can call the [DescribeIncrementBackupList](https://help.aliyun.com/document_detail/2869833.html) operation and view the value of the **BackupSetId*	- parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BackupSetName *string `json:"BackupSetName,omitempty" xml:"BackupSetName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateIncrementBackupSetDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadRequest) GetBackupSetDataFormat() *string {
	return s.BackupSetDataFormat
}

func (s *CreateIncrementBackupSetDownloadRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateIncrementBackupSetDownloadRequest) GetBackupSetName() *string {
	return s.BackupSetName
}

func (s *CreateIncrementBackupSetDownloadRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIncrementBackupSetDownloadRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetDataFormat(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetDataFormat = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetId(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetName(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetName = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetClientToken(v string) *CreateIncrementBackupSetDownloadRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetOwnerId(v string) *CreateIncrementBackupSetDownloadRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) Validate() error {
	return dara.Validate(s)
}
