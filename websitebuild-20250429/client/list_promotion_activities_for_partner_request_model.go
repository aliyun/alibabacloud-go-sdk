// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionActivitiesForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *ListPromotionActivitiesForPartnerRequest
	GetChannel() *string
	SetEmployeeCode(v string) *ListPromotionActivitiesForPartnerRequest
	GetEmployeeCode() *string
	SetMaxResults(v int32) *ListPromotionActivitiesForPartnerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPromotionActivitiesForPartnerRequest
	GetNextToken() *string
}

type ListPromotionActivitiesForPartnerRequest struct {
	// The channel.
	//
	// example:
	//
	// WECHAT
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The employee code.
	//
	// example:
	//
	// 1234
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPromotionActivitiesForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionActivitiesForPartnerRequest) GoString() string {
	return s.String()
}

func (s *ListPromotionActivitiesForPartnerRequest) GetChannel() *string {
	return s.Channel
}

func (s *ListPromotionActivitiesForPartnerRequest) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *ListPromotionActivitiesForPartnerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromotionActivitiesForPartnerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromotionActivitiesForPartnerRequest) SetChannel(v string) *ListPromotionActivitiesForPartnerRequest {
	s.Channel = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerRequest) SetEmployeeCode(v string) *ListPromotionActivitiesForPartnerRequest {
	s.EmployeeCode = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerRequest) SetMaxResults(v int32) *ListPromotionActivitiesForPartnerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerRequest) SetNextToken(v string) *ListPromotionActivitiesForPartnerRequest {
	s.NextToken = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
