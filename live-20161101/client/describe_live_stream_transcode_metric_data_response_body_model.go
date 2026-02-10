// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetNextPageToken() *string
	SetPageSize(v int32) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetStartTime() *string
	SetStreamDetailData(v *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) *DescribeLiveStreamTranscodeMetricDataResponseBody
	GetStreamDetailData() *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData
}

type DescribeLiveStreamTranscodeMetricDataResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-11T02:46:40Z
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
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-11T03:46:40Z
	StartTime        *string                                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamDetailData *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData `json:"StreamDetailData,omitempty" xml:"StreamDetailData,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) GetStreamDetailData() *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData {
	return s.StreamDetailData
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetDomainName(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetEndTime(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetNextPageToken(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetPageSize(v int32) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetRequestId(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetStartTime(v string) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) SetStreamDetailData(v *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) *DescribeLiveStreamTranscodeMetricDataResponseBody {
	s.StreamDetailData = v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBody) Validate() error {
	if s.StreamDetailData != nil {
		if err := s.StreamDetailData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData struct {
	StreamData []*DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData `json:"StreamData,omitempty" xml:"StreamData,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) GetStreamData() []*DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	return s.StreamData
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) SetStreamData(v []*DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData {
	s.StreamData = v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailData) Validate() error {
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

type DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps           *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Resolution    *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TranscodeType *string `json:"TranscodeType,omitempty" xml:"TranscodeType,omitempty"`
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetFps() *string {
	return s.Fps
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) GetTranscodeType() *string {
	return s.TranscodeType
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetAppName(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetDuration(v int64) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetFps(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.Fps = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetRegion(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.Region = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetResolution(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.Resolution = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetStreamName(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetTimeStamp(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) SetTranscodeType(v string) *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData {
	s.TranscodeType = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponseBodyStreamDetailDataStreamData) Validate() error {
	return dara.Validate(s)
}
