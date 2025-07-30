// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *StopMigrationJobRequest
	GetAccountId() *string
	SetClientToken(v string) *StopMigrationJobRequest
	GetClientToken() *string
	SetMigrationJobId(v string) *StopMigrationJobRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *StopMigrationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *StopMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StopMigrationJobRequest
	GetResourceGroupId() *string
}

type StopMigrationJobRequest struct {
	// The IDoftheAlibabaCloudaccount. Youdonotneed to specify this parameter because this parameter will be removed in the future.
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
	// dtsb2c11sxpi3j****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StopMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StopMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *StopMigrationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopMigrationJobRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *StopMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StopMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StopMigrationJobRequest) SetAccountId(v string) *StopMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *StopMigrationJobRequest) SetClientToken(v string) *StopMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StopMigrationJobRequest) SetMigrationJobId(v string) *StopMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StopMigrationJobRequest) SetOwnerId(v string) *StopMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMigrationJobRequest) SetRegionId(v string) *StopMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopMigrationJobRequest) SetResourceGroupId(v string) *StopMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StopMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}
