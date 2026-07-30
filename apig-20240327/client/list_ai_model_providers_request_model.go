// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListAiModelProvidersRequest
	GetGatewayId() *string
	SetPageNumber(v int32) *ListAiModelProvidersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAiModelProvidersRequest
	GetPageSize() *int32
	SetProvider(v string) *ListAiModelProvidersRequest
	GetProvider() *string
}

type ListAiModelProvidersRequest struct {
	// The ID of the AI gateway instance. The target instance must exist, belong to the current account, and be of the AI gateway type.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-8c13d2b4f8a1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
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
	// The fuzzy match condition for the model provider identifier. If left empty, all model providers under the current gateway are queried.
	//
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListAiModelProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListAiModelProvidersRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelProvidersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAiModelProvidersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiModelProvidersRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListAiModelProvidersRequest) SetGatewayId(v string) *ListAiModelProvidersRequest {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelProvidersRequest) SetPageNumber(v int32) *ListAiModelProvidersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAiModelProvidersRequest) SetPageSize(v int32) *ListAiModelProvidersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAiModelProvidersRequest) SetProvider(v string) *ListAiModelProvidersRequest {
	s.Provider = &v
	return s
}

func (s *ListAiModelProvidersRequest) Validate() error {
	return dara.Validate(s)
}
