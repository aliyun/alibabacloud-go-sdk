// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllApply(v bool) *FlightOrderListQueryRequest
	GetAllApply() *bool
	SetApplyId(v int64) *FlightOrderListQueryRequest
	GetApplyId() *int64
	SetDepartId(v string) *FlightOrderListQueryRequest
	GetDepartId() *string
	SetEndTime(v string) *FlightOrderListQueryRequest
	GetEndTime() *string
	SetPage(v int32) *FlightOrderListQueryRequest
	GetPage() *int32
	SetPageSize(v int32) *FlightOrderListQueryRequest
	GetPageSize() *int32
	SetStartTime(v string) *FlightOrderListQueryRequest
	GetStartTime() *string
	SetThirdpartApplyId(v string) *FlightOrderListQueryRequest
	GetThirdpartApplyId() *string
	SetUpdateEndTime(v string) *FlightOrderListQueryRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *FlightOrderListQueryRequest
	GetUpdateStartTime() *string
	SetUserId(v string) *FlightOrderListQueryRequest
	GetUserId() *string
}

type FlightOrderListQueryRequest struct {
	// example:
	//
	// false
	AllApply *bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// example:
	//
	// 175634
	ApplyId  *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// CS-FLIGHT
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateEndTime *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateStartTime *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId          *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryRequest) GetAllApply() *bool {
	return s.AllApply
}

func (s *FlightOrderListQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *FlightOrderListQueryRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *FlightOrderListQueryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *FlightOrderListQueryRequest) GetPage() *int32 {
	return s.Page
}

func (s *FlightOrderListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightOrderListQueryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *FlightOrderListQueryRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *FlightOrderListQueryRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *FlightOrderListQueryRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *FlightOrderListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryRequest) SetAllApply(v bool) *FlightOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetApplyId(v int64) *FlightOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetDepartId(v string) *FlightOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetEndTime(v string) *FlightOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetPage(v int32) *FlightOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetPageSize(v int32) *FlightOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetStartTime(v string) *FlightOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetThirdpartApplyId(v string) *FlightOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUpdateEndTime(v string) *FlightOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUpdateStartTime(v string) *FlightOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *FlightOrderListQueryRequest) SetUserId(v string) *FlightOrderListQueryRequest {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
