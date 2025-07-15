// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectDurationFrom(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectDurationFrom() *int64
	SetConnectDurationTo(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectDurationTo() *int64
	SetConnectEndTimeFrom(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectEndTimeFrom() *int64
	SetConnectEndTimeTo(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectEndTimeTo() *int64
	SetConnectStartTimeFrom(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectStartTimeFrom() *int64
	SetConnectStartTimeTo(v int64) *DescribeUserConnectionRecordsRequest
	GetConnectStartTimeTo() *int64
	SetDesktopGroupId(v string) *DescribeUserConnectionRecordsRequest
	GetDesktopGroupId() *string
	SetDesktopId(v string) *DescribeUserConnectionRecordsRequest
	GetDesktopId() *string
	SetEndUserId(v string) *DescribeUserConnectionRecordsRequest
	GetEndUserId() *string
	SetEndUserType(v string) *DescribeUserConnectionRecordsRequest
	GetEndUserType() *string
	SetMaxResults(v int32) *DescribeUserConnectionRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserConnectionRecordsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeUserConnectionRecordsRequest
	GetRegionId() *string
}

type DescribeUserConnectionRecordsRequest struct {
	// The duration when the cloud desktop is connected, which is the minimum value for condition filtering.
	//
	// example:
	//
	// 100
	ConnectDurationFrom *int64 `json:"ConnectDurationFrom,omitempty" xml:"ConnectDurationFrom,omitempty"`
	// The duration when the cloud desktop is connected, which is the maximum value for condition filtering.
	//
	// example:
	//
	// 100
	ConnectDurationTo *int64 `json:"ConnectDurationTo,omitempty" xml:"ConnectDurationTo,omitempty"`
	// The time when the cloud desktop stops to be connected, which is the minimum value for condition filtering. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1631266200000
	ConnectEndTimeFrom *int64 `json:"ConnectEndTimeFrom,omitempty" xml:"ConnectEndTimeFrom,omitempty"`
	// The time when the cloud desktop stops to be connected, which is the maximum value for condition filtering. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1631268000000
	ConnectEndTimeTo *int64 `json:"ConnectEndTimeTo,omitempty" xml:"ConnectEndTimeTo,omitempty"`
	// The time when the cloud desktop starts to be connected, which is the minimum value for condition filtering. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1631239200000
	ConnectStartTimeFrom *int64 `json:"ConnectStartTimeFrom,omitempty" xml:"ConnectStartTimeFrom,omitempty"`
	// The time when the cloud desktop starts to be connected, which is the maximum value for condition filtering. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1631241000000
	ConnectStartTimeTo *int64 `json:"ConnectStartTimeTo,omitempty" xml:"ConnectStartTimeTo,omitempty"`
	// The ID of the cloud computer pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-138dsptkrt00u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The ID of the authorized user.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The type of the user account.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience account
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- AD_CONNECTOR: enterprise AD account
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SIMPLE
	EndUserType *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUserConnectionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectionRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectDurationFrom() *int64 {
	return s.ConnectDurationFrom
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectDurationTo() *int64 {
	return s.ConnectDurationTo
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectEndTimeFrom() *int64 {
	return s.ConnectEndTimeFrom
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectEndTimeTo() *int64 {
	return s.ConnectEndTimeTo
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectStartTimeFrom() *int64 {
	return s.ConnectStartTimeFrom
}

func (s *DescribeUserConnectionRecordsRequest) GetConnectStartTimeTo() *int64 {
	return s.ConnectStartTimeTo
}

func (s *DescribeUserConnectionRecordsRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeUserConnectionRecordsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeUserConnectionRecordsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUserConnectionRecordsRequest) GetEndUserType() *string {
	return s.EndUserType
}

func (s *DescribeUserConnectionRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserConnectionRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserConnectionRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectDurationFrom(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectDurationFrom = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectDurationTo(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectDurationTo = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectEndTimeFrom(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectEndTimeFrom = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectEndTimeTo(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectEndTimeTo = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectStartTimeFrom(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectStartTimeFrom = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetConnectStartTimeTo(v int64) *DescribeUserConnectionRecordsRequest {
	s.ConnectStartTimeTo = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetDesktopGroupId(v string) *DescribeUserConnectionRecordsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetDesktopId(v string) *DescribeUserConnectionRecordsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserId(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserType(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserType = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetMaxResults(v int32) *DescribeUserConnectionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetNextToken(v string) *DescribeUserConnectionRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetRegionId(v string) *DescribeUserConnectionRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
