// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *CreateDBEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *CreateDBEndpointAddressRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *CreateDBEndpointAddressRequest
	GetDBEndpointId() *string
	SetNetType(v string) *CreateDBEndpointAddressRequest
	GetNetType() *string
	SetOwnerAccount(v string) *CreateDBEndpointAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBEndpointAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDBEndpointAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBEndpointAddressRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *CreateDBEndpointAddressRequest
	GetSecurityGroupId() *string
	SetVPCId(v string) *CreateDBEndpointAddressRequest
	GetVPCId() *string
	SetZoneInfo(v []*CreateDBEndpointAddressRequestZoneInfo) *CreateDBEndpointAddressRequest
	GetZoneInfo() []*CreateDBEndpointAddressRequestZoneInfo
}

type CreateDBEndpointAddressRequest struct {
	// The prefix of the new endpoint. The prefix of the endpoint must meet the following requirements:
	//
	// 	- The prefix can contain lowercase letters, digits, and hyphens (-).
	//
	// 	- The prefix must start with a letter and end with a digit or a letter.
	//
	// 	- The prefix must be 6 to 40 characters in length.
	//
	// example:
	//
	// test-1
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// >  You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query endpoint details.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-**************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type of the endpoint. Set the value to **Public**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the ECS security group.
	//
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The details of the zones.
	ZoneInfo []*CreateDBEndpointAddressRequestZoneInfo `json:"ZoneInfo,omitempty" xml:"ZoneInfo,omitempty" type:"Repeated"`
}

func (s CreateDBEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBEndpointAddressRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBEndpointAddressRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *CreateDBEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *CreateDBEndpointAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBEndpointAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBEndpointAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBEndpointAddressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateDBEndpointAddressRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBEndpointAddressRequest) GetZoneInfo() []*CreateDBEndpointAddressRequestZoneInfo {
	return s.ZoneInfo
}

func (s *CreateDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBClusterId(v string) *CreateDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBEndpointId(v string) *CreateDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetNetType(v string) *CreateDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetSecurityGroupId(v string) *CreateDBEndpointAddressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetVPCId(v string) *CreateDBEndpointAddressRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetZoneInfo(v []*CreateDBEndpointAddressRequestZoneInfo) *CreateDBEndpointAddressRequest {
	s.ZoneInfo = v
	return s
}

func (s *CreateDBEndpointAddressRequest) Validate() error {
	if s.ZoneInfo != nil {
		for _, item := range s.ZoneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBEndpointAddressRequestZoneInfo struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBEndpointAddressRequestZoneInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateDBEndpointAddressRequestZoneInfo) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressRequestZoneInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBEndpointAddressRequestZoneInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBEndpointAddressRequestZoneInfo) SetVSwitchId(v string) *CreateDBEndpointAddressRequestZoneInfo {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBEndpointAddressRequestZoneInfo) SetZoneId(v string) *CreateDBEndpointAddressRequestZoneInfo {
	s.ZoneId = &v
	return s
}

func (s *CreateDBEndpointAddressRequestZoneInfo) Validate() error {
	return dara.Validate(s)
}
