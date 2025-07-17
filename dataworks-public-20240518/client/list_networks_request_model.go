// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListNetworksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworksRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListNetworksRequest
	GetResourceGroupId() *string
	SetSortBy(v string) *ListNetworksRequest
	GetSortBy() *string
}

type ListNetworksRequest struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The fields used for sorting. Fields such as TriggerTime and StartedTime are supported. The value of this parameter is in the Sort field + Sort by (Desc/Asc) format. By default, results are sorted in ascending order. Valid values:
	//
	// 	- Id (Desc/Asc): the network ID
	//
	// 	- Status (Desc/Asc): the network status
	//
	// 	- CreateUser (Desc/Asc): the user who created the network
	//
	// 	- CreateTime (Desc/Asc): the time when the network was created
	//
	// Default value: CreateTime Asc.
	//
	// example:
	//
	// CreateTime Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListNetworksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworksRequest) GoString() string {
	return s.String()
}

func (s *ListNetworksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNetworksRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNetworksRequest) SetPageNumber(v int32) *ListNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworksRequest) SetPageSize(v int32) *ListNetworksRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworksRequest) SetResourceGroupId(v string) *ListNetworksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetworksRequest) SetSortBy(v string) *ListNetworksRequest {
	s.SortBy = &v
	return s
}

func (s *ListNetworksRequest) Validate() error {
	return dara.Validate(s)
}
