// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *ListVpcEndpointsRequest
	GetAddressIpVersion() *string
	SetConnectionStatus(v string) *ListVpcEndpointsRequest
	GetConnectionStatus() *string
	SetEndpointId(v string) *ListVpcEndpointsRequest
	GetEndpointId() *string
	SetEndpointName(v string) *ListVpcEndpointsRequest
	GetEndpointName() *string
	SetEndpointStatus(v string) *ListVpcEndpointsRequest
	GetEndpointStatus() *string
	SetEndpointType(v string) *ListVpcEndpointsRequest
	GetEndpointType() *string
	SetMaxResults(v int32) *ListVpcEndpointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcEndpointsRequest
	GetResourceGroupId() *string
	SetServiceName(v string) *ListVpcEndpointsRequest
	GetServiceName() *string
	SetServiceRegionId(v string) *ListVpcEndpointsRequest
	GetServiceRegionId() *string
	SetTag(v []*ListVpcEndpointsRequestTag) *ListVpcEndpointsRequest
	GetTag() []*ListVpcEndpointsRequestTag
	SetVpcId(v string) *ListVpcEndpointsRequest
	GetVpcId() *string
}

type ListVpcEndpointsRequest struct {
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
	// 	- **Deleting**: The connection is being deleted.
	//
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
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
	// 	- If a next request is to be performed, set the parameter to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
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
	// 1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The tags.
	Tag []*ListVpcEndpointsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the endpoint belongs.
	//
	// example:
	//
	// vpc-fdjkf789dfdfdfde****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListVpcEndpointsRequest) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListVpcEndpointsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointsRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListVpcEndpointsRequest) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *ListVpcEndpointsRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListVpcEndpointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointsRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ListVpcEndpointsRequest) GetTag() []*ListVpcEndpointsRequestTag {
	return s.Tag
}

func (s *ListVpcEndpointsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcEndpointsRequest) SetAddressIpVersion(v string) *ListVpcEndpointsRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetConnectionStatus(v string) *ListVpcEndpointsRequest {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointId(v string) *ListVpcEndpointsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointName(v string) *ListVpcEndpointsRequest {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointStatus(v string) *ListVpcEndpointsRequest {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetEndpointType(v string) *ListVpcEndpointsRequest {
	s.EndpointType = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetMaxResults(v int32) *ListVpcEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetNextToken(v string) *ListVpcEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetRegionId(v string) *ListVpcEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetResourceGroupId(v string) *ListVpcEndpointsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetServiceName(v string) *ListVpcEndpointsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetServiceRegionId(v string) *ListVpcEndpointsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetTag(v []*ListVpcEndpointsRequestTag) *ListVpcEndpointsRequest {
	s.Tag = v
	return s
}

func (s *ListVpcEndpointsRequest) SetVpcId(v string) *ListVpcEndpointsRequest {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointsRequest) Validate() error {
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

type ListVpcEndpointsRequestTag struct {
	// The key of the tag added to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
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

func (s ListVpcEndpointsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointsRequestTag) SetKey(v string) *ListVpcEndpointsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointsRequestTag) SetValue(v string) *ListVpcEndpointsRequestTag {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointsRequestTag) Validate() error {
	return dara.Validate(s)
}
