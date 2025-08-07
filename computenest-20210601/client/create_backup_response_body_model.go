// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CreateBackupResponseBody
	GetBackupId() *string
	SetDescription(v string) *CreateBackupResponseBody
	GetDescription() *string
	SetRequestId(v string) *CreateBackupResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateBackupResponseBody
	GetStatus() *string
}

type CreateBackupResponseBody struct {
	// Backup ID
	//
	// example:
	//
	// backup-cad4a85ff5e340388b93
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// Description
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8DC02957-A0FC-5AB2-8C54-496B636EAF12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Backup status
	//
	// - Creating: In progress
	//
	// - Created: Success
	//
	// - CreateFailed: Failed
	//
	// - Deleting: In progress
	//
	// - Deleted: Success
	//
	// - DeleteFailed: Failed
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateBackupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateBackupResponseBody) SetBackupId(v string) *CreateBackupResponseBody {
	s.BackupId = &v
	return s
}

func (s *CreateBackupResponseBody) SetDescription(v string) *CreateBackupResponseBody {
	s.Description = &v
	return s
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) SetStatus(v string) *CreateBackupResponseBody {
	s.Status = &v
	return s
}

func (s *CreateBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
