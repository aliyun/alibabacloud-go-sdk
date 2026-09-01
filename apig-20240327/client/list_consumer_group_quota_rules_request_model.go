// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupQuotaRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListConsumerGroupQuotaRulesRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListConsumerGroupQuotaRulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListConsumerGroupQuotaRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerGroupQuotaRulesRequest
	GetPageSize() *int32
}

type ListConsumerGroupQuotaRulesRequest struct {
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// daily
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

func (s ListConsumerGroupQuotaRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupQuotaRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupQuotaRulesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListConsumerGroupQuotaRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConsumerGroupQuotaRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupQuotaRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupQuotaRulesRequest) SetGatewayId(v string) *ListConsumerGroupQuotaRulesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesRequest) SetKeyword(v string) *ListConsumerGroupQuotaRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesRequest) SetPageNumber(v int32) *ListConsumerGroupQuotaRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesRequest) SetPageSize(v int32) *ListConsumerGroupQuotaRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesRequest) Validate() error {
	return dara.Validate(s)
}
