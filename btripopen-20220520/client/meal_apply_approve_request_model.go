// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyApproveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperateTime(v string) *MealApplyApproveRequest
	GetOperateTime() *string
	SetRemark(v string) *MealApplyApproveRequest
	GetRemark() *string
	SetStatus(v int32) *MealApplyApproveRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *MealApplyApproveRequest
	GetThirdPartApplyId() *string
	SetUserId(v string) *MealApplyApproveRequest
	GetUserId() *string
}

type MealApplyApproveRequest struct {
	// example:
	//
	// 2022-07-12 16:12:53
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62141
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealApplyApproveRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyApproveRequest) GoString() string {
	return s.String()
}

func (s *MealApplyApproveRequest) GetOperateTime() *string {
	return s.OperateTime
}

func (s *MealApplyApproveRequest) GetRemark() *string {
	return s.Remark
}

func (s *MealApplyApproveRequest) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyApproveRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyApproveRequest) GetUserId() *string {
	return s.UserId
}

func (s *MealApplyApproveRequest) SetOperateTime(v string) *MealApplyApproveRequest {
	s.OperateTime = &v
	return s
}

func (s *MealApplyApproveRequest) SetRemark(v string) *MealApplyApproveRequest {
	s.Remark = &v
	return s
}

func (s *MealApplyApproveRequest) SetStatus(v int32) *MealApplyApproveRequest {
	s.Status = &v
	return s
}

func (s *MealApplyApproveRequest) SetThirdPartApplyId(v string) *MealApplyApproveRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyApproveRequest) SetUserId(v string) *MealApplyApproveRequest {
	s.UserId = &v
	return s
}

func (s *MealApplyApproveRequest) Validate() error {
	return dara.Validate(s)
}
