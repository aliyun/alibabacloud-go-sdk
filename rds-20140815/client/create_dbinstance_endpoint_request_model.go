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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ConnectionStringPrefix        *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceEndpointDescription *string `json:"DBInstanceEndpointDescription,omitempty" xml:"DBInstanceEndpointDescription,omitempty"`
	// This parameter is required.
	DBInstanceEndpointType *string `json:"DBInstanceEndpointType,omitempty" xml:"DBInstanceEndpointType,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	NodeItems []*CreateDBInstanceEndpointRequestNodeItems `json:"NodeItems,omitempty" xml:"NodeItems,omitempty" type:"Repeated"`
	// This parameter is required.
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Weight *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
