// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListListsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListsRequest
	GetPageSize() *int32
	SetQueryArgs(v *ListListsRequestQueryArgs) *ListListsRequest
	GetQueryArgs() *ListListsRequestQueryArgs
}

type ListListsRequest struct {
	// Specifies the page number for paginated results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies the number of results to return per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A JSON object containing query parameters to filter the results.
	//
	// example:
	//
	// ListLists
	QueryArgs *ListListsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListListsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListsRequest) GoString() string {
	return s.String()
}

func (s *ListListsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListsRequest) GetQueryArgs() *ListListsRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListListsRequest) SetPageNumber(v int32) *ListListsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListsRequest) SetPageSize(v int32) *ListListsRequest {
	s.PageSize = &v
	return s
}

func (s *ListListsRequest) SetQueryArgs(v *ListListsRequestQueryArgs) *ListListsRequest {
	s.QueryArgs = v
	return s
}

func (s *ListListsRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListListsRequestQueryArgs struct {
	// If true, sorts the results in descending order.
	//
	// example:
	//
	// true
	Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Filters by a partial match on the list description.
	//
	// example:
	//
	// a custom list
	DescriptionLike *string `json:"DescriptionLike,omitempty" xml:"DescriptionLike,omitempty"`
	// Filters by a partial match on the list ID.
	//
	// example:
	//
	// 40000001
	IdLike *string `json:"IdLike,omitempty" xml:"IdLike,omitempty"`
	// Filters by a partial match on a list item.
	//
	// example:
	//
	// 10.1.1.1
	ItemLike *string `json:"ItemLike,omitempty" xml:"ItemLike,omitempty"`
	// Filters by the list kind.
	//
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// Filters by a partial match on both the list name and list items.
	//
	// example:
	//
	// 10.1.1.1
	NameItemLike *string `json:"NameItemLike,omitempty" xml:"NameItemLike,omitempty"`
	// Filters by a partial match on the list name.
	//
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// Specifies the sort field.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListListsRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListListsRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListListsRequestQueryArgs) GetDesc() *bool {
	return s.Desc
}

func (s *ListListsRequestQueryArgs) GetDescriptionLike() *string {
	return s.DescriptionLike
}

func (s *ListListsRequestQueryArgs) GetIdLike() *string {
	return s.IdLike
}

func (s *ListListsRequestQueryArgs) GetItemLike() *string {
	return s.ItemLike
}

func (s *ListListsRequestQueryArgs) GetKind() *string {
	return s.Kind
}

func (s *ListListsRequestQueryArgs) GetNameItemLike() *string {
	return s.NameItemLike
}

func (s *ListListsRequestQueryArgs) GetNameLike() *string {
	return s.NameLike
}

func (s *ListListsRequestQueryArgs) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListListsRequestQueryArgs) SetDesc(v bool) *ListListsRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetDescriptionLike(v string) *ListListsRequestQueryArgs {
	s.DescriptionLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetIdLike(v string) *ListListsRequestQueryArgs {
	s.IdLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetItemLike(v string) *ListListsRequestQueryArgs {
	s.ItemLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetKind(v string) *ListListsRequestQueryArgs {
	s.Kind = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetNameItemLike(v string) *ListListsRequestQueryArgs {
	s.NameItemLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetNameLike(v string) *ListListsRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListListsRequestQueryArgs) SetOrderBy(v string) *ListListsRequestQueryArgs {
	s.OrderBy = &v
	return s
}

func (s *ListListsRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
