// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplySyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int64) *CommonApplySyncRequest
	GetApplyId() *int64
	SetBizCategory(v int32) *CommonApplySyncRequest
	GetBizCategory() *int32
	SetRemark(v string) *CommonApplySyncRequest
	GetRemark() *string
	SetStatus(v int32) *CommonApplySyncRequest
	GetStatus() *int32
	SetThirdpartyFlowId(v string) *CommonApplySyncRequest
	GetThirdpartyFlowId() *string
	SetUserId(v string) *CommonApplySyncRequest
	GetUserId() *string
}

type CommonApplySyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1003366164
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 3
	BizCategory *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// asfa
	ThirdpartyFlowId *string `json:"thirdparty_flow_id,omitempty" xml:"thirdparty_flow_id,omitempty"`
	// This parameter is required.
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CommonApplySyncRequest) String() string {
	return dara.Prettify(s)
}

func (s CommonApplySyncRequest) GoString() string {
	return s.String()
}

func (s *CommonApplySyncRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *CommonApplySyncRequest) GetBizCategory() *int32 {
	return s.BizCategory
}

func (s *CommonApplySyncRequest) GetRemark() *string {
	return s.Remark
}

func (s *CommonApplySyncRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CommonApplySyncRequest) GetThirdpartyFlowId() *string {
	return s.ThirdpartyFlowId
}

func (s *CommonApplySyncRequest) GetUserId() *string {
	return s.UserId
}

func (s *CommonApplySyncRequest) SetApplyId(v int64) *CommonApplySyncRequest {
	s.ApplyId = &v
	return s
}

func (s *CommonApplySyncRequest) SetBizCategory(v int32) *CommonApplySyncRequest {
	s.BizCategory = &v
	return s
}

func (s *CommonApplySyncRequest) SetRemark(v string) *CommonApplySyncRequest {
	s.Remark = &v
	return s
}

func (s *CommonApplySyncRequest) SetStatus(v int32) *CommonApplySyncRequest {
	s.Status = &v
	return s
}

func (s *CommonApplySyncRequest) SetThirdpartyFlowId(v string) *CommonApplySyncRequest {
	s.ThirdpartyFlowId = &v
	return s
}

func (s *CommonApplySyncRequest) SetUserId(v string) *CommonApplySyncRequest {
	s.UserId = &v
	return s
}

func (s *CommonApplySyncRequest) Validate() error {
	return dara.Validate(s)
}
