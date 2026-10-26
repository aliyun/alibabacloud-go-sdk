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
	// The access control information.
	AclInfo *UpdateInstanceRequestAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// The updated instance name.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The updated network information of the instance.
	NetworkInfo *UpdateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The extension configurations of the instance.
	ProductInfo *UpdateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The updated remarks for the instance.
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
	return nil
}

type UpdateInstanceRequestAclInfo struct {
	// The authentication type of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether password-free access is enabled for VPC connections under the intelligent identity recognition type.
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
	// The public network information. This parameter is valid only when the public network access feature is enabled for the instance.
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
	if s.InternetInfo != nil {
		if err := s.InternetInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceRequestNetworkInfoInternetInfo struct {
	// The public network access IP whitelist.
	//
	// - If no IP whitelist is configured, all IP address ranges are allowed to access the ApsaraMQ for RocketMQ server over the public network by default.
	//
	// - If an IP whitelist is configured, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ server over the public network.
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
	// Specifies whether to enable the adaptive elasticity capability. A value of true indicates enabled, and a value of false indicates shutdown.
	AdaptiveElasticity *bool `json:"adaptiveElasticity,omitempty" xml:"adaptiveElasticity,omitempty"`
	// Specifies whether to enable the burstable elastic capacity beyond the specification.
	//
	// After the burstable elastic capacity is enabled, ApsaraMQ for RocketMQ allows the instance to exceed the TPS limit of the base specification within a certain range. The portion that exceeds the base specification incurs additional elastic specification fees. For more information about billing, see [Computing fee](https://help.aliyun.com/document_detail/427237.html).
	//
	//
	// > Only instances of specific specifications support the burstable elastic specification capability. For more information about supported instances, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The message retention period. Unit: hours.
	//
	// For the valid values, see [Limits](https://help.aliyun.com/document_detail/440347.html) > Resource Quotas > Message retention period limit.
	//
	// The log broker in ApsaraMQ for RocketMQ is fully serverless and elastic. Fees are charged based on the actual storage space used. You can control the storage capacity of messages by adjusting the message retention period. For more information, see [Storage fee](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The ratio of messages sent to messages accepted.
	//
	// Valid values: [0.25, 1]
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether to enable the message trace feature.
	//
	// <props="intl">This parameter does not take effect. Regardless of whether this parameter is enabled, the instance supports the message trace feature by default.
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

func (s UpdateInstanceRequestProductInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestProductInfo) GetAdaptiveElasticity() *bool {
	return s.AdaptiveElasticity
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

func (s *UpdateInstanceRequestProductInfo) SetAdaptiveElasticity(v bool) *UpdateInstanceRequestProductInfo {
	s.AdaptiveElasticity = &v
	return s
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
