// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateInstanceRequest
	GetAutoPay() *bool
	SetChargeType(v string) *CreateInstanceRequest
	GetChargeType() *string
	SetDuration(v int32) *CreateInstanceRequest
	GetDuration() *int32
	SetHttpsPolicy(v string) *CreateInstanceRequest
	GetHttpsPolicy() *string
	SetInstanceCidr(v string) *CreateInstanceRequest
	GetInstanceCidr() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceSpec(v string) *CreateInstanceRequest
	GetInstanceSpec() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetPricingCycle(v string) *CreateInstanceRequest
	GetPricingCycle() *string
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetToken(v string) *CreateInstanceRequest
	GetToken() *string
	SetUserVpcId(v string) *CreateInstanceRequest
	GetUserVpcId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
	SetZoneVSwitchSecurityGroup(v []*CreateInstanceRequestZoneVSwitchSecurityGroup) *CreateInstanceRequest
	GetZoneVSwitchSecurityGroup() []*CreateInstanceRequestZoneVSwitchSecurityGroup
}

type CreateInstanceRequest struct {
	// Whether to automatically pay when renewing. Value:
	//
	// - True: Automatic payment. Please ensure that your account has sufficient balance.
	//
	// - False: Console manual payment. The specific operation is to log in to the console, select Expenses in the upper right corner, enter the Expense Center, and find the target order in Order Management to make payment.
	//
	// Default value: False.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The billing method of the instance. Valid values: PostPaid (pay-as-you-go) and PrePaid (subscription).
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If PricingCycle is set to **Month**, set this parameter to an integer ranges from **1*	- to **9**.
	//
	// 	- If PricingCycle is set to **Year**, set this parameter to an integer ranges from **1*	- to **3**.
	//
	// >  This parameter is valid and required only if the ChargeType parameter is set to **PrePaid**.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The HTTPS policy.
	//
	// example:
	//
	// HTTPS2_TLS1_2
	HttpsPolicy *string `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	// The CIDR block of the VPC integration instance.
	//
	// 	- 192.168.0.0/16
	//
	// 	- 172.16.0.0/12
	//
	// **
	//
	// **Warning*	- The VPC integration instance is connected to the specified CIDR block. Plan your CIDR block carefully to prevent overlaps with the private IP addresses of cloud services.
	//
	// >  This parameter is in invitational preview and not available for public use.
	//
	// example:
	//
	// 172.16.0.0/12
	InstanceCidr *string `json:"InstanceCidr,omitempty" xml:"InstanceCidr,omitempty"`
	// Instance Name
	//
	// This parameter is required.
	//
	// example:
	//
	// ApigatewayInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Instance specifications
	//
	// This parameter is required.
	//
	// example:
	//
	// api.s1.small
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The type of the dedicated instance. Valid values:
	//
	// 	- vpc_connect: a VPC integration instance
	//
	// 	- normal: a conventional dedicated instance
	//
	// >  This parameter is in invitational preview and not available for public use.
	//
	// example:
	//
	// vpc_connect
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The unit of the subscription duration of the subscription instance. Valid values:
	//
	// 	- **year**
	//
	// 	- **month**
	//
	// >  This parameter is required if the ChargeType parameter is set to Prepaid.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The tags that you want to add to the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Passwords are used to prevent duplicate requests from being submitted, please do not reuse them.
	//
	// This parameter is required.
	//
	// example:
	//
	// c20d86c4-1eb3-4d0b-afe9-c586df1e2136
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the user\\"s VPC to be connected by the VPC integration instance.
	//
	// >  This parameter is in invitational preview and not available for public use.
	//
	// example:
	//
	// vpc-m5eo7khlb4h4f8y9egsdg
	UserVpcId *string `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
	// The zone in which you want to create the instance. This parameter is required for a conventional dedicated instance and optional for a virtual private cloud (VPC) integration instance.
	//
	// example:
	//
	// cn-beijing-MAZ3(c,e)
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The network information when the instance is a VPC integration instance, such as the zone, vSwitch, and security group.
	//
	// >  This parameter is in invitational preview and not available for public use.
	ZoneVSwitchSecurityGroup []*CreateInstanceRequestZoneVSwitchSecurityGroup `json:"ZoneVSwitchSecurityGroup,omitempty" xml:"ZoneVSwitchSecurityGroup,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInstanceRequest) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *CreateInstanceRequest) GetInstanceCidr() *string {
	return s.InstanceCidr
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetToken() *string {
	return s.Token
}

func (s *CreateInstanceRequest) GetUserVpcId() *string {
	return s.UserVpcId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) GetZoneVSwitchSecurityGroup() []*CreateInstanceRequestZoneVSwitchSecurityGroup {
	return s.ZoneVSwitchSecurityGroup
}

func (s *CreateInstanceRequest) SetAutoPay(v bool) *CreateInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpsPolicy(v string) *CreateInstanceRequest {
	s.HttpsPolicy = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceCidr(v string) *CreateInstanceRequest {
	s.InstanceCidr = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceSpec(v string) *CreateInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetToken(v string) *CreateInstanceRequest {
	s.Token = &v
	return s
}

func (s *CreateInstanceRequest) SetUserVpcId(v string) *CreateInstanceRequest {
	s.UserVpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneVSwitchSecurityGroup(v []*CreateInstanceRequestZoneVSwitchSecurityGroup) *CreateInstanceRequest {
	s.ZoneVSwitchSecurityGroup = v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneVSwitchSecurityGroup != nil {
		for _, item := range s.ZoneVSwitchSecurityGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// test1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestZoneVSwitchSecurityGroup struct {
	// The IPv4 CIDR block for the vSwitch.
	//
	// example:
	//
	// 192.168.9.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The ID of the security group. Services in the same security group can access each other.
	//
	// example:
	//
	// sg-2ze2ql9nozv8q7kmlt6e
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0xi349n11cxogmvm866tb
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequestZoneVSwitchSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestZoneVSwitchSecurityGroup) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) SetCidrBlock(v string) *CreateInstanceRequestZoneVSwitchSecurityGroup {
	s.CidrBlock = &v
	return s
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) SetSecurityGroupId(v string) *CreateInstanceRequestZoneVSwitchSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) SetVSwitchId(v string) *CreateInstanceRequestZoneVSwitchSecurityGroup {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) SetZoneId(v string) *CreateInstanceRequestZoneVSwitchSecurityGroup {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequestZoneVSwitchSecurityGroup) Validate() error {
	return dara.Validate(s)
}
