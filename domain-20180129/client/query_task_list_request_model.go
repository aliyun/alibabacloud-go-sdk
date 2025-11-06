// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginCreateTime(v int64) *QueryTaskListRequest
	GetBeginCreateTime() *int64
	SetEndCreateTime(v int64) *QueryTaskListRequest
	GetEndCreateTime() *int64
	SetLang(v string) *QueryTaskListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryTaskListRequest
	GetPageSize() *int32
	SetUserClientIp(v string) *QueryTaskListRequest
	GetUserClientIp() *string
}

type QueryTaskListRequest struct {
	// example:
	//
	// 1522080000000
	BeginCreateTime *int64 `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
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
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskListRequest) GetBeginCreateTime() *int64 {
	return s.BeginCreateTime
}

func (s *QueryTaskListRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *QueryTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTaskListRequest) SetBeginCreateTime(v int64) *QueryTaskListRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *QueryTaskListRequest) SetEndCreateTime(v int64) *QueryTaskListRequest {
	s.EndCreateTime = &v
	return s
}

func (s *QueryTaskListRequest) SetLang(v string) *QueryTaskListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskListRequest) SetPageNum(v int32) *QueryTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskListRequest) SetPageSize(v int32) *QueryTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListRequest) SetUserClientIp(v string) *QueryTaskListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTaskListRequest) Validate() error {
	return dara.Validate(s)
}
