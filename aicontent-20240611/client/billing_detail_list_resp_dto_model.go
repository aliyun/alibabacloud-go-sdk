// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingDetailListRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*MetricDefRespDTO) *BillingDetailListRespDTO
	GetColumns() []*MetricDefRespDTO
	SetPage(v int32) *BillingDetailListRespDTO
	GetPage() *int32
	SetPageSize(v int32) *BillingDetailListRespDTO
	GetPageSize() *int32
	SetRows(v []*BillingDetailRowDTO) *BillingDetailListRespDTO
	GetRows() []*BillingDetailRowDTO
	SetTotal(v int64) *BillingDetailListRespDTO
	GetTotal() *int64
}

type BillingDetailListRespDTO struct {
	// The column definitions, which are dynamically generated based on the model types that appear in the query results.
	//
	// example:
	//
	// []
	Columns []*MetricDefRespDTO `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of bill details data.
	//
	// example:
	//
	// []
	Rows []*BillingDetailRowDTO `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s BillingDetailListRespDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingDetailListRespDTO) GoString() string {
	return s.String()
}

func (s *BillingDetailListRespDTO) GetColumns() []*MetricDefRespDTO {
	return s.Columns
}

func (s *BillingDetailListRespDTO) GetPage() *int32 {
	return s.Page
}

func (s *BillingDetailListRespDTO) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BillingDetailListRespDTO) GetRows() []*BillingDetailRowDTO {
	return s.Rows
}

func (s *BillingDetailListRespDTO) GetTotal() *int64 {
	return s.Total
}

func (s *BillingDetailListRespDTO) SetColumns(v []*MetricDefRespDTO) *BillingDetailListRespDTO {
	s.Columns = v
	return s
}

func (s *BillingDetailListRespDTO) SetPage(v int32) *BillingDetailListRespDTO {
	s.Page = &v
	return s
}

func (s *BillingDetailListRespDTO) SetPageSize(v int32) *BillingDetailListRespDTO {
	s.PageSize = &v
	return s
}

func (s *BillingDetailListRespDTO) SetRows(v []*BillingDetailRowDTO) *BillingDetailListRespDTO {
	s.Rows = v
	return s
}

func (s *BillingDetailListRespDTO) SetTotal(v int64) *BillingDetailListRespDTO {
	s.Total = &v
	return s
}

func (s *BillingDetailListRespDTO) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
