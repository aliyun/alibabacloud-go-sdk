// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *ListSchedulesRequest
	GetChannelName() *string
	SetPageNo(v int32) *ListSchedulesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSchedulesRequest
	GetPageSize() *int32
	SetWindowDurationSeconds(v int64) *ListSchedulesRequest
	GetWindowDurationSeconds() *int64
}

type ListSchedulesRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time window of the program schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14400
	WindowDurationSeconds *int64 `json:"WindowDurationSeconds,omitempty" xml:"WindowDurationSeconds,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListSchedulesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSchedulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchedulesRequest) GetWindowDurationSeconds() *int64 {
	return s.WindowDurationSeconds
}

func (s *ListSchedulesRequest) SetChannelName(v string) *ListSchedulesRequest {
	s.ChannelName = &v
	return s
}

func (s *ListSchedulesRequest) SetPageNo(v int32) *ListSchedulesRequest {
	s.PageNo = &v
	return s
}

func (s *ListSchedulesRequest) SetPageSize(v int32) *ListSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesRequest) SetWindowDurationSeconds(v int64) *ListSchedulesRequest {
	s.WindowDurationSeconds = &v
	return s
}

func (s *ListSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
