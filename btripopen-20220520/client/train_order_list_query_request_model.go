// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllApply(v bool) *TrainOrderListQueryRequest
	GetAllApply() *bool
	SetApplyId(v int64) *TrainOrderListQueryRequest
	GetApplyId() *int64
	SetDepartId(v string) *TrainOrderListQueryRequest
	GetDepartId() *string
	SetEndTime(v string) *TrainOrderListQueryRequest
	GetEndTime() *string
	SetPage(v int32) *TrainOrderListQueryRequest
	GetPage() *int32
	SetPageSize(v int32) *TrainOrderListQueryRequest
	GetPageSize() *int32
	SetStartTime(v string) *TrainOrderListQueryRequest
	GetStartTime() *string
	SetThirdpartApplyId(v string) *TrainOrderListQueryRequest
	GetThirdpartApplyId() *string
	SetUpdateEndTime(v string) *TrainOrderListQueryRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *TrainOrderListQueryRequest
	GetUpdateStartTime() *string
	SetUserId(v string) *TrainOrderListQueryRequest
	GetUserId() *string
}

type TrainOrderListQueryRequest struct {
	// example:
	//
	// false
	AllApply *bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// example:
	//
	// 11657
	ApplyId  *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	DepartId *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 3
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 25
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// CS-EDES9898
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	UpdateEndTime *string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	UpdateStartTime *string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	UserId          *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryRequest) GetAllApply() *bool {
	return s.AllApply
}

func (s *TrainOrderListQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *TrainOrderListQueryRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderListQueryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderListQueryRequest) GetPage() *int32 {
	return s.Page
}

func (s *TrainOrderListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TrainOrderListQueryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderListQueryRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *TrainOrderListQueryRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *TrainOrderListQueryRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *TrainOrderListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderListQueryRequest) SetAllApply(v bool) *TrainOrderListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetApplyId(v int64) *TrainOrderListQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetDepartId(v string) *TrainOrderListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetEndTime(v string) *TrainOrderListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetPage(v int32) *TrainOrderListQueryRequest {
	s.Page = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetPageSize(v int32) *TrainOrderListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetStartTime(v string) *TrainOrderListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetThirdpartApplyId(v string) *TrainOrderListQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUpdateEndTime(v string) *TrainOrderListQueryRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUpdateStartTime(v string) *TrainOrderListQueryRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *TrainOrderListQueryRequest) SetUserId(v string) *TrainOrderListQueryRequest {
	s.UserId = &v
	return s
}

func (s *TrainOrderListQueryRequest) Validate() error {
	return dara.Validate(s)
}
