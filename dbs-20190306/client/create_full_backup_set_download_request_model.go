// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullBackupSetDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetDataFormat(v string) *CreateFullBackupSetDownloadRequest
	GetBackupSetDataFormat() *string
	SetBackupSetId(v string) *CreateFullBackupSetDownloadRequest
	GetBackupSetId() *string
	SetClientToken(v string) *CreateFullBackupSetDownloadRequest
	GetClientToken() *string
	SetOwnerId(v string) *CreateFullBackupSetDownloadRequest
	GetOwnerId() *string
}

type CreateFullBackupSetDownloadRequest struct {
	// The format in which the full backup set is downloaded. Valid values:
	//
	// 	- **Native**
	//
	// 	- **SQL**
	//
	// 	- **CSV**(Default value)
	//
	// 	- **JSON**
	//
	// example:
	//
	// SQL
	BackupSetDataFormat *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	// The ID of the full backup set.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1hv****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateFullBackupSetDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFullBackupSetDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadRequest) GetBackupSetDataFormat() *string {
	return s.BackupSetDataFormat
}

func (s *CreateFullBackupSetDownloadRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateFullBackupSetDownloadRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFullBackupSetDownloadRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateFullBackupSetDownloadRequest) SetBackupSetDataFormat(v string) *CreateFullBackupSetDownloadRequest {
	s.BackupSetDataFormat = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetBackupSetId(v string) *CreateFullBackupSetDownloadRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetClientToken(v string) *CreateFullBackupSetDownloadRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetOwnerId(v string) *CreateFullBackupSetDownloadRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) Validate() error {
	return dara.Validate(s)
}
