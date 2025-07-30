// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *SuspendMigrationJobRequest
	GetAccountId() *string
	SetClientToken(v string) *SuspendMigrationJobRequest
	GetClientToken() *string
	SetMigrationJobId(v string) *SuspendMigrationJobRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *SuspendMigrationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *SuspendMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SuspendMigrationJobRequest
	GetResourceGroupId() *string
}

type SuspendMigrationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query all data migration instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsj1x11y51g3b****
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
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s SuspendMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *SuspendMigrationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SuspendMigrationJobRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *SuspendMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SuspendMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SuspendMigrationJobRequest) SetAccountId(v string) *SuspendMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetClientToken(v string) *SuspendMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetMigrationJobId(v string) *SuspendMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetOwnerId(v string) *SuspendMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetRegionId(v string) *SuspendMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetResourceGroupId(v string) *SuspendMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SuspendMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}
