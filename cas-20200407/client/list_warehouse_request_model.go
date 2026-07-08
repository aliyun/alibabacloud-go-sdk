// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWarehouseRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWarehouseRequest
	GetNextToken() *string
	SetWarehouseInstanceIds(v []*string) *ListWarehouseRequest
	GetWarehouseInstanceIds() []*string
	SetWarehouseTypes(v []*string) *ListWarehouseRequest
	GetWarehouseTypes() []*string
}

type ListWarehouseRequest struct {
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
	WarehouseInstanceIds []*string `json:"WarehouseInstanceIds,omitempty" xml:"WarehouseInstanceIds,omitempty" type:"Repeated"`
	// A list of warehouse types.
	//
	// example:
	//
	// pcaCaCert
	WarehouseTypes []*string `json:"WarehouseTypes,omitempty" xml:"WarehouseTypes,omitempty" type:"Repeated"`
}

func (s ListWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ListWarehouseRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWarehouseRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWarehouseRequest) GetWarehouseInstanceIds() []*string {
	return s.WarehouseInstanceIds
}

func (s *ListWarehouseRequest) GetWarehouseTypes() []*string {
	return s.WarehouseTypes
}

func (s *ListWarehouseRequest) SetMaxResults(v int32) *ListWarehouseRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWarehouseRequest) SetNextToken(v string) *ListWarehouseRequest {
	s.NextToken = &v
	return s
}

func (s *ListWarehouseRequest) SetWarehouseInstanceIds(v []*string) *ListWarehouseRequest {
	s.WarehouseInstanceIds = v
	return s
}

func (s *ListWarehouseRequest) SetWarehouseTypes(v []*string) *ListWarehouseRequest {
	s.WarehouseTypes = v
	return s
}

func (s *ListWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
