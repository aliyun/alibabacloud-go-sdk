// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDesktopGroupSessionsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeDesktopGroupSessionsRequest
	GetEndUserId() *string
	SetMaxResults(v int32) *DescribeDesktopGroupSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopGroupSessionsRequest
	GetNextToken() *string
	SetOwnType(v int32) *DescribeDesktopGroupSessionsRequest
	GetOwnType() *int32
	SetRegionId(v string) *DescribeDesktopGroupSessionsRequest
	GetRegionId() *string
	SetSessionStatus(v string) *DescribeDesktopGroupSessionsRequest
	GetSessionStatus() *string
	SetStartTime(v string) *DescribeDesktopGroupSessionsRequest
	GetStartTime() *string
}

type DescribeDesktopGroupSessionsRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// "2022-08-31T06:56:45Z"
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// xianqiu
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// "asdfdfsdfsdfds"
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the session.
	//
	// Valid values:
	//
	// 	- 0: single-session
	//
	// 	- 1: multi-session
	//
	// example:
	//
	// 1
	OwnType *int32 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the session.
	//
	// Valid values:
	//
	// 	- Connected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Disconnected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Connected
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// "2022-08-31T06:56:45Z"
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDesktopGroupSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupSessionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDesktopGroupSessionsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopGroupSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopGroupSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopGroupSessionsRequest) GetOwnType() *int32 {
	return s.OwnType
}

func (s *DescribeDesktopGroupSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopGroupSessionsRequest) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeDesktopGroupSessionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDesktopGroupSessionsRequest) SetEndTime(v string) *DescribeDesktopGroupSessionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetEndUserId(v string) *DescribeDesktopGroupSessionsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetMaxResults(v int32) *DescribeDesktopGroupSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetNextToken(v string) *DescribeDesktopGroupSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetOwnType(v int32) *DescribeDesktopGroupSessionsRequest {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetRegionId(v string) *DescribeDesktopGroupSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetSessionStatus(v string) *DescribeDesktopGroupSessionsRequest {
	s.SessionStatus = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) SetStartTime(v string) *DescribeDesktopGroupSessionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopGroupSessionsRequest) Validate() error {
	return dara.Validate(s)
}
