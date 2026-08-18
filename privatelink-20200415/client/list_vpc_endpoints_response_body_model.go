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
	// The endpoint information.
	Endpoints []*ListVpcEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates whether a next query exists. Valid values:
	//
	// - If **NextToken*	- is empty, no next query exists.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
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
	// The protocol version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The connection bandwidth of the endpoint. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The endpoint connection status. Valid values:
	//
	// - **Pending**: being modified.
	//
	// - **Connecting**: connecting.
	//
	// - **Connected**: connected.
	//
	// - **Disconnecting**: disconnecting.
	//
	// - **Disconnected**: disconnected.
	//
	// - **Deleting**: being deleted.
	//
	// - **ServiceDeleted**: the corresponding endpoint service has been deleted.
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
	// The cross-region bandwidth of the endpoint. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	CrossRegionBandwidth *int32 `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// The business status of the endpoint. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: locked due to overdue payment.
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
	// The endpoint domain name.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The endpoint ID.
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
	// - **Creating**: being created.
	//
	// - **Active**: available.
	//
	// - **Pending**: being modified.
	//
	// - **Deleting**: being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The endpoint type. Valid values:
	//
	// - **Interface**: interface endpoint.
	//
	// - **Reverse**: reverse endpoint.
	//
	// - **GatewayLoadBalancer**: Gateway Load Balancer endpoint (GWLBe). You can create a Gateway Load Balancer endpoint (GWLBe) to connect to a Gateway Load Balancer (GWLB) for load balancing.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The RAM access policy. For more information about policy definitions, see [Basic elements of a permission policy](https://help.aliyun.com/document_detail/93738.html).
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Specifies whether managed protection is enabled. This parameter takes effect only for STS-based calls. Valid values:
	//
	// - **true**: enabled. After managed protection is enabled, only the same user who created the endpoint can modify or delete it through STS.
	//
	// - **false**: disabled.
	ProtectedEnabled *bool `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the resource group.
	//
	// example:
	//
	// 1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same account. Valid values:
	//
	// - **true**: same account.
	//
	// - **false**: different accounts.
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the endpoint service associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service associated with the endpoint.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The region ID of the endpoint service associated with the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The list of tags.
	Tags []*ListVpcEndpointsResponseBodyEndpointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether zone affinity is enabled for the endpoint domain name to resolve to the connected service. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
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

func (s *ListVpcEndpointsResponseBodyEndpoints) GetProtectedEnabled() *bool {
	return s.ProtectedEnabled
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

func (s *ListVpcEndpointsResponseBodyEndpoints) SetProtectedEnabled(v bool) *ListVpcEndpointsResponseBodyEndpoints {
	s.ProtectedEnabled = &v
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
	// The tag key of the instance.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance.
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
