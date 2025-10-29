// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPushMetricDetailDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetNextPageToken() *string
	SetPageSize(v int32) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetStartTime() *string
	SetStreamDetailData(v *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) *DescribeLiveStreamPushMetricDetailDataResponseBody
	GetStreamDetailData() *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData
}

type DescribeLiveStreamPushMetricDetailDataResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range that was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A pagination token. When you call this operation, up to 5,000 rows of data can be returned per query. If the number of rows exceeds 5,000, the response includes a pagination token that is used in the next request to retrieve a new page of results.
	//
	// When you specify the token in the next query, data continues to be obtained from the end of the previous query.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67C***37C0=
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 5000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5EBF2AC3-4B73-40A5-8B32-83F49D5F035E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The data array returned.
	StreamDetailData *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData `json:"StreamDetailData,omitempty" xml:"StreamDetailData,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) GetStreamDetailData() *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData {
	return s.StreamDetailData
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetDomainName(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetEndTime(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetNextPageToken(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetPageSize(v int32) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetRequestId(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetStartTime(v string) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) SetStreamDetailData(v *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) *DescribeLiveStreamPushMetricDetailDataResponseBody {
	s.StreamDetailData = v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBody) Validate() error {
	if s.StreamDetailData != nil {
		if err := s.StreamDetailData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData struct {
	StreamData []*DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData `json:"StreamData,omitempty" xml:"StreamData,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) GetStreamData() []*DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	return s.StreamData
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) SetStreamData(v []*DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData {
	s.StreamData = v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailData) Validate() error {
	if s.StreamData != nil {
		for _, item := range s.StreamData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData struct {
	// The name of the application.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The total bandwidth consumed by the stream per minute. Unit: bit/s.
	//
	// example:
	//
	// 423304182
	ReqBps *float32 `json:"ReqBps,omitempty" xml:"ReqBps,omitempty"`
	// The total amount of traffic consumed by the stream per minute. Unit: bytes.
	//
	// example:
	//
	// 423304182.66
	ReqTraffic *int64 `json:"ReqTraffic,omitempty" xml:"ReqTraffic,omitempty"`
	// The name of the stream.
	//
	// example:
	//
	// test.flv
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2022-09-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GetReqBps() *float32 {
	return s.ReqBps
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GetReqTraffic() *int64 {
	return s.ReqTraffic
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) SetAppName(v string) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) SetReqBps(v float32) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.ReqBps = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) SetReqTraffic(v int64) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.ReqTraffic = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) SetStreamName(v string) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) SetTimeStamp(v string) *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponseBodyStreamDetailDataStreamData) Validate() error {
	return dara.Validate(s)
}
