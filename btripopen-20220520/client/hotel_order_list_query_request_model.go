// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllApply(v bool) *HotelOrderListQueryRequest
	GetAllApply() *bool
	SetApplyId(v int64) *HotelOrderListQueryRequest
	GetApplyId() *int64
	SetCategory(v int32) *HotelOrderListQueryRequest
	GetCategory() *int32
	SetDepartId(v string) *HotelOrderListQueryRequest
	GetDepartId() *string
	SetEndTime(v string) *HotelOrderListQueryRequest
	GetEndTime() *string
	SetPage(v int32) *HotelOrderListQueryRequest
	GetPage() *int32
	SetPageSize(v int32) *HotelOrderListQueryRequest
	GetPageSize() *int32
	SetStartTime(v string) *HotelOrderListQueryRequest
	GetStartTime() *string
	SetThirdpartApplyId(v string) *HotelOrderListQueryRequest
	GetThirdpartApplyId() *string
	SetUpdateEndTime(v string) *HotelOrderListQueryRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *HotelOrderListQueryRequest
	GetUpdateStartTime() *string
	SetUserId(v string) *HotelOrderListQueryRequest
	GetUserId() *string
}

type HotelOrderListQueryRequest struct {
	// example:
	//
	// false
	AllApply *bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// example:
	//
	// 165782
	ApplyId  *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	Category *int32  `json:"category,omitempty" xml:"category,omitempty"`
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
	// 50
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// CS154JKOI
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

func (s HotelOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryRequest) GetAllApply() *bool {
	return s.AllApply
}

func (s *HotelOrderListQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *HotelOrderListQueryRequest) GetCategory() *int32 {
	return s.Category
}

func (s *HotelOrderListQueryRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *HotelOrderListQueryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *HotelOrderListQueryRequest) GetPage() *int32 {
	return s.Page
}

func (s *HotelOrderListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelOrderListQueryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *HotelOrderListQueryRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *HotelOrderListQueryRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *HotelOrderListQueryRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *HotelOrderListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderListQueryRequest) SetAllApply(v bool) *HotelOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetApplyId(v int64) *HotelOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetCategory(v int32) *HotelOrderListQueryRequest {
	s.Category = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetDepartId(v string) *HotelOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetEndTime(v string) *HotelOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetPage(v int32) *HotelOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetPageSize(v int32) *HotelOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetStartTime(v string) *HotelOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetThirdpartApplyId(v string) *HotelOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUpdateEndTime(v string) *HotelOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUpdateStartTime(v string) *HotelOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *HotelOrderListQueryRequest) SetUserId(v string) *HotelOrderListQueryRequest {
	s.UserId = &v
	return s
}

func (s *HotelOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
