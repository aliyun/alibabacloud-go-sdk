// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamRecordContentRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamRecordContentRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamRecordContentRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamRecordContentRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveStreamRecordContentRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeLiveStreamRecordContentRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamRecordContentRequest
	GetStreamName() *string
}

type DescribeLiveStreamRecordContentRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The operation that you want to perform. Set the value to **DescribeLiveStreamRecordContent**.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The beginning of the time range to query. You can only query the recordings in the last 6 months. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamRecordContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamRecordContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamRecordContentRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamRecordContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamRecordContentRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamRecordContentRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamRecordContentRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamRecordContentRequest) SetAppName(v string) *DescribeLiveStreamRecordContentRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetDomainName(v string) *DescribeLiveStreamRecordContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetEndTime(v string) *DescribeLiveStreamRecordContentRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetOwnerId(v int64) *DescribeLiveStreamRecordContentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordContentRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetStartTime(v string) *DescribeLiveStreamRecordContentRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetStreamName(v string) *DescribeLiveStreamRecordContentRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) Validate() error {
	return dara.Validate(s)
}
