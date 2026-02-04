// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenChildInstanceRouteEntriesToAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetCenId() *string
	SetChildInstanceRouteTableId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetChildInstanceRouteTableId() *string
	SetMaxResults(v int32) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetResourceOwnerId() *int64
	SetRouteFilter(v []*ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetRouteFilter() []*ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter
	SetServiceType(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetServiceType() *string
	SetTransitRouterAttachmentId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type ListCenChildInstanceRouteEntriesToAttachmentRequest struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-dc4vwznpwbobrl****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the route table configured on the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp174d1gje79u1g4t****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query and no subsequent queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The filter condition for the destination CIDR block.
	RouteFilter []*ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter `json:"RouteFilter,omitempty" xml:"RouteFilter,omitempty" type:"Repeated"`
	// Specifies whether to host the route. If you leave the parameter empty, the route is not hosted. A value of TR specifies that the route is hosted on a transit router.
	//
	// example:
	//
	// TR
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The ID of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s ListCenChildInstanceRouteEntriesToAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCenChildInstanceRouteEntriesToAttachmentRequest) GoString() string {
	return s.String()
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetRouteFilter() []*ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter {
	return s.RouteFilter
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetCenId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetChildInstanceRouteTableId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetMaxResults(v int32) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetNextToken(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.NextToken = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetOwnerAccount(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetOwnerId(v int64) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetResourceOwnerAccount(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetResourceOwnerId(v int64) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetRouteFilter(v []*ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.RouteFilter = v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetServiceType(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.ServiceType = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) SetTransitRouterAttachmentId(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequest) Validate() error {
	if s.RouteFilter != nil {
		for _, item := range s.RouteFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter struct {
	// The match mode of the route.
	//
	// 	- **prefix-exact-match**: exact match.
	//
	// example:
	//
	// prefix-exact-match
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The destination CIDR blocks.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) GoString() string {
	return s.String()
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) GetKey() *string {
	return s.Key
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) GetValue() []*string {
	return s.Value
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) SetKey(v string) *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter {
	s.Key = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) SetValue(v []*string) *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter {
	s.Value = v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentRequestRouteFilter) Validate() error {
	return dara.Validate(s)
}
