// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMetricDetailDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetNextPageToken() *string
	SetOwnerId(v int64) *DescribeLiveStreamMetricDetailDataRequest
	GetOwnerId() *int64
	SetProtocol(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetProtocol() *string
	SetRegionId(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamMetricDetailDataRequest
	GetStreamName() *string
}

type DescribeLiveStreamMetricDetailDataRequest struct {
	// The name of the application for which you want to query the monitoring data of streams.
	//
	// >  If you specify the StreamName parameter, you must also specify the AppName parameter.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 	- The accelerated domain name. You can specify only one domain name. If you specify multiple domain names, an error occurs.
	//
	// 	- If you do not specify the AppName and StreamName parameters, monitoring data of all streams for the domain name is returned.
	//
	// 	- If you leave this parameter empty, monitoring data of streams under all domain names is returned.
	//
	// 	- If you specify the DomainName parameter and set both the AppName and StreamName parameters to all, monitoring data of all streams in all applications under the specified domain name is returned.
	//
	// 	- When you specify the DomainName parameter, make sure that the domain name is a domain name used for live streaming and that you have the permissions on the domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time, and the maximum time range that can be specified is one day. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The token used to query data by page. Up to 5,000 rows of data can be returned per query. If the number of rows exceeds 5,000, a token that determines the start point of the next query is provided in the response. If you specify this parameter, data continues to be obtained from the end of the previous query.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67CIBKLw*****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The streaming protocol. Valid values: **flv**, **hls**, **rtmp**, **rts**, and **p2p**.
	//
	// You can specify multiple protocols. Separate multiple protocols with commas (,). However, data over multiple protocols is not aggregated and is returned based on the stream.
	//
	// example:
	//
	// flv
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the stream. The stream must belong to the application that is specified by the AppName parameter.
	//
	// >  If you specify the StreamName parameter, you must also specify the AppName parameter.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamMetricDetailDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMetricDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamMetricDetailDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetAppName(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetDomainName(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetEndTime(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetNextPageToken(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetOwnerId(v int64) *DescribeLiveStreamMetricDetailDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetProtocol(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetRegionId(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetStartTime(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) SetStreamName(v string) *DescribeLiveStreamMetricDetailDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataRequest) Validate() error {
	return dara.Validate(s)
}
