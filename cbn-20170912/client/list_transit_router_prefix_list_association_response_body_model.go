// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPrefixListAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTransitRouterPrefixListAssociationResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransitRouterPrefixListAssociationResponseBody
	GetPageSize() *int32
	SetPrefixLists(v []*ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) *ListTransitRouterPrefixListAssociationResponseBody
	GetPrefixLists() []*ListTransitRouterPrefixListAssociationResponseBodyPrefixLists
	SetRequestId(v string) *ListTransitRouterPrefixListAssociationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterPrefixListAssociationResponseBody
	GetTotalCount() *int32
}

type ListTransitRouterPrefixListAssociationResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of prefix lists.
	PrefixLists []*ListTransitRouterPrefixListAssociationResponseBodyPrefixLists `json:"PrefixLists,omitempty" xml:"PrefixLists,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6005CA94-676E-1FEE-985E-7602EFAADD6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTransitRouterPrefixListAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPrefixListAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) GetPrefixLists() []*ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	return s.PrefixLists
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) SetPageNumber(v int32) *ListTransitRouterPrefixListAssociationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) SetPageSize(v int32) *ListTransitRouterPrefixListAssociationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) SetPrefixLists(v []*ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) *ListTransitRouterPrefixListAssociationResponseBody {
	s.PrefixLists = v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) SetRequestId(v string) *ListTransitRouterPrefixListAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) SetTotalCount(v int32) *ListTransitRouterPrefixListAssociationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBody) Validate() error {
	if s.PrefixLists != nil {
		for _, item := range s.PrefixLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterPrefixListAssociationResponseBodyPrefixLists struct {
	// The ID of the next hop.
	//
	// > A value of **BlackHole*	- indicates that all the CIDR blocks in the prefix list are blackhole routes. Packets destined for the CIDR blocks are dropped.
	//
	// example:
	//
	// tr-attach-flbq507rg2ckrj****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The ID of the network instance associated with the next hop connection.
	//
	// example:
	//
	// vpc-6eh7fp9hdqa2wv85t****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **BlackHole**: All the CIDR blocks in the prefix list are blackhole routes. Packets destined for the CIDR blocks are dropped.
	//
	// 	- **VPC**: The next hop of the CIDR blocks in the prefix list is a VPC connection.
	//
	// 	- **VBR**: The next hop of the CIDR blocks in the prefix list is a VBR connection.
	//
	// 	- **TR**: The next hop of the CIDR blocks in the prefix list is an inter-region connection.
	//
	// example:
	//
	// VPC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The ID of the Alibaba Cloud account to which the prefix list belongs.
	//
	// example:
	//
	// 1210123456123456
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-6ehtn5kqxgeyy08fi****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The status of the prefix list. Valid values:
	//
	// 	- **Active**: The prefix list is effective.
	//
	// 	- **Updating**: The prefix list is being updated.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-6ehx7q2jze8ch5ji0****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the route table of the transit router.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouterTableId *string `json:"TransitRouterTableId,omitempty" xml:"TransitRouterTableId,omitempty"`
}

func (s ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetNextHop() *string {
	return s.NextHop
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) GetTransitRouterTableId() *string {
	return s.TransitRouterTableId
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetNextHop(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.NextHop = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetNextHopInstanceId(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.NextHopInstanceId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetNextHopType(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.NextHopType = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetOwnerUid(v int64) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.OwnerUid = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetPrefixListId(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.PrefixListId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetStatus(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.Status = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetTransitRouterId(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) SetTransitRouterTableId(v string) *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists {
	s.TransitRouterTableId = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponseBodyPrefixLists) Validate() error {
	return dara.Validate(s)
}
