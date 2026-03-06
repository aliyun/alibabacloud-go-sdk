// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppCommoditySpecificationsV2ForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAppCommoditySpecificationsV2ForPartnerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppCommoditySpecificationsV2ForPartnerRequest
	GetNextToken() *string
}

type ListAppCommoditySpecificationsV2ForPartnerRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAppCommoditySpecificationsV2ForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerRequest) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppCommoditySpecificationsV2ForPartnerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppCommoditySpecificationsV2ForPartnerRequest) SetMaxResults(v int32) *ListAppCommoditySpecificationsV2ForPartnerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerRequest) SetNextToken(v string) *ListAppCommoditySpecificationsV2ForPartnerRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
