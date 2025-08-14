// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterCidrAllocationResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterCidrAllocationResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterCidrAllocationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterCidrAllocationResponseBody
	GetTotalCount() *int32
	SetTransitRouterCidrAllocations(v []*ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) *ListTransitRouterCidrAllocationResponseBody
	GetTransitRouterCidrAllocations() []*ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations
}

type ListTransitRouterCidrAllocationResponseBody struct {
	// The number of entries returned per page.
	//
	// 	- If no value is specified for **MaxResults**, query results are returned in one batch. The value of **MaxResults*	- indicates the total number of entries.
	//
	// 	- If a value is specified for **MaxResults**, query results are returned in batches. The value of **MaxResults*	- in the response indicates the number of entries in the current batch.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If **NextToken*	- was not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the CIDR blocks that have IP addresses allocated to network instances.
	TransitRouterCidrAllocations []*ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations `json:"TransitRouterCidrAllocations,omitempty" xml:"TransitRouterCidrAllocations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterCidrAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrAllocationResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterCidrAllocationResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterCidrAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterCidrAllocationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterCidrAllocationResponseBody) GetTransitRouterCidrAllocations() []*ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	return s.TransitRouterCidrAllocations
}

func (s *ListTransitRouterCidrAllocationResponseBody) SetMaxResults(v int32) *ListTransitRouterCidrAllocationResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBody) SetNextToken(v string) *ListTransitRouterCidrAllocationResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBody) SetRequestId(v string) *ListTransitRouterCidrAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBody) SetTotalCount(v int32) *ListTransitRouterCidrAllocationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBody) SetTransitRouterCidrAllocations(v []*ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) *ListTransitRouterCidrAllocationResponseBody {
	s.TransitRouterCidrAllocations = v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations struct {
	// The CIDR blocks that have IP addresses allocated to network instances.
	//
	// example:
	//
	// 192.168.10.0/28
	AllocatedCidrBlock *string `json:"AllocatedCidrBlock,omitempty" xml:"AllocatedCidrBlock,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-2nalp6yksc805w****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The name of the network instance connection.
	//
	// example:
	//
	// nametest
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The CIDR block of the transit router.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the CIDR block.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
}

func (s ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GetAllocatedCidrBlock() *string {
	return s.AllocatedCidrBlock
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GetCidr() *string {
	return s.Cidr
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) SetAllocatedCidrBlock(v string) *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	s.AllocatedCidrBlock = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) SetAttachmentId(v string) *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	s.AttachmentId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) SetAttachmentName(v string) *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	s.AttachmentName = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) SetCidr(v string) *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	s.Cidr = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) SetTransitRouterCidrId(v string) *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ListTransitRouterCidrAllocationResponseBodyTransitRouterCidrAllocations) Validate() error {
	return dara.Validate(s)
}
