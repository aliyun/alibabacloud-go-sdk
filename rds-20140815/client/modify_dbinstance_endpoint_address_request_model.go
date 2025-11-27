// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceEndpointAddressRequest
	GetClientToken() *string
	SetConnectionString(v string) *ModifyDBInstanceEndpointAddressRequest
	GetConnectionString() *string
	SetConnectionStringPrefix(v string) *ModifyDBInstanceEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointAddressRequest
	GetDBInstanceEndpointId() *string
	SetDBInstanceId(v string) *ModifyDBInstanceEndpointAddressRequest
	GetDBInstanceId() *string
	SetPort(v string) *ModifyDBInstanceEndpointAddressRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *ModifyDBInstanceEndpointAddressRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointAddressRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ModifyDBInstanceEndpointAddressRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyDBInstanceEndpointAddressRequest
	GetVpcId() *string
}

type ModifyDBInstanceEndpointAddressRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint that you want to modify. The endpoint can be a public endpoint or an internal endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the new endpoint. You can modify only the prefix of the endpoint that is specified by the ConnectionString parameter.
	//
	// example:
	//
	// rm-bp181qoj4s34m33****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The endpoint ID of the instance. You can call the DescribeDBInstanceEndpoints operation to query the endpoint ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-****
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The port number of the new endpoint.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The IP address of the internal endpoint.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitch ID of the internal endpoint.
	//
	// example:
	//
	// vsw-bp12u14ecz****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the internal endpoint.
	//
	// example:
	//
	// vpc-bp17xdic25d****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyDBInstanceEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceEndpointAddressRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetClientToken(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetConnectionString(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetDBInstanceEndpointId(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetDBInstanceId(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetPort(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetVSwitchId(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) SetVpcId(v string) *ModifyDBInstanceEndpointAddressRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
