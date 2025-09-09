// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDrdsInstanceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDrdsInstanceRequest
	GetDescription() *string
	SetDuration(v int32) *CreateDrdsInstanceRequest
	GetDuration() *int32
	SetInstanceSeries(v string) *CreateDrdsInstanceRequest
	GetInstanceSeries() *string
	SetIsAutoRenew(v bool) *CreateDrdsInstanceRequest
	GetIsAutoRenew() *bool
	SetMasterInstId(v string) *CreateDrdsInstanceRequest
	GetMasterInstId() *string
	SetMySQLVersion(v int32) *CreateDrdsInstanceRequest
	GetMySQLVersion() *int32
	SetPayType(v string) *CreateDrdsInstanceRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateDrdsInstanceRequest
	GetPricingCycle() *string
	SetQuantity(v int32) *CreateDrdsInstanceRequest
	GetQuantity() *int32
	SetRegionId(v string) *CreateDrdsInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDrdsInstanceRequest
	GetResourceGroupId() *string
	SetSpecification(v string) *CreateDrdsInstanceRequest
	GetSpecification() *string
	SetType(v string) *CreateDrdsInstanceRequest
	GetType() *string
	SetVpcId(v string) *CreateDrdsInstanceRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateDrdsInstanceRequest
	GetVswitchId() *string
	SetZoneId(v string) *CreateDrdsInstanceRequest
	GetZoneId() *string
	SetIsHa(v bool) *CreateDrdsInstanceRequest
	GetIsHa() *bool
}

type CreateDrdsInstanceRequest struct {
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// c1dd299c-10c6-11ea-bbbb-************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies the description of the instance. The description must meet the following requirements:
	//
	// 	- The description cannot contain the prefix http:// or https://.
	//
	// 	- The description must start with a letter or a Chinese character, and can contain uppercase and lowercase letters, Chinese characters, digits, underscores (_), and hyphens (-).
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies the purchase duration of the subscription instance.
	//
	// 	- If the PricingCycle parameter is set to year, the value range of the Duration parameter is 1 to 3.
	//
	// 	- If the PricingCycle parameter is set to month, the value range of the Duration parameter is 1 to 9.
	//
	// >  This parameter only takes effect when the PayType parameter is set to drdsPre.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies the instance type of the instance. Valid values:
	//
	// 	- **drds.sn2.4c16g**: The instance is of the Starter Edition.
	//
	// 	- **drds.sn2.8c32g**: The instance is of the Standard Edition
	//
	// 	- **drds.sn2.16c64g**: The instance is of the Enterprise Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds.sn2.4c16g
	InstanceSeries *string `json:"InstanceSeries,omitempty" xml:"InstanceSeries,omitempty"`
	// Specifies whether to enable automatic renewal. Valid values:
	//
	// 	- **true**: If the PricingCycle parameter is set to month, the subscription is automatically renewed for one month. If the PricingCycle parameter is set to year, the subscription is automatically renewed for one year.
	//
	// 	- **false**: The auto-renewal feature is disabled for the instance.
	//
	// >  This parameter only takes effect when the PayType parameter is set to drdsPre.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// Specifies the ID of the primary instance. This parameter is only required when you create a read-only instance.
	//
	// example:
	//
	// drds***********
	MasterInstId *string `json:"MasterInstId,omitempty" xml:"MasterInstId,omitempty"`
	// Specifies the MySQL version that is supported by the instance. Valid values:
	//
	// 	- **5**: The instance is fully compatible with MySQL 5.x. This value is the default value.
	//
	// 	- **8**: The instance is fully compatible with MySQL 8.0.
	//
	// >  This parameter only takes effect when you create a primary instance. By default, the MySQL version of the read-only instance is the same as that of the primary instance.
	//
	// example:
	//
	// 5
	MySQLVersion *int32 `json:"MySQLVersion,omitempty" xml:"MySQLVersion,omitempty"`
	// Specifies the billing method of the instance. Valid values:
	//
	// 	- **drdsPre**: The instance uses the subscription billing method.
	//
	// 	- **drdsPost**: The instance uses the pay-as-you-go billing method.
	//
	// 	- **drdsRo**: By default, the pay-as-you-go billing method is used when you create read-only instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsPost
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Specifies the unit of the subscription duration of the subscription instance. Valid values:
	//
	// 	- **year**: The unit of the subscription duration is year.
	//
	// 	- **month**: The unit of the subscription duration is month.
	//
	// >  This parameter is required if you set the PayType parameter to drdsPre.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Specifies the number of instances to be created. You can set the value only to 1. The value specifies that you can create one instance each time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Specifies the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies the specification code of the instance. The value consists of the instance type and the specified instance specification. For example, you can set the value to drds.sn2.4c16g.8c32g.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds.sn2.4c16g.8C32g
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// Specifies the type of the instance. Set the value to PRIVATE. The value PRIVATE specifies that the instance is a dedicated instance.
	//
	// >  You can also set the value to 1 to specify that the instance is a dedicated instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies the ID of the VPC.
	//
	// example:
	//
	// vpc-**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies the ID of the vSwitch.
	//
	// example:
	//
	// vsw-**********
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Specifies the zone ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// Specifies whether the instance is a high-availability instance.
	//
	// example:
	//
	// true
	IsHa *bool `json:"isHa,omitempty" xml:"isHa,omitempty"`
}

func (s CreateDrdsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDrdsInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDrdsInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateDrdsInstanceRequest) GetInstanceSeries() *string {
	return s.InstanceSeries
}

func (s *CreateDrdsInstanceRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *CreateDrdsInstanceRequest) GetMasterInstId() *string {
	return s.MasterInstId
}

func (s *CreateDrdsInstanceRequest) GetMySQLVersion() *int32 {
	return s.MySQLVersion
}

func (s *CreateDrdsInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDrdsInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateDrdsInstanceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateDrdsInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDrdsInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDrdsInstanceRequest) GetSpecification() *string {
	return s.Specification
}

func (s *CreateDrdsInstanceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDrdsInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDrdsInstanceRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateDrdsInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDrdsInstanceRequest) GetIsHa() *bool {
	return s.IsHa
}

func (s *CreateDrdsInstanceRequest) SetClientToken(v string) *CreateDrdsInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDescription(v string) *CreateDrdsInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetDuration(v int32) *CreateDrdsInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetInstanceSeries(v string) *CreateDrdsInstanceRequest {
	s.InstanceSeries = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsAutoRenew(v bool) *CreateDrdsInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMasterInstId(v string) *CreateDrdsInstanceRequest {
	s.MasterInstId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetMySQLVersion(v int32) *CreateDrdsInstanceRequest {
	s.MySQLVersion = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPayType(v string) *CreateDrdsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetPricingCycle(v string) *CreateDrdsInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetQuantity(v int32) *CreateDrdsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetRegionId(v string) *CreateDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetResourceGroupId(v string) *CreateDrdsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetSpecification(v string) *CreateDrdsInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetType(v string) *CreateDrdsInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVpcId(v string) *CreateDrdsInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetVswitchId(v string) *CreateDrdsInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetZoneId(v string) *CreateDrdsInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDrdsInstanceRequest) SetIsHa(v bool) *CreateDrdsInstanceRequest {
	s.IsHa = &v
	return s
}

func (s *CreateDrdsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
