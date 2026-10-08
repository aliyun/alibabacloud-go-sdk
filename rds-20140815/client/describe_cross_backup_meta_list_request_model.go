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
	// The cross-region backup set ID. You can call the DescribeCrossRegionBackups operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The name of the database to query. Exact match is used. The specific database name and the table names within the database are returned.
	//
	// example:
	//
	// testdb1
	GetDbName *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: greater than 0 and up to the maximum value of Integer.
	//
	// >This parameter takes effect only when it is specified together with **PageSize**.
	//
	// example:
	//
	// 1
	PageIndex *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page. Default value: **1**.
	//
	// >This parameter takes effect only when it is specified together with **PageIndex**.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database to query. Fuzzy match is used. Only the matched database names are returned, and table names are not returned.
	//
	// >You can use fuzzy match first. For example, pass in test to match testdb1 and testdb2. After you determine the target database name, use exact match by passing in **GetDbName*	- to view the specific database name and table names.
	//
	// example:
	//
	// test
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
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
