// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListEdgeContainerAppVersionsRequest
	GetAppId() *string
	SetOrderKey(v string) *ListEdgeContainerAppVersionsRequest
	GetOrderKey() *string
	SetOrderType(v string) *ListEdgeContainerAppVersionsRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListEdgeContainerAppVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerAppVersionsRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListEdgeContainerAppVersionsRequest
	GetSearchKey() *string
	SetSearchType(v string) *ListEdgeContainerAppVersionsRequest
	GetSearchType() *string
}

type ListEdgeContainerAppVersionsRequest struct {
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The sorting field. Valid values:
	//
	// 	- Name: the version name.
	//
	// 	- CreateTime: the time when the version was created.
	//
	// 	- UpdateTime: the time when the version was last modified.
	//
	// example:
	//
	// Name
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- ASC: in ascending order.
	//
	// 	- DESC: in descending order.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used for the search.
	//
	// example:
	//
	// ver-100568263967926****
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The parameter used for fuzzy search. Valid values: VersionId and Name.
	//
	// example:
	//
	// VersionId
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
}

func (s ListEdgeContainerAppVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListEdgeContainerAppVersionsRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *ListEdgeContainerAppVersionsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListEdgeContainerAppVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerAppVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerAppVersionsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListEdgeContainerAppVersionsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListEdgeContainerAppVersionsRequest) SetAppId(v string) *ListEdgeContainerAppVersionsRequest {
	s.AppId = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetOrderKey(v string) *ListEdgeContainerAppVersionsRequest {
	s.OrderKey = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetOrderType(v string) *ListEdgeContainerAppVersionsRequest {
	s.OrderType = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetPageNumber(v int32) *ListEdgeContainerAppVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetPageSize(v int32) *ListEdgeContainerAppVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetSearchKey(v string) *ListEdgeContainerAppVersionsRequest {
	s.SearchKey = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) SetSearchType(v string) *ListEdgeContainerAppVersionsRequest {
	s.SearchType = &v
	return s
}

func (s *ListEdgeContainerAppVersionsRequest) Validate() error {
	return dara.Validate(s)
}
