// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// Specifies whether to enable auto-renewal for the instance. This parameter takes effect only if you set paymentType to Subscription. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal cycle of the instance. This parameter takes effect only if you set autoRenew to true. Unit: months.
	//
	// Valid values:
	//
	// 	- Monthly renewal: 1, 2, 3, 6, and 12
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The commodity code. Valid values:
	//
	// 	- ons_rmqpost_public_intl: pay-as-you-go
	//
	// 	- ons_rmqsub_public_intl: subscription
	//
	// 	- ons_rmqsrvlesspost_public_intl: serverless instance
	//
	// serverless instance requires this parameter
	//
	// example:
	//
	// ons_ rmqpost_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The name of the instance that you want to create.
	//
	// If you leave this parameter empty, the instance ID is used as the instance name.
	//
	// example:
	//
	// rmq-cn-72u3048uxxx
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The information about the network.
	//
	// This parameter is required.
	NetworkInfo *CreateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The billing method of the instance. ApsaraMQ for RocketMQ supports the subscription and pay-as-you-go billing methods.
	//
	// Valid values:
	//
	// 	- PayAsYouGo: This billing method allows you to use resources before you pay for the resources.
	//
	// 	- Subscription: This billing method allows you to use resources after you pay for the resources.
	//
	// For more information, see [Billing methods](https://help.aliyun.com/document_detail/427234.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The subscription duration of the instance. This parameter takes effect only if you set paymentType to Subscription.
	//
	// Valid values:
	//
	// 	- Monthly subscription: 1, 2, 3, 4, 5, and 6
	//
	// 	- Yearly subscription: 1, 2, and 3
	//
	// example:
	//
	// 3
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	// The information about instance specifications.
	ProductInfo *CreateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The instance description.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzy6pist7uuna
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance. For information about the differences among primary edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	//
	// >  After you create an instance, you can only upgrade the primary edition of the instance. The following editions are sorted in ascending order: Standard Edition, Professional Edition, Enterprise Platinum Edition. For example, you can upgrade an instance from Standard Edition to Professional Edition, but you cannot downgrade an instance from Professional Edition to Standard Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The code of the service to which the instance belongs. The service code of ApsaraMQ for RocketMQ is rmq.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The sub-category edition of the instance. For information about the differences among sub-category edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// 	- serverless: serverless
	//
	// If you set seriesCode to ultimate, you can set this parameter only to cluster_ha.
	//
	// >  After you create an instance, you cannot change the sub-category edition of the instance.
	//
	// Valid values:
	//
	// 	- serverless: serverless
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The tags that you want to add to the instance.
	Tags []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value of this parameter, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	return dara.Validate(s)
}

type CreateInstanceRequestNetworkInfo struct {
	// The Internet-related configurations.
	//
	// This parameter is required.
	InternetInfo *CreateInstanceRequestNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The virtual private cloud (VPC)-related configurations.
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
	return dara.Validate(s)
}

type CreateInstanceRequestNetworkInfoInternetInfo struct {
	// The Internet bandwidth. Unit: MB/s.
	//
	// This parameter is required only if you set flowOutType to payByBandwidth.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The metering method of Internet usage.
	//
	// Valid values:
	//
	// 	- payByBandwidth: pay-by-bandwidth. This value is valid only if you enable Internet access.
	//
	// 	- payByTraffic: pay-by-traffic. This value is valid only if you enable Internet access.
	//
	// 	- uninvolved: No metering method is involved. This value is valid only if you disable Internet access.
	//
	// This parameter is required.
	//
	// example:
	//
	// uninvolved
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Specifies whether to enable the Internet access feature.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, ApsaraMQ for RocketMQ allows you to access instances in VPCs. If you enable Internet access for an instance, you can access the instance over the Internet. After you enable this feature, you are charged for outbound Internet traffic. For more information, see [Internet access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// Deprecated
	//
	// The whitelist that includes the CIDR blocks that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet. This parameter can be configured only if you use the public endpoint to access the instance.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the CIDR blocks in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
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
	// The ID of the security group to which the instance belongs.
	//
	// example:
	//
	// sg-bp17hpmgz96tvnsdy6so
	SecurityGroupIds *string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch with which the instance is associated. If you want to specify multiple vSwitches, separate the vSwitches with vertical bars (|).
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the vSwitch with which the instance is associated. If you want to change the vSwitch with which the instance is associated, you must release the instance and purchase a new instance.
	//
	// >  We recommend that you configure vSwitches instead of this parameter.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The vSwitches.
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the vSwitch with which the instance is associated. If you want to change the vSwitch with which the instance is associated, you must release the instance and purchase a new instance.
	//
	// >  This parameter is required. We recommend that you configure this parameter instead of vSwitchId.
	VSwitches []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC with which the instance to be created is associated.
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the VPC with which the instance is associated. If you want to change the VPC with which the instance is associated, you must release the instance and create a new instance.
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
	return dara.Validate(s)
}

type CreateInstanceRequestNetworkInfoVpcInfoVSwitches struct {
	// The ID of the vSwitch with which the instance is associated.
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
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// After you enable the elastic TPS feature for an ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	// >  The elastic TPS feature is supported only by instances of specific editions. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section of the [Limits](https://help.aliyun.com/document_detail/440347.html) topic.
	//
	// ApsaraMQ for RocketMQ supports serverless scaling of message storage. You are charged storage fees based on your actual storage usage. You can change the retention period of messages to manage storage capacity. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that specifies the messaging transactions per second (TPS) of the instance. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio of the message sending TPS to the messaging TPS of the instance.
	//
	// For example, if the maximum messaging TPS of an instance is 1,000 and the ratio of the message sending TPS to the messaging TPS of the instance is 0.8, the maximum message sending TPS of the instance is 800 and the maximum message receiving TPS is 200.
	//
	// Valid values: 0 to 1. Default value: 0.5.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether to enable the encryption at rest feature.
	//
	// example:
	//
	// false
	StorageEncryption *bool `json:"storageEncryption,omitempty" xml:"storageEncryption,omitempty"`
	// The key for encryption at rest.
	//
	// example:
	//
	// xxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
}

func (s CreateInstanceRequestProductInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestProductInfo) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *CreateInstanceRequestProductInfo) GetMessageRetentionTime() *int32 {
	return s.MessageRetentionTime
}

func (s *CreateInstanceRequestProductInfo) GetMsgProcessSpec() *string {
	return s.MsgProcessSpec
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

func (s *CreateInstanceRequestProductInfo) SetAutoScaling(v bool) *CreateInstanceRequestProductInfo {
	s.AutoScaling = &v
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

func (s *CreateInstanceRequestProductInfo) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	// The `key` of the tag.
	//
	// example:
	//
	// aaa
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The `value` of the tag.
	//
	// example:
	//
	// bbb
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
