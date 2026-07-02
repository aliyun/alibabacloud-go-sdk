// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairSkvDdbWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTairSkvDdbWorkspaceRequest
	GetClientToken() *string
	SetInstanceName(v string) *CreateTairSkvDdbWorkspaceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateTairSkvDdbWorkspaceRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *CreateTairSkvDdbWorkspaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTairSkvDdbWorkspaceRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateTairSkvDdbWorkspaceRequest
	GetPassword() *string
	SetPort(v int32) *CreateTairSkvDdbWorkspaceRequest
	GetPort() *int32
	SetRegionId(v string) *CreateTairSkvDdbWorkspaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTairSkvDdbWorkspaceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTairSkvDdbWorkspaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTairSkvDdbWorkspaceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateTairSkvDdbWorkspaceRequest
	GetSecurityToken() *string
	SetVSwitchId(v string) *CreateTairSkvDdbWorkspaceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateTairSkvDdbWorkspaceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateTairSkvDdbWorkspaceRequest
	GetZoneId() *string
}

type CreateTairSkvDdbWorkspaceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token is case-sensitive and can contain up to 64 ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz**
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. The name cannot contain the following characters: @/:="<>{} and spaces.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Set the value to tair_skv_ddb_ws.
	//
	// This parameter is required.
	//
	// example:
	//
	// tair_skv_ddb_ws
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must meet the following requirements:
	//
	// 	- The password is 8 to 32 characters in length.
	//
	// 	- The password contains at least three of the following character types: uppercase letters, lowercase letters, special characters, and digits. Supported special characters are `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port of the instance. Valid values: 1 to 65535. Default value: 443.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/61012.htm) to query available regions. Use this parameter to specify the region in which to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d**
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc**
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query available zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairSkvDdbWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTairSkvDdbWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTairSkvDdbWorkspaceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetClientToken(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetInstanceName(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetInstanceType(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetOwnerAccount(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetOwnerId(v int64) *CreateTairSkvDdbWorkspaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetPassword(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.Password = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetPort(v int32) *CreateTairSkvDdbWorkspaceRequest {
	s.Port = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetRegionId(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetResourceGroupId(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetResourceOwnerAccount(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetResourceOwnerId(v int64) *CreateTairSkvDdbWorkspaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetSecurityToken(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetVSwitchId(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetVpcId(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) SetZoneId(v string) *CreateTairSkvDdbWorkspaceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTairSkvDdbWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
