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
	// The application name.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name. Only a single domain name can be queried at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The end time must be later than the start time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06-11T03:46:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The paged query token. Each query returns a maximum of 5,000 rows of data. If the data to be queried exceeds 5,000 rows, the response provides the start index for the next query.
	//
	// Pass this token in the request to continue querying data from the row after the last row returned in the previous query. This token is used for paging.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67C***37C0=
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06-11T02:46:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name.
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
