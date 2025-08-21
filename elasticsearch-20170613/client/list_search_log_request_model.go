// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListSearchLogRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListSearchLogRequest
	GetEndTime() *int64
	SetPage(v int32) *ListSearchLogRequest
	GetPage() *int32
	SetQuery(v string) *ListSearchLogRequest
	GetQuery() *string
	SetSize(v int32) *ListSearchLogRequest
	GetSize() *int32
	SetType(v string) *ListSearchLogRequest
	GetType() *string
}

type ListSearchLogRequest struct {
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
	// The header of the response.
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
	// host:``172.16.**.**`` AND content:netty
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The number of entries returned per page.
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
	// INSTANCELOG
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSearchLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLogRequest) GoString() string {
	return s.String()
}

func (s *ListSearchLogRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListSearchLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSearchLogRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSearchLogRequest) GetQuery() *string {
	return s.Query
}

func (s *ListSearchLogRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListSearchLogRequest) GetType() *string {
	return s.Type
}

func (s *ListSearchLogRequest) SetBeginTime(v int64) *ListSearchLogRequest {
	s.BeginTime = &v
	return s
}

func (s *ListSearchLogRequest) SetEndTime(v int64) *ListSearchLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListSearchLogRequest) SetPage(v int32) *ListSearchLogRequest {
	s.Page = &v
	return s
}

func (s *ListSearchLogRequest) SetQuery(v string) *ListSearchLogRequest {
	s.Query = &v
	return s
}

func (s *ListSearchLogRequest) SetSize(v int32) *ListSearchLogRequest {
	s.Size = &v
	return s
}

func (s *ListSearchLogRequest) SetType(v string) *ListSearchLogRequest {
	s.Type = &v
	return s
}

func (s *ListSearchLogRequest) Validate() error {
	return dara.Validate(s)
}
