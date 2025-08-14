// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenChildInstanceRouteEntriesToAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody
	GetRequestId() *string
	SetRouteEntry(v []*ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody
	GetRouteEntry() []*ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry
}

type ListCenChildInstanceRouteEntriesToAttachmentResponseBody struct {
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If **NextToken*	- is not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// 530BC816-F575-412A-AAB2-435125D26328
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the route.
	RouteEntry []*ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry `json:"RouteEntry,omitempty" xml:"RouteEntry,omitempty" type:"Repeated"`
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) GetRouteEntry() []*ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	return s.RouteEntry
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) SetNextToken(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) SetRequestId(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) SetRouteEntry(v []*ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) *ListCenChildInstanceRouteEntriesToAttachmentResponseBody {
	s.RouteEntry = v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-9adwg6ghpq8oq4dp7q
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the route table configured on the network instance.
	//
	// example:
	//
	// vtb-bp1tlaj1c4nxr2t3e****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 10.0.0.0/8
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Indicates whether the route is hosted. If the parameter is empty, the route is not hosted. A value of TR indicates that the route is hosted on a transit router.
	//
	// example:
	//
	// TR
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the route. Valid values:
	//
	// 	- **Available**: The route is available.
	//
	// 	- **Pending**: The route is being configured.
	//
	// 	- **Modifying**: the route is being modified.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-y463sghkkv1loe****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GoString() string {
	return s.String()
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetCenId() *string {
	return s.CenId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetCenId(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.CenId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetChildInstanceRouteTableId(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetDestinationCidrBlock(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetServiceType(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.ServiceType = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetStatus(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.Status = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) SetTransitRouterAttachmentId(v string) *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponseBodyRouteEntry) Validate() error {
	return dara.Validate(s)
}
