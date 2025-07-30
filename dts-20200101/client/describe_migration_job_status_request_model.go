// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMigrationJobStatusRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeMigrationJobStatusRequest
	GetClientToken() *string
	SetMigrationJobId(v string) *DescribeMigrationJobStatusRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *DescribeMigrationJobStatusRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeMigrationJobStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeMigrationJobStatusRequest
	GetResourceGroupId() *string
}

type DescribeMigrationJobStatusRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsz2v12jfo309****
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

func (s DescribeMigrationJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMigrationJobStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMigrationJobStatusRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DescribeMigrationJobStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeMigrationJobStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMigrationJobStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrationJobStatusRequest) SetAccountId(v string) *DescribeMigrationJobStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetClientToken(v string) *DescribeMigrationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetMigrationJobId(v string) *DescribeMigrationJobStatusRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetOwnerId(v string) *DescribeMigrationJobStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetRegionId(v string) *DescribeMigrationJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetResourceGroupId(v string) *DescribeMigrationJobStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
