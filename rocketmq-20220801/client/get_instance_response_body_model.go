// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetDynamicCode(v string) *GetInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response result.
	Data *GetInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Each request has a unique ID, which you can use to troubleshoot and locate issues.
	//
	// example:
	//
	// 05AB7FBD-F1D3-5D87-BF78-BD782249****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicCode(v string) *GetInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicMessage(v string) *GetInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyData struct {
	// The account information.
	AccountInfo *GetInstanceResponseBodyDataAccountInfo `json:"accountInfo,omitempty" xml:"accountInfo,omitempty" type:"Struct"`
	// The access control information.
	AclInfo *GetInstanceResponseBodyDataAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// Indicates whether auto-renewal is enabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// Indicates whether auto-renewal is enabled.
	//
	// example:
	//
	// false
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The BID of the sales channel for the commodity.
	//
	// example:
	//
	// 26842
	Bid *string `json:"bid,omitempty" xml:"bid,omitempty"`
	// The commodity code of an MSMQ for RocketMQ 5.x series instance, such as ons_rmqsub_public_cn.
	//
	// example:
	//
	// ons_rmqsub_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created. The value is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the instance expires. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-09-01 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The extended configuration. This field is no longer recommended. Use the productInfo, internetInfo, and aclInfo fields instead.
	ExtConfig *GetInstanceResponseBodyDataExtConfig `json:"extConfig,omitempty" xml:"extConfig,omitempty" type:"Struct"`
	// The number of consumer groups.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The capacity information.
	InstanceQuotas []*GetInstanceResponseBodyDataInstanceQuotas `json:"instanceQuotas,omitempty" xml:"instanceQuotas,omitempty" type:"Repeated"`
	// The network information.
	NetworkInfo *GetInstanceResponseBodyDataNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The payment type of the instance.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The extended configuration of the instance.
	ProductInfo *GetInstanceResponseBodyDataProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time when the instance was released. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-09-07 00:00:00
	ReleaseTime *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	// The remarks for the instance.
	//
	// example:
	//
	// This is remark for instance.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm3tmjruyribi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary series code of the instance. For differences between primary series, see [Product selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The service code of the instance. The service code for MSMQ for RocketMQ is rmq.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The instance software information.
	Software *GetInstanceResponseBodyDataSoftware `json:"software,omitempty" xml:"software,omitempty" type:"Struct"`
	// The time when the instance was started. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-series code of the instance. For differences between sub-series, see [Product selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The resource tags.
	Tags []*GetInstanceResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The number of topics.
	//
	// example:
	//
	// 10
	TopicCount *int64 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	// The time when the instance was last modified. The value is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-02 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user to whom the instance belongs.
	//
	// example:
	//
	// 111111111111111
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetAccountInfo() *GetInstanceResponseBodyDataAccountInfo {
	return s.AccountInfo
}

func (s *GetInstanceResponseBodyData) GetAclInfo() *GetInstanceResponseBodyDataAclInfo {
	return s.AclInfo
}

func (s *GetInstanceResponseBodyData) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetInstanceResponseBodyData) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *GetInstanceResponseBodyData) GetBid() *string {
	return s.Bid
}

func (s *GetInstanceResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetInstanceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetInstanceResponseBodyData) GetExtConfig() *GetInstanceResponseBodyDataExtConfig {
	return s.ExtConfig
}

func (s *GetInstanceResponseBodyData) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *GetInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBodyData) GetInstanceQuotas() []*GetInstanceResponseBodyDataInstanceQuotas {
	return s.InstanceQuotas
}

func (s *GetInstanceResponseBodyData) GetNetworkInfo() *GetInstanceResponseBodyDataNetworkInfo {
	return s.NetworkInfo
}

func (s *GetInstanceResponseBodyData) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetInstanceResponseBodyData) GetProductInfo() *GetInstanceResponseBodyDataProductInfo {
	return s.ProductInfo
}

func (s *GetInstanceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *GetInstanceResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBodyData) GetSeriesCode() *string {
	return s.SeriesCode
}

func (s *GetInstanceResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetInstanceResponseBodyData) GetSoftware() *GetInstanceResponseBodyDataSoftware {
	return s.Software
}

