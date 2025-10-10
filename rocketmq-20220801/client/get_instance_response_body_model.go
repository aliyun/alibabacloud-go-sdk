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
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
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
	// The HTTP status code returned.
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
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 05AB7FBD-F1D3-5D87-BF78-BD782249****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
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
	return dara.Validate(s)
}

type GetInstanceResponseBodyData struct {
	// The account information.
	AccountInfo *GetInstanceResponseBodyDataAccountInfo `json:"accountInfo,omitempty" xml:"accountInfo,omitempty" type:"Struct"`
	// The information about access control.
	AclInfo *GetInstanceResponseBodyDataAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// The business ID (BID) of the commodity.
	//
	// example:
	//
	// 26842
	Bid *string `json:"bid,omitempty" xml:"bid,omitempty"`
	// The commodity code of the instance. The commodity code of a ApsaraMQ for RocketMQ 5.0 instance has a similar format as ons_rmqsub_public_cn.
	//
	// example:
	//
	// ons_rmqsub_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-09-01 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// Deprecated
	//
	// The extended configurations. We recommend you configure productInfo, internetInfo, or aclInfo instead of this parameter.
	ExtConfig *GetInstanceResponseBodyDataExtConfig `json:"extConfig,omitempty" xml:"extConfig,omitempty" type:"Struct"`
	// The number of groups.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The instance quotas.
	InstanceQuotas []*GetInstanceResponseBodyDataInstanceQuotas `json:"instanceQuotas,omitempty" xml:"instanceQuotas,omitempty" type:"Repeated"`
	// The network information.
	NetworkInfo *GetInstanceResponseBodyDataNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The billing method of the instance.
	//
	// Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The extended configurations of the instance.
	ProductInfo *GetInstanceResponseBodyDataProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time when the instance was released.
	//
	// example:
	//
	// 2022-09-07 00:00:00
	ReleaseTime *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// This is remark for instance.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3tmjruyribi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance. For information about the differences between primary edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The code of the service to which the instance belongs. The service code of ApsaraMQ for RocketMQ is rmq.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The instance software information.
	Software *GetInstanceResponseBodyDataSoftware `json:"software,omitempty" xml:"software,omitempty" type:"Struct"`
	// The time when the instance was started.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- RELEASED
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- CHANGING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-category edition of the instance. For information about the differences between sub-category edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
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
	// The time when the instance was last modified.
	//
	// example:
	//
	// 2022-08-02 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user who owns the instance.
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
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataAccountInfo struct {
	// The username of the instance. If you access a ApsaraMQ for RocketMQ instance over the Internet, you must configure the username and password of the instance in the SDK code for authentication.
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
	// The authentication type of the instance. This parameter is no longer in use. We recommend that you configure aclTypes.
	//
	// Valid values:
	//
	// - default: intelligent identity authentication
	//
	// - apache_acl:access control list (ACL) identity authentication**
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The authentication types of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	// Valid value:
	//
	// 	- default: intelligent authentication
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// This parameter is valid only when the supportAutoScaling parameter is set to enable.
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The Internet bandwidth. Unit: MB/s.
	//
	// example:
	//
	// 10
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The metering method of Internet usage.
	//
	// Valid values:
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// 	- paybybandwidth: pay-by-bandwidth
	//
	// 	- uninvolved: N/A
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, you can access ApsaraMQ for RocketMQ instances in virtual private clouds (VPCs). If you enable the Internet access feature, you are charged for Internet outbound bandwidth. For more information, see [Internet access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section in [Usage limits](https://help.aliyun.com/document_detail/440347.html).
	//
	// The storage of messages in ApsaraMQ for RocketMQ is serverless and scalable. You are charged for message storage based on your actual usage. You can change the retention period of messages to adjust storage capacity. For more information, see [Storage fee](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that is used to send and receive messages. For information about the upper limit of TPS, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio between sent messages and received messages in the instance.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether the elastic TPS feature is supported by the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// After you enable the elastic TPS feature for a ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fee](https://help.aliyun.com/document_detail/427237.html).
	//
	// > The elastic TPS feature is supported only by specific instance editions. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
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
	// The number of topics that are free of charge on the instance.
	//
	// example:
	//
	// 20
	FreeCount *float64 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	// The quota name.
	//
	// Valid value:
	//
	// 	- TOPIC_COUNT: the number of topics that can be created on the instance
	//
	// example:
	//
	// TOPIC_COUNT
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The total number of topics on the instance.
	//
	// example:
	//
	// 100
	TotalCount *float64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The number of used topics on the instance.
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
	// The endpoints.
	Endpoints []*GetInstanceResponseBodyDataNetworkInfoEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The information about the Internet.
	InternetInfo *GetInstanceResponseBodyDataNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The virtual private cloud (VPC) information.
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
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataNetworkInfoEndpoints struct {
	// The type of the endpoint that is used to access the instance.
	//
	// Valid values:
	//
	// 	- TCP_VPC: VPC endpoint
	//
	// 	- TCP_INTERNET: public endpoint
	//
	// example:
	//
	// TCP_INTERNET
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The endpoint that is used to access the instance.
	//
	// example:
	//
	// rmq-cn-c4d2tbk****-vpc.cn-hangzhou.rmq.aliyuncs.com:8080
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet. This parameter can be configured only if you use the public endpoint to access the instance.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// We recommend that you configure internetInfo.ipWhitelist instead of this parameter.
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

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) GetIpWhitelist() []*string {
	return s.IpWhitelist
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
	// The Internet bandwidth. Unit: MB/s.
	//
	// example:
	//
	// 1
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The metering method for Internet usage.
	//
	// Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth. If the Internet access feature is enabled, specify this value for the parameter.
	//
	// 	- uninvolved: N/A. If the Internet access feature is not enabled, specify this value for the parameter.
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Specifies whether to enable the Internet access feature.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, ApsaraMQ for RocketMQ instances are accessed in virtual private clouds (VPCs). If you enable the Internet access feature, you are charged for Internet outbound bandwidth. For more information, see [Internet access fee](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker.
	//
	// 	- If this parameter is not configured, all IP addresses are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If this parameter is configured, only the IP addresses that are included in the whitelist can access the ApsaraMQ for RocketMQ broker over the Internet.
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
	// The ID of the vSwitch with which the instance is associated.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The vSwitches.
	VSwitches []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC with which the instance is associated.
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
	return dara.Validate(s)
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
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// This parameter is valid only when the supportAutoScaling parameter is set to enable.
	//
	// example:
	//
	// true
	AutoScaling  *bool   `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	CapacityType *string `json:"capacityType,omitempty" xml:"capacityType,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section in [Usage limits](https://help.aliyun.com/document_detail/440347.html).
	//
	// The storage of messages in ApsaraMQ for RocketMQ is serverless and scalable. You are charged for message storage based on your actual usage. You can change the retention period of messages to adjust storage capacity. For more information, see [Storage fee](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that is used to send and receive messages. For information about the upper limit of TPS, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec      *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	ProvisionedCapacity *int64  `json:"provisionedCapacity,omitempty" xml:"provisionedCapacity,omitempty"`
	// The ratio between sent messages and received messages in the instance.
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
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// After you enable the elastic TPS feature for a ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fee](https://help.aliyun.com/document_detail/427237.html).
	//
	// > The elastic TPS feature is supported by only specific instance editions. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	SupportAutoScaling *bool `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
	// Indicates whether the message trace feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is not in use. By default, the message trace feature is enabled for ApsaraMQ for RocketMQ instances, regardless of whether this parameter is configured.
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

func (s *GetInstanceResponseBodyDataProductInfo) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *GetInstanceResponseBodyDataProductInfo) GetCapacityType() *string {
	return s.CapacityType
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

func (s *GetInstanceResponseBodyDataProductInfo) SetAutoScaling(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetCapacityType(v string) *GetInstanceResponseBodyDataProductInfo {
	s.CapacityType = &v
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
	// The period of upgrade time.
	//
	// example:
	//
	// 02:00-06:00
	MaintainTime *string `json:"maintainTime,omitempty" xml:"maintainTime,omitempty"`
	// The version of software.
	//
	// example:
	//
	// 5.0-rmq-20230619-1
	SoftwareVersion *string `json:"softwareVersion,omitempty" xml:"softwareVersion,omitempty"`
	// The upgrade method.
	//
	// Valid values:
	//
	// - Auto: automatic upgrade
	//
	// - Manual: manual upgrade
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
	// The tag key of the resource.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value of the resource.
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
