// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListLogstashLogRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListLogstashLogRequest
	GetEndTime() *int64
	SetPage(v int32) *ListLogstashLogRequest
	GetPage() *int32
	SetQuery(v string) *ListLogstashLogRequest
	GetQuery() *string
	SetSize(v int32) *ListLogstashLogRequest
	GetSize() *int32
	SetType(v string) *ListLogstashLogRequest
	GetType() *string
}

type ListLogstashLogRequest struct {
	// 20
	//
	// example:
	//
	// 1531910852074
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1531910852074
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 1
	//
	// This parameter is required.
	//
	// example:
	//
	// host:10.7.xx.xx AND level:info AND content:opening
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The severity level of the log entry. Including trace, debug, info, warn, error, etc. (GC logs have no level).
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// 1531910852074
	//
	// This parameter is required.
	//
	// example:
	//
	// LOGSTASH_INSTANCE_LOG
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListLogstashLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashLogRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashLogRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListLogstashLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLogstashLogRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListLogstashLogRequest) GetQuery() *string {
	return s.Query
}

func (s *ListLogstashLogRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListLogstashLogRequest) GetType() *string {
	return s.Type
}

func (s *ListLogstashLogRequest) SetBeginTime(v int64) *ListLogstashLogRequest {
	s.BeginTime = &v
	return s
}

func (s *ListLogstashLogRequest) SetEndTime(v int64) *ListLogstashLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListLogstashLogRequest) SetPage(v int32) *ListLogstashLogRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashLogRequest) SetQuery(v string) *ListLogstashLogRequest {
	s.Query = &v
	return s
}

func (s *ListLogstashLogRequest) SetSize(v int32) *ListLogstashLogRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashLogRequest) SetType(v string) *ListLogstashLogRequest {
	s.Type = &v
	return s
}

func (s *ListLogstashLogRequest) Validate() error {
	return dara.Validate(s)
}
