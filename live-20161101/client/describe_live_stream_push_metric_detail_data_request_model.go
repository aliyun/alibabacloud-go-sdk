// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPushMetricDetailDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetNextPageToken() *string
	SetOwnerId(v int64) *DescribeLiveStreamPushMetricDetailDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamPushMetricDetailDataRequest
	GetStreamName() *string
}

type DescribeLiveStreamPushMetricDetailDataRequest struct {
	// The name of the application to which the live stream belongs. The stream-level data of this application is returned.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 	- The accelerated domain name. You can specify only one domain name. If you specify multiple domain names, an error occurs.
	//
	// 	- If you do not specify the AppName and StreamName parameters, data of all streams under the specified domain name is returned. The data is not aggregated.
	//
	// 	- If you specify the DomainName parameter and set both the AppName and StreamName parameters to all, data of all streams in all applications under the specified domain name is aggregated and returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time, and the maximum time range that can be specified is one day. Specify the time in the ISO 8601 standard
	//
	// in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
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
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. The data of the stream in the specified application is returned. If the StreamName parameter is specified, the AppName parameter must also be specified.
	//
	// example:
	//
	// test
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamPushMetricDetailDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPushMetricDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetAppName(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetDomainName(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetEndTime(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetNextPageToken(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetOwnerId(v int64) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetRegionId(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetStartTime(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) SetStreamName(v string) *DescribeLiveStreamPushMetricDetailDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataRequest) Validate() error {
	return dara.Validate(s)
}
