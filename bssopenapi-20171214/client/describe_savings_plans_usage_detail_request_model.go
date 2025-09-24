// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailRequest
	GetBillOwnerId() *int64
	SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailRequest
	GetEndPeriod() *string
	SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailRequest
	GetMaxResults() *int32
	SetPeriodType(v string) *DescribeSavingsPlansUsageDetailRequest
	GetPeriodType() *string
	SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailRequest
	GetStartPeriod() *string
	SetToken(v string) *DescribeSavingsPlansUsageDetailRequest
	GetToken() *string
}

type DescribeSavingsPlansUsageDetailRequest struct {
	// The ID of the account for which you want to query usage details. If you do not set this parameter, the data of the current Alibaba Cloud account and its RAM users is queried. To query the data of a RAM user, specify the ID of the RAM user.
	//
	// example:
	//
	// 123745698925000
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The end of the time range to query. The end is excluded from the time range. If you do not set this parameter, the end time is the current time. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-02 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The time granularity at which usage details are queried. Valid values: MONTH, DAY, and HOUR.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The beginning of the time range to query. The beginning is included in the time range. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The token that is used to retrieve the next page of results. You do not need to set this parameter if you query usage details within a specific time range for the first time. The response returns a token that you can use to query usage details that are displayed on the next page. If a null value is returned for the NextToken parameter, no more coverage details can be queried.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageDetailRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetBillOwnerId(v int64) *DescribeSavingsPlansUsageDetailRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetMaxResults(v int32) *DescribeSavingsPlansUsageDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetPeriodType(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) SetToken(v string) *DescribeSavingsPlansUsageDetailRequest {
	s.Token = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailRequest) Validate() error {
	return dara.Validate(s)
}
