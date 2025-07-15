// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveStreamsNotifyRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveStreamsNotifyRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveStreamsNotifyRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetStatus() *string
	SetStreamName(v string) *DescribeLiveStreamsNotifyRecordsRequest
	GetStreamName() *string
}

type DescribeLiveStreamsNotifyRecordsRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// push.example1.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 500. Valid values: integers from 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The start time must be in the last seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-10T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to return the records of successful or failed callbacks. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamsNotifyRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetAppName(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetDomainName(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetEndTime(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetOwnerId(v int64) *DescribeLiveStreamsNotifyRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetPageNumber(v int32) *DescribeLiveStreamsNotifyRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetPageSize(v int32) *DescribeLiveStreamsNotifyRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetRegionId(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetStartTime(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetStatus(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) SetStreamName(v string) *DescribeLiveStreamsNotifyRecordsRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyRecordsRequest) Validate() error {
	return dara.Validate(s)
}
