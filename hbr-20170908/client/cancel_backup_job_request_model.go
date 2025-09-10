// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBackupJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *CancelBackupJobRequest
	GetEdition() *string
	SetJobId(v string) *CancelBackupJobRequest
	GetJobId() *string
	SetVaultId(v string) *CancelBackupJobRequest
	GetVaultId() *string
}

type CancelBackupJobRequest struct {
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the backup job.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-******************************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*****************************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CancelBackupJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelBackupJobRequest) GoString() string {
	return s.String()
}

func (s *CancelBackupJobRequest) GetEdition() *string {
	return s.Edition
}

func (s *CancelBackupJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelBackupJobRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CancelBackupJobRequest) SetEdition(v string) *CancelBackupJobRequest {
	s.Edition = &v
	return s
}

func (s *CancelBackupJobRequest) SetJobId(v string) *CancelBackupJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelBackupJobRequest) SetVaultId(v string) *CancelBackupJobRequest {
	s.VaultId = &v
	return s
}

func (s *CancelBackupJobRequest) Validate() error {
	return dara.Validate(s)
}
