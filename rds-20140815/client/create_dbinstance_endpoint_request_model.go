// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBInstanceEndpointRequest
	GetClientToken() *string
	SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceEndpointDescription(v string) *CreateDBInstanceEndpointRequest
	GetDBInstanceEndpointDescription() *string
	SetDBInstanceEndpointType(v string) *CreateDBInstanceEndpointRequest
	GetDBInstanceEndpointType() *string
	SetDBInstanceId(v string) *CreateDBInstanceEndpointRequest
	GetDBInstanceId() *string
	SetNodeItems(v []*CreateDBInstanceEndpointRequestNodeItems) *CreateDBInstanceEndpointRequest
	GetNodeItems() []*CreateDBInstanceEndpointRequestNodeItems
	SetPort(v string) *CreateDBInstanceEndpointRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateDBInstanceEndpointRequest
	GetPrivateIpAddress() *string
	SetResourceGroupId(v string) *CreateDBInstanceEndpointRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceEndpointRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *CreateDBInstanceEndpointRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBInstanceEndpointRequest
	GetVpcId() *string
}

type CreateDBInstanceEndpointRequest struct {
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
	NodeItems []*CreateDBInstanceEndpointRequestNodeItems `json:"NodeItems,omitempty" xml:"NodeItems,omitempty" type:"Repeated"`
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

func (s CreateDBInstanceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceEndpointRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBInstanceEndpointRequest) GetDBInstanceEndpointDescription() *string {
	return s.DBInstanceEndpointDescription
}

func (s *CreateDBInstanceEndpointRequest) GetDBInstanceEndpointType() *string {
	return s.DBInstanceEndpointType
}

func (s *CreateDBInstanceEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceEndpointRequest) GetNodeItems() []*CreateDBInstanceEndpointRequestNodeItems {
	return s.NodeItems
}

func (s *CreateDBInstanceEndpointRequest) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceEndpointRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDBInstanceEndpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceEndpointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceEndpointRequest) SetClientToken(v string) *CreateDBInstanceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetConnectionStringPrefix(v string) *CreateDBInstanceEndpointRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetDBInstanceEndpointDescription(v string) *CreateDBInstanceEndpointRequest {
	s.DBInstanceEndpointDescription = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetDBInstanceEndpointType(v string) *CreateDBInstanceEndpointRequest {
	s.DBInstanceEndpointType = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetDBInstanceId(v string) *CreateDBInstanceEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetNodeItems(v []*CreateDBInstanceEndpointRequestNodeItems) *CreateDBInstanceEndpointRequest {
	s.NodeItems = v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetPort(v string) *CreateDBInstanceEndpointRequest {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetPrivateIpAddress(v string) *CreateDBInstanceEndpointRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetResourceGroupId(v string) *CreateDBInstanceEndpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetResourceOwnerId(v int64) *CreateDBInstanceEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetVSwitchId(v string) *CreateDBInstanceEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) SetVpcId(v string) *CreateDBInstanceEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequest) Validate() error {
	if s.NodeItems != nil {
		for _, item := range s.NodeItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBInstanceEndpointRequestNodeItems struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rn-xxxx-****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The weight of the node. Read requests are distributed based on the weight.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateDBInstanceEndpointRequestNodeItems) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointRequestNodeItems) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointRequestNodeItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceEndpointRequestNodeItems) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateDBInstanceEndpointRequestNodeItems) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateDBInstanceEndpointRequestNodeItems) SetDBInstanceId(v string) *CreateDBInstanceEndpointRequestNodeItems {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequestNodeItems) SetNodeId(v string) *CreateDBInstanceEndpointRequestNodeItems {
	s.NodeId = &v
	return s
}

func (s *CreateDBInstanceEndpointRequestNodeItems) SetWeight(v int64) *CreateDBInstanceEndpointRequestNodeItems {
	s.Weight = &v
	return s
}

func (s *CreateDBInstanceEndpointRequestNodeItems) Validate() error {
	return dara.Validate(s)
}
