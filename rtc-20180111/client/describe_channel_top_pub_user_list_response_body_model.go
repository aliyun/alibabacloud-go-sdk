// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelTopPubUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeChannelTopPubUserListResponseBody
	GetRequestId() *string
	SetTopPubUserDetailList(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) *DescribeChannelTopPubUserListResponseBody
	GetTopPubUserDetailList() []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList
}

type DescribeChannelTopPubUserListResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopPubUserDetailList []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList `json:"TopPubUserDetailList,omitempty" xml:"TopPubUserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeChannelTopPubUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelTopPubUserListResponseBody) GetTopPubUserDetailList() []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	return s.TopPubUserDetailList
}

func (s *DescribeChannelTopPubUserListResponseBody) SetRequestId(v string) *DescribeChannelTopPubUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBody) SetTopPubUserDetailList(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) *DescribeChannelTopPubUserListResponseBody {
	s.TopPubUserDetailList = v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList struct {
	// example:
	//
	// 1615893327
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615893442
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                        `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetLocation() *string {
	return s.Location
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetOnlineDuration() *int64 {
	return s.OnlineDuration
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetOnlinePeriods() []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	return s.OnlinePeriods
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GetUserId() *string {
	return s.UserId
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetCreatedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetLocation(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlineDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlinePeriods(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetUserId(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) Validate() error {
	return dara.Validate(s)
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1615893327
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1615893442
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) GetJoinTs() *int64 {
	return s.JoinTs
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) GetLeaveTs() *int64 {
	return s.LeaveTs
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) Validate() error {
	return dara.Validate(s)
}
