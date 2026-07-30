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
	// The ID of the AI gateway instance. The target instance must exist, belong to the current account, and be of the AI gateway type.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The fuzzy match keyword for the model provider identifier or model name. If left empty, all model cards under the current gateway are queried.
	//
	// example:
	//
	// qwen
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 500.
	//
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
