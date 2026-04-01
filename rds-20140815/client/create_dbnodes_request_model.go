// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBNodesRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreateDBNodesRequest
	GetDBInstanceId() *string
	SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest
	GetDBNode() []*CreateDBNodesRequestDBNode
	SetOwnerAccount(v string) *CreateDBNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBNodesRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateDBNodesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBNodesRequest
	GetResourceOwnerId() *int64
}

type CreateDBNodesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze450g4ctg6t****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the node.
	//
	// This parameter is required.
	DBNode       []*CreateDBNodesRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	OwnerAccount *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBNodesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBNodesRequest) GetDBNode() []*CreateDBNodesRequestDBNode {
	return s.DBNode
}

func (s *CreateDBNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBNodesRequest) SetClientToken(v string) *CreateDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBInstanceId(v string) *CreateDBNodesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest {
	s.DBNode = v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerAccount(v string) *CreateDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerId(v int64) *CreateDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceGroupId(v string) *CreateDBNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerAccount(v string) *CreateDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerId(v int64) *CreateDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) Validate() error {
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBNodesRequestDBNode struct {
	// The specification information of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.n2.medium.xc
	ClassCode *string `json:"classCode,omitempty" xml:"classCode,omitempty"`
	// The vSwitch ID of the node.
	//
	// example:
	//
	// vsw-bp1sxxsodv28ey5dl****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// The ID of the zone in which the node is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou-a
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateDBNodesRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesRequestDBNode) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequestDBNode) GetClassCode() *string {
	return s.ClassCode
}

func (s *CreateDBNodesRequestDBNode) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateDBNodesRequestDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBNodesRequestDBNode) SetClassCode(v string) *CreateDBNodesRequestDBNode {
	s.ClassCode = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) SetVswitchId(v string) *CreateDBNodesRequestDBNode {
	s.VswitchId = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) SetZoneId(v string) *CreateDBNodesRequestDBNode {
	s.ZoneId = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) Validate() error {
	return dara.Validate(s)
}
