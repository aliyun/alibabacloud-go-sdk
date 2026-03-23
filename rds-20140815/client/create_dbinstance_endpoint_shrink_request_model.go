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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ConnectionStringPrefix        *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBInstanceEndpointDescription *string `json:"DBInstanceEndpointDescription,omitempty" xml:"DBInstanceEndpointDescription,omitempty"`
	// This parameter is required.
	DBInstanceEndpointType *string `json:"DBInstanceEndpointType,omitempty" xml:"DBInstanceEndpointType,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	NodeItemsShrink *string `json:"NodeItems,omitempty" xml:"NodeItems,omitempty"`
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
