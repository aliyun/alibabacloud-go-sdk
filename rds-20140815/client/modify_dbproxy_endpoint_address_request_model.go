// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetDBProxyNewConnectString(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyNewConnectString() *string
	SetDBProxyNewConnectStringPort(v string) *ModifyDBProxyEndpointAddressRequest
	GetDBProxyNewConnectStringPort() *string
	SetOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type ModifyDBProxyEndpointAddressRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the database proxy endpoint. Valid values:
	//
	// 	- **Public**
	//
	// 	- **VPC*	- (default)
	//
	// >  If the RDS instance runs MySQL, this parameter is required.
	//
	// example:
	//
	// Public
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// The ID of the database proxy endpoint. You can call the DescribeDBProxyEndpoint operation to query the ID of the database proxy endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta9um4****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The prefix of the new database proxy endpoint. A custom value is supported.
	//
	// >  You must specify at least one of the **DBProxyNewConnectString*	- and **DBProxyNewConnectStringPort*	- parameters.
	//
	// example:
	//
	// test123456
	DBProxyNewConnectString *string `json:"DBProxyNewConnectString,omitempty" xml:"DBProxyNewConnectString,omitempty"`
	// The port number that is associated with the database proxy endpoint. A custom value is supported.
	//
	// >  You must specify at least one of the **DBProxyNewConnectString*	- and **DBProxyNewConnectStringPort*	- parameters.
	//
	// example:
	//
	// 3307
	DBProxyNewConnectStringPort *string `json:"DBProxyNewConnectStringPort,omitempty" xml:"DBProxyNewConnectStringPort,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyNewConnectString() *string {
	return s.DBProxyNewConnectString
}

func (s *ModifyDBProxyEndpointAddressRequest) GetDBProxyNewConnectStringPort() *string {
	return s.DBProxyNewConnectStringPort
}

func (s *ModifyDBProxyEndpointAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyEndpointAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyNewConnectString(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectString = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetDBProxyNewConnectStringPort(v string) *ModifyDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectStringPort = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBProxyEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
