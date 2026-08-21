// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeDnsClusters(v []*AddResolverRuleRequestEdgeDnsClusters) *AddResolverRuleRequest
	GetEdgeDnsClusters() []*AddResolverRuleRequestEdgeDnsClusters
	SetEndpointId(v string) *AddResolverRuleRequest
	GetEndpointId() *string
	SetForwardIp(v []*AddResolverRuleRequestForwardIp) *AddResolverRuleRequest
	GetForwardIp() []*AddResolverRuleRequestForwardIp
	SetLang(v string) *AddResolverRuleRequest
	GetLang() *string
	SetName(v string) *AddResolverRuleRequest
	GetName() *string
	SetType(v string) *AddResolverRuleRequest
	GetType() *string
	SetVpcs(v []*AddResolverRuleRequestVpcs) *AddResolverRuleRequest
	GetVpcs() []*AddResolverRuleRequestVpcs
	SetZoneName(v string) *AddResolverRuleRequest
	GetZoneName() *string
}

type AddResolverRuleRequest struct {
	EdgeDnsClusters []*AddResolverRuleRequestEdgeDnsClusters `json:"EdgeDnsClusters,omitempty" xml:"EdgeDnsClusters,omitempty" type:"Repeated"`
	// The ID of the outbound endpoint. The outbound endpoint forwards DNS queries to the specified destination IP addresses.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The IP addresses and ports of the destination servers in the external DNS system to which DNS queries are forwarded. You can specify up to **6*	- destination servers. Both private and public IP addresses are supported.
	//
	// > If you specify public IP addresses for the external DNS servers, and the Elastic Compute Service (ECS) instances in the VPC of the outbound endpoint do not have public IP addresses, enable a NAT Gateway and configure SNAT entries.
	//
	// This parameter is required.
	ForwardIp []*AddResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the forwarding rule. Name the rule as needed.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the forwarding rule. The only valid value is OUTBOUND. This value indicates that DNS queries are forwarded to an external IP address.
	//
	// > You cannot change this value after the forwarding rule is created.
	//
	// example:
	//
	// OUTBOUND
	Type *string                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Vpcs []*AddResolverRuleRequestVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
	// The domain name (zone) for which you want to forward DNS queries.
	//
	// > You cannot change this value after the forwarding rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s AddResolverRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequest) GetEdgeDnsClusters() []*AddResolverRuleRequestEdgeDnsClusters {
	return s.EdgeDnsClusters
}

func (s *AddResolverRuleRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *AddResolverRuleRequest) GetForwardIp() []*AddResolverRuleRequestForwardIp {
	return s.ForwardIp
}

func (s *AddResolverRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *AddResolverRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddResolverRuleRequest) GetType() *string {
	return s.Type
}

func (s *AddResolverRuleRequest) GetVpcs() []*AddResolverRuleRequestVpcs {
	return s.Vpcs
}

func (s *AddResolverRuleRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *AddResolverRuleRequest) SetEdgeDnsClusters(v []*AddResolverRuleRequestEdgeDnsClusters) *AddResolverRuleRequest {
	s.EdgeDnsClusters = v
	return s
}

func (s *AddResolverRuleRequest) SetEndpointId(v string) *AddResolverRuleRequest {
	s.EndpointId = &v
	return s
}

func (s *AddResolverRuleRequest) SetForwardIp(v []*AddResolverRuleRequestForwardIp) *AddResolverRuleRequest {
	s.ForwardIp = v
	return s
}

func (s *AddResolverRuleRequest) SetLang(v string) *AddResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *AddResolverRuleRequest) SetName(v string) *AddResolverRuleRequest {
	s.Name = &v
	return s
}

func (s *AddResolverRuleRequest) SetType(v string) *AddResolverRuleRequest {
	s.Type = &v
	return s
}

func (s *AddResolverRuleRequest) SetVpcs(v []*AddResolverRuleRequestVpcs) *AddResolverRuleRequest {
	s.Vpcs = v
	return s
}

func (s *AddResolverRuleRequest) SetZoneName(v string) *AddResolverRuleRequest {
	s.ZoneName = &v
	return s
}

func (s *AddResolverRuleRequest) Validate() error {
	if s.EdgeDnsClusters != nil {
		for _, item := range s.EdgeDnsClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ForwardIp != nil {
		for _, item := range s.ForwardIp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Vpcs != nil {
		for _, item := range s.Vpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddResolverRuleRequestEdgeDnsClusters struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s AddResolverRuleRequestEdgeDnsClusters) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleRequestEdgeDnsClusters) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequestEdgeDnsClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddResolverRuleRequestEdgeDnsClusters) SetClusterId(v string) *AddResolverRuleRequestEdgeDnsClusters {
	s.ClusterId = &v
	return s
}

func (s *AddResolverRuleRequestEdgeDnsClusters) Validate() error {
	return dara.Validate(s)
}

type AddResolverRuleRequestForwardIp struct {
	// The IP address of the destination server.
	//
	// > The IP addresses in the following ranges are reserved by the system and cannot be used as the IP addresses of external DNS systems: 100.100.2.136 to 100.100.2.138 and 100.100.2.116 to 100.100.2.118.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number of the destination server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s AddResolverRuleRequestForwardIp) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleRequestForwardIp) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequestForwardIp) GetIp() *string {
	return s.Ip
}

func (s *AddResolverRuleRequestForwardIp) GetPort() *int32 {
	return s.Port
}

func (s *AddResolverRuleRequestForwardIp) SetIp(v string) *AddResolverRuleRequestForwardIp {
	s.Ip = &v
	return s
}

func (s *AddResolverRuleRequestForwardIp) SetPort(v int32) *AddResolverRuleRequestForwardIp {
	s.Port = &v
	return s
}

func (s *AddResolverRuleRequestForwardIp) Validate() error {
	return dara.Validate(s)
}

type AddResolverRuleRequestVpcs struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType   *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	VpcUserId *int64  `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s AddResolverRuleRequestVpcs) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleRequestVpcs) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequestVpcs) GetRegionId() *string {
	return s.RegionId
}

func (s *AddResolverRuleRequestVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *AddResolverRuleRequestVpcs) GetVpcType() *string {
	return s.VpcType
}

func (s *AddResolverRuleRequestVpcs) GetVpcUserId() *int64 {
	return s.VpcUserId
}

func (s *AddResolverRuleRequestVpcs) SetRegionId(v string) *AddResolverRuleRequestVpcs {
	s.RegionId = &v
	return s
}

func (s *AddResolverRuleRequestVpcs) SetVpcId(v string) *AddResolverRuleRequestVpcs {
	s.VpcId = &v
	return s
}

func (s *AddResolverRuleRequestVpcs) SetVpcType(v string) *AddResolverRuleRequestVpcs {
	s.VpcType = &v
	return s
}

func (s *AddResolverRuleRequestVpcs) SetVpcUserId(v int64) *AddResolverRuleRequestVpcs {
	s.VpcUserId = &v
	return s
}

func (s *AddResolverRuleRequestVpcs) Validate() error {
	return dara.Validate(s)
}
