// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeComputeResourceRequest
	GetDBClusterId() *string
	SetDBClusterVersion(v string) *DescribeComputeResourceRequest
	GetDBClusterVersion() *string
	SetMigrate(v bool) *DescribeComputeResourceRequest
	GetMigrate() *bool
	SetOwnerAccount(v string) *DescribeComputeResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeComputeResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeComputeResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeComputeResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeComputeResourceRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeComputeResourceRequest
	GetZoneId() *string
}

type DescribeComputeResourceRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The version of the AnalyticDB for MySQL Data Warehouse Edition cluster. Set the value to **3.0**.
	//
	// example:
	//
	// 3.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// Specifies whether to query the specifications of the available computing resources that are migrated from AnalyticDB for MySQL Data Warehouse Edition to Data Lakehouse Edition. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	Migrate      *bool   `json:"Migrate,omitempty" xml:"Migrate,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/129857.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeComputeResourceRequest) GetDBClusterVersion() *string {
	return s.DBClusterVersion
}

func (s *DescribeComputeResourceRequest) GetMigrate() *bool {
	return s.Migrate
}

func (s *DescribeComputeResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeComputeResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeComputeResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeComputeResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeComputeResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeComputeResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeComputeResourceRequest) SetDBClusterId(v string) *DescribeComputeResourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetDBClusterVersion(v string) *DescribeComputeResourceRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetMigrate(v bool) *DescribeComputeResourceRequest {
	s.Migrate = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetOwnerAccount(v string) *DescribeComputeResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetOwnerId(v int64) *DescribeComputeResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetRegionId(v string) *DescribeComputeResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetResourceOwnerAccount(v string) *DescribeComputeResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetResourceOwnerId(v int64) *DescribeComputeResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetZoneId(v string) *DescribeComputeResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
