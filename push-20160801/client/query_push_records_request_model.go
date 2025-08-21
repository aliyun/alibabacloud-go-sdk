// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryPushRecordsRequest
	GetAppKey() *int64
	SetEndTime(v string) *QueryPushRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *QueryPushRecordsRequest
	GetKeyword() *string
	SetNextToken(v string) *QueryPushRecordsRequest
	GetNextToken() *string
	SetPage(v int32) *QueryPushRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *QueryPushRecordsRequest
	GetPageSize() *int32
	SetPushType(v string) *QueryPushRecordsRequest
	GetPushType() *string
	SetSource(v string) *QueryPushRecordsRequest
	GetSource() *string
	SetStartTime(v string) *QueryPushRecordsRequest
	GetStartTime() *string
	SetTarget(v string) *QueryPushRecordsRequest
	GetTarget() *string
}

type QueryPushRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 333526247
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-29T06:24:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// FFPpkmhCPm*****************xjk=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 8
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// NOTICE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// example:
	//
	// API
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-15T02:05:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryPushRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryPushRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPushRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryPushRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryPushRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *QueryPushRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPushRecordsRequest) GetPushType() *string {
	return s.PushType
}

func (s *QueryPushRecordsRequest) GetSource() *string {
	return s.Source
}

func (s *QueryPushRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPushRecordsRequest) GetTarget() *string {
	return s.Target
}

func (s *QueryPushRecordsRequest) SetAppKey(v int64) *QueryPushRecordsRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushRecordsRequest) SetEndTime(v string) *QueryPushRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetKeyword(v string) *QueryPushRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *QueryPushRecordsRequest) SetNextToken(v string) *QueryPushRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPage(v int32) *QueryPushRecordsRequest {
	s.Page = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPageSize(v int32) *QueryPushRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPushType(v string) *QueryPushRecordsRequest {
	s.PushType = &v
	return s
}

func (s *QueryPushRecordsRequest) SetSource(v string) *QueryPushRecordsRequest {
	s.Source = &v
	return s
}

func (s *QueryPushRecordsRequest) SetStartTime(v string) *QueryPushRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetTarget(v string) *QueryPushRecordsRequest {
	s.Target = &v
	return s
}

func (s *QueryPushRecordsRequest) Validate() error {
	return dara.Validate(s)
}
