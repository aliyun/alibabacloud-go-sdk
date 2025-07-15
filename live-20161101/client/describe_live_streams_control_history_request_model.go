// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsControlHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamsControlHistoryRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamsControlHistoryRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamsControlHistoryRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamsControlHistoryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamsControlHistoryRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamsControlHistoryRequest
	GetStartTime() *string
}

type DescribeLiveStreamsControlHistoryRequest struct {
	// The application name. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  The time range specified by the StartTime and EndTime parameters cannot exceed seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveStreamsControlHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsControlHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetAppName(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetDomainName(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetEndTime(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetOwnerId(v int64) *DescribeLiveStreamsControlHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetRegionId(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetStartTime(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) Validate() error {
	return dara.Validate(s)
}
