// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventSdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSdlEventSdListRequest
	GetCurrentPage() *int32
	SetDstIp(v string) *DescribeSdlEventSdListRequest
	GetDstIp() *string
	SetEndTime(v int64) *DescribeSdlEventSdListRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlEventSdListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeSdlEventSdListRequest
	GetPageSize() *int32
	SetSrcIp(v string) *DescribeSdlEventSdListRequest
	GetSrcIp() *string
	SetStartTime(v int64) *DescribeSdlEventSdListRequest
	GetStartTime() *int64
	SetUuid(v string) *DescribeSdlEventSdListRequest
	GetUuid() *string
}

type DescribeSdlEventSdListRequest struct {
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
	// 1761185080
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
	// 172.16.0.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 1656664560
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// aa58cdf6-6cf8-493c-912d-97619a24****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSdlEventSdListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventSdListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventSdListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSdlEventSdListRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeSdlEventSdListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlEventSdListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlEventSdListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSdlEventSdListRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlEventSdListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventSdListRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSdlEventSdListRequest) SetCurrentPage(v int32) *DescribeSdlEventSdListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetDstIp(v string) *DescribeSdlEventSdListRequest {
	s.DstIp = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetEndTime(v int64) *DescribeSdlEventSdListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetLang(v string) *DescribeSdlEventSdListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetPageSize(v int32) *DescribeSdlEventSdListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetSrcIp(v string) *DescribeSdlEventSdListRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetStartTime(v int64) *DescribeSdlEventSdListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) SetUuid(v string) *DescribeSdlEventSdListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeSdlEventSdListRequest) Validate() error {
	return dara.Validate(s)
}
