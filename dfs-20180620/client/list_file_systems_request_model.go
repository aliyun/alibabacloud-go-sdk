// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputRegionId(v string) *ListFileSystemsRequest
	GetInputRegionId() *string
	SetLimit(v int32) *ListFileSystemsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFileSystemsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListFileSystemsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListFileSystemsRequest
	GetOrderType() *string
	SetStartOffset(v int32) *ListFileSystemsRequest
	GetStartOffset() *int32
}

type ListFileSystemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 10
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

func (s ListFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListFileSystemsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListFileSystemsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFileSystemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFileSystemsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListFileSystemsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListFileSystemsRequest) GetStartOffset() *int32 {
	return s.StartOffset
}

func (s *ListFileSystemsRequest) SetInputRegionId(v string) *ListFileSystemsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListFileSystemsRequest) SetLimit(v int32) *ListFileSystemsRequest {
	s.Limit = &v
	return s
}

func (s *ListFileSystemsRequest) SetNextToken(v string) *ListFileSystemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFileSystemsRequest) SetOrderBy(v string) *ListFileSystemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFileSystemsRequest) SetOrderType(v string) *ListFileSystemsRequest {
	s.OrderType = &v
	return s
}

func (s *ListFileSystemsRequest) SetStartOffset(v int32) *ListFileSystemsRequest {
	s.StartOffset = &v
	return s
}

func (s *ListFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
