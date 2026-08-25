// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateReplicationVaultRequest
	GetDescription() *string
	SetEncryptType(v string) *CreateReplicationVaultRequest
	GetEncryptType() *string
	SetKmsKeyId(v string) *CreateReplicationVaultRequest
	GetKmsKeyId() *string
	SetRedundancyType(v string) *CreateReplicationVaultRequest
	GetRedundancyType() *string
	SetReplicationSourceRegionId(v string) *CreateReplicationVaultRequest
	GetReplicationSourceRegionId() *string
	SetReplicationSourceVaultId(v string) *CreateReplicationVaultRequest
	GetReplicationSourceVaultId() *string
	SetVaultName(v string) *CreateReplicationVaultRequest
	GetVaultName() *string
	SetVaultRegionId(v string) *CreateReplicationVaultRequest
	GetVaultRegionId() *string
	SetVaultStorageClass(v string) *CreateReplicationVaultRequest
	GetVaultStorageClass() *string
}

type CreateReplicationVaultRequest struct {
	// The description of the backup vault. The description must be 0 to 255 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption type of the replication target vault. This parameter is valid only when VaultType is set to STANDARD. The encryption type must be the same as that of the source backup repository. Valid values:
	//
	// - **HBR_PRIVATE**: fully managed by Cloud Backup. The built-in secret key encryption method of the backup service is used.
	//
	// - **KMS**: uses a custom key from Alibaba Cloud Key Management Service (KMS) for encryption.
	//
	// example:
	//
	// HBR_PRIVATE
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The custom key or alias from Alibaba Cloud KMS. This parameter is required only when EncryptType is set to KMS.
	//
	// example:
	//
	// alias/test
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The data redundancy storage method of the backup vault. Valid values:
	//
	// - LRS: locally redundant storage (LRS). The data redundancy storage mechanism is used to store redundant copies of each object on multiple devices across multiple facilities within the same zone, ensuring data durability and availability in the event of hardware failure.
	//
	// - ZRS: zone-redundant storage (ZRS). The multi-zone mechanism is used to distribute user data across three zones in the same region. If one zone becomes unavailable, the data can still be accessed normally.
	//
	// example:
	//
	// LRS
	RedundancyType *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	// The region ID of the source vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ReplicationSourceRegionId *string `json:"ReplicationSourceRegionId,omitempty" xml:"ReplicationSourceRegionId,omitempty"`
	// The ID of the source vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-*********************
	ReplicationSourceVaultId *string `json:"ReplicationSourceVaultId,omitempty" xml:"ReplicationSourceVaultId,omitempty"`
	// The name of the backup vault. The name must be 1 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// mirrorvaultname
	VaultName *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	// The region ID of the backup vault.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	VaultRegionId *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	// The storage class of the backup vault. The value can only be **STANDARD**, which indicates standard storage.
	//
	// example:
	//
	// STANDARD
	VaultStorageClass *string `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
}

func (s CreateReplicationVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationVaultRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateReplicationVaultRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *CreateReplicationVaultRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateReplicationVaultRequest) GetRedundancyType() *string {
	return s.RedundancyType
}

func (s *CreateReplicationVaultRequest) GetReplicationSourceRegionId() *string {
	return s.ReplicationSourceRegionId
}

func (s *CreateReplicationVaultRequest) GetReplicationSourceVaultId() *string {
	return s.ReplicationSourceVaultId
}

func (s *CreateReplicationVaultRequest) GetVaultName() *string {
	return s.VaultName
}

func (s *CreateReplicationVaultRequest) GetVaultRegionId() *string {
	return s.VaultRegionId
}

func (s *CreateReplicationVaultRequest) GetVaultStorageClass() *string {
	return s.VaultStorageClass
}

func (s *CreateReplicationVaultRequest) SetDescription(v string) *CreateReplicationVaultRequest {
	s.Description = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetEncryptType(v string) *CreateReplicationVaultRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetKmsKeyId(v string) *CreateReplicationVaultRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetRedundancyType(v string) *CreateReplicationVaultRequest {
	s.RedundancyType = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetReplicationSourceRegionId(v string) *CreateReplicationVaultRequest {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetReplicationSourceVaultId(v string) *CreateReplicationVaultRequest {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultName(v string) *CreateReplicationVaultRequest {
	s.VaultName = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultRegionId(v string) *CreateReplicationVaultRequest {
	s.VaultRegionId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultStorageClass(v string) *CreateReplicationVaultRequest {
	s.VaultStorageClass = &v
	return s
}

func (s *CreateReplicationVaultRequest) Validate() error {
	return dara.Validate(s)
}
