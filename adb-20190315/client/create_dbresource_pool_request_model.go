// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourcePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateDBResourcePoolRequest
	GetDBClusterId() *string
	SetNodeNum(v int32) *CreateDBResourcePoolRequest
	GetNodeNum() *int32
	SetOwnerAccount(v string) *CreateDBResourcePoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBResourcePoolRequest
	GetOwnerId() *int64
	SetPoolName(v string) *CreateDBResourcePoolRequest
	GetPoolName() *string
	SetQueryType(v string) *CreateDBResourcePoolRequest
	GetQueryType() *string
	SetResourceOwnerAccount(v string) *CreateDBResourcePoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBResourcePoolRequest
	GetResourceOwnerId() *int64
}

type CreateDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes. Default value: 0.
	//
	// 	- Each node provides 16 cores and 64 GB memory.
	//
	// 	- The total amount of resources provided by the nodes (number of nodes × 16 cores, number of nodes × 64 GB memory) cannot exceed the total amount of resources in the cluster. Set this parameter to a proper value.
	//
	// example:
	//
	// 2
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// 	- The name can be up to 255 characters in length.
	//
	// 	- The name must start with a letter or a digit.
	//
	// 	- The name can contain letters, digits, hyphens (_), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The mode in which to execute SQL statements.
	//
	// 	- **batch**
	//
	// 	- **interactive**
	//
	// > For more information, see [Query execution modes](https://help.aliyun.com/document_detail/189502.html).
	//
	// example:
	//
	// interactive
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBResourcePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBResourcePoolRequest) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *CreateDBResourcePoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBResourcePoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBResourcePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *CreateDBResourcePoolRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *CreateDBResourcePoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBResourcePoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBResourcePoolRequest) SetDBClusterId(v string) *CreateDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetNodeNum(v int32) *CreateDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetPoolName(v string) *CreateDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetQueryType(v string) *CreateDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) Validate() error {
	return dara.Validate(s)
}
