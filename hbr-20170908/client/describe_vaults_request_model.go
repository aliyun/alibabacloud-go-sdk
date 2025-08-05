// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVaultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVaultsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeVaultsRequest
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeVaultsRequest
	GetStatus() *string
	SetTag(v []*DescribeVaultsRequestTag) *DescribeVaultsRequest
	GetTag() []*DescribeVaultsRequestTag
	SetVaultId(v string) *DescribeVaultsRequest
	GetVaultId() *string
	SetVaultName(v string) *DescribeVaultsRequest
	GetVaultName() *string
	SetVaultRegionId(v string) *DescribeVaultsRequest
	GetVaultRegionId() *string
	SetVaultType(v string) *DescribeVaultsRequest
	GetVaultType() *string
}

type DescribeVaultsRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-*********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the backup vault. Valid values:
	//
	// 	- **UNKNOWN**: The backup vault is in an unknown state.
	//
	// 	- **INITIALIZING**: The backup vault is being initialized.
	//
	// 	- **CREATED**: The backup vault is created.
	//
	// 	- **ERROR**: An error occurs on the backup vault.
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Tag information. Supports up to 20 tags.
	//
	// example:
	//
	// 6a745bceffb042959b3b5206d6f12ad1
	Tag []*DescribeVaultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Backup vault ID.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	// The name of the backup vault. The name must be 1 to 64 characters in length.
	//
	// example:
	//
	// vaultname
	VaultName *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	// The region ID to which the backup vault belongs.
	//
	// example:
	//
	// cn-shanghai
	VaultRegionId *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	// Backup repository type. The values are as follows:
	//
	// - **STANDARD**: Represents a standard repository, which can be used for ECS file backups, OSS backups, NAS backups, etc.
	//
	// - **OTS_BACKUP**: Represents a TableStore repository, which is only used for TableStore backups, and TableStore must use this type of repository.
	//
	// example:
	//
	// STANDARD
	VaultType *string `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
}

func (s DescribeVaultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVaultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVaultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVaultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVaultsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVaultsRequest) GetTag() []*DescribeVaultsRequestTag {
	return s.Tag
}

func (s *DescribeVaultsRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeVaultsRequest) GetVaultName() *string {
	return s.VaultName
}

func (s *DescribeVaultsRequest) GetVaultRegionId() *string {
	return s.VaultRegionId
}

func (s *DescribeVaultsRequest) GetVaultType() *string {
	return s.VaultType
}

func (s *DescribeVaultsRequest) SetPageNumber(v int32) *DescribeVaultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVaultsRequest) SetPageSize(v int32) *DescribeVaultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVaultsRequest) SetResourceGroupId(v string) *DescribeVaultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVaultsRequest) SetStatus(v string) *DescribeVaultsRequest {
	s.Status = &v
	return s
}

func (s *DescribeVaultsRequest) SetTag(v []*DescribeVaultsRequestTag) *DescribeVaultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVaultsRequest) SetVaultId(v string) *DescribeVaultsRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeVaultsRequest) SetVaultName(v string) *DescribeVaultsRequest {
	s.VaultName = &v
	return s
}

func (s *DescribeVaultsRequest) SetVaultRegionId(v string) *DescribeVaultsRequest {
	s.VaultRegionId = &v
	return s
}

func (s *DescribeVaultsRequest) SetVaultType(v string) *DescribeVaultsRequest {
	s.VaultType = &v
	return s
}

func (s *DescribeVaultsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Value of the tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVaultsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVaultsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVaultsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVaultsRequestTag) SetKey(v string) *DescribeVaultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVaultsRequestTag) SetValue(v string) *DescribeVaultsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVaultsRequestTag) Validate() error {
	return dara.Validate(s)
}
