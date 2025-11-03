// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMirrorFiltersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListTrafficMirrorFiltersResponseBody
	GetCount() *int32
	SetMaxResults(v int32) *ListTrafficMirrorFiltersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrafficMirrorFiltersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTrafficMirrorFiltersResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrafficMirrorFiltersResponseBody
	GetTotalCount() *string
	SetTrafficMirrorFilters(v []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) *ListTrafficMirrorFiltersResponseBody
	GetTrafficMirrorFilters() []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters
}

type ListTrafficMirrorFiltersResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 739CA01C-92EB-4C69-BCC0-280149C6F41E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the filters.
	TrafficMirrorFilters []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters `json:"TrafficMirrorFilters,omitempty" xml:"TrafficMirrorFilters,omitempty" type:"Repeated"`
}

func (s ListTrafficMirrorFiltersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListTrafficMirrorFiltersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrafficMirrorFiltersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMirrorFiltersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficMirrorFiltersResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrafficMirrorFiltersResponseBody) GetTrafficMirrorFilters() []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	return s.TrafficMirrorFilters
}

func (s *ListTrafficMirrorFiltersResponseBody) SetCount(v int32) *ListTrafficMirrorFiltersResponseBody {
	s.Count = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) SetMaxResults(v int32) *ListTrafficMirrorFiltersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) SetNextToken(v string) *ListTrafficMirrorFiltersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) SetRequestId(v string) *ListTrafficMirrorFiltersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) SetTotalCount(v string) *ListTrafficMirrorFiltersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) SetTrafficMirrorFilters(v []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) *ListTrafficMirrorFiltersResponseBody {
	s.TrafficMirrorFilters = v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBody) Validate() error {
	if s.TrafficMirrorFilters != nil {
		for _, item := range s.TrafficMirrorFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters struct {
	// The time when the filter is created.
	//
	// example:
	//
	// 2023-09-05T15:26Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The information about the outbound rules.
	EgressRules []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules `json:"EgressRules,omitempty" xml:"EgressRules,omitempty" type:"Repeated"`
	// The information about the inbound rules.
	IngressRules []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules `json:"IngressRules,omitempty" xml:"IngressRules,omitempty" type:"Repeated"`
	// The ID of the resource group to which the traffic mirror session belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the filter.
	//
	// example:
	//
	// This is a filter.
	TrafficMirrorFilterDescription *string `json:"TrafficMirrorFilterDescription,omitempty" xml:"TrafficMirrorFilterDescription,omitempty"`
	// The ID of the filter.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The filter name.
	//
	// example:
	//
	// abc
	TrafficMirrorFilterName *string `json:"TrafficMirrorFilterName,omitempty" xml:"TrafficMirrorFilterName,omitempty"`
	// The status of the filter. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	TrafficMirrorFilterStatus *string `json:"TrafficMirrorFilterStatus,omitempty" xml:"TrafficMirrorFilterStatus,omitempty"`
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetEgressRules() []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	return s.EgressRules
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetIngressRules() []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	return s.IngressRules
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetTags() []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags {
	return s.Tags
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetTrafficMirrorFilterDescription() *string {
	return s.TrafficMirrorFilterDescription
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetTrafficMirrorFilterName() *string {
	return s.TrafficMirrorFilterName
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) GetTrafficMirrorFilterStatus() *string {
	return s.TrafficMirrorFilterStatus
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetCreationTime(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.CreationTime = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetEgressRules(v []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.EgressRules = v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetIngressRules(v []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.IngressRules = v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetResourceGroupId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetTags(v []*ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.Tags = v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetTrafficMirrorFilterDescription(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.TrafficMirrorFilterDescription = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetTrafficMirrorFilterId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetTrafficMirrorFilterName(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.TrafficMirrorFilterName = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) SetTrafficMirrorFilterStatus(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters {
	s.TrafficMirrorFilterStatus = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFilters) Validate() error {
	if s.EgressRules != nil {
		for _, item := range s.EgressRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressRules != nil {
		for _, item := range s.IngressRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules struct {
	// The action of the outbound rule. Valid values:
	//
	// 	- **accept**
	//
	// 	- **drop**
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the outbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the outbound traffic.
	//
	// example:
	//
	// 22/40
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The version of IP protocol.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the outbound rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol that is used by the outbound traffic to be mirrored. Valid values:
	//
	// 	- **ALL**
	//
	// 	- **ICMP**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the outbound traffic.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the outbound traffic.
	//
	// example:
	//
	// 22/40
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The direction of the network traffic. Valid values:
	//
	// 	- **egress**
	//
	// 	- **ingress**
	//
	// example:
	//
	// egress
	TrafficDirection *string `json:"TrafficDirection,omitempty" xml:"TrafficDirection,omitempty"`
	// The ID of the filter associated with the outbound rule.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The ID of the outbound rule.
	//
	// example:
	//
	// tmr-j6c89rzmtd3hhdugq****
	TrafficMirrorFilterRuleId *string `json:"TrafficMirrorFilterRuleId,omitempty" xml:"TrafficMirrorFilterRuleId,omitempty"`
	// The status of the outbound rule. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	TrafficMirrorFilterRuleStatus *string `json:"TrafficMirrorFilterRuleStatus,omitempty" xml:"TrafficMirrorFilterRuleStatus,omitempty"`
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetAction() *string {
	return s.Action
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetTrafficDirection() *string {
	return s.TrafficDirection
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetTrafficMirrorFilterRuleId() *string {
	return s.TrafficMirrorFilterRuleId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) GetTrafficMirrorFilterRuleStatus() *string {
	return s.TrafficMirrorFilterRuleStatus
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetAction(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.Action = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetDestinationCidrBlock(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetDestinationPortRange(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetIpVersion(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.IpVersion = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetPriority(v int32) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.Priority = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetProtocol(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.Protocol = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetSourceCidrBlock(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetSourcePortRange(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.SourcePortRange = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetTrafficDirection(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.TrafficDirection = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetTrafficMirrorFilterId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetTrafficMirrorFilterRuleId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.TrafficMirrorFilterRuleId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) SetTrafficMirrorFilterRuleStatus(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules {
	s.TrafficMirrorFilterRuleStatus = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersEgressRules) Validate() error {
	return dara.Validate(s)
}

type ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules struct {
	// The action of the inbound rule. Valid values:
	//
	// 	- **accept**
	//
	// 	- **drop**
	//
	// example:
	//
	// accept
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The destination CIDR block of the inbound traffic.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the inbound traffic.
	//
	// example:
	//
	// -1/-1
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// The version of IP protocol.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The priority of the inbound rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol that is used by the inbound traffic to be mirrored. Valid values:
	//
	// 	- **ALL**
	//
	// 	- **ICMP**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// example:
	//
	// ALL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block of the inbound traffic.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The destination port range of the inbound traffic.
	//
	// example:
	//
	// -1/-1
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The direction of the network traffic. Valid values:
	//
	// 	- **egress**
	//
	// 	- **ingress**
	//
	// example:
	//
	// ingress
	TrafficDirection *string `json:"TrafficDirection,omitempty" xml:"TrafficDirection,omitempty"`
	// The ID of the filter associated with the inbound rule.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The ID of the inbound rule.
	//
	// example:
	//
	// tmr-j6cezu8e68rnpepet****
	TrafficMirrorFilterRuleId *string `json:"TrafficMirrorFilterRuleId,omitempty" xml:"TrafficMirrorFilterRuleId,omitempty"`
	// The status of the inbound rule. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	TrafficMirrorFilterRuleStatus *string `json:"TrafficMirrorFilterRuleStatus,omitempty" xml:"TrafficMirrorFilterRuleStatus,omitempty"`
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetAction() *string {
	return s.Action
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetTrafficDirection() *string {
	return s.TrafficDirection
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetTrafficMirrorFilterRuleId() *string {
	return s.TrafficMirrorFilterRuleId
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) GetTrafficMirrorFilterRuleStatus() *string {
	return s.TrafficMirrorFilterRuleStatus
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetAction(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.Action = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetDestinationCidrBlock(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetDestinationPortRange(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.DestinationPortRange = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetIpVersion(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.IpVersion = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetPriority(v int32) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.Priority = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetProtocol(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.Protocol = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetSourceCidrBlock(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.SourceCidrBlock = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetSourcePortRange(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.SourcePortRange = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetTrafficDirection(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.TrafficDirection = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetTrafficMirrorFilterId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetTrafficMirrorFilterRuleId(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.TrafficMirrorFilterRuleId = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) SetTrafficMirrorFilterRuleStatus(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules {
	s.TrafficMirrorFilterRuleStatus = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersIngressRules) Validate() error {
	return dara.Validate(s)
}

type ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) GoString() string {
	return s.String()
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) GetKey() *string {
	return s.Key
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) GetValue() *string {
	return s.Value
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) SetKey(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags {
	s.Key = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) SetValue(v string) *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags {
	s.Value = &v
	return s
}

func (s *ListTrafficMirrorFiltersResponseBodyTrafficMirrorFiltersTags) Validate() error {
	return dara.Validate(s)
}
