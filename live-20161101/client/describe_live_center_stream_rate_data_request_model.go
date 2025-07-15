// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterStreamRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveCenterStreamRateDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveCenterStreamRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveCenterStreamRateDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeLiveCenterStreamRateDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveCenterStreamRateDataRequest
	GetStreamName() *string
}

type DescribeLiveCenterStreamRateDataRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-03-05T18:01:03Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-03-05T18:00:53Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveCenterStreamRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterStreamRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterStreamRateDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveCenterStreamRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveCenterStreamRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveCenterStreamRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveCenterStreamRateDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveCenterStreamRateDataRequest) SetAppName(v string) *DescribeLiveCenterStreamRateDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataRequest) SetDomainName(v string) *DescribeLiveCenterStreamRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataRequest) SetEndTime(v string) *DescribeLiveCenterStreamRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataRequest) SetStartTime(v string) *DescribeLiveCenterStreamRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataRequest) SetStreamName(v string) *DescribeLiveCenterStreamRateDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveCenterStreamRateDataRequest) Validate() error {
	return dara.Validate(s)
}
