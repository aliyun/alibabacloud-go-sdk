// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v int64) *ListRoutesRequest
	GetNetworkId() *int64
	SetPageNumber(v int32) *ListRoutesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRoutesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListRoutesRequest
	GetResourceGroupId() *string
	SetSortBy(v string) *ListRoutesRequest
	GetSortBy() *string
}

type ListRoutesRequest struct {
	// The ID of the network resource.
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the general quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of sort fields. Fields such as scheduled time and start time are supported. The format is "sort field + sort order (Desc/Asc)" (Asc is the default if omitted). Valid values:
	//
	// - Id (Desc/Asc): route ID
	//
	// - DestinationCidr (Desc/Asc): destination CIDR
	//
	// - CreateTime (Desc/Asc): creation time
	//
	// Default value: CreateTime Asc.
	//
	// example:
	//
	// CreateTime Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListRoutesRequest) GetNetworkId() *int64 {
	return s.NetworkId
}

func (s *ListRoutesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRoutesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRoutesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRoutesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRoutesRequest) SetNetworkId(v int64) *ListRoutesRequest {
	s.NetworkId = &v
	return s
}

func (s *ListRoutesRequest) SetPageNumber(v int32) *ListRoutesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoutesRequest) SetPageSize(v int32) *ListRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutesRequest) SetResourceGroupId(v string) *ListRoutesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRoutesRequest) SetSortBy(v string) *ListRoutesRequest {
	s.SortBy = &v
	return s
}

func (s *ListRoutesRequest) Validate() error {
	return dara.Validate(s)
}
