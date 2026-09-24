// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthUserConnectDurationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *QueryAuthUserConnectDurationListRequest
	GetComparisonOperator() *string
	SetDataDate(v string) *QueryAuthUserConnectDurationListRequest
	GetDataDate() *string
	SetIsAdUser(v bool) *QueryAuthUserConnectDurationListRequest
	GetIsAdUser() *bool
	SetNextToken(v string) *QueryAuthUserConnectDurationListRequest
	GetNextToken() *string
	SetPageNum(v int32) *QueryAuthUserConnectDurationListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryAuthUserConnectDurationListRequest
	GetPageSize() *int32
	SetStatisticType(v string) *QueryAuthUserConnectDurationListRequest
	GetStatisticType() *string
	SetThreshold(v int32) *QueryAuthUserConnectDurationListRequest
	GetThreshold() *int32
	SetUserName(v string) *QueryAuthUserConnectDurationListRequest
	GetUserName() *string
	SetWithDetail(v bool) *QueryAuthUserConnectDurationListRequest
	GetWithDetail() *bool
}

type QueryAuthUserConnectDurationListRequest struct {
	// The comparison operator for connection duration. This parameter is used together with Threshold to filter users by a threshold. Valid values:
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// example:
	//
	// GreaterThanThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The date for the statistics, in the yyyy-MM-dd format. If this parameter is left empty, statistics from the previous day are returned by default.
	//
	// example:
	//
	// 2026-06-24
	DataDate *string `json:"DataDate,omitempty" xml:"DataDate,omitempty"`
	// Specifies whether to query only Active Directory (AD) users. This parameter is required when UserName or WithDetail is specified.
	IsAdUser *bool `json:"IsAdUser,omitempty" xml:"IsAdUser,omitempty"`
	// The paging token. This parameter is used only when statistics are collected by individual session details (StatisticType=SingleSession). You do not need to specify this parameter for the first request. For subsequent requests, set this parameter to the NextToken value returned in the previous response to retrieve the next page.
	//
	// example:
	//
	// d129c6c0e8c04c8a9f0e2b7c1a3f5e6d
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number, starting from 1. Default value: 1. This parameter takes effect only when statistics are collected by daily cumulative duration (StatisticType=Daily).
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The statistics type. Valid values:
	//
	// - Daily: collects statistics by daily cumulative connection duration. This is the default value.
	//
	// - SingleSession: collects statistics by individual session details.
	//
	// example:
	//
	// Daily
	StatisticType *string `json:"StatisticType,omitempty" xml:"StatisticType,omitempty"`
	// The connection duration threshold, in seconds. This parameter must be used together with ComparisonOperator.
	//
	// example:
	//
	// 3600
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The name of the end user. Fuzzy match is supported. When you use this parameter, you must also specify IsAdUser.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Specifies whether to backfill user details such as display name and nickname. This parameter supports both AD users and convenience users. When you use this parameter, you must also specify IsAdUser.
	WithDetail *bool `json:"WithDetail,omitempty" xml:"WithDetail,omitempty"`
}

func (s QueryAuthUserConnectDurationListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthUserConnectDurationListRequest) GoString() string {
	return s.String()
}

func (s *QueryAuthUserConnectDurationListRequest) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *QueryAuthUserConnectDurationListRequest) GetDataDate() *string {
	return s.DataDate
}

func (s *QueryAuthUserConnectDurationListRequest) GetIsAdUser() *bool {
	return s.IsAdUser
}

func (s *QueryAuthUserConnectDurationListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAuthUserConnectDurationListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAuthUserConnectDurationListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAuthUserConnectDurationListRequest) GetStatisticType() *string {
	return s.StatisticType
}

func (s *QueryAuthUserConnectDurationListRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *QueryAuthUserConnectDurationListRequest) GetUserName() *string {
	return s.UserName
}

func (s *QueryAuthUserConnectDurationListRequest) GetWithDetail() *bool {
	return s.WithDetail
}

func (s *QueryAuthUserConnectDurationListRequest) SetComparisonOperator(v string) *QueryAuthUserConnectDurationListRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetDataDate(v string) *QueryAuthUserConnectDurationListRequest {
	s.DataDate = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetIsAdUser(v bool) *QueryAuthUserConnectDurationListRequest {
	s.IsAdUser = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetNextToken(v string) *QueryAuthUserConnectDurationListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetPageNum(v int32) *QueryAuthUserConnectDurationListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetPageSize(v int32) *QueryAuthUserConnectDurationListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetStatisticType(v string) *QueryAuthUserConnectDurationListRequest {
	s.StatisticType = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetThreshold(v int32) *QueryAuthUserConnectDurationListRequest {
	s.Threshold = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetUserName(v string) *QueryAuthUserConnectDurationListRequest {
	s.UserName = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) SetWithDetail(v bool) *QueryAuthUserConnectDurationListRequest {
	s.WithDetail = &v
	return s
}

func (s *QueryAuthUserConnectDurationListRequest) Validate() error {
	return dara.Validate(s)
}
