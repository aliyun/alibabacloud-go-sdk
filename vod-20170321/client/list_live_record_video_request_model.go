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
	// The name of the application used during recording.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name used during recording.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the recording operation (when the live stream recording begins). The end time must be later than the start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting method for results. Valid values:
	//
	// - **CreationTime:Desc*	- (default): sorted by creation time in descending order.
	//
	// - **CreationTime:Asc**: sorted by creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the recording operation (when the live stream recording begins). Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream used during recording.
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
