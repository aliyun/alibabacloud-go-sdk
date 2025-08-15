// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclInfo(v *UpdateInstanceRequestAclInfo) *UpdateInstanceRequest
	GetAclInfo() *UpdateInstanceRequestAclInfo
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetNetworkInfo(v *UpdateInstanceRequestNetworkInfo) *UpdateInstanceRequest
	GetNetworkInfo() *UpdateInstanceRequestNetworkInfo
	SetProductInfo(v *UpdateInstanceRequestProductInfo) *UpdateInstanceRequest
	GetProductInfo() *UpdateInstanceRequestProductInfo
	SetRemark(v string) *UpdateInstanceRequest
	GetRemark() *string
}

type UpdateInstanceRequest struct {
	// The access control list for the instance.
	AclInfo *UpdateInstanceRequestAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// The updated name of the instance.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The updated network information about the instance.
	NetworkInfo *UpdateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// Additional configurations of the instance.
	ProductInfo *UpdateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The updated description of the instance.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAclInfo() *UpdateInstanceRequestAclInfo {
	return s.AclInfo
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetNetworkInfo() *UpdateInstanceRequestNetworkInfo {
	return s.NetworkInfo
}

func (s *UpdateInstanceRequest) GetProductInfo() *UpdateInstanceRequestProductInfo {
	return s.ProductInfo
}

func (s *UpdateInstanceRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateInstanceRequest) SetAclInfo(v *UpdateInstanceRequestAclInfo) *UpdateInstanceRequest {
	s.AclInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkInfo(v *UpdateInstanceRequestNetworkInfo) *UpdateInstanceRequest {
	s.NetworkInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetProductInfo(v *UpdateInstanceRequestProductInfo) *UpdateInstanceRequest {
	s.ProductInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetRemark(v string) *UpdateInstanceRequest {
	s.Remark = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestAclInfo struct {
	// The authentication type of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	DefaultVpcAuthFree *bool `json:"defaultVpcAuthFree,omitempty" xml:"defaultVpcAuthFree,omitempty"`
}

func (s UpdateInstanceRequestAclInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestAclInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestAclInfo) GetAclTypes() []*string {
	return s.AclTypes
}

func (s *UpdateInstanceRequestAclInfo) GetDefaultVpcAuthFree() *bool {
	return s.DefaultVpcAuthFree
}

func (s *UpdateInstanceRequestAclInfo) SetAclTypes(v []*string) *UpdateInstanceRequestAclInfo {
	s.AclTypes = v
	return s
}

func (s *UpdateInstanceRequestAclInfo) SetDefaultVpcAuthFree(v bool) *UpdateInstanceRequestAclInfo {
	s.DefaultVpcAuthFree = &v
	return s
}

func (s *UpdateInstanceRequestAclInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestNetworkInfo struct {
	// The information about the Internet over which the instance is accessed. This parameter takes effect only if the Internet access feature is enabled for the instance.
	InternetInfo *UpdateInstanceRequestNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
}

func (s UpdateInstanceRequestNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfo) GetInternetInfo() *UpdateInstanceRequestNetworkInfoInternetInfo {
	return s.InternetInfo
}

func (s *UpdateInstanceRequestNetworkInfo) SetInternetInfo(v *UpdateInstanceRequestNetworkInfoInternetInfo) *UpdateInstanceRequestNetworkInfo {
	s.InternetInfo = v
	return s
}

func (s *UpdateInstanceRequestNetworkInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestNetworkInfoInternetInfo struct {
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequestNetworkInfoInternetInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfoInternetInfo) GetIpWhitelist() []*string {
	return s.IpWhitelist
}

func (s *UpdateInstanceRequestNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *UpdateInstanceRequestNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

func (s *UpdateInstanceRequestNetworkInfoInternetInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestProductInfo struct {
	// Specifies whether to enable the elastic transactions per second (TPS) feature for the instance.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// After you enable the elastic TPS feature for an ApsaraMQ for RocketMQ instance, you can use a specific number of TPS that exceeds the specification limit. You are charged for using the elastic TPS feature. For more information, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	// >  The elastic TPS feature is supported only by specific instance editions. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
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
	// The ratio of the number of messages that you can send to the number of messages that you can receive on the instance.
	//
	// Value values: 0.25 to 1.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether to enable the message trace feature.
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

func (s UpdateInstanceRequestProductInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestProductInfo) GetAutoScaling() *bool {
	return s.AutoScaling
}

func (s *UpdateInstanceRequestProductInfo) GetMessageRetentionTime() *int32 {
	return s.MessageRetentionTime
}

func (s *UpdateInstanceRequestProductInfo) GetSendReceiveRatio() *float32 {
	return s.SendReceiveRatio
}

func (s *UpdateInstanceRequestProductInfo) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *UpdateInstanceRequestProductInfo) SetAutoScaling(v bool) *UpdateInstanceRequestProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetMessageRetentionTime(v int32) *UpdateInstanceRequestProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetSendReceiveRatio(v float32) *UpdateInstanceRequestProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetTraceOn(v bool) *UpdateInstanceRequestProductInfo {
	s.TraceOn = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) Validate() error {
	return dara.Validate(s)
}
