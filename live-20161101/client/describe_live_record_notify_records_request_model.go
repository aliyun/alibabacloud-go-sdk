// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveRecordNotifyRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveRecordNotifyRecordsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveRecordNotifyRecordsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetStatus() *string
	SetStreamName(v string) *DescribeLiveRecordNotifyRecordsRequest
	GetStreamName() *string
}

type DescribeLiveRecordNotifyRecordsRequest struct {
	// The name of the application to which the live stream belongs.
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
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-10T21:03:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1. Valid values: 1 to 100000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 500. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  You can query data within the last seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-10T21:03:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether the callback is successful. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveRecordNotifyRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveRecordNotifyRecordsRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetAppName(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetDomainName(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetEndTime(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetOwnerId(v int64) *DescribeLiveRecordNotifyRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetPageNumber(v int64) *DescribeLiveRecordNotifyRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetPageSize(v int64) *DescribeLiveRecordNotifyRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetRegionId(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetStartTime(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetStatus(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) SetStreamName(v string) *DescribeLiveRecordNotifyRecordsRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordNotifyRecordsRequest) Validate() error {
	return dara.Validate(s)
}
