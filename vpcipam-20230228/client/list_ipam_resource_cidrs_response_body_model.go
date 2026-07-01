// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceCidrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamResourceCidrsResponseBody
	GetCount() *int64
	SetIpamResourceCidrs(v []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs) *ListIpamResourceCidrsResponseBody
	GetIpamResourceCidrs() []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs
	SetMaxResults(v int64) *ListIpamResourceCidrsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamResourceCidrsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamResourceCidrsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamResourceCidrsResponseBody
	GetTotalCount() *int64
}

type ListIpamResourceCidrsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resource information.
	IpamResourceCidrs []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs `json:"IpamResourceCidrs,omitempty" xml:"IpamResourceCidrs,omitempty" type:"Repeated"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// - If **NextToken*	- is empty, no more results exist.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 49A9DE56-B68C-5FFC-BC06-509D086F287C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned under the current query conditions.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceCidrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamResourceCidrsResponseBody) GetIpamResourceCidrs() []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	return s.IpamResourceCidrs
}

func (s *ListIpamResourceCidrsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamResourceCidrsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceCidrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamResourceCidrsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamResourceCidrsResponseBody) SetCount(v int64) *ListIpamResourceCidrsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetIpamResourceCidrs(v []*ListIpamResourceCidrsResponseBodyIpamResourceCidrs) *ListIpamResourceCidrsResponseBody {
	s.IpamResourceCidrs = v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetMaxResults(v int64) *ListIpamResourceCidrsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetNextToken(v string) *ListIpamResourceCidrsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetRequestId(v string) *ListIpamResourceCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) SetTotalCount(v int64) *ListIpamResourceCidrsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBody) Validate() error {
	if s.IpamResourceCidrs != nil {
		for _, item := range s.IpamResourceCidrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrs struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 132193271328****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The CIDR block of the resource.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The compliance status of the resource.
	//
	// - **Compliant**: The CIDR block of the managed resource complies with the allocation rules of the IPAM pool.
	//
	// - **Noncompliant**: The CIDR block of the managed resource does not comply with one or more allocation rules of the IPAM pool.
	//
	// - **Ignored**: The resource has been excluded from monitoring. Ignored resources are not evaluated for overlap or allocation rule compliance.
	//
	// - **Unmanaged**: The resource does not have a CIDR block allocated from an IPAM pool. IPAM does not monitor whether the CIDR block of the resource complies with the allocation rules of the pool, but monitors the CIDR block for overlap.
	//
	// example:
	//
	// Compliant
	ComplianceStatus *string `json:"ComplianceStatus,omitempty" xml:"ComplianceStatus,omitempty"`
	// The details of the resource IP address count.
	IpCountDetail *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail `json:"IpCountDetail,omitempty" xml:"IpCountDetail,omitempty" type:"Struct"`
	// The IP utilization rate, expressed as a decimal.
	//
	// example:
	//
	// 0
	IpUsage *string `json:"IpUsage,omitempty" xml:"IpUsage,omitempty"`
	// The instance ID of the IPAM pool CIDR allocation.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamAllocationId *string `json:"IpamAllocationId,omitempty" xml:"IpamAllocationId,omitempty"`
	// The instance ID of the IPAM.
	//
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The instance ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t***
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The instance ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The management status of the resource.
	//
	// - **Managed**: The resource has a CIDR block allocated from an IPAM pool. IPAM monitors the resource for potential CIDR overlap and compliance with pool allocation rules.
	//
	// - **Unmanaged**: The resource does not have a CIDR block allocated from an IPAM pool. IPAM monitors the resource for potential CIDRs that comply with pool allocation rules and monitors CIDRs for overlap.
	//
	// - **Ignored**: The resource has been excluded from monitoring. Ignored resources are not evaluated for overlap or allocation rule compliance. When a resource is ignored, any space allocated to it from an IPAM pool is returned to the pool, and the resource is not re-imported through automatic import (if an automatic import allocation rule is configured for the pool).
	//
	// example:
	//
	// Managed
	ManagementStatus *string `json:"ManagementStatus,omitempty" xml:"ManagementStatus,omitempty"`
	// The list of resources that overlap with the current resource.
	OverlapDetail []*ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail `json:"OverlapDetail,omitempty" xml:"OverlapDetail,omitempty" type:"Repeated"`
	// The overlap status of the resource.
	//
	// - **Nonoverlapping**: The CIDR block of the resource does not overlap with other CIDR blocks within the same scope.
	//
	// - **Overlapping**: The CIDR block of the resource overlaps with another CIDR block within the same scope.
	//
	// - **Ignored**: The resource has been excluded from monitoring. Ignored resources are not evaluated for overlap or allocation rule compliance.
	//
	// example:
	//
	// Nonoverlapping
	OverlapStatus *string `json:"OverlapStatus,omitempty" xml:"OverlapStatus,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The Alibaba Cloud account ID of the resource ownership.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the region where the resource takes effect.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type. Valid values:
	//
	// - **VPC**: The resource type is VPC.
	//
	// - **VSwitch**: The resource type is vSwitch.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The status of the resource in the IPAM pool. Valid values:
	//
	// - **Created**: created.
	//
	// - **Deleted**: deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance ID of the VPC-connected instance to which the resource belongs.
	//
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrs) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetComplianceStatus() *string {
	return s.ComplianceStatus
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpCountDetail() *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	return s.IpCountDetail
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpUsage() *string {
	return s.IpUsage
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpamAllocationId() *string {
	return s.IpamAllocationId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetManagementStatus() *string {
	return s.ManagementStatus
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetOverlapDetail() []*ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	return s.OverlapDetail
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetOverlapStatus() *string {
	return s.OverlapStatus
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetStatus() *string {
	return s.Status
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetAliUid(v int64) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.AliUid = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.Cidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetComplianceStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ComplianceStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpCountDetail(v *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpCountDetail = v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpUsage(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpUsage = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamAllocationId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamAllocationId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamPoolId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetIpamScopeId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetManagementStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ManagementStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetOverlapDetail(v []*ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.OverlapDetail = v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetOverlapStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.OverlapStatus = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceName(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceName = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceOwnerId(v int64) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceRegionId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetResourceType(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.ResourceType = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetSourceCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetStatus(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.Status = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) SetVpcId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrs {
	s.VpcId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrs) Validate() error {
	if s.IpCountDetail != nil {
		if err := s.IpCountDetail.Validate(); err != nil {
			return err
		}
	}
	if s.OverlapDetail != nil {
		for _, item := range s.OverlapDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail struct {
	// The number of available IP addresses.
	//
	// example:
	//
	// 252
	FreeIpCount *string `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	// The total number of IP addresses.
	//
	// example:
	//
	// 256
	TotalIpCount *string `json:"TotalIpCount,omitempty" xml:"TotalIpCount,omitempty"`
	// The number of allocated IP addresses.
	//
	// example:
	//
	// 4
	UsedIpCount *string `json:"UsedIpCount,omitempty" xml:"UsedIpCount,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) GetFreeIpCount() *string {
	return s.FreeIpCount
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) GetTotalIpCount() *string {
	return s.TotalIpCount
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) GetUsedIpCount() *string {
	return s.UsedIpCount
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetFreeIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.FreeIpCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetTotalIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.TotalIpCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) SetUsedIpCount(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail {
	s.UsedIpCount = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsIpCountDetail) Validate() error {
	return dara.Validate(s)
}

type ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail struct {
	// The CIDR block of the resource that overlaps with the current resource.
	//
	// example:
	//
	// 192.168.1.0/24
	OverlapResourceCidr *string `json:"OverlapResourceCidr,omitempty" xml:"OverlapResourceCidr,omitempty"`
	// The instance ID of the resource that overlaps with the current resource.
	//
	// example:
	//
	// vpc-aq3fjgnig5av6jb8d****
	OverlapResourceId *string `json:"OverlapResourceId,omitempty" xml:"OverlapResourceId,omitempty"`
	// The region of the instance that overlaps with the current resource.
	//
	// example:
	//
	// cn-hangzhou
	OverlapResourceRegion *string `json:"OverlapResourceRegion,omitempty" xml:"OverlapResourceRegion,omitempty"`
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) GetOverlapResourceCidr() *string {
	return s.OverlapResourceCidr
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) GetOverlapResourceId() *string {
	return s.OverlapResourceId
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) GetOverlapResourceRegion() *string {
	return s.OverlapResourceRegion
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceCidr(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceCidr = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceId(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceId = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) SetOverlapResourceRegion(v string) *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail {
	s.OverlapResourceRegion = &v
	return s
}

func (s *ListIpamResourceCidrsResponseBodyIpamResourceCidrsOverlapDetail) Validate() error {
	return dara.Validate(s)
}
