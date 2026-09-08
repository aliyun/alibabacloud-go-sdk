// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablePropagationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTablePropagationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterRouteTablePropagationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterRouteTablePropagationsResponseBody
	GetTotalCount() *int32
	SetTransitRouterPropagations(v []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) *ListTransitRouterRouteTablePropagationsResponseBody
	GetTransitRouterPropagations() []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations
}

type ListTransitRouterRouteTablePropagationsResponseBody struct {
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query in a paged query.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04C81E0D-945E-4D61-A561-3DEA322F243B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of route learning relationships.
	TransitRouterPropagations []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations `json:"TransitRouterPropagations,omitempty" xml:"TransitRouterPropagations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTablePropagationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) GetTransitRouterPropagations() []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	return s.TransitRouterPropagations
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetNextToken(v string) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetRequestId(v string) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) SetTransitRouterPropagations(v []*ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) *ListTransitRouterRouteTablePropagationsResponseBody {
	s.TransitRouterPropagations = v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBody) Validate() error {
	if s.TransitRouterPropagations != nil {
		for _, item := range s.TransitRouterPropagations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations struct {
	// The network instance ID.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The network instance type. Valid values:
	//
	// - **VPC**: virtual private cloud (VPC) instance.
	//
	// - **VBR**: virtual border router (VBR) instance.
	//
	// - **TR**: transit router instance.
	//
	// - **VPN**: VPN connection.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the route learning relationship. Valid values:
	//
	// - **Enabling**: being enabled.
	//
	// - **Disabling**: being disabled.
	//
	// - **Active**: active.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The network instance connection ID.
	//
	// example:
	//
	// tr-attach-vx6iwhjr1x1j78****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the Enterprise Edition transit router route table.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetResourceId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetResourceType(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetStatus(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponseBodyTransitRouterPropagations) Validate() error {
	return dara.Validate(s)
}
