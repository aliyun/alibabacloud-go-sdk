// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterFailedRequests(v bool) *GetGatewayQuotaRuleSubjectUsageRequest
	GetFilterFailedRequests() *bool
	SetPageNumber(v int32) *GetGatewayQuotaRuleSubjectUsageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetGatewayQuotaRuleSubjectUsageRequest
	GetPageSize() *int32
}

type GetGatewayQuotaRuleSubjectUsageRequest struct {
	// Specifies whether to filter out zero values.
	//
	// example:
	//
	// true
	FilterFailedRequests *bool `json:"filterFailedRequests,omitempty" xml:"filterFailedRequests,omitempty"`
	// The page number of the detailed consumption (request) records of the subject within the cycle.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of detailed consumption (request) records per page for the subject within the cycle. Maximum value: 10.
	//
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

func (s *GetGatewayQuotaRuleSubjectUsageRequest) GetFilterFailedRequests() *bool {
	return s.FilterFailedRequests
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetGatewayQuotaRuleSubjectUsageRequest) SetFilterFailedRequests(v bool) *GetGatewayQuotaRuleSubjectUsageRequest {
	s.FilterFailedRequests = &v
	return s
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
