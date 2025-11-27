// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBInstanceEndpointShrinkRequest
	GetClientToken() *string
	SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointShrinkRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceEndpointDescription(v string) *CreateDBInstanceEndpointShrinkRequest
	GetDBInstanceEndpointDescription() *string
	SetDBInstanceEndpointType(v string) *CreateDBInstanceEndpointShrinkRequest
	GetDBInstanceEndpointType() *string
	SetDBInstanceId(v string) *CreateDBInstanceEndpointShrinkRequest
	GetDBInstanceId() *string
	SetNodeItemsShrink(v string) *CreateDBInstanceEndpointShrinkRequest
	GetNodeItemsShrink() *string
	SetPort(v string) *CreateDBInstanceEndpointShrinkRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateDBInstanceEndpointShrinkRequest
	GetPrivateIpAddress() *string
	SetResourceGroupId(v string) *CreateDBInstanceEndpointShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceEndpointShrinkRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *CreateDBInstanceEndpointShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBInstanceEndpointShrinkRequest
	GetVpcId() *string
}

type CreateDBInstanceEndpointShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The prefix of the internal endpoint.
	//
	// When you create any type of endpoint, an internal endpoint is automatically created for the endpoint. This parameter specifies the prefix of the internal endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****-ro
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The user-defined description of the endpoint.
	//
	// example:
	//
	// for readonly business
	DBInstanceEndpointDescription *string `json:"DBInstanceEndpointDescription,omitempty" xml:"DBInstanceEndpointDescription,omitempty"`
	// The endpoint type. Valid values:
	//
	// 	- Primary: read/write endpoint of the instance
	//
	// 	- Readonly: read-only endpoint of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// Readonly
	DBInstanceEndpointType *string `json:"DBInstanceEndpointType,omitempty" xml:"DBInstanceEndpointType,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The information about the endpoint.
	//
	// This parameter is required.
	NodeItemsShrink *string `json:"NodeItems,omitempty" xml:"NodeItems,omitempty"`
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.
	//
	// Valid values: 3000 to 5999.
	//
	// This parameter is required.
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
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitch ID of the internal endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1kqp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the internal endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxxmmxjqqi****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDBInstanceEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetDBInstanceEndpointDescription() *string {
	return s.DBInstanceEndpointDescription
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetDBInstanceEndpointType() *string {
	return s.DBInstanceEndpointType
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetNodeItemsShrink() *string {
	return s.NodeItemsShrink
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceEndpointShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetClientToken(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetDBInstanceEndpointDescription(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.DBInstanceEndpointDescription = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetDBInstanceEndpointType(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.DBInstanceEndpointType = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetDBInstanceId(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetNodeItemsShrink(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.NodeItemsShrink = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetPort(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetPrivateIpAddress(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetResourceOwnerId(v int64) *CreateDBInstanceEndpointShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) SetVpcId(v string) *CreateDBInstanceEndpointShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
