// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKVvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKVvDataResponseBody
	GetDataInterval() *string
	SetEndTime(v string) *DescribeRTSNativeSDKVvDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeRTSNativeSDKVvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeRTSNativeSDKVvDataResponseBody
	GetStartTime() *string
	SetVvData(v []*DescribeRTSNativeSDKVvDataResponseBodyVvData) *DescribeRTSNativeSDKVvDataResponseBody
	GetVvData() []*DescribeRTSNativeSDKVvDataResponseBodyVvData
}

type DescribeRTSNativeSDKVvDataResponseBody struct {
	// The time granularity.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The end of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of playbacks and the number of successful playbacks at each interval.
	VvData []*DescribeRTSNativeSDKVvDataResponseBodyVvData `json:"VvData,omitempty" xml:"VvData,omitempty" type:"Repeated"`
}

func (s DescribeRTSNativeSDKVvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKVvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) GetVvData() []*DescribeRTSNativeSDKVvDataResponseBodyVvData {
	return s.VvData
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) SetDataInterval(v string) *DescribeRTSNativeSDKVvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) SetEndTime(v string) *DescribeRTSNativeSDKVvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) SetRequestId(v string) *DescribeRTSNativeSDKVvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) SetStartTime(v string) *DescribeRTSNativeSDKVvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) SetVvData(v []*DescribeRTSNativeSDKVvDataResponseBodyVvData) *DescribeRTSNativeSDKVvDataResponseBody {
	s.VvData = v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBody) Validate() error {
	if s.VvData != nil {
		for _, item := range s.VvData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRTSNativeSDKVvDataResponseBodyVvData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The number of successful playbacks within the period of time.
	//
	// example:
	//
	// 99
	VvSuccess *string `json:"VvSuccess,omitempty" xml:"VvSuccess,omitempty"`
	// The total number of playbacks within the period of time.
	//
	// example:
	//
	// 100
	VvTotal *string `json:"VvTotal,omitempty" xml:"VvTotal,omitempty"`
}

func (s DescribeRTSNativeSDKVvDataResponseBodyVvData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKVvDataResponseBodyVvData) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) GetVvSuccess() *string {
	return s.VvSuccess
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) GetVvTotal() *string {
	return s.VvTotal
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) SetTimeStamp(v string) *DescribeRTSNativeSDKVvDataResponseBodyVvData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) SetVvSuccess(v string) *DescribeRTSNativeSDKVvDataResponseBodyVvData {
	s.VvSuccess = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) SetVvTotal(v string) *DescribeRTSNativeSDKVvDataResponseBodyVvData {
	s.VvTotal = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponseBodyVvData) Validate() error {
	return dara.Validate(s)
}
