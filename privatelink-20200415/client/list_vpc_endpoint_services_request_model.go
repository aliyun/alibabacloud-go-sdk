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
	// The protocol. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **1000**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The endpoint service ID.
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
	// 	- **slb**: a Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The tags.
	Tag []*ListVpcEndpointServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
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
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
