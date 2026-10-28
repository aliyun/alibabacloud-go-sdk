// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateMigrationJobRequest
	GetAccountId() *string
	SetClientToken(v string) *CreateMigrationJobRequest
	GetClientToken() *string
	SetMigrationJobClass(v string) *CreateMigrationJobRequest
	GetMigrationJobClass() *string
	SetOwnerId(v string) *CreateMigrationJobRequest
	GetOwnerId() *string
	SetRegion(v string) *CreateMigrationJobRequest
	GetRegion() *string
	SetRegionId(v string) *CreateMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateMigrationJobRequest
	GetResourceGroupId() *string
}

type CreateMigrationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Ensures the idempotence of the request. You can generate the value on your client to ensure that it is unique across different requests. The **ClientToken*	- parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The specifications of the data migration instance. Valid values: **small**, **medium**, **large**, **xlarge**, and **2xlarge**.
	//
	// > - For the test performance of each specification, see [Data migration specifications](https://help.aliyun.com/document_detail/26606.html).
	//
	// - For instance specifications and pricing, see [Pricing](https://help.aliyun.com/document_detail/117780.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2xlarge
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the data migration instance resides. This is the region where the destination database instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region where the data migration instance resides. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateMigrationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMigrationJobRequest) GetMigrationJobClass() *string {
	return s.MigrationJobClass
}

func (s *CreateMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateMigrationJobRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateMigrationJobRequest) SetAccountId(v string) *CreateMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *CreateMigrationJobRequest) SetClientToken(v string) *CreateMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMigrationJobRequest) SetMigrationJobClass(v string) *CreateMigrationJobRequest {
	s.MigrationJobClass = &v
	return s
}

func (s *CreateMigrationJobRequest) SetOwnerId(v string) *CreateMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMigrationJobRequest) SetRegion(v string) *CreateMigrationJobRequest {
	s.Region = &v
	return s
}

func (s *CreateMigrationJobRequest) SetRegionId(v string) *CreateMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMigrationJobRequest) SetResourceGroupId(v string) *CreateMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}
