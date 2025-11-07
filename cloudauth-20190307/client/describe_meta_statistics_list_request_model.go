// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *DescribeMetaStatisticsListRequest
	GetApi() *string
	SetEndDate(v int64) *DescribeMetaStatisticsListRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeMetaStatisticsListRequest
	GetStartDate() *int64
}

type DescribeMetaStatisticsListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MOBILE_ONLINE_LENGTH
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1750694399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1760112000000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeMetaStatisticsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsListRequest) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaStatisticsListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeMetaStatisticsListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeMetaStatisticsListRequest) SetApi(v string) *DescribeMetaStatisticsListRequest {
	s.Api = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) SetEndDate(v int64) *DescribeMetaStatisticsListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) SetStartDate(v int64) *DescribeMetaStatisticsListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMetaStatisticsListRequest) Validate() error {
	return dara.Validate(s)
}
