// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBClusterId(v string) *ModifyClusterConnectionStringRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyClusterConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyClusterConnectionStringRequest
	GetOwnerId() *int64
	SetPort(v int32) *ModifyClusterConnectionStringRequest
	GetPort() *int32
	SetResourceOwnerAccount(v string) *ModifyClusterConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClusterConnectionStringRequest
	GetResourceOwnerId() *int64
}

type ModifyClusterConnectionStringRequest struct {
	// The prefix of the public endpoint.
	//
	// 	- The prefix can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	//
	// 	- The prefix can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-123
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The current public endpoint of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusterNetInfo](https://help.aliyun.com/document_detail/143384.html) operation to query the public endpoint of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp18934i73vb****.ads.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp18934i73vb****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. Set the value to **3306**.
	//
	// example:
	//
	// 3306
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyClusterConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyClusterConnectionStringRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyClusterConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyClusterConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClusterConnectionStringRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyClusterConnectionStringRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClusterConnectionStringRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
