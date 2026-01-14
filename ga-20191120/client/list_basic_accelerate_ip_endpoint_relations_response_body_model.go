// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpEndpointRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpEndpointRelations(v []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) *ListBasicAccelerateIpEndpointRelationsResponseBody
	GetAccelerateIpEndpointRelations() []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations
	SetMaxResults(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody
	GetTotalCount() *string
}

type ListBasicAccelerateIpEndpointRelationsResponseBody struct {
	// A list of accelerated IP addresses and the endpoints with which the accelerated IP addresses are associated.
	AccelerateIpEndpointRelations []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations `json:"AccelerateIpEndpointRelations,omitempty" xml:"AccelerateIpEndpointRelations,omitempty" type:"Repeated"`
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

func (s ListBasicAccelerateIpEndpointRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) GetAccelerateIpEndpointRelations() []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	return s.AccelerateIpEndpointRelations
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetAccelerateIpEndpointRelations(v []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.AccelerateIpEndpointRelations = v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetMaxResults(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetRequestId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetTotalCount(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) Validate() error {
	if s.AccelerateIpEndpointRelations != nil {
		for _, item := range s.AccelerateIpEndpointRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations struct {
	// The ID of the accelerated IP address.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp1miyl2kn2racucvegr5
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// ep01
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
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
	// The type of the secondary address of the endpoint.
	//
	// 	- **primary**: a primary private IP address.
	//
	// 	- **secondary**: a secondary private IP address.
	//
	// This parameter is returned if the endpoint type is **ECS**, **ENI**, or **NLB**. If the endpoint type is set to **NLB**, **primary*	- is returned.
	//
	// example:
	//
	// primary
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	// The type of endpoint. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// 	- **NLB**: Network Load Balancer (NLB) instance.
	//
	// example:
	//
	// NLB
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the zone where the endpoint is created.
	//
	// This parameter is returned only when the endpoint type is **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// 116.132.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The association status between the accelerated IP address and endpoint.
	//
	// A value of **active*	- indicates that the accelerated IP address is associated with the endpoint.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GetState() *string {
	return s.State
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetAccelerateIpId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetAcceleratorId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointName(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointName = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointSubAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointSubAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointSubAddressType(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointSubAddressType = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointType(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointType = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointZoneId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointZoneId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetIpAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.IpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetState(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.State = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) Validate() error {
	return dara.Validate(s)
}
