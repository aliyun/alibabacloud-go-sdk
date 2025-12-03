// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputRegionId(v string) *ListAccessGroupsRequest
	GetInputRegionId() *string
	SetLimit(v int32) *ListAccessGroupsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListAccessGroupsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListAccessGroupsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListAccessGroupsRequest
	GetOrderType() *string
	SetStartOffset(v int32) *ListAccessGroupsRequest
	GetStartOffset() *int32
}

type ListAccessGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 100
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CreateTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 10
	StartOffset *int32 `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListAccessGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListAccessGroupsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAccessGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessGroupsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListAccessGroupsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAccessGroupsRequest) GetStartOffset() *int32 {
	return s.StartOffset
}

func (s *ListAccessGroupsRequest) SetInputRegionId(v string) *ListAccessGroupsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListAccessGroupsRequest) SetLimit(v int32) *ListAccessGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListAccessGroupsRequest) SetNextToken(v string) *ListAccessGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessGroupsRequest) SetOrderBy(v string) *ListAccessGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAccessGroupsRequest) SetOrderType(v string) *ListAccessGroupsRequest {
	s.OrderType = &v
	return s
}

func (s *ListAccessGroupsRequest) SetStartOffset(v int32) *ListAccessGroupsRequest {
	s.StartOffset = &v
	return s
}

func (s *ListAccessGroupsRequest) Validate() error {
	return dara.Validate(s)
}
