// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryActiveUserStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryActiveUserStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryActiveUserStatisticResponseBody
	GetTotalCount() *int32
	SetUserStatisticList(v []*QueryActiveUserStatisticResponseBodyUserStatisticList) *QueryActiveUserStatisticResponseBody
	GetUserStatisticList() []*QueryActiveUserStatisticResponseBodyUserStatisticList
}

type QueryActiveUserStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data points.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of active user statistics.
	UserStatisticList []*QueryActiveUserStatisticResponseBodyUserStatisticList `json:"UserStatisticList,omitempty" xml:"UserStatisticList,omitempty" type:"Repeated"`
}

func (s QueryActiveUserStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryActiveUserStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryActiveUserStatisticResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryActiveUserStatisticResponseBody) GetUserStatisticList() []*QueryActiveUserStatisticResponseBodyUserStatisticList {
	return s.UserStatisticList
}

func (s *QueryActiveUserStatisticResponseBody) SetRequestId(v string) *QueryActiveUserStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryActiveUserStatisticResponseBody) SetTotalCount(v int32) *QueryActiveUserStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryActiveUserStatisticResponseBody) SetUserStatisticList(v []*QueryActiveUserStatisticResponseBodyUserStatisticList) *QueryActiveUserStatisticResponseBody {
	s.UserStatisticList = v
	return s
}

func (s *QueryActiveUserStatisticResponseBody) Validate() error {
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

type QueryActiveUserStatisticResponseBodyUserStatisticList struct {
	// The number of deduplicated active users within the statistical period.
	//
	// example:
	//
	// 56
	ActiveUserCount *int32 `json:"ActiveUserCount,omitempty" xml:"ActiveUserCount,omitempty"`
	// The formatted date string in UTC+8, in the format of yyyy-MM-dd.
	//
	// example:
	//
	// 2020-11-30
	FormatDate *string `json:"FormatDate,omitempty" xml:"FormatDate,omitempty"`
	// The epoch timestamp in milliseconds corresponding to the data point.
	//
	// example:
	//
	// 1606723951000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s QueryActiveUserStatisticResponseBodyUserStatisticList) String() string {
	return dara.Prettify(s)
}

func (s QueryActiveUserStatisticResponseBodyUserStatisticList) GoString() string {
	return s.String()
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) GetActiveUserCount() *int32 {
	return s.ActiveUserCount
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) GetFormatDate() *string {
	return s.FormatDate
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) SetActiveUserCount(v int32) *QueryActiveUserStatisticResponseBodyUserStatisticList {
	s.ActiveUserCount = &v
	return s
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) SetFormatDate(v string) *QueryActiveUserStatisticResponseBodyUserStatisticList {
	s.FormatDate = &v
	return s
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) SetTimeStamp(v int64) *QueryActiveUserStatisticResponseBodyUserStatisticList {
	s.TimeStamp = &v
	return s
}

func (s *QueryActiveUserStatisticResponseBodyUserStatisticList) Validate() error {
	return dara.Validate(s)
}
