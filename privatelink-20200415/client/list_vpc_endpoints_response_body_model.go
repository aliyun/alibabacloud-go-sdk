// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ListVpcEndpointsResponseBodyEndpoints) *ListVpcEndpointsResponseBody
	GetEndpoints() []*ListVpcEndpointsResponseBodyEndpoints
	SetMaxResults(v int32) *ListVpcEndpointsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVpcEndpointsResponseBody
	GetTotalCount() *int32
}

type ListVpcEndpointsResponseBody struct {
	// The information about the endpoints.
	Endpoints []*ListVpcEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBody) GetEndpoints() []*ListVpcEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListVpcEndpointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpcEndpointsResponseBody) SetEndpoints(v []*ListVpcEndpointsResponseBodyEndpoints) *ListVpcEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetMaxResults(v int32) *ListVpcEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetNextToken(v string) *ListVpcEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetRequestId(v string) *ListVpcEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetTotalCount(v int32) *ListVpcEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointsResponseBodyEndpoints struct {
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
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The endpoint connection is being modified.
	//
	// 	- **Connecting**: The endpoint connection is being established.
	//
	// 	- **Connected**: The endpoint connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The endpoint connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding service is deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2021-09-24T18:00:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The domain name of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Interface**: interface endpoint
	//
	// 	- **Reverse**: reverse endpoint
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The Resource Access Management (RAM) policy. For more information about policy definitions, see [Policy elements](https://help.aliyun.com/document_detail/93738.html).
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the endpoint service that is associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service that is associated with the endpoint.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tags added to the resource.
	Tags []*ListVpcEndpointsResponseBodyEndpointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointBusinessStatus() *string {
	return s.EndpointBusinessStatus
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetResourceOwner() *bool {
	return s.ResourceOwner
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetTags() []*ListVpcEndpointsResponseBodyEndpointsTags {
	return s.Tags
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcEndpointsResponseBodyEndpoints) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetAddressIpVersion(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.AddressIpVersion = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetBandwidth(v int64) *ListVpcEndpointsResponseBodyEndpoints {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetConnectionStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetCreateTime(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointBusinessStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointDescription(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointDescription = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointDomain(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointDomain = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointName(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointStatus(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetEndpointType(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetPolicyDocument(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.PolicyDocument = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetRegionId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetResourceGroupId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetResourceOwner(v bool) *ListVpcEndpointsResponseBodyEndpoints {
	s.ResourceOwner = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetServiceId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetServiceName(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetTags(v []*ListVpcEndpointsResponseBodyEndpointsTags) *ListVpcEndpointsResponseBodyEndpoints {
	s.Tags = v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetVpcId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) SetZoneAffinityEnabled(v bool) *ListVpcEndpointsResponseBodyEndpoints {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpoints) Validate() error {
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

type ListVpcEndpointsResponseBodyEndpointsTags struct {
	// The key of the tag added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointsResponseBodyEndpointsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyEndpointsTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) SetKey(v string) *ListVpcEndpointsResponseBodyEndpointsTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) SetValue(v string) *ListVpcEndpointsResponseBodyEndpointsTags {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyEndpointsTags) Validate() error {
	return dara.Validate(s)
}
