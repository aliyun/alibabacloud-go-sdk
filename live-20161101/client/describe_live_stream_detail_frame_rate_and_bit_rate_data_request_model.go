// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDetailFrameRateAndBitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest
	GetStreamName() *string
}

type DescribeLiveStreamDetailFrameRateAndBitRateDataRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain or streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If the StartTime and EndTime parameters are invalid, or if the StartTime and EndTime parameters are not specified, data in the last hour is queried by default.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleStreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetAppName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetDomainName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetEndTime(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetOwnerId(v int64) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetRegionId(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetStartTime(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) SetStreamName(v string) *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
