// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelCardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListAiModelCardsRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListAiModelCardsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListAiModelCardsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAiModelCardsRequest
	GetPageSize() *int32
}

type ListAiModelCardsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// qwen
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAiModelCardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelCardsRequest) GoString() string {
	return s.String()
}

func (s *ListAiModelCardsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelCardsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAiModelCardsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAiModelCardsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiModelCardsRequest) SetGatewayId(v string) *ListAiModelCardsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelCardsRequest) SetKeyword(v string) *ListAiModelCardsRequest {
	s.Keyword = &v
	return s
}

func (s *ListAiModelCardsRequest) SetPageNumber(v int32) *ListAiModelCardsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAiModelCardsRequest) SetPageSize(v int32) *ListAiModelCardsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAiModelCardsRequest) Validate() error {
	return dara.Validate(s)
}
