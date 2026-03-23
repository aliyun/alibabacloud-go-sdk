// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossBackupMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *DescribeCrossBackupMetaListRequest
	GetBackupSetId() *string
	SetGetDbName(v string) *DescribeCrossBackupMetaListRequest
	GetGetDbName() *string
	SetOwnerId(v int64) *DescribeCrossBackupMetaListRequest
	GetOwnerId() *int64
	SetPageIndex(v string) *DescribeCrossBackupMetaListRequest
	GetPageIndex() *string
	SetPageSize(v string) *DescribeCrossBackupMetaListRequest
	GetPageSize() *string
	SetPattern(v string) *DescribeCrossBackupMetaListRequest
	GetPattern() *string
	SetRegion(v string) *DescribeCrossBackupMetaListRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeCrossBackupMetaListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossBackupMetaListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossBackupMetaListRequest
	GetResourceOwnerId() *int64
}

type DescribeCrossBackupMetaListRequest struct {
	// This parameter is required.
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	GetDbName            *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageIndex            *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pattern              *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCrossBackupMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossBackupMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossBackupMetaListRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeCrossBackupMetaListRequest) GetGetDbName() *string {
	return s.GetDbName
}

func (s *DescribeCrossBackupMetaListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossBackupMetaListRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *DescribeCrossBackupMetaListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCrossBackupMetaListRequest) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeCrossBackupMetaListRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeCrossBackupMetaListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCrossBackupMetaListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossBackupMetaListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossBackupMetaListRequest) SetBackupSetId(v string) *DescribeCrossBackupMetaListRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetGetDbName(v string) *DescribeCrossBackupMetaListRequest {
	s.GetDbName = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetOwnerId(v int64) *DescribeCrossBackupMetaListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetPageIndex(v string) *DescribeCrossBackupMetaListRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetPageSize(v string) *DescribeCrossBackupMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetPattern(v string) *DescribeCrossBackupMetaListRequest {
	s.Pattern = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetRegion(v string) *DescribeCrossBackupMetaListRequest {
	s.Region = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetResourceGroupId(v string) *DescribeCrossBackupMetaListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetResourceOwnerAccount(v string) *DescribeCrossBackupMetaListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) SetResourceOwnerId(v int64) *DescribeCrossBackupMetaListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossBackupMetaListRequest) Validate() error {
	return dara.Validate(s)
}
