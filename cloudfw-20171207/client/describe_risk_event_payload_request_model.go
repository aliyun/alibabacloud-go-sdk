// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventPayloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstIP(v string) *DescribeRiskEventPayloadRequest
	GetDstIP() *string
	SetDstVpcId(v string) *DescribeRiskEventPayloadRequest
	GetDstVpcId() *string
	SetEndTime(v string) *DescribeRiskEventPayloadRequest
	GetEndTime() *string
	SetFirewallType(v string) *DescribeRiskEventPayloadRequest
	GetFirewallType() *string
	SetPublicIP(v string) *DescribeRiskEventPayloadRequest
	GetPublicIP() *string
	SetSrcIP(v string) *DescribeRiskEventPayloadRequest
	GetSrcIP() *string
	SetSrcVpcId(v string) *DescribeRiskEventPayloadRequest
	GetSrcVpcId() *string
	SetStartTime(v string) *DescribeRiskEventPayloadRequest
	GetStartTime() *string
	SetUUID(v string) *DescribeRiskEventPayloadRequest
	GetUUID() *string
}

type DescribeRiskEventPayloadRequest struct {
	// The destination IP address to query. If you specify this parameter, all intrusion events with the specified destination IP address are queried.
	//
	// example:
	//
	// 203.0.113.2
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The ID of the destination VPC to query. If you specify this parameter, all intrusion events that contain the specified ID of the destination VPC are queried.
	//
	// example:
	//
	// vpc-uf6jr1klwqb60dyn2****
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// The end of the time range to query. The value is a timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1681288980
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the firewall. Valid values:
	//
	// 	- **VpcFirewall**: virtual private cloud (VPC) firewall
	//
	// 	- **InternetFirewall*	- (default): Internet firewall
	//
	// example:
	//
	// InternetFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The public IP address. If you specify this parameter, all intrusion events that contain the specified public IP address are queried.
	//
	// example:
	//
	// 203.0.113.3
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The source IP address to query. If you specify this parameter, all intrusion events from the specified source IP address are queried.
	//
	// example:
	//
	// 203.0.113.1
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The ID of the source VPC to query. If you specify this parameter, all intrusion events that contain the specified ID of the source VPC are queried.
	//
	// example:
	//
	// vpc-wz9j2lqyo15udw5nt****
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The beginning of the time range to query. The value is a timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669533617
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The UUID of the intrusion event.
	//
	// This parameter is required.
	//
	// example:
	//
	// e62c25e0-1073-46bd-9567-b8f12b3d****
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s DescribeRiskEventPayloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventPayloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeRiskEventPayloadRequest) GetDstVpcId() *string {
	return s.DstVpcId
}

func (s *DescribeRiskEventPayloadRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRiskEventPayloadRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeRiskEventPayloadRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeRiskEventPayloadRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeRiskEventPayloadRequest) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeRiskEventPayloadRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRiskEventPayloadRequest) GetUUID() *string {
	return s.UUID
}

func (s *DescribeRiskEventPayloadRequest) SetDstIP(v string) *DescribeRiskEventPayloadRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetDstVpcId(v string) *DescribeRiskEventPayloadRequest {
	s.DstVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetEndTime(v string) *DescribeRiskEventPayloadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetFirewallType(v string) *DescribeRiskEventPayloadRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetPublicIP(v string) *DescribeRiskEventPayloadRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetSrcIP(v string) *DescribeRiskEventPayloadRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetSrcVpcId(v string) *DescribeRiskEventPayloadRequest {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetStartTime(v string) *DescribeRiskEventPayloadRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetUUID(v string) *DescribeRiskEventPayloadRequest {
	s.UUID = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) Validate() error {
	return dara.Validate(s)
}
