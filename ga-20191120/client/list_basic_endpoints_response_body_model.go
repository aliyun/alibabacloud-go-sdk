// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ListBasicEndpointsResponseBodyEndpoints) *ListBasicEndpointsResponseBody
	GetEndpoints() []*ListBasicEndpointsResponseBodyEndpoints
	SetMaxResults(v string) *ListBasicEndpointsResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListBasicEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBasicEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListBasicEndpointsResponseBody
	GetTotalCount() *string
}

type ListBasicEndpointsResponseBody struct {
	// The endpoints that are associated with the basic GA instance.
	Endpoints []*ListBasicEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- is not returned, it indicates that no additional results exist.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBasicEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponseBody) GetEndpoints() []*ListBasicEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListBasicEndpointsResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListBasicEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBasicEndpointsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListBasicEndpointsResponseBody) SetEndpoints(v []*ListBasicEndpointsResponseBodyEndpoints) *ListBasicEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetMaxResults(v string) *ListBasicEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetNextToken(v string) *ListBasicEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetRequestId(v string) *ListBasicEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetTotalCount(v string) *ListBasicEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) Validate() error {
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

type ListBasicEndpointsResponseBodyEndpoints struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint that is associated with the basic GA instance.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The secondary address of the endpoint.
	//
	// This parameter is returned if the endpoint type is **ECS**, **ENI**, or **NLB**.
	//
	// 	- If the endpoint type is **ECS**, **EndpointSubAddress*	- returns the primary or secondary private IP address of the primary ENI.
	//
	// 	- If the endpoint type is **ENI**, **EndpointSubAddress*	- returns the primary or secondary private IP address of the secondary ENI.
	//
	// 	- If the endpoint type is **NLB**, **EndpointSubAddress*	- returns the primary private IP address of the NLB backend server.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The secondary address type of the endpoint.
	//
	// 	- **primary**: a primary private IP address.
	//
	// 	- **secondary**: a secondary private IP address.
	//
	// This parameter is returned if the type of the endpoint is set to **ECS**, **ENI**, or **NLB**. If the endpoint type is set to **NLB**, **primary*	- is returned.
	//
	// example:
	//
	// primary
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	// The type of endpoint. Valid values:
	//
	// 	- **ENI**: ENI.
	//
	// 	- **SLB**: CLB instance.
	//
	// 	- **ECS**: ECS instance.
	//
	// 	- **NLB**: NLB instance.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The zone ID of the endpoint.
	//
	// This parameter is returned only when the endpoint type is set to **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- **init**: The endpoint is being initialized.
	//
	// 	- **active**: The endpoint is available.
	//
	// 	- **updating**: The endpoint is being configured.
	//
	// 	- **binding**: The endpoint is being associated.
	//
	// 	- **unbinding**: The endpoint is being disassociated.
	//
	// 	- **deleting**: The endpoint is being deleted.
	//
	// 	- **bound**: The endpoint is associated.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListBasicEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetName() *string {
	return s.Name
}

func (s *ListBasicEndpointsResponseBodyEndpoints) GetState() *string {
	return s.State
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetAcceleratorId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointAddress(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointGroupId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointGroupId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddress(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddressType(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddressType = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointType(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointZoneId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointZoneId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetName(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetState(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.State = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}
