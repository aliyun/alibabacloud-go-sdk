// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBClusterSSLRequest
	GetConnectionString() *string
	SetDBClusterId(v string) *ModifyDBClusterSSLRequest
	GetDBClusterId() *string
	SetEnableSSL(v bool) *ModifyDBClusterSSLRequest
	GetEnableSSL() *bool
	SetOwnerAccount(v string) *ModifyDBClusterSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterSSLRequest struct {
	// The internal or public endpoint for which the server certificate needs to be created or updated.
	//
	// example:
	//
	// am-d7oualxo05x4o5be872***.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable SSL encryption. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSSL            *bool   `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBClusterSSLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterSSLRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *ModifyDBClusterSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterSSLRequest) SetConnectionString(v string) *ModifyDBClusterSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBClusterId(v string) *ModifyDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetEnableSSL(v bool) *ModifyDBClusterSSLRequest {
	s.EnableSSL = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) Validate() error {
	return dara.Validate(s)
}
