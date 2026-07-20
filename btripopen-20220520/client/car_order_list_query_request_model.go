// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllApply(v bool) *CarOrderListQueryRequest
	GetAllApply() *bool
	SetApplyId(v int64) *CarOrderListQueryRequest
	GetApplyId() *int64
	SetDepartId(v string) *CarOrderListQueryRequest
	GetDepartId() *string
	SetEndTime(v string) *CarOrderListQueryRequest
	GetEndTime() *string
	SetPage(v int32) *CarOrderListQueryRequest
	GetPage() *int32
	SetPageSize(v int32) *CarOrderListQueryRequest
	GetPageSize() *int32
	SetStartTime(v string) *CarOrderListQueryRequest
	GetStartTime() *string
	SetThirdpartApplyId(v string) *CarOrderListQueryRequest
	GetThirdpartApplyId() *string
	SetUpdateEndTime(v string) *CarOrderListQueryRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *CarOrderListQueryRequest
	GetUpdateStartTime() *string
	SetUserId(v string) *CarOrderListQueryRequest
	GetUserId() *string
}

type CarOrderListQueryRequest struct {
	// example:
	//
	// false
	AllApply *bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// example:
	//
	// 117429516
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// departId
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
	// cs2NH_n1QTC3R6hB9m-BAQ08221658314273
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateEndTime *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// example:
	//
	// 2022-07-01 00:00:00
	UpdateStartTime *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	// example:
	//
	// userId
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryRequest) GetAllApply() *bool {
	return s.AllApply
}

func (s *CarOrderListQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *CarOrderListQueryRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *CarOrderListQueryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CarOrderListQueryRequest) GetPage() *int32 {
	return s.Page
}

func (s *CarOrderListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CarOrderListQueryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CarOrderListQueryRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *CarOrderListQueryRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *CarOrderListQueryRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *CarOrderListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *CarOrderListQueryRequest) SetAllApply(v bool) *CarOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *CarOrderListQueryRequest) SetApplyId(v int64) *CarOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetDepartId(v string) *CarOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetEndTime(v string) *CarOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetPage(v int32) *CarOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *CarOrderListQueryRequest) SetPageSize(v int32) *CarOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarOrderListQueryRequest) SetStartTime(v string) *CarOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetThirdpartApplyId(v string) *CarOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUpdateEndTime(v string) *CarOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUpdateStartTime(v string) *CarOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *CarOrderListQueryRequest) SetUserId(v string) *CarOrderListQueryRequest {
	s.UserId = &v
	return s
}

func (s *CarOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
