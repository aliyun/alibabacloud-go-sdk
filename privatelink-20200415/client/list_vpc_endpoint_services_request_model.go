// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *ListVpcEndpointServicesRequest
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *ListVpcEndpointServicesRequest
	GetAutoAcceptEnabled() *bool
	SetMaxResults(v int32) *ListVpcEndpointServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServicesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointServicesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcEndpointServicesRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListVpcEndpointServicesRequest
	GetResourceId() *string
	SetServiceBusinessStatus(v string) *ListVpcEndpointServicesRequest
	GetServiceBusinessStatus() *string
	SetServiceId(v string) *ListVpcEndpointServicesRequest
	GetServiceId() *string
	SetServiceName(v string) *ListVpcEndpointServicesRequest
	GetServiceName() *string
	SetServiceResourceType(v string) *ListVpcEndpointServicesRequest
	GetServiceResourceType() *string
	SetServiceStatus(v string) *ListVpcEndpointServicesRequest
	GetServiceStatus() *string
	SetTag(v []*ListVpcEndpointServicesRequestTag) *ListVpcEndpointServicesRequest
	GetTag() []*ListVpcEndpointServicesRequestTag
	SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesRequest
	GetZoneAffinityEnabled() *bool
}

type ListVpcEndpointServicesRequest struct {
	// The IP address version. Valid values:
	//
	// - **IPv4**: IPv4 type.
	//
	// - **DualStack**: Dual-stack type.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connections. Valid values:
	//
	// - **true**: Automatically accept endpoint connections.
	//
	// - **false**: Do not automatically accept endpoint connections.
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The number of entries to return per page. Valid values: **1*	- to **1000**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token for the next query. Valid values:
	//
	// - Leave this parameter empty for the first query or when no further results exist.
	//
	// - If another query is needed, set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the endpoint service is deployed.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the service resource.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The business status of the endpoint service. Valid values:
	//
	// - **Normal**: The endpoint service is running as expected.
	//
	// - **FinancialLocked**: The endpoint service is locked due to an overdue payment.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The ID of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// - **slb**: The service resource is a Classic Load Balancer (CLB) instance.
	//
	// - **alb**: The service resource is an Application Load Balancer (ALB) instance.
	//
	// - **nlb**: The service resource is a Network Load Balancer (NLB) instance.
	//
	// - **gwlb**: The service resource is a Gateway Load Balancer (GWLB) instance.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The status of the endpoint service. Valid values:
	//
	// - **Creating**: The endpoint service is being created.
	//
	// - **Pending**: The endpoint service is being modified.
	//
	// - **Active**: The endpoint service is available.
	//
	// - **Deleting**: The endpoint service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The list of tags.
	Tag []*ListVpcEndpointServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether zonal affinity is enabled for domain name resolution. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListVpcEndpointServicesRequest) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *ListVpcEndpointServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServicesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointServicesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListVpcEndpointServicesRequest) GetServiceBusinessStatus() *string {
	return s.ServiceBusinessStatus
}

func (s *ListVpcEndpointServicesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServicesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesRequest) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *ListVpcEndpointServicesRequest) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListVpcEndpointServicesRequest) GetTag() []*ListVpcEndpointServicesRequestTag {
	return s.Tag
}

func (s *ListVpcEndpointServicesRequest) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *ListVpcEndpointServicesRequest) SetAddressIpVersion(v string) *ListVpcEndpointServicesRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetAutoAcceptEnabled(v bool) *ListVpcEndpointServicesRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetMaxResults(v int32) *ListVpcEndpointServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetNextToken(v string) *ListVpcEndpointServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetRegionId(v string) *ListVpcEndpointServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetResourceGroupId(v string) *ListVpcEndpointServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetResourceId(v string) *ListVpcEndpointServicesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceBusinessStatus(v string) *ListVpcEndpointServicesRequest {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceId(v string) *ListVpcEndpointServicesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceName(v string) *ListVpcEndpointServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceResourceType(v string) *ListVpcEndpointServicesRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetServiceStatus(v string) *ListVpcEndpointServicesRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetTag(v []*ListVpcEndpointServicesRequestTag) *ListVpcEndpointServicesRequest {
	s.Tag = v
	return s
}

func (s *ListVpcEndpointServicesRequest) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointServicesRequestTag struct {
	// The tag key of the instance. You can specify up to 20 tag keys. The key cannot be an empty string.
	//
	// The key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. You can specify up to 20 tag values. The value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointServicesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointServicesRequestTag) SetKey(v string) *ListVpcEndpointServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesRequestTag) SetValue(v string) *ListVpcEndpointServicesRequestTag {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointServicesRequestTag) Validate() error {
	return dara.Validate(s)
}
