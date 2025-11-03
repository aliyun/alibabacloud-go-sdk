// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcGatewayEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ListVpcGatewayEndpointsResponseBodyEndpoints) *ListVpcGatewayEndpointsResponseBody
	GetEndpoints() []*ListVpcGatewayEndpointsResponseBodyEndpoints
	SetMaxResults(v int64) *ListVpcGatewayEndpointsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListVpcGatewayEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcGatewayEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVpcGatewayEndpointsResponseBody
	GetTotalCount() *int64
}

type ListVpcGatewayEndpointsResponseBody struct {
	// The list of gateway endpoints.
	Endpoints []*ListVpcGatewayEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AB1129F-32C1-5E4D-9E22-E4A859CA46EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcGatewayEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsResponseBody) GetEndpoints() []*ListVpcGatewayEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListVpcGatewayEndpointsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVpcGatewayEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcGatewayEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcGatewayEndpointsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVpcGatewayEndpointsResponseBody) SetEndpoints(v []*ListVpcGatewayEndpointsResponseBodyEndpoints) *ListVpcGatewayEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBody) SetMaxResults(v int64) *ListVpcGatewayEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBody) SetNextToken(v string) *ListVpcGatewayEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBody) SetRequestId(v string) *ListVpcGatewayEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBody) SetTotalCount(v int64) *ListVpcGatewayEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBody) Validate() error {
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

type ListVpcGatewayEndpointsResponseBodyEndpoints struct {
	// The ID of the route table associated with the gateway endpoint.
	AssociatedRouteTables []*string `json:"AssociatedRouteTables,omitempty" xml:"AssociatedRouteTables,omitempty" type:"Repeated"`
	// The time when the endpoint was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-09-08T08:43:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the gateway endpoint.
	//
	// example:
	//
	// test_description
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The ID of the gateway endpoint.
	//
	// example:
	//
	// vpce-bp1i1212ss2whuwyw****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the gateway endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The status of the gateway endpoint. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Associating**
	//
	// 	- **Dissociating**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The access policy for the cloud service.
	//
	// For more information about the syntax and structure of the access policy, see [Policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the resource group to which the gateway endpoint belongs.
	//
	// example:
	//
	// rg-acfmxvfvazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tag list.
	Tags []*ListVpcGatewayEndpointsResponseBodyEndpointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the gateway endpoint belongs.
	//
	// example:
	//
	// vpc-bp1gsk7h12ew7oegk****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcGatewayEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetAssociatedRouteTables() []*string {
	return s.AssociatedRouteTables
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetTags() []*ListVpcGatewayEndpointsResponseBodyEndpointsTags {
	return s.Tags
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetAssociatedRouteTables(v []*string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.AssociatedRouteTables = v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetCreationTime(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.CreationTime = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetEndpointDescription(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.EndpointDescription = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetEndpointName(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.EndpointName = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetEndpointStatus(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetPolicyDocument(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.PolicyDocument = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetResourceGroupId(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetServiceName(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.ServiceName = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetTags(v []*ListVpcGatewayEndpointsResponseBodyEndpointsTags) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.Tags = v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) SetVpcId(v string) *ListVpcGatewayEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpoints) Validate() error {
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

type ListVpcGatewayEndpointsResponseBodyEndpointsTags struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcGatewayEndpointsResponseBodyEndpointsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsResponseBodyEndpointsTags) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpointsTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpointsTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpointsTags) SetKey(v string) *ListVpcGatewayEndpointsResponseBodyEndpointsTags {
	s.Key = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpointsTags) SetValue(v string) *ListVpcGatewayEndpointsResponseBodyEndpointsTags {
	s.Value = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponseBodyEndpointsTags) Validate() error {
	return dara.Validate(s)
}
