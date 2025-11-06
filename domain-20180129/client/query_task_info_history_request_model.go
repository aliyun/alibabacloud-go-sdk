// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskInfoHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginCreateTime(v int64) *QueryTaskInfoHistoryRequest
	GetBeginCreateTime() *int64
	SetCreateTimeCursor(v int64) *QueryTaskInfoHistoryRequest
	GetCreateTimeCursor() *int64
	SetEndCreateTime(v int64) *QueryTaskInfoHistoryRequest
	GetEndCreateTime() *int64
	SetLang(v string) *QueryTaskInfoHistoryRequest
	GetLang() *string
	SetPageSize(v int32) *QueryTaskInfoHistoryRequest
	GetPageSize() *int32
	SetTaskNoCursor(v string) *QueryTaskInfoHistoryRequest
	GetTaskNoCursor() *string
	SetUserClientIp(v string) *QueryTaskInfoHistoryRequest
	GetUserClientIp() *string
}

type QueryTaskInfoHistoryRequest struct {
	// example:
	//
	// 1522080000000
	BeginCreateTime *int64 `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	// example:
	//
	// 1522080000000
	CreateTimeCursor *int64 `json:"CreateTimeCursor,omitempty" xml:"CreateTimeCursor,omitempty"`
	// example:
	//
	// 1522080000000
	EndCreateTime *int64 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// aa634d3f-927e-4d17-9d2c-test
	TaskNoCursor *string `json:"TaskNoCursor,omitempty" xml:"TaskNoCursor,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskInfoHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryRequest) GetBeginCreateTime() *int64 {
	return s.BeginCreateTime
}

func (s *QueryTaskInfoHistoryRequest) GetCreateTimeCursor() *int64 {
	return s.CreateTimeCursor
}

func (s *QueryTaskInfoHistoryRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *QueryTaskInfoHistoryRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTaskInfoHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskInfoHistoryRequest) GetTaskNoCursor() *string {
	return s.TaskNoCursor
}

func (s *QueryTaskInfoHistoryRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTaskInfoHistoryRequest) SetBeginCreateTime(v int64) *QueryTaskInfoHistoryRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetCreateTimeCursor(v int64) *QueryTaskInfoHistoryRequest {
	s.CreateTimeCursor = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetEndCreateTime(v int64) *QueryTaskInfoHistoryRequest {
	s.EndCreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetLang(v string) *QueryTaskInfoHistoryRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetPageSize(v int32) *QueryTaskInfoHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetTaskNoCursor(v string) *QueryTaskInfoHistoryRequest {
	s.TaskNoCursor = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetUserClientIp(v string) *QueryTaskInfoHistoryRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) Validate() error {
	return dara.Validate(s)
}
