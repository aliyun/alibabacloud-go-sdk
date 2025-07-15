// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteVpcIpBlock(v string) *ModifyInstanceAttributeShrinkRequest
	GetDeleteVpcIpBlock() *string
	SetEgressIpv6Enable(v string) *ModifyInstanceAttributeShrinkRequest
	GetEgressIpv6Enable() *string
	SetHttpsPolicy(v string) *ModifyInstanceAttributeShrinkRequest
	GetHttpsPolicy() *string
	SetIPV6Enabled(v string) *ModifyInstanceAttributeShrinkRequest
	GetIPV6Enabled() *string
	SetInstanceId(v string) *ModifyInstanceAttributeShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeShrinkRequest
	GetInstanceName() *string
	SetIntranetSegments(v string) *ModifyInstanceAttributeShrinkRequest
	GetIntranetSegments() *string
	SetMaintainEndTime(v string) *ModifyInstanceAttributeShrinkRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyInstanceAttributeShrinkRequest
	GetMaintainStartTime() *string
	SetToConnectVpcIpBlockShrink(v string) *ModifyInstanceAttributeShrinkRequest
	GetToConnectVpcIpBlockShrink() *string
	SetToken(v string) *ModifyInstanceAttributeShrinkRequest
	GetToken() *string
	SetVpcSlbIntranetEnable(v string) *ModifyInstanceAttributeShrinkRequest
	GetVpcSlbIntranetEnable() *string
}

type ModifyInstanceAttributeShrinkRequest struct {
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
	ToConnectVpcIpBlockShrink *string `json:"ToConnectVpcIpBlock,omitempty" xml:"ToConnectVpcIpBlock,omitempty"`
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

func (s ModifyInstanceAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeShrinkRequest) GetDeleteVpcIpBlock() *string {
	return s.DeleteVpcIpBlock
}

func (s *ModifyInstanceAttributeShrinkRequest) GetEgressIpv6Enable() *string {
	return s.EgressIpv6Enable
}

func (s *ModifyInstanceAttributeShrinkRequest) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *ModifyInstanceAttributeShrinkRequest) GetIPV6Enabled() *string {
	return s.IPV6Enabled
}

func (s *ModifyInstanceAttributeShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeShrinkRequest) GetIntranetSegments() *string {
	return s.IntranetSegments
}

func (s *ModifyInstanceAttributeShrinkRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyInstanceAttributeShrinkRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyInstanceAttributeShrinkRequest) GetToConnectVpcIpBlockShrink() *string {
	return s.ToConnectVpcIpBlockShrink
}

func (s *ModifyInstanceAttributeShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *ModifyInstanceAttributeShrinkRequest) GetVpcSlbIntranetEnable() *string {
	return s.VpcSlbIntranetEnable
}

func (s *ModifyInstanceAttributeShrinkRequest) SetDeleteVpcIpBlock(v string) *ModifyInstanceAttributeShrinkRequest {
	s.DeleteVpcIpBlock = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetEgressIpv6Enable(v string) *ModifyInstanceAttributeShrinkRequest {
	s.EgressIpv6Enable = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetHttpsPolicy(v string) *ModifyInstanceAttributeShrinkRequest {
	s.HttpsPolicy = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetIPV6Enabled(v string) *ModifyInstanceAttributeShrinkRequest {
	s.IPV6Enabled = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetInstanceId(v string) *ModifyInstanceAttributeShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetInstanceName(v string) *ModifyInstanceAttributeShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetIntranetSegments(v string) *ModifyInstanceAttributeShrinkRequest {
	s.IntranetSegments = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetMaintainEndTime(v string) *ModifyInstanceAttributeShrinkRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetMaintainStartTime(v string) *ModifyInstanceAttributeShrinkRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetToConnectVpcIpBlockShrink(v string) *ModifyInstanceAttributeShrinkRequest {
	s.ToConnectVpcIpBlockShrink = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetToken(v string) *ModifyInstanceAttributeShrinkRequest {
	s.Token = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) SetVpcSlbIntranetEnable(v string) *ModifyInstanceAttributeShrinkRequest {
	s.VpcSlbIntranetEnable = &v
	return s
}

func (s *ModifyInstanceAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
