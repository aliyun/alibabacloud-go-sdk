// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *StartMigrationJobRequest
	GetAccountId() *string
	SetMigrationJobId(v string) *StartMigrationJobRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *StartMigrationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *StartMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartMigrationJobRequest
	GetResourceGroupId() *string
}

type StartMigrationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter is about to be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Instance ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtss0611o8vv90****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StartMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StartMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *StartMigrationJobRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *StartMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StartMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartMigrationJobRequest) SetAccountId(v string) *StartMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *StartMigrationJobRequest) SetMigrationJobId(v string) *StartMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StartMigrationJobRequest) SetOwnerId(v string) *StartMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMigrationJobRequest) SetRegionId(v string) *StartMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartMigrationJobRequest) SetResourceGroupId(v string) *StartMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}
