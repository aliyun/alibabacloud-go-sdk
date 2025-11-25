// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSdlEventDetailRequest
	GetCurrentPage() *int32
	SetDstIp(v string) *DescribeSdlEventDetailRequest
	GetDstIp() *string
	SetEndTime(v int64) *DescribeSdlEventDetailRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlEventDetailRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeSdlEventDetailRequest
	GetPageSize() *int32
	SetSrcIp(v string) *DescribeSdlEventDetailRequest
	GetSrcIp() *string
	SetStartTime(v int64) *DescribeSdlEventDetailRequest
	GetStartTime() *int64
	SetUuid(v string) *DescribeSdlEventDetailRequest
	GetUuid() *string
}

type DescribeSdlEventDetailRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 47.100.102.XXX
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// example:
	//
	// 1753755251
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 121.40.84.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 1656664560
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 326ce10e-5e17-4235-879a-6f2502cd****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSdlEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSdlEventDetailRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeSdlEventDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlEventDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSdlEventDetailRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlEventDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSdlEventDetailRequest) SetCurrentPage(v int32) *DescribeSdlEventDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetDstIp(v string) *DescribeSdlEventDetailRequest {
	s.DstIp = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetEndTime(v int64) *DescribeSdlEventDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetLang(v string) *DescribeSdlEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetPageSize(v int32) *DescribeSdlEventDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetSrcIp(v string) *DescribeSdlEventDetailRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetStartTime(v int64) *DescribeSdlEventDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) SetUuid(v string) *DescribeSdlEventDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeSdlEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
