// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBClusterVipRequest
	GetConnectionString() *string
	SetDBClusterId(v string) *ModifyDBClusterVipRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBClusterVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterVipRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBClusterVipRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBClusterVipRequest
	GetVSwitchId() *string
}

type ModifyDBClusterVipRequest struct {
	// The internal or public endpoint for which the server certificate needs to be created or updated.
	//
	// example:
	//
	// am-2ze8mbuai974s4y2500000169.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-2ze8mbuai97*****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1at5ze0t5u3xtqn****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1aadw9k19x6cis9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBClusterVipRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVipRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVipRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBClusterVipRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterVipRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBClusterVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBClusterVipRequest) SetConnectionString(v string) *ModifyDBClusterVipRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetDBClusterId(v string) *ModifyDBClusterVipRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetOwnerAccount(v string) *ModifyDBClusterVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetOwnerId(v int64) *ModifyDBClusterVipRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetRegionId(v string) *ModifyDBClusterVipRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetResourceOwnerId(v int64) *ModifyDBClusterVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetVPCId(v string) *ModifyDBClusterVipRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) SetVSwitchId(v string) *ModifyDBClusterVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBClusterVipRequest) Validate() error {
	return dara.Validate(s)
}
