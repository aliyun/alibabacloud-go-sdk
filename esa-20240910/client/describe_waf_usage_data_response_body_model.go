// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeWafUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeWafUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeWafUsageDataResponseBody
	GetStartTime() *string
	SetUsageData(v []*DescribeWafUsageDataResponseBodyUsageData) *DescribeWafUsageDataResponseBody
	GetUsageData() []*DescribeWafUsageDataResponseBodyUsageData
}

type DescribeWafUsageDataResponseBody struct {
	// The end of the time range for the returned data. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC+0.
	//
	// example:
	//
	// 2022-08-10T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2022-08-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The returned data.
	UsageData []*DescribeWafUsageDataResponseBodyUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeWafUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeWafUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWafUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeWafUsageDataResponseBody) GetUsageData() []*DescribeWafUsageDataResponseBodyUsageData {
	return s.UsageData
}

func (s *DescribeWafUsageDataResponseBody) SetEndTime(v string) *DescribeWafUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeWafUsageDataResponseBody) SetRequestId(v string) *DescribeWafUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafUsageDataResponseBody) SetStartTime(v string) *DescribeWafUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeWafUsageDataResponseBody) SetUsageData(v []*DescribeWafUsageDataResponseBodyUsageData) *DescribeWafUsageDataResponseBody {
	s.UsageData = v
	return s
}

func (s *DescribeWafUsageDataResponseBody) Validate() error {
	if s.UsageData != nil {
		for _, item := range s.UsageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWafUsageDataResponseBodyUsageData struct {
	// The number of requests with normal access.
	//
	// example:
	//
	// 123
	AccessCount *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The number of blocked requests.
	//
	// example:
	//
	// 123
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// The number of observed requests.
	//
	// example:
	//
	// 123
	ObserveCount *int64 `json:"ObserveCount,omitempty" xml:"ObserveCount,omitempty"`
	// The domain record name.
	//
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The beginning of the time interval.
	//
	// example:
	//
	// 2022-08-10T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeWafUsageDataResponseBodyUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafUsageDataResponseBodyUsageData) GoString() string {
	return s.String()
}

func (s *DescribeWafUsageDataResponseBodyUsageData) GetAccessCount() *int64 {
	return s.AccessCount
}

func (s *DescribeWafUsageDataResponseBodyUsageData) GetBlockCount() *int64 {
	return s.BlockCount
}

func (s *DescribeWafUsageDataResponseBodyUsageData) GetObserveCount() *int64 {
	return s.ObserveCount
}

func (s *DescribeWafUsageDataResponseBodyUsageData) GetRecordName() *string {
	return s.RecordName
}

func (s *DescribeWafUsageDataResponseBodyUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeWafUsageDataResponseBodyUsageData) SetAccessCount(v int64) *DescribeWafUsageDataResponseBodyUsageData {
	s.AccessCount = &v
	return s
}

func (s *DescribeWafUsageDataResponseBodyUsageData) SetBlockCount(v int64) *DescribeWafUsageDataResponseBodyUsageData {
	s.BlockCount = &v
	return s
}

func (s *DescribeWafUsageDataResponseBodyUsageData) SetObserveCount(v int64) *DescribeWafUsageDataResponseBodyUsageData {
	s.ObserveCount = &v
	return s
}

func (s *DescribeWafUsageDataResponseBodyUsageData) SetRecordName(v string) *DescribeWafUsageDataResponseBodyUsageData {
	s.RecordName = &v
	return s
}

func (s *DescribeWafUsageDataResponseBodyUsageData) SetTimeStamp(v string) *DescribeWafUsageDataResponseBodyUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeWafUsageDataResponseBodyUsageData) Validate() error {
	return dara.Validate(s)
}
