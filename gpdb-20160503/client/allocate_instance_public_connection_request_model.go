// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *AllocateInstancePublicConnectionRequest
	GetAddressType() *string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateInstancePublicConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateInstancePublicConnectionRequest struct {
	// The network type of the endpoint. Valid values:
	//
	// 	- **primary**: primary endpoint
	//
	// 	- **cluster**: instance endpoint. This value is supported only for an instance that contains multiple coordinator nodes.
	//
	// >  The default value is primary.
	//
	// example:
	//
	// primary
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The prefix of the endpoint.
	//
	// Specify a prefix for the endpoint. Example: `gp-bp12ga6v69h86****`. In this example, the endpoint is `gp-bp12ga6v69h86****.gpdb.rds.aliyuncs.com`.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. Example: 5432.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5432
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) SetAddressType(v string) *AllocateInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
