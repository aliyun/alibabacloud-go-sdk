// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerAuthorizationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameLike(v string) *ListConsumerAuthorizationRulesRequest
	GetApiNameLike() *string
	SetPageNumber(v int32) *ListConsumerAuthorizationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerAuthorizationRulesRequest
	GetPageSize() *int32
}

type ListConsumerAuthorizationRulesRequest struct {
	// example:
	//
	// qwen3
	ApiNameLike *string `json:"apiNameLike,omitempty" xml:"apiNameLike,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConsumerAuthorizationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerAuthorizationRulesRequest) GetApiNameLike() *string {
	return s.ApiNameLike
}

func (s *ListConsumerAuthorizationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerAuthorizationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerAuthorizationRulesRequest) SetApiNameLike(v string) *ListConsumerAuthorizationRulesRequest {
	s.ApiNameLike = &v
	return s
}

func (s *ListConsumerAuthorizationRulesRequest) SetPageNumber(v int32) *ListConsumerAuthorizationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerAuthorizationRulesRequest) SetPageSize(v int32) *ListConsumerAuthorizationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerAuthorizationRulesRequest) Validate() error {
	return dara.Validate(s)
}
