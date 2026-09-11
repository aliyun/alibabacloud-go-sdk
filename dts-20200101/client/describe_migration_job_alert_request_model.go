// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMigrationJobAlertRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeMigrationJobAlertRequest
	GetClientToken() *string
	SetMigrationJobId(v string) *DescribeMigrationJobAlertRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *DescribeMigrationJobAlertRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeMigrationJobAlertRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeMigrationJobAlertRequest
	GetResourceGroupId() *string
}

type DescribeMigrationJobAlertRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter is about to be discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data migration instance. You can call the DescribeMigrationJobs operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtslb9113qq11n****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeMigrationJobAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMigrationJobAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMigrationJobAlertRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DescribeMigrationJobAlertRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeMigrationJobAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMigrationJobAlertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrationJobAlertRequest) SetAccountId(v string) *DescribeMigrationJobAlertRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetClientToken(v string) *DescribeMigrationJobAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetMigrationJobId(v string) *DescribeMigrationJobAlertRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetOwnerId(v string) *DescribeMigrationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetRegionId(v string) *DescribeMigrationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetResourceGroupId(v string) *DescribeMigrationJobAlertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) Validate() error {
	return dara.Validate(s)
}
