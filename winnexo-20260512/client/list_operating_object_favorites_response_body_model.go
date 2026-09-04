// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperatingObjectFavoritesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOperatingObjectFavoritesResponseBody
	GetCode() *string
	SetGraphName(v string) *ListOperatingObjectFavoritesResponseBody
	GetGraphName() *string
	SetHasMore(v bool) *ListOperatingObjectFavoritesResponseBody
	GetHasMore() *bool
	SetItems(v []*ListOperatingObjectFavoritesResponseBodyItems) *ListOperatingObjectFavoritesResponseBody
	GetItems() []*ListOperatingObjectFavoritesResponseBodyItems
	SetMessage(v string) *ListOperatingObjectFavoritesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListOperatingObjectFavoritesResponseBody
	GetNextToken() *string
	SetObjectType(v string) *ListOperatingObjectFavoritesResponseBody
	GetObjectType() *string
	SetOperatingObjectName(v string) *ListOperatingObjectFavoritesResponseBody
	GetOperatingObjectName() *string
	SetPageSize(v int64) *ListOperatingObjectFavoritesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListOperatingObjectFavoritesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListOperatingObjectFavoritesResponseBody
	GetTotal() *int64
}

type ListOperatingObjectFavoritesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The graph name. Call listGraphs to retrieve available graphs.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// Indicates whether more pages are available.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The MCP card list.
	Items []*ListOperatingObjectFavoritesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJ2IjoxLCJpZCI6OTAyfQ.c2lnbmF0dXJlX2V4YW1wbGU
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The digital employee name (operating object name).
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page size.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of results.
	//
	// example:
	//
	// 1001
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOperatingObjectFavoritesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperatingObjectFavoritesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperatingObjectFavoritesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOperatingObjectFavoritesResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *ListOperatingObjectFavoritesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListOperatingObjectFavoritesResponseBody) GetItems() []*ListOperatingObjectFavoritesResponseBodyItems {
	return s.Items
}

func (s *ListOperatingObjectFavoritesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOperatingObjectFavoritesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOperatingObjectFavoritesResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListOperatingObjectFavoritesResponseBody) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOperatingObjectFavoritesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOperatingObjectFavoritesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperatingObjectFavoritesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListOperatingObjectFavoritesResponseBody) SetCode(v string) *ListOperatingObjectFavoritesResponseBody {
	s.Code = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetGraphName(v string) *ListOperatingObjectFavoritesResponseBody {
	s.GraphName = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetHasMore(v bool) *ListOperatingObjectFavoritesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetItems(v []*ListOperatingObjectFavoritesResponseBodyItems) *ListOperatingObjectFavoritesResponseBody {
	s.Items = v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetMessage(v string) *ListOperatingObjectFavoritesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetNextToken(v string) *ListOperatingObjectFavoritesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetObjectType(v string) *ListOperatingObjectFavoritesResponseBody {
	s.ObjectType = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetOperatingObjectName(v string) *ListOperatingObjectFavoritesResponseBody {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetPageSize(v int64) *ListOperatingObjectFavoritesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetRequestId(v string) *ListOperatingObjectFavoritesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) SetTotal(v int64) *ListOperatingObjectFavoritesResponseBody {
	s.Total = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperatingObjectFavoritesResponseBodyItems struct {
	// The to-do card type description.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The follow time. The value is a Unix timestamp in seconds.
	//
	// example:
	//
	// 1787822400
	FollowedAt *int64 `json:"followedAt,omitempty" xml:"followedAt,omitempty"`
	// The graph name.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the recommended item. The value can be a **feedId*	- or a micro-application ID.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object name.
	//
	// example:
	//
	// 469ac312-403c-41fb-aae3-de5260e30906
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The bound object type, such as customer or project.
	//
	// example:
	//
	// table
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ListOperatingObjectFavoritesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListOperatingObjectFavoritesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetFollowedAt() *int64 {
	return s.FollowedAt
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetGraphName() *string {
	return s.GraphName
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetDescription(v string) *ListOperatingObjectFavoritesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetFollowedAt(v int64) *ListOperatingObjectFavoritesResponseBodyItems {
	s.FollowedAt = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetGraphName(v string) *ListOperatingObjectFavoritesResponseBodyItems {
	s.GraphName = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetObjectId(v string) *ListOperatingObjectFavoritesResponseBodyItems {
	s.ObjectId = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetObjectName(v string) *ListOperatingObjectFavoritesResponseBodyItems {
	s.ObjectName = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) SetObjectType(v string) *ListOperatingObjectFavoritesResponseBodyItems {
	s.ObjectType = &v
	return s
}

func (s *ListOperatingObjectFavoritesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
