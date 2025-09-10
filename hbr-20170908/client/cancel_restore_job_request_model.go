// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestoreJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *CancelRestoreJobRequest
	GetEdition() *string
	SetRestoreId(v string) *CancelRestoreJobRequest
	GetRestoreId() *string
	SetVaultId(v string) *CancelRestoreJobRequest
	GetVaultId() *string
}

type CancelRestoreJobRequest struct {
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the restore job.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-*********************
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CancelRestoreJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobRequest) GetEdition() *string {
	return s.Edition
}

func (s *CancelRestoreJobRequest) GetRestoreId() *string {
	return s.RestoreId
}

func (s *CancelRestoreJobRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CancelRestoreJobRequest) SetEdition(v string) *CancelRestoreJobRequest {
	s.Edition = &v
	return s
}

func (s *CancelRestoreJobRequest) SetRestoreId(v string) *CancelRestoreJobRequest {
	s.RestoreId = &v
	return s
}

func (s *CancelRestoreJobRequest) SetVaultId(v string) *CancelRestoreJobRequest {
	s.VaultId = &v
	return s
}

func (s *CancelRestoreJobRequest) Validate() error {
	return dara.Validate(s)
}
