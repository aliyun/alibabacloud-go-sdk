// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *ListAccessRulesRequest
	GetAccessGroupId() *string
	SetInputRegionId(v string) *ListAccessRulesRequest
	GetInputRegionId() *string
	SetLimit(v int32) *ListAccessRulesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListAccessRulesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListAccessRulesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListAccessRulesRequest
	GetOrderType() *string
	SetStartOffset(v int32) *ListAccessRulesRequest
	GetStartOffset() *int32
}

type ListAccessRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
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
	// Priority
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

func (s ListAccessRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAccessRulesRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ListAccessRulesRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListAccessRulesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAccessRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessRulesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListAccessRulesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAccessRulesRequest) GetStartOffset() *int32 {
	return s.StartOffset
}

func (s *ListAccessRulesRequest) SetAccessGroupId(v string) *ListAccessRulesRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessRulesRequest) SetInputRegionId(v string) *ListAccessRulesRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListAccessRulesRequest) SetLimit(v int32) *ListAccessRulesRequest {
	s.Limit = &v
	return s
}

func (s *ListAccessRulesRequest) SetNextToken(v string) *ListAccessRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessRulesRequest) SetOrderBy(v string) *ListAccessRulesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAccessRulesRequest) SetOrderType(v string) *ListAccessRulesRequest {
	s.OrderType = &v
	return s
}

func (s *ListAccessRulesRequest) SetStartOffset(v int32) *ListAccessRulesRequest {
	s.StartOffset = &v
	return s
}

func (s *ListAccessRulesRequest) Validate() error {
	return dara.Validate(s)
}
