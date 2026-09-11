// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationMode(v *DescribeMigrationJobDetailRequestMigrationMode) *DescribeMigrationJobDetailRequest
	GetMigrationMode() *DescribeMigrationJobDetailRequestMigrationMode
	SetAccountId(v string) *DescribeMigrationJobDetailRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeMigrationJobDetailRequest
	GetClientToken() *string
	SetMigrationJobId(v string) *DescribeMigrationJobDetailRequest
	GetMigrationJobId() *string
	SetOwnerId(v string) *DescribeMigrationJobDetailRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeMigrationJobDetailRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeMigrationJobDetailRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMigrationJobDetailRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeMigrationJobDetailRequest
	GetResourceGroupId() *string
}

type DescribeMigrationJobDetailRequest struct {
	MigrationMode *DescribeMigrationJobDetailRequestMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a value from your client to ensure that the value is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Instance ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsta7w132u12h****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer greater than **0*	- and must not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records per page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s DescribeMigrationJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailRequest) GetMigrationMode() *DescribeMigrationJobDetailRequestMigrationMode {
	return s.MigrationMode
}

func (s *DescribeMigrationJobDetailRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMigrationJobDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMigrationJobDetailRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DescribeMigrationJobDetailRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeMigrationJobDetailRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeMigrationJobDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMigrationJobDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMigrationJobDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMigrationJobDetailRequest) SetMigrationMode(v *DescribeMigrationJobDetailRequestMigrationMode) *DescribeMigrationJobDetailRequest {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetAccountId(v string) *DescribeMigrationJobDetailRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetClientToken(v string) *DescribeMigrationJobDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetMigrationJobId(v string) *DescribeMigrationJobDetailRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetOwnerId(v string) *DescribeMigrationJobDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageNum(v int32) *DescribeMigrationJobDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageSize(v int32) *DescribeMigrationJobDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetRegionId(v string) *DescribeMigrationJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetResourceGroupId(v string) *DescribeMigrationJobDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) Validate() error {
	if s.MigrationMode != nil {
		if err := s.MigrationMode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobDetailRequestMigrationMode struct {
	// Specifies whether to query the details of full data migration. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// > Default value: **false**.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Specifies whether to query the details of incremental data migration. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// > Default value: **false**.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Specifies whether to query the details of schema migration. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// > Default value: **false**.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeMigrationJobDetailRequestMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) Validate() error {
	return dara.Validate(s)
}
