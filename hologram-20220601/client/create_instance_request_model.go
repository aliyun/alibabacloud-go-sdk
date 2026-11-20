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
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *CreateInstanceRequest
	GetChargeType() *string
	SetColdStorageSize(v int64) *CreateInstanceRequest
	GetColdStorageSize() *int64
	SetCpu(v int64) *CreateInstanceRequest
	GetCpu() *int64
	SetDuration(v int64) *CreateInstanceRequest
	GetDuration() *int64
	SetEnableServerlessComputing(v bool) *CreateInstanceRequest
	GetEnableServerlessComputing() *bool
	SetGatewayCount(v int64) *CreateInstanceRequest
	GetGatewayCount() *int64
	SetInitialDatabases(v string) *CreateInstanceRequest
	GetInitialDatabases() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetLeaderInstanceId(v string) *CreateInstanceRequest
	GetLeaderInstanceId() *string
	SetPricingCycle(v string) *CreateInstanceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetStorageSize(v int64) *CreateInstanceRequest
	GetStorageSize() *int64
	SetStorageType(v string) *CreateInstanceRequest
	GetStorageType() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
}

type CreateInstanceRequest struct {
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// - true: Automatic payment is enabled.
	//
	// - false: Only an order is generated, and no payment is made.
	//
	// > Default value: true. If your payment method has insufficient balance, you can set the AutoPay parameter to false. In this case, an unpaid order is generated, and you can log on to the Expenses and Costs console to complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	// Specifies whether to enable monthly auto-renewal. Default value: false. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The billing method. Valid values:
	//
	// - PrePaid: Subscription.
	//
	// - PostPaid: Pay-as-you-go.
	//
	// > This parameter is ignored for Shared instances. Shared is a defined specifications type, and its billing method defaults to PostPaid.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The low-frequency storage capacity of the instance. Unit: GB.
	//
	// > This parameter is ignored for pay-as-you-go instances.
	//
	// example:
	//
	// 500
	ColdStorageSize *int64 `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	// The instance specifications. Valid values:
	//
	// - 32 cores, 128 GB (Number of compute nodes: 2)
	//
	// - 64 cores, 256 GB (Number of compute nodes: 4)
	//
	// - 96 cores, 384 GB (Number of compute nodes: 6)
	//
	// - 128 cores, 512 GB (Number of compute nodes: 8)
	//
	// - Other specifications
	//
	// >
	//
	//  >- Specify the number of cores.
	//
	//  >- To purchase an instance with more than 1,024 cores, submit a ticket.
	//
	//  >- You do not need to specify the specifications for a Shared instance.
	//
	// example:
	//
	// 64
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The subscription duration. For example, set this parameter to 2 for a two-month subscription.
	//
	// > This parameter is not required if the billing method is PostPaid.
	//
	// example:
	//
	// 2
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Specifies whether to enable Serverless Computing.
	//
	// example:
	//
	// true
	EnableServerlessComputing *bool `json:"enableServerlessComputing,omitempty" xml:"enableServerlessComputing,omitempty"`
	// The number of gateways. Valid values: 2 to 50.
	//
	// > This parameter is required only for Warehouse instances.
	//
	// example:
	//
	// 4
	GatewayCount *int64 `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	// The initial databases.
	//
	// example:
	//
	// chatbot
	InitialDatabases *string `json:"initialDatabases,omitempty" xml:"initialDatabases,omitempty"`
	// The name of the instance. The name must be 2 to 64 characters in length. Each Chinese character or letter counts as one character.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_holo
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The instance type. Valid values:
	//
	// - Standard: General-purpose.
	//
	// - Follower: Read-only secondary instance.
	//
	// - Warehouse: Compute group.
	//
	// - Shared: Shared.
	//
	// - Severless: Serverless.
	//
	// This parameter is required.
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The ID of the primary instance. This parameter is required for Follower instances.
	//
	// > The primary and secondary instances must meet the following conditions:
	//
	// > - The primary instance is in the Running state.
	//
	// > - The primary and secondary instances are in the same region.
	//
	// > - The primary and secondary instances are in the same zone.
	//
	// > - The primary instance is associated with fewer than 10 secondary instances.
	//
	// > - The primary and secondary instances belong to the same Alibaba Cloud account.
	//
	// example:
	//
	// hgpostcn-cn-lbj3aworq112
	LeaderInstanceId *string `json:"leaderInstanceId,omitempty" xml:"leaderInstanceId,omitempty"`
	// The billing cycle. Valid values:
	//
	// - Month: Monthly billing.
	//
	// - Hour: Hourly billing.
	//
	// >
	//
	// > - PrePaid supports only Month.
	//
	// > - PostPaid supports only Hour.
	//
	// > - For Shared instances, this parameter is automatically set to Hour and does not need to be specified.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
	// The region ID. You can go to <props="china">[service registration endpoints](https://www.alibabacloud.com/help/en/hologres/developer-reference/api-hologram-2022-06-01-endpoint)
	//
	// <props="intl">[service registration endpoints](https://www.alibabacloud.com/help/zh/hologres/developer-reference/api-hologram-2022-06-01-endpoint) to view the supported regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group. If this parameter is left empty, the default resource group of the account is used.
	//
	// example:
	//
	// ""
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The standard storage capacity of the instance. Unit: GB.
	//
	// > This parameter is ignored for pay-as-you-go instances.
	//
	// example:
	//
	// 500
	StorageSize *int64 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The storage type.
	//
	// example:
	//
	// redundant
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The ID of the vSwitch. The zone ID of the vSwitch must be the same as that of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2vccsiymtxxxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC). The region ID of the VPC must be the same as that of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-t4netc3y5xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The zone ID. Refer to the usage notes.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequest) GetColdStorageSize() *int64 {
	return s.ColdStorageSize
}

func (s *CreateInstanceRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *CreateInstanceRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateInstanceRequest) GetEnableServerlessComputing() *bool {
	return s.EnableServerlessComputing
}

func (s *CreateInstanceRequest) GetGatewayCount() *int64 {
	return s.GatewayCount
}

func (s *CreateInstanceRequest) GetInitialDatabases() *string {
	return s.InitialDatabases
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetLeaderInstanceId() *string {
	return s.LeaderInstanceId
}

func (s *CreateInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreateInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) SetAutoPay(v bool) *CreateInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetColdStorageSize(v int64) *CreateInstanceRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetCpu(v int64) *CreateInstanceRequest {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int64) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetEnableServerlessComputing(v bool) *CreateInstanceRequest {
	s.EnableServerlessComputing = &v
	return s
}

func (s *CreateInstanceRequest) SetGatewayCount(v int64) *CreateInstanceRequest {
	s.GatewayCount = &v
	return s
}

func (s *CreateInstanceRequest) SetInitialDatabases(v string) *CreateInstanceRequest {
	s.InitialDatabases = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetLeaderInstanceId(v string) *CreateInstanceRequest {
	s.LeaderInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSize(v int64) *CreateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageType(v string) *CreateInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
