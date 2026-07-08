// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWarehouseShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWarehouseShrinkRequest
	GetNextToken() *string
	SetWarehouseInstanceIdsShrink(v string) *ListWarehouseShrinkRequest
	GetWarehouseInstanceIdsShrink() *string
	SetWarehouseTypesShrink(v string) *ListWarehouseShrinkRequest
	GetWarehouseTypesShrink() *string
}

type ListWarehouseShrinkRequest struct {
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token from a previous response. Use this token to retrieve the next page of results. Omit this parameter for the first request.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A list of warehouse instance IDs.
	//
	// example:
	//
	// cas-wh-uc-gl2bsq
	WarehouseInstanceIdsShrink *string `json:"WarehouseInstanceIds,omitempty" xml:"WarehouseInstanceIds,omitempty"`
	// A list of warehouse types.
	//
	// example:
	//
	// pcaCaCert
	WarehouseTypesShrink *string `json:"WarehouseTypes,omitempty" xml:"WarehouseTypes,omitempty"`
}

func (s ListWarehouseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWarehouseShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWarehouseShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWarehouseShrinkRequest) GetWarehouseInstanceIdsShrink() *string {
	return s.WarehouseInstanceIdsShrink
}

func (s *ListWarehouseShrinkRequest) GetWarehouseTypesShrink() *string {
	return s.WarehouseTypesShrink
}

func (s *ListWarehouseShrinkRequest) SetMaxResults(v int32) *ListWarehouseShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWarehouseShrinkRequest) SetNextToken(v string) *ListWarehouseShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWarehouseShrinkRequest) SetWarehouseInstanceIdsShrink(v string) *ListWarehouseShrinkRequest {
	s.WarehouseInstanceIdsShrink = &v
	return s
}

func (s *ListWarehouseShrinkRequest) SetWarehouseTypesShrink(v string) *ListWarehouseShrinkRequest {
	s.WarehouseTypesShrink = &v
	return s
}

func (s *ListWarehouseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
