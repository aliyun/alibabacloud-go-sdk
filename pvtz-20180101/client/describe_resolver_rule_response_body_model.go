// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindEdgeDnsClusters(v []*DescribeResolverRuleResponseBodyBindEdgeDnsClusters) *DescribeResolverRuleResponseBody
	GetBindEdgeDnsClusters() []*DescribeResolverRuleResponseBodyBindEdgeDnsClusters
	SetBindVpcs(v []*DescribeResolverRuleResponseBodyBindVpcs) *DescribeResolverRuleResponseBody
	GetBindVpcs() []*DescribeResolverRuleResponseBodyBindVpcs
	SetCreateTime(v string) *DescribeResolverRuleResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeResolverRuleResponseBody
	GetCreateTimestamp() *int64
	SetEndpointId(v string) *DescribeResolverRuleResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *DescribeResolverRuleResponseBody
	GetEndpointName() *string
	SetForwardIps(v []*DescribeResolverRuleResponseBodyForwardIps) *DescribeResolverRuleResponseBody
	GetForwardIps() []*DescribeResolverRuleResponseBodyForwardIps
	SetId(v string) *DescribeResolverRuleResponseBody
	GetId() *string
	SetName(v string) *DescribeResolverRuleResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeResolverRuleResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeResolverRuleResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeResolverRuleResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeResolverRuleResponseBody
	GetUpdateTimestamp() *int64
	SetZoneName(v string) *DescribeResolverRuleResponseBody
	GetZoneName() *string
}

type DescribeResolverRuleResponseBody struct {
	BindEdgeDnsClusters []*DescribeResolverRuleResponseBodyBindEdgeDnsClusters `json:"BindEdgeDnsClusters,omitempty" xml:"BindEdgeDnsClusters,omitempty" type:"Repeated"`
	// The virtual private clouds (VPCs) that are associated with the forwarding rule.
	BindVpcs []*DescribeResolverRuleResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	// The time when the forwarding rule was created.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the forwarding rule was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The destination IP addresses.
	ForwardIps []*DescribeResolverRuleResponseBodyForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// hr****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// forward rule-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13D5113B-7E34-407F-A9C1-D96CD2B04277
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the forwarding rule. Valid value:
	//
	// OUTBOUND: outbound forwarding rule. This type of rule forwards Domain Name System (DNS) requests to one or more external IP addresses.
	//
	// example:
	//
	// OUTBOUND
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the forwarding rule was updated.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the forwarding rule was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The name of the forward zone.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResolverRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBody) GetBindEdgeDnsClusters() []*DescribeResolverRuleResponseBodyBindEdgeDnsClusters {
	return s.BindEdgeDnsClusters
}

func (s *DescribeResolverRuleResponseBody) GetBindVpcs() []*DescribeResolverRuleResponseBodyBindVpcs {
	return s.BindVpcs
}

func (s *DescribeResolverRuleResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResolverRuleResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeResolverRuleResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeResolverRuleResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *DescribeResolverRuleResponseBody) GetForwardIps() []*DescribeResolverRuleResponseBodyForwardIps {
	return s.ForwardIps
}

func (s *DescribeResolverRuleResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeResolverRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeResolverRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResolverRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeResolverRuleResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeResolverRuleResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeResolverRuleResponseBody) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeResolverRuleResponseBody) SetBindEdgeDnsClusters(v []*DescribeResolverRuleResponseBodyBindEdgeDnsClusters) *DescribeResolverRuleResponseBody {
	s.BindEdgeDnsClusters = v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetBindVpcs(v []*DescribeResolverRuleResponseBodyBindVpcs) *DescribeResolverRuleResponseBody {
	s.BindVpcs = v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetCreateTime(v string) *DescribeResolverRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetCreateTimestamp(v int64) *DescribeResolverRuleResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetEndpointId(v string) *DescribeResolverRuleResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetEndpointName(v string) *DescribeResolverRuleResponseBody {
	s.EndpointName = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetForwardIps(v []*DescribeResolverRuleResponseBodyForwardIps) *DescribeResolverRuleResponseBody {
	s.ForwardIps = v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetId(v string) *DescribeResolverRuleResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetName(v string) *DescribeResolverRuleResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetRequestId(v string) *DescribeResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetType(v string) *DescribeResolverRuleResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetUpdateTime(v string) *DescribeResolverRuleResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetUpdateTimestamp(v int64) *DescribeResolverRuleResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetZoneName(v string) *DescribeResolverRuleResponseBody {
	s.ZoneName = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) Validate() error {
	if s.BindEdgeDnsClusters != nil {
		for _, item := range s.BindEdgeDnsClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BindVpcs != nil {
		for _, item := range s.BindVpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ForwardIps != nil {
		for _, item := range s.ForwardIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResolverRuleResponseBodyBindEdgeDnsClusters struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName   *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterUserId *int64  `json:"ClusterUserId,omitempty" xml:"ClusterUserId,omitempty"`
}

func (s DescribeResolverRuleResponseBodyBindEdgeDnsClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleResponseBodyBindEdgeDnsClusters) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) GetClusterUserId() *int64 {
	return s.ClusterUserId
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) SetClusterId(v string) *DescribeResolverRuleResponseBodyBindEdgeDnsClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) SetClusterName(v string) *DescribeResolverRuleResponseBodyBindEdgeDnsClusters {
	s.ClusterName = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) SetClusterUserId(v int64) *DescribeResolverRuleResponseBodyBindEdgeDnsClusters {
	s.ClusterUserId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindEdgeDnsClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverRuleResponseBodyBindVpcs struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// vpc-name-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The ID of the user to which the VPC belongs.
	//
	// example:
	//
	// 32454****
	VpcUserId *string `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeResolverRuleResponseBodyBindVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleResponseBodyBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) GetVpcUserId() *string {
	return s.VpcUserId
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetRegionId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.RegionId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetRegionName(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.RegionName = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcName(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcType(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcType = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcUserId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcUserId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverRuleResponseBodyForwardIps struct {
	// The destination IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeResolverRuleResponseBodyForwardIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleResponseBodyForwardIps) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBodyForwardIps) GetIp() *string {
	return s.Ip
}

func (s *DescribeResolverRuleResponseBodyForwardIps) GetPort() *int32 {
	return s.Port
}

func (s *DescribeResolverRuleResponseBodyForwardIps) SetIp(v string) *DescribeResolverRuleResponseBodyForwardIps {
	s.Ip = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyForwardIps) SetPort(v int32) *DescribeResolverRuleResponseBodyForwardIps {
	s.Port = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyForwardIps) Validate() error {
	return dara.Validate(s)
}
