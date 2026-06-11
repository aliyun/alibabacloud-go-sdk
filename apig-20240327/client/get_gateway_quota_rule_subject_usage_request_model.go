// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetGatewayQuotaRuleSubjectUsageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetGatewayQuotaRuleSubjectUsageRequest
	GetPageSize() *int32
}

type GetGatewayQuotaRuleSubjectUsageRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) SetPageNumber(v int32) *GetGatewayQuotaRuleSubjectUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) SetPageSize(v int32) *GetGatewayQuotaRuleSubjectUsageRequest {
	s.PageSize = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) Validate() error {
	return dara.Validate(s)
}
