// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralBillPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetGeneralBills(v []*GeneralBill) *GeneralBillPageResult
	GetGeneralBills() []*GeneralBill
	SetPageNumber(v int32) *GeneralBillPageResult
	GetPageNumber() *int32
	SetPageSize(v int32) *GeneralBillPageResult
	GetPageSize() *int32
	SetRequestId(v string) *GeneralBillPageResult
	GetRequestId() *string
	SetTotal(v int32) *GeneralBillPageResult
	GetTotal() *int32
}

type GeneralBillPageResult struct {
	// The list of bills.
	GeneralBills []*GeneralBill `json:"generalBills,omitempty" xml:"generalBills,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the API request.
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total count.
	//
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GeneralBillPageResult) String() string {
	return dara.Prettify(s)
}

func (s GeneralBillPageResult) GoString() string {
	return s.String()
}

func (s *GeneralBillPageResult) GetGeneralBills() []*GeneralBill {
	return s.GeneralBills
}

func (s *GeneralBillPageResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GeneralBillPageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GeneralBillPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GeneralBillPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *GeneralBillPageResult) SetGeneralBills(v []*GeneralBill) *GeneralBillPageResult {
	s.GeneralBills = v
	return s
}

func (s *GeneralBillPageResult) SetPageNumber(v int32) *GeneralBillPageResult {
	s.PageNumber = &v
	return s
}

func (s *GeneralBillPageResult) SetPageSize(v int32) *GeneralBillPageResult {
	s.PageSize = &v
	return s
}

func (s *GeneralBillPageResult) SetRequestId(v string) *GeneralBillPageResult {
	s.RequestId = &v
	return s
}

func (s *GeneralBillPageResult) SetTotal(v int32) *GeneralBillPageResult {
	s.Total = &v
	return s
}

func (s *GeneralBillPageResult) Validate() error {
	if s.GeneralBills != nil {
		for _, item := range s.GeneralBills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
