// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamHistoryUserNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamHistoryUserNumRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamHistoryUserNumRequest
	GetStreamName() *string
}

type DescribeLiveStreamHistoryUserNumRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  The time range specified by the StartTime and EndTime parameters cannot exceed one day. The end time must not be later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  You can query data in the last **30*	- days.
	//
	// This parameter is required.
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

func (s DescribeLiveStreamHistoryUserNumRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamHistoryUserNumRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetAppName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetDomainName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetEndTime(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetOwnerId(v int64) *DescribeLiveStreamHistoryUserNumRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetSecurityToken(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetStartTime(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetStreamName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) Validate() error {
	return dara.Validate(s)
}
