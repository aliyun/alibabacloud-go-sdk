// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListLiveRecordVideoRequest
	GetAppName() *string
	SetDomainName(v string) *ListLiveRecordVideoRequest
	GetDomainName() *string
	SetEndTime(v string) *ListLiveRecordVideoRequest
	GetEndTime() *string
	SetPageNo(v int32) *ListLiveRecordVideoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveRecordVideoRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLiveRecordVideoRequest
	GetSortBy() *string
	SetStartTime(v string) *ListLiveRecordVideoRequest
	GetStartTime() *string
	SetStreamName(v string) *ListLiveRecordVideoRequest
	GetStreamName() *string
}

type ListLiveRecordVideoRequest struct {
	// The name of the application that was used to record the live stream.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name of the recorded live stream.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The query is performed based on the time range during which the required live streams were recorded. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting rule of the results. Valid values:
	//
	// 	- **CreationTime:Desc**: The results are sorted in reverse chronological order based on the creation time.
	//
	// 	- **CreationTime:Asc**: The results are sorted in chronological order based on the creation time.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. The query is performed based on the time range during which the required live streams were recorded. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the recorded live stream.
	//
	// example:
	//
	// live-test
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ListLiveRecordVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordVideoRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListLiveRecordVideoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListLiveRecordVideoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLiveRecordVideoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveRecordVideoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveRecordVideoRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordVideoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLiveRecordVideoRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *ListLiveRecordVideoRequest) SetAppName(v string) *ListLiveRecordVideoRequest {
	s.AppName = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetDomainName(v string) *ListLiveRecordVideoRequest {
	s.DomainName = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetEndTime(v string) *ListLiveRecordVideoRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetPageNo(v int32) *ListLiveRecordVideoRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetPageSize(v int32) *ListLiveRecordVideoRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetSortBy(v string) *ListLiveRecordVideoRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetStartTime(v string) *ListLiveRecordVideoRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetStreamName(v string) *ListLiveRecordVideoRequest {
	s.StreamName = &v
	return s
}

func (s *ListLiveRecordVideoRequest) Validate() error {
	return dara.Validate(s)
}
