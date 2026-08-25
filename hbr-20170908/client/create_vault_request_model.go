// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVaultRequest
	GetDescription() *string
	SetEncryptType(v string) *CreateVaultRequest
	GetEncryptType() *string
	SetKmsKeyId(v string) *CreateVaultRequest
	GetKmsKeyId() *string
	SetReplication(v bool) *CreateVaultRequest
	GetReplication() *bool
	SetVaultName(v string) *CreateVaultRequest
	GetVaultName() *string
	SetVaultRegionId(v string) *CreateVaultRequest
	GetVaultRegionId() *string
	SetVaultStorageClass(v string) *CreateVaultRequest
	GetVaultStorageClass() *string
	SetVaultType(v string) *CreateVaultRequest
	GetVaultType() *string
	SetWormEnabled(v bool) *CreateVaultRequest
	GetWormEnabled() *bool
}

type CreateVaultRequest struct {
	// The description of the backup vault. The description can be 0 to 255 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption type of the source data. This parameter is valid only if you set VaultType to STANDARD or OTS_BACKUP. Valid values:
	//
	// - **HBR_PRIVATE**: The backup vault is encrypted using the built-in encryption method of Cloud Backup.
	//
	// - **KMS**: The backup vault is encrypted using a customer master key (CMK) from Key Management Service (KMS).
	//
	// example:
	//
	// KMS
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The ID or alias of the KMS key. This parameter is required only if you set EncryptType to KMS.
	//
	// example:
	//
	// alias/yzs-hhht
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// Specifies whether to create a replication vault.
	//
	// example:
	//
	// true
	Replication *bool `json:"Replication,omitempty" xml:"Replication,omitempty"`
	// The name of the backup vault. The name must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// backupvaultname
	VaultName *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	// The region ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	VaultRegionId *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	// The storage class of the backup vault.
	//
	// - **STANDARD**: Standard.
	//
	// - **ARCHIVE**: This value is deprecated.
	//
	// - **COLD_ARCHIVE**: This value is deprecated.
	//
	// - **IA**: This value is deprecated.
	//
	// example:
	//
	// STANDARD
	VaultStorageClass *string `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
	// The type of the backup vault. Valid values:
	//
	// - **STANDARD**: a standard backup vault.
	//
	// - **OTS_BACKUP**: a Tablestore backup vault.
	//
	// example:
	//
	// STANDARD
	VaultType *string `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
	// Specifies whether to enable backup locking.
	//
	// example:
	//
	// false
	WormEnabled *bool `json:"WormEnabled,omitempty" xml:"WormEnabled,omitempty"`
}

func (s CreateVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultRequest) GoString() string {
	return s.String()
}

func (s *CreateVaultRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVaultRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *CreateVaultRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateVaultRequest) GetReplication() *bool {
	return s.Replication
}

func (s *CreateVaultRequest) GetVaultName() *string {
	return s.VaultName
}

func (s *CreateVaultRequest) GetVaultRegionId() *string {
	return s.VaultRegionId
}

func (s *CreateVaultRequest) GetVaultStorageClass() *string {
	return s.VaultStorageClass
}

func (s *CreateVaultRequest) GetVaultType() *string {
	return s.VaultType
}

func (s *CreateVaultRequest) GetWormEnabled() *bool {
	return s.WormEnabled
}

func (s *CreateVaultRequest) SetDescription(v string) *CreateVaultRequest {
	s.Description = &v
	return s
}

func (s *CreateVaultRequest) SetEncryptType(v string) *CreateVaultRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateVaultRequest) SetKmsKeyId(v string) *CreateVaultRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateVaultRequest) SetReplication(v bool) *CreateVaultRequest {
	s.Replication = &v
	return s
}

func (s *CreateVaultRequest) SetVaultName(v string) *CreateVaultRequest {
	s.VaultName = &v
	return s
}

func (s *CreateVaultRequest) SetVaultRegionId(v string) *CreateVaultRequest {
	s.VaultRegionId = &v
	return s
}

func (s *CreateVaultRequest) SetVaultStorageClass(v string) *CreateVaultRequest {
	s.VaultStorageClass = &v
	return s
}

func (s *CreateVaultRequest) SetVaultType(v string) *CreateVaultRequest {
	s.VaultType = &v
	return s
}

func (s *CreateVaultRequest) SetWormEnabled(v bool) *CreateVaultRequest {
	s.WormEnabled = &v
	return s
}

func (s *CreateVaultRequest) Validate() error {
	return dara.Validate(s)
}
