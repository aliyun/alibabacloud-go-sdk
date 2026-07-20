// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApproveIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest
	GetApproveIdShrink() *string
	SetBookerIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest
	GetBookerIdShrink() *string
	SetDepartIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest
	GetDepartIdShrink() *string
	SetEndDate(v string) *FlightOrderListQueryV2ShrinkRequest
	GetEndDate() *string
	SetPageSize(v int32) *FlightOrderListQueryV2ShrinkRequest
	GetPageSize() *int32
	SetScrollId(v string) *FlightOrderListQueryV2ShrinkRequest
	GetScrollId() *string
	SetStartDate(v string) *FlightOrderListQueryV2ShrinkRequest
	GetStartDate() *string
	SetSupplierShrink(v string) *FlightOrderListQueryV2ShrinkRequest
	GetSupplierShrink() *string
	SetThirdpartApproveIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest
	GetThirdpartApproveIdShrink() *string
	SetUpdateEndDate(v string) *FlightOrderListQueryV2ShrinkRequest
	GetUpdateEndDate() *string
	SetUpdateStartDate(v string) *FlightOrderListQueryV2ShrinkRequest
	GetUpdateStartDate() *string
}

type FlightOrderListQueryV2ShrinkRequest struct {
	ApproveIdShrink *string `json:"approve_id,omitempty" xml:"approve_id,omitempty"`
	BookerIdShrink  *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	DepartIdShrink  *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_Size,omitempty" xml:"page_Size,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4ZGViODFlYmM3MYzM4
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	StartDate                *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	SupplierShrink           *string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	ThirdpartApproveIdShrink *string `json:"thirdpart_approve_id,omitempty" xml:"thirdpart_approve_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateEndDate *string `json:"update_end_date,omitempty" xml:"update_end_date,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateStartDate *string `json:"update_start_date,omitempty" xml:"update_start_date,omitempty"`
}

func (s FlightOrderListQueryV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetApproveIdShrink() *string {
	return s.ApproveIdShrink
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetBookerIdShrink() *string {
	return s.BookerIdShrink
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetDepartIdShrink() *string {
	return s.DepartIdShrink
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetSupplierShrink() *string {
	return s.SupplierShrink
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetThirdpartApproveIdShrink() *string {
	return s.ThirdpartApproveIdShrink
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetUpdateEndDate() *string {
	return s.UpdateEndDate
}

func (s *FlightOrderListQueryV2ShrinkRequest) GetUpdateStartDate() *string {
	return s.UpdateStartDate
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetApproveIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.ApproveIdShrink = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetBookerIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.BookerIdShrink = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetDepartIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.DepartIdShrink = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetEndDate(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetPageSize(v int32) *FlightOrderListQueryV2ShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetScrollId(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.ScrollId = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetStartDate(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetSupplierShrink(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.SupplierShrink = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetThirdpartApproveIdShrink(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.ThirdpartApproveIdShrink = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetUpdateEndDate(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.UpdateEndDate = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) SetUpdateStartDate(v string) *FlightOrderListQueryV2ShrinkRequest {
	s.UpdateStartDate = &v
	return s
}

func (s *FlightOrderListQueryV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
