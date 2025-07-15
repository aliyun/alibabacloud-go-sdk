// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteVpcIpBlock(v string) *ModifyInstanceAttributeRequest
	GetDeleteVpcIpBlock() *string
	SetEgressIpv6Enable(v string) *ModifyInstanceAttributeRequest
	GetEgressIpv6Enable() *string
	SetHttpsPolicy(v string) *ModifyInstanceAttributeRequest
	GetHttpsPolicy() *string
	SetIPV6Enabled(v string) *ModifyInstanceAttributeRequest
	GetIPV6Enabled() *string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeRequest
	GetInstanceName() *string
	SetIntranetSegments(v string) *ModifyInstanceAttributeRequest
	GetIntranetSegments() *string
	SetMaintainEndTime(v string) *ModifyInstanceAttributeRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyInstanceAttributeRequest
	GetMaintainStartTime() *string
	SetToConnectVpcIpBlock(v *ModifyInstanceAttributeRequestToConnectVpcIpBlock) *ModifyInstanceAttributeRequest
	GetToConnectVpcIpBlock() *ModifyInstanceAttributeRequestToConnectVpcIpBlock
	SetToken(v string) *ModifyInstanceAttributeRequest
	GetToken() *string
	SetVpcSlbIntranetEnable(v string) *ModifyInstanceAttributeRequest
	GetVpcSlbIntranetEnable() *string
}

type ModifyInstanceAttributeRequest struct {
	// If delete VPC Ip block.
	//
	// example:
	//
	// true
	DeleteVpcIpBlock *string `json:"DeleteVpcIpBlock,omitempty" xml:"DeleteVpcIpBlock,omitempty"`
	// If enable outbound IPv6 Traffic.
	//
	// example:
	//
	// true
	EgressIpv6Enable *string `json:"EgressIpv6Enable,omitempty" xml:"EgressIpv6Enable,omitempty"`
	// The HTTPS policy.
	//
	// example:
	//
	// HTTPS2_TLS1_0
	HttpsPolicy *string `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	// If enable inbound IPv6 Traffic.
	//
	// example:
	//
	// true
	IPV6Enabled *string `json:"IPV6Enabled,omitempty" xml:"IPV6Enabled,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-ht-8xxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance Name
	//
	// example:
	//
	// apigatewayInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Custom private CIDR block.
	//
	// example:
	//
	// 172.0.0.1/24
	IntranetSegments *string `json:"IntranetSegments,omitempty" xml:"IntranetSegments,omitempty"`
	// Maintainable end time.
	//
	// example:
	//
	// 23:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// Maintainable start time.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The information about the CIDR block that API Gateway can use to access the virtual private cloud (VPC) of the backend service.
	ToConnectVpcIpBlock *ModifyInstanceAttributeRequestToConnectVpcIpBlock `json:"ToConnectVpcIpBlock,omitempty" xml:"ToConnectVpcIpBlock,omitempty" type:"Struct"`
	// The token of the request.
	//
	// example:
	//
	// c20d86c4-1eb3-4d0b-afe9-c586df1e2136
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// Specifies whether to enable the self-calling domain name.
	//
	// example:
	//
	// false
	VpcSlbIntranetEnable *string `json:"VpcSlbIntranetEnable,omitempty" xml:"VpcSlbIntranetEnable,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetDeleteVpcIpBlock() *string {
	return s.DeleteVpcIpBlock
}

func (s *ModifyInstanceAttributeRequest) GetEgressIpv6Enable() *string {
	return s.EgressIpv6Enable
}

func (s *ModifyInstanceAttributeRequest) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *ModifyInstanceAttributeRequest) GetIPV6Enabled() *string {
	return s.IPV6Enabled
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeRequest) GetIntranetSegments() *string {
	return s.IntranetSegments
}

func (s *ModifyInstanceAttributeRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyInstanceAttributeRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyInstanceAttributeRequest) GetToConnectVpcIpBlock() *ModifyInstanceAttributeRequestToConnectVpcIpBlock {
	return s.ToConnectVpcIpBlock
}

func (s *ModifyInstanceAttributeRequest) GetToken() *string {
	return s.Token
}

func (s *ModifyInstanceAttributeRequest) GetVpcSlbIntranetEnable() *string {
	return s.VpcSlbIntranetEnable
}

func (s *ModifyInstanceAttributeRequest) SetDeleteVpcIpBlock(v string) *ModifyInstanceAttributeRequest {
	s.DeleteVpcIpBlock = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetEgressIpv6Enable(v string) *ModifyInstanceAttributeRequest {
	s.EgressIpv6Enable = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetHttpsPolicy(v string) *ModifyInstanceAttributeRequest {
	s.HttpsPolicy = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetIPV6Enabled(v string) *ModifyInstanceAttributeRequest {
	s.IPV6Enabled = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetIntranetSegments(v string) *ModifyInstanceAttributeRequest {
	s.IntranetSegments = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetMaintainEndTime(v string) *ModifyInstanceAttributeRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetMaintainStartTime(v string) *ModifyInstanceAttributeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetToConnectVpcIpBlock(v *ModifyInstanceAttributeRequestToConnectVpcIpBlock) *ModifyInstanceAttributeRequest {
	s.ToConnectVpcIpBlock = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetToken(v string) *ModifyInstanceAttributeRequest {
	s.Token = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetVpcSlbIntranetEnable(v string) *ModifyInstanceAttributeRequest {
	s.VpcSlbIntranetEnable = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceAttributeRequestToConnectVpcIpBlock struct {
	// The CIDR block of the VSwitch.
	//
	// example:
	//
	// 172.16.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// Specifies whether the CIDR block is a custom CIDR block.
	//
	// example:
	//
	// false
	Customized *bool `json:"Customized,omitempty" xml:"Customized,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz94cqvaoe1ipxxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyInstanceAttributeRequestToConnectVpcIpBlock) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestToConnectVpcIpBlock) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) GetCustomized() *bool {
	return s.Customized
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) SetCidrBlock(v string) *ModifyInstanceAttributeRequestToConnectVpcIpBlock {
	s.CidrBlock = &v
	return s
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) SetCustomized(v bool) *ModifyInstanceAttributeRequestToConnectVpcIpBlock {
	s.Customized = &v
	return s
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) SetVswitchId(v string) *ModifyInstanceAttributeRequestToConnectVpcIpBlock {
	s.VswitchId = &v
	return s
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) SetZoneId(v string) *ModifyInstanceAttributeRequestToConnectVpcIpBlock {
	s.ZoneId = &v
	return s
}

func (s *ModifyInstanceAttributeRequestToConnectVpcIpBlock) Validate() error {
	return dara.Validate(s)
}
