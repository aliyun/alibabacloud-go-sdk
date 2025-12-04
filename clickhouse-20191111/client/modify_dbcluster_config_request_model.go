// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterConfigRequest
	GetOwnerId() *int64
	SetReason(v string) *ModifyDBClusterConfigRequest
	GetReason() *string
	SetRegionId(v string) *ModifyDBClusterConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterConfigRequest
	GetResourceOwnerId() *int64
	SetUserConfig(v string) *ModifyDBClusterConfigRequest
	GetUserConfig() *string
}

type ModifyDBClusterConfigRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1t9lbb7a4z7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reason for the change.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The names of the parameters and the new values that you want to specify for the parameters.
	//
	// >  You can change the value of a single parameter. The values of parameters that are not specified will not be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"keep_alive_timeout":"301"}
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ModifyDBClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterConfigRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyDBClusterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterConfigRequest) GetUserConfig() *string {
	return s.UserConfig
}

func (s *ModifyDBClusterConfigRequest) SetDBClusterId(v string) *ModifyDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetOwnerAccount(v string) *ModifyDBClusterConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetOwnerId(v int64) *ModifyDBClusterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetReason(v string) *ModifyDBClusterConfigRequest {
	s.Reason = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetRegionId(v string) *ModifyDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetResourceOwnerId(v int64) *ModifyDBClusterConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetUserConfig(v string) *ModifyDBClusterConfigRequest {
	s.UserConfig = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
