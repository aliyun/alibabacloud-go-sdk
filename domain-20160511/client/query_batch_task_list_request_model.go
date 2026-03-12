// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginCreateTime(v string) *QueryBatchTaskListRequest
	GetBeginCreateTime() *string
	SetEndCreateTime(v string) *QueryBatchTaskListRequest
	GetEndCreateTime() *string
	SetLang(v string) *QueryBatchTaskListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryBatchTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryBatchTaskListRequest
	GetPageSize() *int32
	SetUserClientIp(v string) *QueryBatchTaskListRequest
	GetUserClientIp() *string
}

type QueryBatchTaskListRequest struct {
	BeginCreateTime *string `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	EndCreateTime   *string `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryBatchTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskListRequest) GetBeginCreateTime() *string {
	return s.BeginCreateTime
}

func (s *QueryBatchTaskListRequest) GetEndCreateTime() *string {
	return s.EndCreateTime
}

func (s *QueryBatchTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryBatchTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryBatchTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBatchTaskListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryBatchTaskListRequest) SetBeginCreateTime(v string) *QueryBatchTaskListRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *QueryBatchTaskListRequest) SetEndCreateTime(v string) *QueryBatchTaskListRequest {
	s.EndCreateTime = &v
	return s
}

func (s *QueryBatchTaskListRequest) SetLang(v string) *QueryBatchTaskListRequest {
	s.Lang = &v
	return s
}

func (s *QueryBatchTaskListRequest) SetPageNum(v int32) *QueryBatchTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryBatchTaskListRequest) SetPageSize(v int32) *QueryBatchTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTaskListRequest) SetUserClientIp(v string) *QueryBatchTaskListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryBatchTaskListRequest) Validate() error {
	return dara.Validate(s)
}
