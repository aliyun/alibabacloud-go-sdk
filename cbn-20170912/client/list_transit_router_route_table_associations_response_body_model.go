// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTableAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTableAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterRouteTableAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterRouteTableAssociationsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAssociations(v []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) *ListTransitRouterRouteTableAssociationsResponseBody
	GetTransitRouterAssociations() []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations
}

type ListTransitRouterRouteTableAssociationsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- was not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// a415****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F6B1D9AB-176D-4399-801D-8BC576F4EB0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of associated forwarding correlations.
	TransitRouterAssociations []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations `json:"TransitRouterAssociations,omitempty" xml:"TransitRouterAssociations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTableAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) GetTransitRouterAssociations() []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	return s.TransitRouterAssociations
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetNextToken(v string) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetRequestId(v string) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) SetTransitRouterAssociations(v []*ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) *ListTransitRouterRouteTableAssociationsResponseBody {
	s.TransitRouterAssociations = v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBody) Validate() error {
	if s.TransitRouterAssociations != nil {
		for _, item := range s.TransitRouterAssociations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations struct {
	// The ID of the next hop.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of next hop. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **TR**: transit router
	//
	// 	- **VPN*	- :VPN attachment
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the associated forwarding correlation. Valid values:
	//
	// 	- **Active**: The associated forwarding correlation is available.
	//
	// 	- **Associating**: The associated forwarding correlation is being created.
	//
	// 	- **Dissociating**: The associated forwarding correlation is being deleted.
	//
	// 	- **Deleted**: The associated forwarding correlation is deleted.
	//
	// example:
	//
	// Associating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetResourceId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetResourceType(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetStatus(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponseBodyTransitRouterAssociations) Validate() error {
	return dara.Validate(s)
}
