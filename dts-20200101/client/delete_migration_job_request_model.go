// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteMigrationJobRequest
	GetAccountId() *string
	SetMigrationJobId(v string) *DeleteMigrationJobRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *DeleteMigrationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteMigrationJobRequest
	GetResourceGroupId() *string
}

type DeleteMigrationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query all data migration instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsyiwe9b0gp2p****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMigrationJobRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DeleteMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteMigrationJobRequest) SetAccountId(v string) *DeleteMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetMigrationJobId(v string) *DeleteMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetOwnerId(v string) *DeleteMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetRegionId(v string) *DeleteMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetResourceGroupId(v string) *DeleteMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}
