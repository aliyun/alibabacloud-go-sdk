// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSdlEventListRequest
	GetCurrentPage() *int32
	SetDstIp(v string) *DescribeSdlEventListRequest
	GetDstIp() *string
	SetEndTime(v int64) *DescribeSdlEventListRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlEventListRequest
	GetLang() *string
	SetLocation(v string) *DescribeSdlEventListRequest
	GetLocation() *string
	SetOnlyAiEvt(v int32) *DescribeSdlEventListRequest
	GetOnlyAiEvt() *int32
	SetOrder(v string) *DescribeSdlEventListRequest
	GetOrder() *string
	SetPageSize(v int32) *DescribeSdlEventListRequest
	GetPageSize() *int32
	SetSensitiveLevel(v string) *DescribeSdlEventListRequest
	GetSensitiveLevel() *string
	SetSort(v string) *DescribeSdlEventListRequest
	GetSort() *string
	SetSrcIp(v string) *DescribeSdlEventListRequest
	GetSrcIp() *string
	SetStartTime(v int64) *DescribeSdlEventListRequest
	GetStartTime() *int64
	SetUuid(v string) *DescribeSdlEventListRequest
	GetUuid() *string
}

type DescribeSdlEventListRequest struct {
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
	// 1756433077
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 0
	OnlyAiEvt *int32 `json:"OnlyAiEvt,omitempty" xml:"OnlyAiEvt,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// example:
	//
	// TotalBytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 172.16.0.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 1759198702
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// aa58cdf6-6cf8-493c-912d-97619a24****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSdlEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSdlEventListRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeSdlEventListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlEventListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlEventListRequest) GetLocation() *string {
	return s.Location
}

func (s *DescribeSdlEventListRequest) GetOnlyAiEvt() *int32 {
	return s.OnlyAiEvt
}

func (s *DescribeSdlEventListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeSdlEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSdlEventListRequest) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSdlEventListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeSdlEventListRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlEventListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlEventListRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSdlEventListRequest) SetCurrentPage(v int32) *DescribeSdlEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetDstIp(v string) *DescribeSdlEventListRequest {
	s.DstIp = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetEndTime(v int64) *DescribeSdlEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetLang(v string) *DescribeSdlEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetLocation(v string) *DescribeSdlEventListRequest {
	s.Location = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetOnlyAiEvt(v int32) *DescribeSdlEventListRequest {
	s.OnlyAiEvt = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetOrder(v string) *DescribeSdlEventListRequest {
	s.Order = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetPageSize(v int32) *DescribeSdlEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetSensitiveLevel(v string) *DescribeSdlEventListRequest {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetSort(v string) *DescribeSdlEventListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetSrcIp(v string) *DescribeSdlEventListRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetStartTime(v int64) *DescribeSdlEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlEventListRequest) SetUuid(v string) *DescribeSdlEventListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeSdlEventListRequest) Validate() error {
	return dara.Validate(s)
}
