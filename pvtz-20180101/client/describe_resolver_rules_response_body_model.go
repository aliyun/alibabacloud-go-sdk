// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeResolverRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResolverRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeResolverRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeResolverRulesResponseBodyRules) *DescribeResolverRulesResponseBody
	GetRules() []*DescribeResolverRulesResponseBodyRules
	SetTotalItems(v int32) *DescribeResolverRulesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeResolverRulesResponseBody
	GetTotalPages() *int32
}

type DescribeResolverRulesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A10E03D7-808C-422D-9144-F8586C2E2297
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rules.
	Rules []*DescribeResolverRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeResolverRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResolverRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResolverRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResolverRulesResponseBody) GetRules() []*DescribeResolverRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeResolverRulesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeResolverRulesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeResolverRulesResponseBody) SetPageNumber(v int32) *DescribeResolverRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetPageSize(v int32) *DescribeResolverRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetRequestId(v string) *DescribeResolverRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetRules(v []*DescribeResolverRulesResponseBodyRules) *DescribeResolverRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetTotalItems(v int32) *DescribeResolverRulesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetTotalPages(v int32) *DescribeResolverRulesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResolverRulesResponseBodyRules struct {
	BindEdgeDnsClusters []*DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters `json:"BindEdgeDnsClusters,omitempty" xml:"BindEdgeDnsClusters,omitempty" type:"Repeated"`
	// The VPCs associated with the forwarding rule.
	BindVpcs []*DescribeResolverRulesResponseBodyRulesBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	// The time when the forwarding was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
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
	// The IP addresses and ports of the external DNS servers. Enter the IP addresses and ports of the destination servers to which the DNS requests are forwarded.
	ForwardIps []*DescribeResolverRulesResponseBodyRulesForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
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
	Name                   *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PriorityForwardConfigs []*DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs `json:"PriorityForwardConfigs,omitempty" xml:"PriorityForwardConfigs,omitempty" type:"Repeated"`
	// The type of the forwarding rule.
	//
	// The parameter value can only be OUTBOUND, which indicates that Domain Name System (DNS) requests are forwarded to one or more external IP addresses.
	//
	// example:
	//
	// OUTBOUND
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the forwarding rule was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
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
	// The zone for which you want to forward DNS requests.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRules) GetBindEdgeDnsClusters() []*DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters {
	return s.BindEdgeDnsClusters
}

func (s *DescribeResolverRulesResponseBodyRules) GetBindVpcs() []*DescribeResolverRulesResponseBodyRulesBindVpcs {
	return s.BindVpcs
}

func (s *DescribeResolverRulesResponseBodyRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResolverRulesResponseBodyRules) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeResolverRulesResponseBodyRules) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeResolverRulesResponseBodyRules) GetEndpointName() *string {
	return s.EndpointName
}

func (s *DescribeResolverRulesResponseBodyRules) GetForwardIps() []*DescribeResolverRulesResponseBodyRulesForwardIps {
	return s.ForwardIps
}

func (s *DescribeResolverRulesResponseBodyRules) GetId() *string {
	return s.Id
}

func (s *DescribeResolverRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *DescribeResolverRulesResponseBodyRules) GetPriorityForwardConfigs() []*DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	return s.PriorityForwardConfigs
}

func (s *DescribeResolverRulesResponseBodyRules) GetType() *string {
	return s.Type
}

func (s *DescribeResolverRulesResponseBodyRules) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeResolverRulesResponseBodyRules) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeResolverRulesResponseBodyRules) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeResolverRulesResponseBodyRules) SetBindEdgeDnsClusters(v []*DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) *DescribeResolverRulesResponseBodyRules {
	s.BindEdgeDnsClusters = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetBindVpcs(v []*DescribeResolverRulesResponseBodyRulesBindVpcs) *DescribeResolverRulesResponseBodyRules {
	s.BindVpcs = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetCreateTime(v string) *DescribeResolverRulesResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetCreateTimestamp(v int64) *DescribeResolverRulesResponseBodyRules {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetEndpointId(v string) *DescribeResolverRulesResponseBodyRules {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetEndpointName(v string) *DescribeResolverRulesResponseBodyRules {
	s.EndpointName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetForwardIps(v []*DescribeResolverRulesResponseBodyRulesForwardIps) *DescribeResolverRulesResponseBodyRules {
	s.ForwardIps = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetId(v string) *DescribeResolverRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetName(v string) *DescribeResolverRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetPriorityForwardConfigs(v []*DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) *DescribeResolverRulesResponseBodyRules {
	s.PriorityForwardConfigs = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetType(v string) *DescribeResolverRulesResponseBodyRules {
	s.Type = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetUpdateTime(v string) *DescribeResolverRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetUpdateTimestamp(v int64) *DescribeResolverRulesResponseBodyRules {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetZoneName(v string) *DescribeResolverRulesResponseBodyRules {
	s.ZoneName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) Validate() error {
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
	if s.PriorityForwardConfigs != nil {
		for _, item := range s.PriorityForwardConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName   *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterUserId *int64  `json:"ClusterUserId,omitempty" xml:"ClusterUserId,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) GetClusterUserId() *int64 {
	return s.ClusterUserId
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) SetClusterId(v string) *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) SetClusterName(v string) *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters {
	s.ClusterName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) SetClusterUserId(v int64) *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters {
	s.ClusterUserId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindEdgeDnsClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverRulesResponseBodyRulesBindVpcs struct {
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region to which the VPC belongs.
	//
	// example:
	//
	// ap-southeast-1
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID. This ID uniquely identifies the VPC.
	//
	// example:
	//
	// vpc-0jl96awrjt75ezglc****
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
	// The user ID to which the VPC belongs.
	//
	// example:
	//
	// 141339776561****
	VpcUserId *string `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesBindVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) GetVpcUserId() *string {
	return s.VpcUserId
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetRegionId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.RegionId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetRegionName(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.RegionName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcName(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcType(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcType = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcUserId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcUserId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverRulesResponseBodyRulesForwardIps struct {
	// The IP address of the destination server.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port of the destination server.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesForwardIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesForwardIps) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) GetIp() *string {
	return s.Ip
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) GetPort() *int32 {
	return s.Port
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) SetIp(v string) *DescribeResolverRulesResponseBodyRulesForwardIps {
	s.Ip = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) SetPort(v int32) *DescribeResolverRulesResponseBodyRulesForwardIps {
	s.Port = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs struct {
	AlidnsServiceAddresses []*string `json:"AlidnsServiceAddresses,omitempty" xml:"AlidnsServiceAddresses,omitempty" type:"Repeated"`
	CustomAddresses        []*string `json:"CustomAddresses,omitempty" xml:"CustomAddresses,omitempty" type:"Repeated"`
	EnableStatus           *string   `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	Priority               *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Protocol               *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GetAlidnsServiceAddresses() []*string {
	return s.AlidnsServiceAddresses
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GetCustomAddresses() []*string {
	return s.CustomAddresses
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) SetAlidnsServiceAddresses(v []*string) *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	s.AlidnsServiceAddresses = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) SetCustomAddresses(v []*string) *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	s.CustomAddresses = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) SetEnableStatus(v string) *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	s.EnableStatus = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) SetPriority(v int32) *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	s.Priority = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) SetProtocol(v string) *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs {
	s.Protocol = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesPriorityForwardConfigs) Validate() error {
	return dara.Validate(s)
}
