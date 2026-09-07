// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryHistoryActiveUserStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryHistoryActiveUserStatisticResponseBody
	GetTotalCount() *int32
	SetUserStatisticList(v []*QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) *QueryHistoryActiveUserStatisticResponseBody
	GetUserStatisticList() []*QueryHistoryActiveUserStatisticResponseBodyUserStatisticList
}

type QueryHistoryActiveUserStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5CC5E450-FC43-4F5B-B540-9964BD313427
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of daily active user statistics.
	UserStatisticList []*QueryHistoryActiveUserStatisticResponseBodyUserStatisticList `json:"UserStatisticList,omitempty" xml:"UserStatisticList,omitempty" type:"Repeated"`
}

func (s QueryHistoryActiveUserStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoryActiveUserStatisticResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryHistoryActiveUserStatisticResponseBody) GetUserStatisticList() []*QueryHistoryActiveUserStatisticResponseBodyUserStatisticList {
	return s.UserStatisticList
}

func (s *QueryHistoryActiveUserStatisticResponseBody) SetRequestId(v string) *QueryHistoryActiveUserStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBody) SetTotalCount(v int32) *QueryHistoryActiveUserStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBody) SetUserStatisticList(v []*QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) *QueryHistoryActiveUserStatisticResponseBody {
	s.UserStatisticList = v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBody) Validate() error {
	if s.UserStatisticList != nil {
		for _, item := range s.UserStatisticList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryActiveUserStatisticResponseBodyUserStatisticList struct {
	// The number of deduplicated active users on the day.
	//
	// example:
	//
	// 150
	ActiveUserCount *int32 `json:"ActiveUserCount,omitempty" xml:"ActiveUserCount,omitempty"`
	// The date in the standard yyyy-MM-dd format, in the UTC+8 time zone.
	//
	// example:
	//
	// 2024-12-01
	FormatDate *string `json:"FormatDate,omitempty" xml:"FormatDate,omitempty"`
	// The timestamp of the date, in milliseconds.
	//
	// example:
	//
	// 1735689600000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) GetActiveUserCount() *int32 {
	return s.ActiveUserCount
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) GetFormatDate() *string {
	return s.FormatDate
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) SetActiveUserCount(v int32) *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList {
	s.ActiveUserCount = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) SetFormatDate(v string) *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList {
	s.FormatDate = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) SetTimeStamp(v int64) *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList {
	s.TimeStamp = &v
	return s
}

func (s *QueryHistoryActiveUserStatisticResponseBodyUserStatisticList) Validate() error {
	return dara.Validate(s)
}
