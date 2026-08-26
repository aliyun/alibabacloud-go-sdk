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
	// The application name. Specify this parameter to query stream-level data for a specific application.
	//
	// > If you specify StreamName, you must also specify AppName.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// - The accelerated domain name to query. Only a single domain name can be queried at a time. An error is returned if multiple domain names are specified.
	//
	// - If AppName and StreamName are not specified, stream-level data for all streams under the domain name is returned.
	//
	// - If the domain name is left empty, aggregate data for all accelerated domain names under the account is returned.
	//
	// - If DomainName is specified and both AppName and StreamName are set to all, aggregate data for the specified accelerated domain name is returned.
	//
	// - When you specify DomainName, make sure the domain name is a live streaming domain and the user calling this operation has the required permissions on the domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time, and the difference cannot exceed 1 day. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The paged query token. A maximum of 5,000 rows of data can be returned per query. If the data to query exceeds 5,000 rows, the response includes the starting index for the next paging request. Pass this token in the request to continue querying data from where the previous query ended.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67CIBKLw*****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The stream protocol. Valid values: **flv**, **hls**, **rtmp**, **rts**, and **p2p**.
	//
	// You can query data for multiple protocols by separating them with commas (,). Data for multiple protocols is not aggregated and is output at the stream level.
	//
	// > The **rts*	- option queries Real-Time Streaming (RTS) streams that use the ARTC protocol.
	//
	// > - When using rts, you may need to additionally collect statistics for the xxx_AliRTS-opus transcoding stream. This is because when playing an RTS stream on the web, a transcoding stream with the _AliRTS-opus suffix appended to the stream name is automatically generated. For more information, see [RTS sub-second latency automatic transcoding](https://help.aliyun.com/document_detail/2948703.html).
	//
	// example:
	//
	// flv
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. Specify this parameter together with AppName to return stream-level data.
	//
	// > If you specify StreamName, you must also specify AppName.
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