func (s *GetInstanceResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyData) GetSubSeriesCode() *string {
	return s.SubSeriesCode
}

func (s *GetInstanceResponseBodyData) GetTags() []*GetInstanceResponseBodyDataTags {
	return s.Tags
}

func (s *GetInstanceResponseBodyData) GetTopicCount() *int64 {
	return s.TopicCount
}

func (s *GetInstanceResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetInstanceResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceResponseBodyData) SetAccountInfo(v *GetInstanceResponseBodyDataAccountInfo) *GetInstanceResponseBodyData {
	s.AccountInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetAclInfo(v *GetInstanceResponseBodyDataAclInfo) *GetInstanceResponseBodyData {
	s.AclInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetAutoRenew(v bool) *GetInstanceResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetAutoRenewPeriod(v int32) *GetInstanceResponseBodyData {
	s.AutoRenewPeriod = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetBid(v string) *GetInstanceResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCommodityCode(v string) *GetInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreateTime(v string) *GetInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExpireTime(v string) *GetInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExtConfig(v *GetInstanceResponseBodyDataExtConfig) *GetInstanceResponseBodyData {
	s.ExtConfig = v
	return s
}

func (s *GetInstanceResponseBodyData) SetGroupCount(v int64) *GetInstanceResponseBodyData {
	s.GroupCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceQuotas(v []*GetInstanceResponseBodyDataInstanceQuotas) *GetInstanceResponseBodyData {
	s.InstanceQuotas = v
	return s
}

func (s *GetInstanceResponseBodyData) SetNetworkInfo(v *GetInstanceResponseBodyDataNetworkInfo) *GetInstanceResponseBodyData {
	s.NetworkInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetPaymentType(v string) *GetInstanceResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetProductInfo(v *GetInstanceResponseBodyDataProductInfo) *GetInstanceResponseBodyData {
	s.ProductInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetRegionId(v string) *GetInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetReleaseTime(v string) *GetInstanceResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRemark(v string) *GetInstanceResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetResourceGroupId(v string) *GetInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServiceCode(v string) *GetInstanceResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSoftware(v *GetInstanceResponseBodyDataSoftware) *GetInstanceResponseBodyData {
	s.Software = v
	return s
}

func (s *GetInstanceResponseBodyData) SetStartTime(v string) *GetInstanceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSubSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SubSeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTags(v []*GetInstanceResponseBodyDataTags) *GetInstanceResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyData) SetTopicCount(v int64) *GetInstanceResponseBodyData {
	s.TopicCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUpdateTime(v string) *GetInstanceResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUserId(v string) *GetInstanceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
	if s.AccountInfo != nil {
		if err := s.AccountInfo.Validate(); err != nil {
			return err
		}
	}
	if s.AclInfo != nil {
		if err := s.AclInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ExtConfig != nil {
		if err := s.ExtConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceQuotas != nil {
		for _, item := range s.InstanceQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.Software != nil {
		if err := s.Software.Validate(); err != nil {
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

type GetInstanceResponseBodyDataAccountInfo struct {
	// The username of the instance. When you use public network access, you must configure the username and password of the instance in the SDK code for identity verification.
	//
	// example:
	//
	// 6W0xz2uPfiwp****
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceResponseBodyDataAccountInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataAccountInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAccountInfo) GetUsername() *string {
	return s.Username
}

func (s *GetInstanceResponseBodyDataAccountInfo) SetUsername(v string) *GetInstanceResponseBodyDataAccountInfo {
	s.Username = &v
	return s
}

func (s *GetInstanceResponseBodyDataAccountInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataAclInfo struct {
	// Deprecated
	//
	// The authentication type of the instance. This parameter is deprecated. Use the aclTypes parameter instead.
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The authentication types of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether the authentication-free in VPC feature is enabled.
	//
	// example:
	//
	// true
	DefaultVpcAuthFree *bool `json:"defaultVpcAuthFree,omitempty" xml:"defaultVpcAuthFree,omitempty"`
}

func (s GetInstanceResponseBodyDataAclInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataAclInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAclInfo) GetAclType() *string {
	return s.AclType
}

func (s *GetInstanceResponseBodyDataAclInfo) GetAclTypes() []*string {
	return s.AclTypes
}

func (s *GetInstanceResponseBodyDataAclInfo) GetDefaultVpcAuthFree() *bool {
	return s.DefaultVpcAuthFree
}

func (s *GetInstanceResponseBodyDataAclInfo) SetAclType(v string) *GetInstanceResponseBodyDataAclInfo {
	s.AclType = &v
	return s
}

func (s *GetInstanceResponseBodyDataAclInfo) SetAclTypes(v []*string) *GetInstanceResponseBodyDataAclInfo {
	s.AclTypes = v
	return s
}

func (s *GetInstanceResponseBodyDataAclInfo) SetDefaultVpcAuthFree(v bool) *GetInstanceResponseBodyDataAclInfo {
	s.DefaultVpcAuthFree = &v
	return s
}

func (s *GetInstanceResponseBodyDataAclInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataExtConfig struct {
	// The authentication type of the instance.
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// Indicates whether burstable elastic capacity beyond the base specification is enabled.
	//
	// This parameter takes effect only when the value of supportAutoScaling is enable, which means the instance supports enabling elastic capacity.
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The Internet bandwidth specification. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The billing type for public network usage.
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Indicates whether public network access is enabled.
	//
	// The instance supports VPC access by default. If public network access is enabled, ApsaraMQ for RocketMQ charges fees for public downstream bandwidth. For more information about billing, see [Public network access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The message retention period. Unit: hours.
	//
	// For the value range, see [Limits](https://help.aliyun.com/document_detail/440347.html) > Resource Quotas > Message retention period.
	//
	// The message storage of ApsaraMQ for RocketMQ is fully serverless and elastic. You are charged based on the actual storage space used. You can control the storage capacity by adjusting the message retention period. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The message processing specification. For the TPS upper limit of message sending and receiving, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio of message sending to message receiving.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Indicates whether the instance supports enabling burstable elastic capacity beyond the base specification.
	//
	// After burstable elastic capacity is enabled, ApsaraMQ for RocketMQ allows the instance to exceed the TPS limit of the base specification within a certain range. Additional elastic specification fees are charged for the portion that exceeds the base specification. For more information about billing, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	// > Only instances of specific specifications support burstable elastic capacity. For more information about supported instances, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	SupportAutoScaling *bool `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
}

func (s GetInstanceResponseBodyDataExtConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataExtConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataExtConfig) GetAclType() *string {
	return s.AclType
}

func (s *GetInstanceResponseBodyDataExtConfig) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *GetInstanceResponseBodyDataExtConfig) GetFlowOutBandwidth() *int32 {
	return s.FlowOutBandwidth
}

func (s *GetInstanceResponseBodyDataExtConfig) GetFlowOutType() *string {
	return s.FlowOutType
}

func (s *GetInstanceResponseBodyDataExtConfig) GetInternetSpec() *string {
	return s.InternetSpec
}

func (s *GetInstanceResponseBodyDataExtConfig) GetMessageRetentionTime() *int32 {
	return s.MessageRetentionTime
}

func (s *GetInstanceResponseBodyDataExtConfig) GetMsgProcessSpec() *string {
	return s.MsgProcessSpec
}

func (s *GetInstanceResponseBodyDataExtConfig) GetSendReceiveRatio() *float32 {
	return s.SendReceiveRatio
}

func (s *GetInstanceResponseBodyDataExtConfig) GetSupportAutoScaling() *bool {
	return s.SupportAutoScaling
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAclType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.AclType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutBandwidth(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutBandwidth = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetInternetSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.InternetSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMessageRetentionTime(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.MessageRetentionTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMsgProcessSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.MsgProcessSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSendReceiveRatio(v float32) *GetInstanceResponseBodyDataExtConfig {
	s.SendReceiveRatio = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSupportAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.SupportAutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataInstanceQuotas struct {
	// The free quota. Unit: count.
	//
	// example:
	//
	// 20
	FreeCount *float64 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	// The quota name of the instance.
	//
	// example:
	//
	// TOPIC_COUNT
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The total quota. Unit: count.
	//
	// example:
	//
	// 100
	TotalCount *float64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The used quota. Unit: count.
	//
	// example:
	//
	// 10
	UsedCount *float64 `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
}

func (s GetInstanceResponseBodyDataInstanceQuotas) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataInstanceQuotas) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) GetFreeCount() *float64 {
	return s.FreeCount
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) GetTotalCount() *float64 {
	return s.TotalCount
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) GetUsedCount() *float64 {
	return s.UsedCount
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetFreeCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.FreeCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetQuotaName(v string) *GetInstanceResponseBodyDataInstanceQuotas {
	s.QuotaName = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetTotalCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetUsedCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.UsedCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNetworkInfo struct {
	// The endpoint information.
	Endpoints []*GetInstanceResponseBodyDataNetworkInfoEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The public network information of the instance.
	InternetInfo *GetInstanceResponseBodyDataNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The VPC information.
	VpcInfo *GetInstanceResponseBodyDataNetworkInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBodyDataNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfo) GetEndpoints() []*GetInstanceResponseBodyDataNetworkInfoEndpoints {
	return s.Endpoints
}

func (s *GetInstanceResponseBodyDataNetworkInfo) GetInternetInfo() *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	return s.InternetInfo
}

func (s *GetInstanceResponseBodyDataNetworkInfo) GetVpcInfo() *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	return s.VpcInfo
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetEndpoints(v []*GetInstanceResponseBodyDataNetworkInfoEndpoints) *GetInstanceResponseBodyDataNetworkInfo {
	s.Endpoints = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetInternetInfo(v *GetInstanceResponseBodyDataNetworkInfoInternetInfo) *GetInstanceResponseBodyDataNetworkInfo {
	s.InternetInfo = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetVpcInfo(v *GetInstanceResponseBodyDataNetworkInfoVpcInfo) *GetInstanceResponseBodyDataNetworkInfo {
	s.VpcInfo = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetInstanceResponseBodyDataNetworkInfoEndpoints struct {
	// The associated PrivateLink endpoint instance.
	//
	// example:
	//
	// ep-bpxxx
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	// The endpoint type of the instance.
	//
	// example:
	//
	// TCP_INTERNET
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// rmq-cn-c4d2tbk****-vpc.cn-hangzhou.rmq.aliyuncs.com:8080
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The IP whitelist for public network access. IP whitelists are supported only for Internet endpoints and are not supported for VPC endpoints.
	//
	// - If no IP whitelist is configured, all IP address ranges are allowed to access the MSMQ for RocketMQ queue broker over the Internet by default.
	//
	// - If an IP whitelist is configured, only IP addresses in the whitelist are allowed to access the MSMQ for RocketMQ queue broker over the Internet.
	//
	// This field is no longer recommended. Use the internetInfo.ipWhitelist field instead.
	//
	// example:
	//
	// 192.168.x.x/24
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetIpWhitelist() []*string {
	return s.IpWhitelist
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointId(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointType(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointUrl(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointUrl = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetIpWhitelist(v []*string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.IpWhitelist = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNetworkInfoInternetInfo struct {
	// The Internet bandwidth specification. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The billing type for public network usage.
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Indicates whether public network access is enabled.
	//
	// The instance supports VPC access by default. If public network access is enabled, ApsaraMQ for RocketMQ charges fees for public downstream bandwidth. For more information about billing, see [Public network access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The IP whitelist for public network access.
	//
	// - If the IP whitelist is not configured, all IP address ranges are allowed to access the ApsaraMQ for RocketMQ broker over the public network by default.
	//
	// - If the IP whitelist is configured, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the public network.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyDataNetworkInfoInternetInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) GetFlowOutBandwidth() *int32 {
	return s.FlowOutBandwidth
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) GetFlowOutType() *string {
	return s.FlowOutType
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) GetInternetSpec() *string {
	return s.InternetSpec
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) GetIpWhitelist() []*string {
	return s.IpWhitelist
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetFlowOutBandwidth(v int32) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.FlowOutBandwidth = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetFlowOutType(v string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.FlowOutType = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetInternetSpec(v string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.InternetSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNetworkInfoVpcInfo struct {
	// The security group ID.
	//
	// example:
	//
	// sg-hp35r2hc3a3sv8q2sb16
	SecurityGroupIds *string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch associated with the instance.
	//
	// <notice>This field is not recommended. Use the vSwitches field instead. For compatibility, this field currently returns only one vSwitch.</notice>
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The list of vSwitches.
	VSwitches []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC associated with the instance.
	//
	// example:
	//
	// vpc-uf6of9452b2pba82c****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) GetVSwitches() []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches {
	return s.VSwitches
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetSecurityGroupIds(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.SecurityGroupIds = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVSwitchId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVSwitches(v []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VSwitches = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVpcId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) Validate() error {
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

type GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) SetVSwitchId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) SetZoneId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataProductInfo struct {
	// The adaptive elasticity.
	//
	// example:
	//
	// false
	AdaptiveElasticity *bool `json:"adaptiveElasticity,omitempty" xml:"adaptiveElasticity,omitempty"`
	// Indicates whether burstable elastic capacity beyond the base specification is enabled for seamless scaling.
	//
	// This parameter takes effect only when the value of supportAutoScaling is enable, which means the instance supports enabling elastic capacity.
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The capacity mode.
	//
	// example:
	//
	// provisioned
	CapacityType *string `json:"capacityType,omitempty" xml:"capacityType,omitempty"`
	// The cross-region replication mode.
	//
	// example:
	//
	// async_cross_az
	DrReplicationMode *string `json:"drReplicationMode,omitempty" xml:"drReplicationMode,omitempty"`
	// The maximum interval for scheduled messages. Unit: seconds.
	//
	// example:
	//
	// 3456000
	MaxDelayTime *int32 `json:"maxDelayTime,omitempty" xml:"maxDelayTime,omitempty"`
	// The message retention period. Unit: hours.
	//
	// For the value range, see [Limits](https://help.aliyun.com/document_detail/440347.html) > Resource Quotas > Message retention period.
	//
	// The message storage of ApsaraMQ for RocketMQ is fully serverless and elastic. You are charged based on the actual storage space used. You can control the storage capacity by adjusting the message retention period. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The message processing specification. For the TPS upper limit of message sending and receiving, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The provisioned capacity.
	//
	// example:
	//
	// 3000
	ProvisionedCapacity *int64 `json:"provisionedCapacity,omitempty" xml:"provisionedCapacity,omitempty"`
	// The ratio of message sending to message receiving.
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
	// The storage encryption key.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// Indicates whether the burstable elastic capability beyond the base specification is enabled.
	//
	//
	// After the burstable elastic capability is enabled, ApsaraMQ for RocketMQ allows the instance to exceed the TPS limit of the base specification within a certain range. Additional elastic specification fees are charged for the portion that exceeds the base specification. For more information about billing, see [Computing capability fees](https://help.aliyun.com/document_detail/427237.html).
	//
	//
	// > Only instances of specific specifications support the burstable elastic capability. For more information about supported instances, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	SupportAutoScaling *bool `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
	// Indicates whether the message trace feature is enabled.
	//
	// - true: The message trace feature is enabled.
	//
	// - false: The message trace feature is disabled.
	//
	// <props="intl">This parameter does not take effect. Instances support the message trace feature by default regardless of whether this parameter is enabled.
	//
	// <props="china">This parameter takes effect only for Serverless instances. If the message trace feature is enabled, message trace fees are incurred. For more information about billing, see [Serverless billing](https://help.aliyun.com/document_detail/2586280.html).
	//
	// <props="china">For pay-as-you-go and subscription instances, this parameter does not take effect. The message trace feature is supported by default regardless of whether this parameter is enabled.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s GetInstanceResponseBodyDataProductInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataProductInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataProductInfo) GetAdaptiveElasticity() *bool {
	return s.AdaptiveElasticity
}

func (s *GetInstanceResponseBodyDataProductInfo) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *GetInstanceResponseBodyDataProductInfo) GetCapacityType() *string {
	return s.CapacityType
}

func (s *GetInstanceResponseBodyDataProductInfo) GetDrReplicationMode() *string {
	return s.DrReplicationMode
}

func (s *GetInstanceResponseBodyDataProductInfo) GetMaxDelayTime() *int32 {
	return s.MaxDelayTime
}

func (s *GetInstanceResponseBodyDataProductInfo) GetMessageRetentionTime() *int32 {
	return s.MessageRetentionTime
}

func (s *GetInstanceResponseBodyDataProductInfo) GetMsgProcessSpec() *string {
	return s.MsgProcessSpec
}

func (s *GetInstanceResponseBodyDataProductInfo) GetProvisionedCapacity() *int64 {
	return s.ProvisionedCapacity
}

func (s *GetInstanceResponseBodyDataProductInfo) GetSendReceiveRatio() *float32 {
	return s.SendReceiveRatio
}

func (s *GetInstanceResponseBodyDataProductInfo) GetStorageEncryption() *bool {
	return s.StorageEncryption
}

func (s *GetInstanceResponseBodyDataProductInfo) GetStorageSecretKey() *string {
	return s.StorageSecretKey
}

func (s *GetInstanceResponseBodyDataProductInfo) GetSupportAutoScaling() *bool {
	return s.SupportAutoScaling
}

func (s *GetInstanceResponseBodyDataProductInfo) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *GetInstanceResponseBodyDataProductInfo) SetAdaptiveElasticity(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.AdaptiveElasticity = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetAutoScaling(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetCapacityType(v string) *GetInstanceResponseBodyDataProductInfo {
	s.CapacityType = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetDrReplicationMode(v string) *GetInstanceResponseBodyDataProductInfo {
	s.DrReplicationMode = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetMaxDelayTime(v int32) *GetInstanceResponseBodyDataProductInfo {
	s.MaxDelayTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetMessageRetentionTime(v int32) *GetInstanceResponseBodyDataProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetMsgProcessSpec(v string) *GetInstanceResponseBodyDataProductInfo {
	s.MsgProcessSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetProvisionedCapacity(v int64) *GetInstanceResponseBodyDataProductInfo {
	s.ProvisionedCapacity = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetSendReceiveRatio(v float32) *GetInstanceResponseBodyDataProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetStorageEncryption(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.StorageEncryption = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetStorageSecretKey(v string) *GetInstanceResponseBodyDataProductInfo {
	s.StorageSecretKey = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetSupportAutoScaling(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.SupportAutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetTraceOn(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.TraceOn = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataSoftware struct {
	// Deprecated
	//
	// The upgrade time window, in the format of HH:mm-HH:mm.
	//
	// example:
	//
	// 02:00-06:00
	MaintainTime *string `json:"maintainTime,omitempty" xml:"maintainTime,omitempty"`
	// The software version.
	//
	// example:
	//
	// 5.0-rmq-20230619-1
	SoftwareVersion *string `json:"softwareVersion,omitempty" xml:"softwareVersion,omitempty"`
	// The upgrade method.
	//
	// example:
	//
	// auto
	UpgradeMethod *string `json:"upgradeMethod,omitempty" xml:"upgradeMethod,omitempty"`
}

func (s GetInstanceResponseBodyDataSoftware) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataSoftware) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataSoftware) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *GetInstanceResponseBodyDataSoftware) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *GetInstanceResponseBodyDataSoftware) GetUpgradeMethod() *string {
	return s.UpgradeMethod
}

func (s *GetInstanceResponseBodyDataSoftware) SetMaintainTime(v string) *GetInstanceResponseBodyDataSoftware {
	s.MaintainTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataSoftware) SetSoftwareVersion(v string) *GetInstanceResponseBodyDataSoftware {
	s.SoftwareVersion = &v
	return s
}

func (s *GetInstanceResponseBodyDataSoftware) SetUpgradeMethod(v string) *GetInstanceResponseBodyDataSoftware {
	s.UpgradeMethod = &v
	return s
}

func (s *GetInstanceResponseBodyDataSoftware) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the resource tag.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetInstanceResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyDataTags) SetKey(v string) *GetInstanceResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyDataTags) SetValue(v string) *GetInstanceResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
