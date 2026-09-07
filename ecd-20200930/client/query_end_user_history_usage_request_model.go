// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEndUserHistoryUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryEndUserHistoryUsageRequest
	GetEndDate() *string
	SetIsAdUser(v bool) *QueryEndUserHistoryUsageRequest
	GetIsAdUser() *bool
	SetPageNum(v int32) *QueryEndUserHistoryUsageRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryEndUserHistoryUsageRequest
	GetPageSize() *int32
	SetStartDate(v string) *QueryEndUserHistoryUsageRequest
	GetStartDate() *string
}

type QueryEndUserHistoryUsageRequest struct {
	// The end date of the query. Format: yyyy-MM-dd. The date cannot be later than yesterday or earlier than StartDate. Default value: yesterday.
	//
	// example:
	//
	// 2024-01-15
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to query Active Directory (AD) domain users. If this parameter is set to true, AD domain users are queried. If this parameter is set to false or not specified, convenience account users are queried.
	IsAdUser *bool `json:"IsAdUser,omitempty" xml:"IsAdUser,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start date of the query. Format: yyyy-MM-dd. The date cannot be earlier than 32 days ago. Default value: yesterday.
	//
	// example:
	//
	// 2024-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryEndUserHistoryUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEndUserHistoryUsageRequest) GoString() string {
	return s.String()
}

func (s *QueryEndUserHistoryUsageRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryEndUserHistoryUsageRequest) GetIsAdUser() *bool {
	return s.IsAdUser
}

func (s *QueryEndUserHistoryUsageRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryEndUserHistoryUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryEndUserHistoryUsageRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryEndUserHistoryUsageRequest) SetEndDate(v string) *QueryEndUserHistoryUsageRequest {
	s.EndDate = &v
	return s
}

func (s *QueryEndUserHistoryUsageRequest) SetIsAdUser(v bool) *QueryEndUserHistoryUsageRequest {
	s.IsAdUser = &v
	return s
}

func (s *QueryEndUserHistoryUsageRequest) SetPageNum(v int32) *QueryEndUserHistoryUsageRequest {
	s.PageNum = &v
	return s
}

func (s *QueryEndUserHistoryUsageRequest) SetPageSize(v int32) *QueryEndUserHistoryUsageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEndUserHistoryUsageRequest) SetStartDate(v string) *QueryEndUserHistoryUsageRequest {
	s.StartDate = &v
	return s
}

func (s *QueryEndUserHistoryUsageRequest) Validate() error {
	return dara.Validate(s)
}
