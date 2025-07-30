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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The specification of the data migration instance. Valid values: **small**, **medium**, **large**, **xlarge**, and **2xlarge**.
	//
	// >
	//
	// 	- For more information about the test performance of each specification, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html).
	//
	// 	- For more information about the pricing of data migration instances, see [Pricing](https://help.aliyun.com/document_detail/117780.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2xlarge
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data migration instance resides. The region ID of the data migration instance is the same as that of the destination database. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region where the data migration instance resides. You do not need to specify this parameter because this parameter will be removed in the future.
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
