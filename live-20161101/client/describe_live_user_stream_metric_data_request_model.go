// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserStreamMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveUserStreamMetricDataRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveUserStreamMetricDataRequest
	GetPageSize() *int64
	SetProtocol(v string) *DescribeLiveUserStreamMetricDataRequest
	GetProtocol() *string
	SetStartTime(v string) *DescribeLiveUserStreamMetricDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveUserStreamMetricDataRequest
	GetStreamName() *string
}

type DescribeLiveUserStreamMetricDataRequest struct {
	// The application name. Specify the application name to query stream-level data for the corresponding application. If `StreamName` is specified, `AppName` must also be specified.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain to query.
	//
	//
	// > Only a single domain name is supported. An error is returned if multiple domain names are specified. If the domain name is empty, aggregate data for all streaming domains under the user is queried. If `AppName` and `StreamName` are not specified, stream-level data for all streams under the domain is returned.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time and the difference cannot exceed 1 day. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 5000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The stream protocol name. Specify the protocol name to query data for the corresponding protocol. Supported protocols: `flv`, `hls`, `rtmp`, `rts`, `p2p`. You can query data for multiple protocols by separating them with commas (,). Data for multiple protocols is not aggregated and is output at the stream level.
	//
	// > The **rts*	- option queries Real-Time Streaming (RTS) streams using the ARTC protocol.
	//
	// > - When using rts, you may need to additionally count the xxx_AliRTS-opus transcoding stream. This is because when playing an RTS stream on the web, a transcoding stream with the _AliRTS-opus suffix appended to the stream name is automatically generated, producing transcoding stream data. For more information, see [RTS sub-second latency automatic transcoding](https://help.aliyun.com/document_detail/2948703.html).
	//
	// example:
	//
	// flv
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. If `StreamName` is specified, stream-level data for the specified `StreamName` under the specified `AppName` is returned. If `StreamName` is specified, `AppName` must also be specified.
	//
	// example:
	//
	// test.flv
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveUserStreamMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserStreamMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserStreamMetricDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetAppName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetDomainName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetEndTime(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetPageSize(v int64) *DescribeLiveUserStreamMetricDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetProtocol(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetStartTime(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) SetStreamName(v string) *DescribeLiveUserStreamMetricDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
