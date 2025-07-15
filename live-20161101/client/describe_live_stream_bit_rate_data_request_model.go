// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamBitRateDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamBitRateDataRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamBitRateDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamBitRateDataRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamBitRateDataRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveStreamBitRateDataRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeLiveStreamBitRateDataRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamBitRateDataRequest
	GetStreamName() *string
}

type DescribeLiveStreamBitRateDataRequest struct {
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
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
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
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamBitRateDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamBitRateDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamBitRateDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamBitRateDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamBitRateDataRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamBitRateDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamBitRateDataRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamBitRateDataRequest) SetAppName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetDomainName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetEndTime(v string) *DescribeLiveStreamBitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetOwnerId(v int64) *DescribeLiveStreamBitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetSecurityToken(v string) *DescribeLiveStreamBitRateDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetStartTime(v string) *DescribeLiveStreamBitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetStreamName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) Validate() error {
	return dara.Validate(s)
}
