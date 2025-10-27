// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourcePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBResourcePoolRequest
	GetDBClusterId() *string
	SetNodeNum(v int32) *ModifyDBResourcePoolRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *ModifyDBResourcePoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBResourcePoolRequest
	GetOwnerId() *int64
	SetPoolName(v string) *ModifyDBResourcePoolRequest
	GetPoolName() *string
	SetQueryType(v string) *ModifyDBResourcePoolRequest
	GetQueryType() *string
	SetResourceOwnerAccount(v string) *ModifyDBResourcePoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBResourcePoolRequest
	GetResourceOwnerId() *int64
}

type ModifyDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes.
	//
	// 	- Each node provides 16 cores and 64 GB memory.
	//
	// 	- The amount of resources that you want to add to or remove from the cluster cannot exceed the total amount of resources in the cluster.
	//
	// > - If you do not specify this parameter, the original value is retained.
	//
	// > - You must specify at least one of the QueryType and NodeNum parameters.
	//
	// example:
	//
	// 2
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The mode in which SQL statements are executed. Valid values:
	//
	// 	- **batch**
	//
	// 	- **interactive**
	//
	// > If you do not specify this parameter, the original value is retained.
	//
	// example:
	//
	// batch
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourcePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBResourcePoolRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ModifyDBResourcePoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBResourcePoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBResourcePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *ModifyDBResourcePoolRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ModifyDBResourcePoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBResourcePoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBResourcePoolRequest) SetDBClusterId(v string) *ModifyDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetNodeNum(v int32) *ModifyDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetPoolName(v string) *ModifyDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetQueryType(v string) *ModifyDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) Validate() error {
	return dara.Validate(s)
}
