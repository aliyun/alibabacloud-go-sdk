// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChangeLogListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryChangeLogListRequest
	GetDomainName() *string
	SetEndDate(v int64) *QueryChangeLogListRequest
	GetEndDate() *int64
	SetLang(v string) *QueryChangeLogListRequest
	GetLang() *string
	SetPageNum(v int32) *QueryChangeLogListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryChangeLogListRequest
	GetPageSize() *int32
	SetStartDate(v int64) *QueryChangeLogListRequest
	GetStartDate() *int64
	SetUserClientIp(v string) *QueryChangeLogListRequest
	GetUserClientIp() *string
}

type QueryChangeLogListRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1522080000000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1522080000000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryChangeLogListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChangeLogListRequest) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryChangeLogListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *QueryChangeLogListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryChangeLogListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryChangeLogListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryChangeLogListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *QueryChangeLogListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryChangeLogListRequest) SetDomainName(v string) *QueryChangeLogListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryChangeLogListRequest) SetEndDate(v int64) *QueryChangeLogListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryChangeLogListRequest) SetLang(v string) *QueryChangeLogListRequest {
	s.Lang = &v
	return s
}

func (s *QueryChangeLogListRequest) SetPageNum(v int32) *QueryChangeLogListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryChangeLogListRequest) SetPageSize(v int32) *QueryChangeLogListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryChangeLogListRequest) SetStartDate(v int64) *QueryChangeLogListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryChangeLogListRequest) SetUserClientIp(v string) *QueryChangeLogListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryChangeLogListRequest) Validate() error {
	return dara.Validate(s)
}
