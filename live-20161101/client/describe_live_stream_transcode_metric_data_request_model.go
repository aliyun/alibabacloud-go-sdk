// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetNextPageToken() *string
	SetOwnerId(v int64) *DescribeLiveStreamTranscodeMetricDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamTranscodeMetricDataRequest
	GetStreamName() *string
}

type DescribeLiveStreamTranscodeMetricDataRequest struct {
	// The name of the application.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06-11T03:46:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A pagination token. When you call this operation, up to 5,000 rows of data can be returned per query. If the number of rows exceeds 5,000, the response includes a pagination token that is used in the next request to retrieve a new page of results.
	//
	// When you specify the token in the next query, data continues to be obtained from the end of the previous query.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67C***37C0=
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06-11T02:46:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the stream.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamTranscodeMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetAppName(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetDomainName(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetEndTime(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetNextPageToken(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetOwnerId(v int64) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetRegionId(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetStartTime(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) SetStreamName(v string) *DescribeLiveStreamTranscodeMetricDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
