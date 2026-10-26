// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclInfo(v *CreateInstanceRequestAclInfo) *CreateInstanceRequest
	GetAclInfo() *CreateInstanceRequestAclInfo
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceRequest
	GetAutoRenewPeriod() *int32
	SetCommodityCode(v string) *CreateInstanceRequest
	GetCommodityCode() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetNetworkInfo(v *CreateInstanceRequestNetworkInfo) *CreateInstanceRequest
	GetNetworkInfo() *CreateInstanceRequestNetworkInfo
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateInstanceRequest
	GetPeriodUnit() *string
	SetProductInfo(v *CreateInstanceRequestProductInfo) *CreateInstanceRequest
	GetProductInfo() *CreateInstanceRequestProductInfo
	SetRemark(v string) *CreateInstanceRequest
	GetRemark() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetSeriesCode(v string) *CreateInstanceRequest
	GetSeriesCode() *string
	SetServiceCode(v string) *CreateInstanceRequest
	GetServiceCode() *string
	SetSubSeriesCode(v string) *CreateInstanceRequest
	GetSubSeriesCode() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
}

type CreateInstanceRequest struct {
	// The access control information.
	AclInfo *CreateInstanceRequestAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when the billing method of the instance is set to Subscription.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The Unified Auto Renewal Cycle. This parameter takes effect only when auto-renewal is enabled. Unit: months.
	//
	// Valid values:
	//
	// - Monthly renewal: 1, 2, 3, 6, 12
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// Deprecated
	//
	// The commodity code.
	//
	// <props="china">- ons_rmqsub_public_cn: subscription instance.
	//
	// <props="china">- ons_rmqpost_public_cn: pay-as-you-go instance.
	//
	// <props="china">- ons_rmqsrvlesspost_public_cn: serverless instance.
	//
	// <props="intl">- ons_rmqpost_public_intl: pay-as-you-go.
	//
	// <props="intl">- ons_rmqsub_public_intl: subscription.
	//
	// <props="intl">- ons_rmqsrvlesspost_public_intl: serverless instance.
	//
	// This parameter is required for serverless instances.
	//
	// example:
	//
	// ons_ rmqpost_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The name of the instance to be created.
	//
	// If you do not specify this parameter, the instance ID is used as the instance name.
	//
	// example:
	//
	// rmq-cn-72u3048uxxx
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The network configuration information.
	//
	// This parameter is required.
	NetworkInfo *CreateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The billing method of the instance. ApsaraMQ for RocketMQ supports subscription and pay-as-you-go billing methods.
	//
	// For more information, see [Billable methods](https://help.aliyun.com/document_detail/427234.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The subscription duration. This parameter takes effect only when the billing method of the instance is set to Subscription.
	//
	// Valid values:
	//
	// - Monthly subscription: 1, 2, 3, 4, 5, 6
	//
	// - Yearly subscription: 1, 2, 3
	//
	// example:
	//
	// 3
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The minimum epoch unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	// The instance specification information.
	ProductInfo *CreateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The remarks of the instance.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzy6pist7uuna
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary series encoding of the instance. For more information about the differences between primary series, see [Product selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The service encoding of the instance. The service encoding for MSMQ for RocketMQ is rmq.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The sub-series encoding of the instance. For more information about the differences between sub-series, see [Product selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// If the primary series is set to ultimate (Platinum Edition), the sub-series can only be set to cluster_ha (Cluster High Availability Edition).
	//
	// 	Notice: After the instance is created, the sub-series cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The list of resource tags.
	Tags []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. The client generates the value of this parameter, which must be unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// c2c5d1274a8d4317a13bc5b0d4******
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAclInfo() *CreateInstanceRequestAclInfo {
	return s.AclInfo
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetNetworkInfo() *CreateInstanceRequestNetworkInfo {
	return s.NetworkInfo
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateInstanceRequest) GetProductInfo() *CreateInstanceRequestProductInfo {
	return s.ProductInfo
}

func (s *CreateInstanceRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetSeriesCode() *string {
	return s.SeriesCode
}

func (s *CreateInstanceRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreateInstanceRequest) GetSubSeriesCode() *string {
	return s.SubSeriesCode
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) SetAclInfo(v *CreateInstanceRequestAclInfo) *CreateInstanceRequest {
	s.AclInfo = v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetCommodityCode(v string) *CreateInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkInfo(v *CreateInstanceRequestNetworkInfo) *CreateInstanceRequest {
	s.NetworkInfo = v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetProductInfo(v *CreateInstanceRequestProductInfo) *CreateInstanceRequest {
	s.ProductInfo = v
	return s
}

func (s *CreateInstanceRequest) SetRemark(v string) *CreateInstanceRequest {
	s.Remark = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSeriesCode(v string) *CreateInstanceRequest {
	s.SeriesCode = &v
	return s
}

func (s *CreateInstanceRequest) SetServiceCode(v string) *CreateInstanceRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreateInstanceRequest) SetSubSeriesCode(v string) *CreateInstanceRequest {
	s.SubSeriesCode = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.AclInfo != nil {
		if err := s.AclInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInfo != nil {
		if err := s.NetworkInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProductInfo != nil {
		if err := s.ProductInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestAclInfo struct {
	// Specifies whether to enable authentication-free access for VPC endpoints when the intelligent identity recognition type is used. Valid values:
	//
	// - true: Enable authentication-free access.
	//
	// - false: Disable authentication-free access.
	//
	// example:
	//
	// false
	DefaultVpcAuthFree *bool `json:"defaultVpcAuthFree,omitempty" xml:"defaultVpcAuthFree,omitempty"`
}

func (s CreateInstanceRequestAclInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestAclInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestAclInfo) GetDefaultVpcAuthFree() *bool {
	return s.DefaultVpcAuthFree
}

func (s *CreateInstanceRequestAclInfo) SetDefaultVpcAuthFree(v bool) *CreateInstanceRequestAclInfo {
	s.DefaultVpcAuthFree = &v
	return s
}

func (s *CreateInstanceRequestAclInfo) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestNetworkInfo struct {
	// The Internet configuration information.
	//
	// This parameter is required.
	InternetInfo *CreateInstanceRequestNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The VPC configuration information.
	//
	// This parameter is required.
	VpcInfo *CreateInstanceRequestNetworkInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfo) GetInternetInfo() *CreateInstanceRequestNetworkInfoInternetInfo {
	return s.InternetInfo
}

func (s *CreateInstanceRequestNetworkInfo) GetVpcInfo() *CreateInstanceRequestNetworkInfoVpcInfo {
	return s.VpcInfo
}

func (s *CreateInstanceRequestNetworkInfo) SetInternetInfo(v *CreateInstanceRequestNetworkInfoInternetInfo) *CreateInstanceRequestNetworkInfo {
	s.InternetInfo = v
	return s
}

func (s *CreateInstanceRequestNetworkInfo) SetVpcInfo(v *CreateInstanceRequestNetworkInfoVpcInfo) *CreateInstanceRequestNetworkInfo {
	s.VpcInfo = v
	return s
}

func (s *CreateInstanceRequestNetworkInfo) Validate() error {
	if s.InternetInfo != nil {
		if err := s.InternetInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VpcInfo != nil {
		if err := s.VpcInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceRequestNetworkInfoInternetInfo struct {
	// The Internet bandwidth specification. Unit: Mb/s.
	//
	// This parameter is required only when the Internet billing type is set to fixed bandwidth billing.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The Internet billing type.
	//
	// This parameter is required.
	//
	// example:
	//
	// uninvolved
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Specifies whether to enable public network access.
	//
	// By default, instances support VPC access. If you enable public network access, ApsaraMQ for RocketMQ charges fees for public network downstream bandwidth. For more information about billing, see [Public network access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// Deprecated
	//
	// The public network access IP whitelist. Only public endpoints support IP whitelist configuration. VPC endpoints do not support this feature.
	//
	// - If no IP whitelist is configured, all IP address ranges are allowed to access the ApsaraMQ for RocketMQ server over the Internet by default.
	//
	// - If an IP whitelist is configured, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ server over the Internet.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequestNetworkInfoInternetInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) GetFlowOutBandwidth() *int32 {
	return s.FlowOutBandwidth
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) GetFlowOutType() *string {
	return s.FlowOutType
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) GetInternetSpec() *string {
	return s.InternetSpec
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) GetIpWhitelist() []*string {
	return s.IpWhitelist
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetFlowOutBandwidth(v int32) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.FlowOutBandwidth = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetFlowOutType(v string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.FlowOutType = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetInternetSpec(v string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.InternetSpec = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestNetworkInfoVpcInfo struct {
	// The security group ID.
	//
	// example:
	//
	// sg-bp17hpmgz96tvnsdy6so
	SecurityGroupIds *string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch associated with the instance. If multiple vSwitches exist, separate them with vertical bars (|).
	//
	// > After an ApsaraMQ for RocketMQ instance is created, the selected vSwitch cannot be changed. To modify the associated vSwitch, release the existing ApsaraMQ for RocketMQ instance and purchase a new one.
	//
	// <notice>This field is not recommended. Use the vSwitches field instead.</notice>
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The list of vSwitches.
	//
	// > After an ApsaraMQ for RocketMQ instance is created, the selected vSwitch cannot be changed. To modify the associated vSwitch, release the existing ApsaraMQ for RocketMQ instance and purchase a new one.
	//
	// <notice>This is a required field. Because the vSwitchId field is no longer recommended, use this field to configure vSwitches.</notice>
	VSwitches []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC to associate with the instance.
	//
	// > After an ApsaraMQ for RocketMQ instance is created, the selected VPC cannot be changed. To modify the associated VPC, release the existing ApsaraMQ for RocketMQ instance and purchase a new one.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9qt50xhtj9krb******
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateInstanceRequestNetworkInfoVpcInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) GetVSwitches() []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches {
	return s.VSwitches
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetSecurityGroupIds(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.SecurityGroupIds = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVSwitchId(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVSwitches(v []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VSwitches = v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVpcId(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestNetworkInfoVpcInfoVSwitches struct {
	// The ID of the vSwitch associated with the instance.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s CreateInstanceRequestNetworkInfoVpcInfoVSwitches) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoVpcInfoVSwitches) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoVpcInfoVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequestNetworkInfoVpcInfoVSwitches) SetVSwitchId(v string) *CreateInstanceRequestNetworkInfoVpcInfoVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfoVSwitches) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestProductInfo struct {
	// Specifies whether to enable adaptive elasticity. A value of true indicates that adaptive elasticity is enabled. A value of false indicates that adaptive elasticity is shutdown.
	AdaptiveElasticity *bool `json:"adaptiveElasticity,omitempty" xml:"adaptiveElasticity,omitempty"`
	// Specifies whether to enable burstable elastic capacity beyond the specification.
	//
	// After burstable elastic capacity is enabled, MSMQ for RocketMQ allows the instance to exceed the TPS limit of the base specification within a certain range. The portion that exceeds the base specification incurs additional elastic specification fees. For more information about billing, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	//
	// > Only instances of certain specifications support burstable elastic capacity. For more information about supported instances, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The capacity pattern.
	//
	// example:
	//
	// provisioned
	CapacityType *string `json:"capacityType,omitempty" xml:"capacityType,omitempty"`
	// The cross-region replication pattern.
	//
	// example:
	//
	// async_cross_az
	DrReplicationMode *string `json:"drReplicationMode,omitempty" xml:"drReplicationMode,omitempty"`
	// The message retention period. Unit: hours.
	//
	// For the valid value range, see [Limits](https://help.aliyun.com/document_detail/440347.html) > Resource Quotas > Message Retention Period.
	//
	// The log broker in MSMQ for RocketMQ is fully serverless and elastic. Fees are charged based on the actual storage space used. You can control the storage capacity by adjusting the message retention period. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The message sending and receiving computing specification. For more information about the TPS upper limit for message sending and receiving, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The provisioned capacity.
	//
	// example:
	//
	// 5000
	ProvisionedCapacity *int64 `json:"provisionedCapacity,omitempty" xml:"provisionedCapacity,omitempty"`
	// The ratio of message sending TPS to the total message sending and accepting TPS of the instance.
	//
	// For example, if the TPS upper limit for message sending and accepting of the instance is 1000 per second and the message sending ratio is 0.8, the TPS upper limit for message sending is 800 per second and the TPS upper limit for message accepting is 200 per second.
	//
	// Valid values: [0.05, 0.95]. Default value: 0.5.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Indicates whether storage encryption is enabled.
	//
	// example:
	//
	// false
	StorageEncryption *bool `json:"storageEncryption,omitempty" xml:"storageEncryption,omitempty"`
	// The key for storage encryption.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// Specifies whether to enable the message trace feature.
	//
	// true: Enable the message trace feature.
	//
	// false: Shutdown the message trace feature.
	//
	// This parameter takes effect only for Serverless instances. If the message trace feature is enabled, message trace fees are incurred. For more information about billing, see the Serverless billing description. For pay-as-you-go and subscription instances, this parameter does not take effect. The message trace feature is supported by default regardless of whether this parameter is enabled.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s CreateInstanceRequestProductInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestProductInfo) GetAdaptiveElasticity() *bool {
	return s.AdaptiveElasticity
}

func (s *CreateInstanceRequestProductInfo) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *CreateInstanceRequestProductInfo) GetCapacityType() *string {
	return s.CapacityType
}

func (s *CreateInstanceRequestProductInfo) GetDrReplicationMode() *string {
	return s.DrReplicationMode
}

func (s *CreateInstanceRequestProductInfo) GetMessageRetentionTime() *int32 {
	return s.MessageRetentionTime
}

func (s *CreateInstanceRequestProductInfo) GetMsgProcessSpec() *string {
	return s.MsgProcessSpec
}

func (s *CreateInstanceRequestProductInfo) GetProvisionedCapacity() *int64 {
	return s.ProvisionedCapacity
}

func (s *CreateInstanceRequestProductInfo) GetSendReceiveRatio() *float32 {
	return s.SendReceiveRatio
}

func (s *CreateInstanceRequestProductInfo) GetStorageEncryption() *bool {
	return s.StorageEncryption
}

func (s *CreateInstanceRequestProductInfo) GetStorageSecretKey() *string {
	return s.StorageSecretKey
}

func (s *CreateInstanceRequestProductInfo) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *CreateInstanceRequestProductInfo) SetAdaptiveElasticity(v bool) *CreateInstanceRequestProductInfo {
	s.AdaptiveElasticity = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetAutoScaling(v bool) *CreateInstanceRequestProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetCapacityType(v string) *CreateInstanceRequestProductInfo {
	s.CapacityType = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetDrReplicationMode(v string) *CreateInstanceRequestProductInfo {
	s.DrReplicationMode = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetMessageRetentionTime(v int32) *CreateInstanceRequestProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetMsgProcessSpec(v string) *CreateInstanceRequestProductInfo {
	s.MsgProcessSpec = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetProvisionedCapacity(v int64) *CreateInstanceRequestProductInfo {
	s.ProvisionedCapacity = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetSendReceiveRatio(v float32) *CreateInstanceRequestProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetStorageEncryption(v bool) *CreateInstanceRequestProductInfo {
	s.StorageEncryption = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetStorageSecretKey(v string) *CreateInstanceRequestProductInfo {
	s.StorageSecretKey = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetTraceOn(v bool) *CreateInstanceRequestProductInfo {
	s.TraceOn = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// xxxKey
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the resource tag.
	//
	// example:
	//
	// xxxValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
