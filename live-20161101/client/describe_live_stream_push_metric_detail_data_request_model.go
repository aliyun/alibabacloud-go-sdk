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
	// The name of the application to which the live stream belongs. Specify the application name to query data at the stream granularity for the corresponding application.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// - The accelerated domain name to query. Only a single domain name can be queried. An error is returned if you specify multiple domain names.
	//
	// - If you do not specify AppName or StreamName, data at the stream granularity for all streams under the specified accelerated domain name is returned without aggregation.
	//
	// - If DomainName is specified and both AppName and StreamName are set to all, aggregate data at the stream granularity for all streams under the specified accelerated domain name is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the UTC time zone. The end time must be later than the start time, and the difference cannot exceed 1 day.
	//
	// Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The paging query token. Each query returns a maximum of 5,000 rows of data. If the data to be queried exceeds 5,000 rows, the response includes the start index for the next query.
	//
	// Pass this token in the request to continue querying data from the row after the last row returned in the previous query.
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the UTC time zone.
	//
	// Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. If you specify StreamName, data at the stream granularity is returned for the specified StreamName under the specified AppName. You must specify AppName when you specify StreamName.
	//
	// example:
	//
	// liveStream****
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
