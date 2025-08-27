// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperateTime(v string) *CarApplyModifyRequest
	GetOperateTime() *string
	SetRemark(v string) *CarApplyModifyRequest
	GetRemark() *string
	SetStatus(v int32) *CarApplyModifyRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *CarApplyModifyRequest
	GetThirdPartApplyId() *string
	SetUserId(v string) *CarApplyModifyRequest
	GetUserId() *string
}

type CarApplyModifyRequest struct {
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
	// IRGS1413
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open62141
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CarApplyModifyRequest) GoString() string {
	return s.String()
}

func (s *CarApplyModifyRequest) GetOperateTime() *string {
	return s.OperateTime
}

func (s *CarApplyModifyRequest) GetRemark() *string {
	return s.Remark
}

func (s *CarApplyModifyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CarApplyModifyRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *CarApplyModifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyModifyRequest) SetOperateTime(v string) *CarApplyModifyRequest {
	s.OperateTime = &v
	return s
}

func (s *CarApplyModifyRequest) SetRemark(v string) *CarApplyModifyRequest {
	s.Remark = &v
	return s
}

func (s *CarApplyModifyRequest) SetStatus(v int32) *CarApplyModifyRequest {
	s.Status = &v
	return s
}

func (s *CarApplyModifyRequest) SetThirdPartApplyId(v string) *CarApplyModifyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyModifyRequest) SetUserId(v string) *CarApplyModifyRequest {
	s.UserId = &v
	return s
}

func (s *CarApplyModifyRequest) Validate() error {
	return dara.Validate(s)
}
