// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetApproveId(v []*string) *FlightOrderListQueryV2Request
	GetApproveId() []*string
	SetBookerId(v []*string) *FlightOrderListQueryV2Request
	GetBookerId() []*string
	SetDepartId(v []*string) *FlightOrderListQueryV2Request
	GetDepartId() []*string
	SetEndDate(v string) *FlightOrderListQueryV2Request
	GetEndDate() *string
	SetPageSize(v int32) *FlightOrderListQueryV2Request
	GetPageSize() *int32
	SetScrollId(v string) *FlightOrderListQueryV2Request
	GetScrollId() *string
	SetStartDate(v string) *FlightOrderListQueryV2Request
	GetStartDate() *string
	SetSupplier(v []*string) *FlightOrderListQueryV2Request
	GetSupplier() []*string
	SetThirdpartApproveId(v []*string) *FlightOrderListQueryV2Request
	GetThirdpartApproveId() []*string
	SetUpdateEndDate(v string) *FlightOrderListQueryV2Request
	GetUpdateEndDate() *string
	SetUpdateStartDate(v string) *FlightOrderListQueryV2Request
	GetUpdateStartDate() *string
}

type FlightOrderListQueryV2Request struct {
	ApproveId []*string `json:"approve_id,omitempty" xml:"approve_id,omitempty" type:"Repeated"`
	BookerId  []*string `json:"booker_id,omitempty" xml:"booker_id,omitempty" type:"Repeated"`
	DepartId  []*string `json:"depart_id,omitempty" xml:"depart_id,omitempty" type:"Repeated"`
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
	StartDate          *string   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	Supplier           []*string `json:"supplier,omitempty" xml:"supplier,omitempty" type:"Repeated"`
	ThirdpartApproveId []*string `json:"thirdpart_approve_id,omitempty" xml:"thirdpart_approve_id,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateEndDate *string `json:"update_end_date,omitempty" xml:"update_end_date,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateStartDate *string `json:"update_start_date,omitempty" xml:"update_start_date,omitempty"`
}

func (s FlightOrderListQueryV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2Request) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2Request) GetApproveId() []*string {
	return s.ApproveId
}

func (s *FlightOrderListQueryV2Request) GetBookerId() []*string {
	return s.BookerId
}

func (s *FlightOrderListQueryV2Request) GetDepartId() []*string {
	return s.DepartId
}

func (s *FlightOrderListQueryV2Request) GetEndDate() *string {
	return s.EndDate
}

func (s *FlightOrderListQueryV2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightOrderListQueryV2Request) GetScrollId() *string {
	return s.ScrollId
}

func (s *FlightOrderListQueryV2Request) GetStartDate() *string {
	return s.StartDate
}

func (s *FlightOrderListQueryV2Request) GetSupplier() []*string {
	return s.Supplier
}

func (s *FlightOrderListQueryV2Request) GetThirdpartApproveId() []*string {
	return s.ThirdpartApproveId
}

func (s *FlightOrderListQueryV2Request) GetUpdateEndDate() *string {
	return s.UpdateEndDate
}

func (s *FlightOrderListQueryV2Request) GetUpdateStartDate() *string {
	return s.UpdateStartDate
}

func (s *FlightOrderListQueryV2Request) SetApproveId(v []*string) *FlightOrderListQueryV2Request {
	s.ApproveId = v
	return s
}

func (s *FlightOrderListQueryV2Request) SetBookerId(v []*string) *FlightOrderListQueryV2Request {
	s.BookerId = v
	return s
}

func (s *FlightOrderListQueryV2Request) SetDepartId(v []*string) *FlightOrderListQueryV2Request {
	s.DepartId = v
	return s
}

func (s *FlightOrderListQueryV2Request) SetEndDate(v string) *FlightOrderListQueryV2Request {
	s.EndDate = &v
	return s
}

func (s *FlightOrderListQueryV2Request) SetPageSize(v int32) *FlightOrderListQueryV2Request {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryV2Request) SetScrollId(v string) *FlightOrderListQueryV2Request {
	s.ScrollId = &v
	return s
}

func (s *FlightOrderListQueryV2Request) SetStartDate(v string) *FlightOrderListQueryV2Request {
	s.StartDate = &v
	return s
}

func (s *FlightOrderListQueryV2Request) SetSupplier(v []*string) *FlightOrderListQueryV2Request {
	s.Supplier = v
	return s
}

func (s *FlightOrderListQueryV2Request) SetThirdpartApproveId(v []*string) *FlightOrderListQueryV2Request {
	s.ThirdpartApproveId = v
	return s
}

func (s *FlightOrderListQueryV2Request) SetUpdateEndDate(v string) *FlightOrderListQueryV2Request {
	s.UpdateEndDate = &v
	return s
}

func (s *FlightOrderListQueryV2Request) SetUpdateStartDate(v string) *FlightOrderListQueryV2Request {
	s.UpdateStartDate = &v
	return s
}

func (s *FlightOrderListQueryV2Request) Validate() error {
	return dara.Validate(s)
}
