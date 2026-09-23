// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentMonitorMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCustomAgentMonitorMetricsResponseBodyData) *DescribeCustomAgentMonitorMetricsResponseBody
	GetData() *DescribeCustomAgentMonitorMetricsResponseBodyData
	SetErrorCode(v string) *DescribeCustomAgentMonitorMetricsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeCustomAgentMonitorMetricsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeCustomAgentMonitorMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomAgentMonitorMetricsResponseBody
	GetSuccess() *bool
}

type DescribeCustomAgentMonitorMetricsResponseBody struct {
	// The response struct.
	Data *DescribeCustomAgentMonitorMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomAgentMonitorMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentMonitorMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) GetData() *DescribeCustomAgentMonitorMetricsResponseBodyData {
	return s.Data
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) SetData(v *DescribeCustomAgentMonitorMetricsResponseBodyData) *DescribeCustomAgentMonitorMetricsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) SetErrorCode(v string) *DescribeCustomAgentMonitorMetricsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) SetErrorMessage(v string) *DescribeCustomAgentMonitorMetricsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) SetRequestId(v string) *DescribeCustomAgentMonitorMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) SetSuccess(v bool) *DescribeCustomAgentMonitorMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomAgentMonitorMetricsResponseBodyData struct {
	// The number of active users.
	//
	// example:
	//
	// 6
	ActiveUserCount *int64 `json:"ActiveUserCount,omitempty" xml:"ActiveUserCount,omitempty"`
	// The custom agent ID.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The total number of dislikes.
	//
	// example:
	//
	// 0
	DislikeCount *int64 `json:"DislikeCount,omitempty" xml:"DislikeCount,omitempty"`
	// The end time of the statistical period (epoch millis).
	//
	// example:
	//
	// 1756742400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The aggregation granularity: DAY / HOUR.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The total number of likes.
	//
	// example:
	//
	// 10
	LikeCount *int64 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// The total number of sessions.
	//
	// example:
	//
	// 102
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The start time of the statistical period (epoch millis).
	//
	// example:
	//
	// 1782835200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trend data aggregated by the specified granularity. Time points without data are filled with 0. The data is sorted in chronological order.
	Trend []*DescribeCustomAgentMonitorMetricsResponseBodyDataTrend `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
}

func (s DescribeCustomAgentMonitorMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentMonitorMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetActiveUserCount() *int64 {
	return s.ActiveUserCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetDislikeCount() *int64 {
	return s.DislikeCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetLikeCount() *int64 {
	return s.LikeCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) GetTrend() []*DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	return s.Trend
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetActiveUserCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.ActiveUserCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetCustomAgentId(v string) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetDislikeCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.DislikeCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetEndTime(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetGranularity(v string) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.Granularity = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetLikeCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.LikeCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetSessionCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.SessionCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetStartTime(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) SetTrend(v []*DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) *DescribeCustomAgentMonitorMetricsResponseBodyData {
	s.Trend = v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyData) Validate() error {
	if s.Trend != nil {
		for _, item := range s.Trend {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomAgentMonitorMetricsResponseBodyDataTrend struct {
	// The number of active users within the statistical period.
	//
	// example:
	//
	// 2
	ActiveUserCount *int64 `json:"ActiveUserCount,omitempty" xml:"ActiveUserCount,omitempty"`
	// The number of dislikes within the statistical period.
	//
	// example:
	//
	// 0
	DislikeCount *int64 `json:"DislikeCount,omitempty" xml:"DislikeCount,omitempty"`
	// The number of likes within the statistical period.
	//
	// example:
	//
	// 1
	LikeCount *int64 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// The number of sessions within the statistical period.
	//
	// example:
	//
	// 10
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The statistical time. For daily granularity, the format is 2026-09-01. For hourly granularity, the format is 2026-09-01 13:00.
	//
	// example:
	//
	// 2026-09-01
	StatTime *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
	// The start timestamp of the statistical period (epoch millis).
	//
	// example:
	//
	// 1782835200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetActiveUserCount() *int64 {
	return s.ActiveUserCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetDislikeCount() *int64 {
	return s.DislikeCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetLikeCount() *int64 {
	return s.LikeCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetStatTime() *string {
	return s.StatTime
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetActiveUserCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.ActiveUserCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetDislikeCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.DislikeCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetLikeCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.LikeCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetSessionCount(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.SessionCount = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetStatTime(v string) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.StatTime = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) SetTimestamp(v int64) *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend {
	s.Timestamp = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsResponseBodyDataTrend) Validate() error {
	return dara.Validate(s)
}
