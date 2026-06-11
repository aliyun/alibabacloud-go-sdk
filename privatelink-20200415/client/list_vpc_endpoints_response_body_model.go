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
	// A list of endpoints.
	Endpoints []*ListVpcEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results.
	//
	// - If **NextToken*	- is empty, no next page exists.
	//
	// - If a value is returned, use it in your next request to retrieve the next page of results.
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
	// The total number of entries that match the query.
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
	// The IP address family. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The connection bandwidth of the endpoint, in Mbps.
	//
	// example:
	//
	// 1024
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The connection status of the endpoint. Valid values:
	//
	// - **Pending**: The connection is being modified.
	//
	// - **Connecting**: The endpoint is being connected.
	//
	// - **Connected**: The endpoint is connected.
	//
	// - **Disconnecting**: The endpoint is being disconnected.
	//
	// - **Disconnected**: The endpoint is disconnected.
	//
	// - **Deleting**: The endpoint is being deleted.
	//
	// - **ServiceDeleted**: The associated endpoint service has been deleted.
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
	// The bandwidth of the cross-region connection, in Mbps.
	//
	// example:
	//
	// 1000
	CrossRegionBandwidth *int32 `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// The business status of the endpoint. Valid values:
	//
	// - **Normal**: The endpoint is running as expected.
	//
	// - **FinancialLocked**: The endpoint is locked due to an overdue payment.
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
	// The status of the endpoint. Valid values:
	//
	// - **Creating**: The endpoint is being created.
	//
	// - **Active**: The endpoint is available.
	//
	// - **Pending**: The endpoint is being modified.
	//
	// - **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Interface**: an interface endpoint.
	//
	// - **Reverse**: a reverse endpoint.
	//
	// - **GatewayLoadBalancer**: a gateway load balancer endpoint.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The RAM access policy. For details on the policy syntax, see [Basic elements of a RAM policy](https://help.aliyun.com/document_detail/93738.html).
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the region that contains the endpoint.
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
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the associated endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the associated endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The ID of the region where the associated endpoint service is deployed.
	//
	// example:
	//
	// cn-huhehaote
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// A list of tags.
	Tags []*ListVpcEndpointsResponseBodyEndpointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether zone-aware DNS resolution is enabled. Valid values:
	//
	// - **true**
	//
	// - **false**
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

func (s *ListVpcEndpointsResponseBodyEndpoints) GetCrossRegionBandwidth() *int32 {
	return s.CrossRegionBandwidth
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

func (s *ListVpcEndpointsResponseBodyEndpoints) GetServiceRegionId() *string {
	return s.ServiceRegionId
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

func (s *ListVpcEndpointsResponseBodyEndpoints) SetCrossRegionBandwidth(v int32) *ListVpcEndpointsResponseBodyEndpoints {
	s.CrossRegionBandwidth = &v
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

func (s *ListVpcEndpointsResponseBodyEndpoints) SetServiceRegionId(v string) *ListVpcEndpointsResponseBodyEndpoints {
	s.ServiceRegionId = &v
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
