// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderKey(v string) *ListEdgeContainerAppsRequest
	GetOrderKey() *string
	SetOrderType(v string) *ListEdgeContainerAppsRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListEdgeContainerAppsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppsRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListEdgeContainerAppsRequest
	GetSearchKey() *string
	SetSearchType(v string) *ListEdgeContainerAppsRequest
	GetSearchType() *string
}

type ListEdgeContainerAppsRequest struct {
	// The sorting field. This parameter is left empty by default. Valid values:
	//
	// 	- Name: the version name.
	//
	// 	- CreateTime: the time when the version was created.
	//
	// 	- UpdateTime: the time when the version was last modified.
	//
	// example:
	//
	// CreateTime
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The order in which you want to sort the query results. This parameter is left empty by default. Valid values:
	//
	// 	- ASC: in ascending order.
	//
	// 	- DESC: in descending order.
	//
	// example:
	//
	// Asc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: **1**. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// ver-1005682639679266816
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The search criterion based on which you want to perform fuzzy search. Valid values:
	//
	// 	- Appid: the application ID.
	//
	// 	- Name: the application name.
	//
	// example:
	//
	// Appid
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
}

func (s ListEdgeContainerAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppsRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *ListEdgeContainerAppsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListEdgeContainerAppsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListEdgeContainerAppsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListEdgeContainerAppsRequest) SetOrderKey(v string) *ListEdgeContainerAppsRequest {
	s.OrderKey = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) SetOrderType(v string) *ListEdgeContainerAppsRequest {
	s.OrderType = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) SetPageNumber(v int32) *ListEdgeContainerAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) SetPageSize(v int32) *ListEdgeContainerAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) SetSearchKey(v string) *ListEdgeContainerAppsRequest {
	s.SearchKey = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) SetSearchType(v string) *ListEdgeContainerAppsRequest {
	s.SearchType = &v
	return s
}

func (s *ListEdgeContainerAppsRequest) Validate() error {
	return dara.Validate(s)
}
