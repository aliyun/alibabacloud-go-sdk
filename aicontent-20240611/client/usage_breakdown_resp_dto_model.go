// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageBreakdownRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetGranularity(v string) *UsageBreakdownRespDTO
	GetGranularity() *string
	SetPage(v int32) *UsageBreakdownRespDTO
	GetPage() *int32
	SetPageSize(v int32) *UsageBreakdownRespDTO
	GetPageSize() *int32
	SetRows(v []*UsageBreakdownRowDTO) *UsageBreakdownRespDTO
	GetRows() []*UsageBreakdownRowDTO
	SetTotal(v int64) *UsageBreakdownRespDTO
	GetTotal() *int64
}

type UsageBreakdownRespDTO struct {
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// []
	Rows []*UsageBreakdownRowDTO `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s UsageBreakdownRespDTO) String() string {
	return dara.Prettify(s)
}

func (s UsageBreakdownRespDTO) GoString() string {
	return s.String()
}

func (s *UsageBreakdownRespDTO) GetGranularity() *string {
	return s.Granularity
}

func (s *UsageBreakdownRespDTO) GetPage() *int32 {
	return s.Page
}

func (s *UsageBreakdownRespDTO) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UsageBreakdownRespDTO) GetRows() []*UsageBreakdownRowDTO {
	return s.Rows
}

func (s *UsageBreakdownRespDTO) GetTotal() *int64 {
	return s.Total
}

func (s *UsageBreakdownRespDTO) SetGranularity(v string) *UsageBreakdownRespDTO {
	s.Granularity = &v
	return s
}

func (s *UsageBreakdownRespDTO) SetPage(v int32) *UsageBreakdownRespDTO {
	s.Page = &v
	return s
}

func (s *UsageBreakdownRespDTO) SetPageSize(v int32) *UsageBreakdownRespDTO {
	s.PageSize = &v
	return s
}

func (s *UsageBreakdownRespDTO) SetRows(v []*UsageBreakdownRowDTO) *UsageBreakdownRespDTO {
	s.Rows = v
	return s
}

func (s *UsageBreakdownRespDTO) SetTotal(v int64) *UsageBreakdownRespDTO {
	s.Total = &v
	return s
}

func (s *UsageBreakdownRespDTO) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
