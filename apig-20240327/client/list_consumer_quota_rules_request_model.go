// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerQuotaRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListConsumerQuotaRulesRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListConsumerQuotaRulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListConsumerQuotaRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerQuotaRulesRequest
	GetPageSize() *int32
}

type ListConsumerQuotaRulesRequest struct {
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

func (s ListConsumerQuotaRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerQuotaRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerQuotaRulesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListConsumerQuotaRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConsumerQuotaRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerQuotaRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerQuotaRulesRequest) SetGatewayId(v string) *ListConsumerQuotaRulesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListConsumerQuotaRulesRequest) SetKeyword(v string) *ListConsumerQuotaRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListConsumerQuotaRulesRequest) SetPageNumber(v int32) *ListConsumerQuotaRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerQuotaRulesRequest) SetPageSize(v int32) *ListConsumerQuotaRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerQuotaRulesRequest) Validate() error {
	return dara.Validate(s)
}
