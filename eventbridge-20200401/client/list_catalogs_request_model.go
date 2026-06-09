// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListCatalogsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListCatalogsRequest
	GetNextToken() *string
}

type ListCatalogsRequest struct {
	// 每页返回的最大数据条数。不传时默认 10，最大 100
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 分页查询的起始Token。首次查询不传或传 "0"；后续翻页使用上一次响应中返回的 NextToken 值
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListCatalogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCatalogsRequest) SetLimit(v int32) *ListCatalogsRequest {
	s.Limit = &v
	return s
}

func (s *ListCatalogsRequest) SetNextToken(v string) *ListCatalogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
