// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckThirdRightSendPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizGroup(v string) *CheckThirdRightSendPlanShrinkRequest
	GetBizGroup() *string
	SetBizType(v string) *CheckThirdRightSendPlanShrinkRequest
	GetBizType() *string
	SetExtendInfoShrink(v string) *CheckThirdRightSendPlanShrinkRequest
	GetExtendInfoShrink() *string
	SetSn(v string) *CheckThirdRightSendPlanShrinkRequest
	GetSn() *string
	SetSupplierId(v int64) *CheckThirdRightSendPlanShrinkRequest
	GetSupplierId() *int64
}

type CheckThirdRightSendPlanShrinkRequest struct {
	// example:
	//
	// cc
	BizGroup *string `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	// example:
	//
	// ailabs
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// {}
	ExtendInfoShrink *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 01000019100307010000600
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 1
	SupplierId *int64 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CheckThirdRightSendPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanShrinkRequest) GetBizGroup() *string {
	return s.BizGroup
}

func (s *CheckThirdRightSendPlanShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *CheckThirdRightSendPlanShrinkRequest) GetExtendInfoShrink() *string {
	return s.ExtendInfoShrink
}

func (s *CheckThirdRightSendPlanShrinkRequest) GetSn() *string {
	return s.Sn
}

func (s *CheckThirdRightSendPlanShrinkRequest) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetBizGroup(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.BizGroup = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetBizType(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetExtendInfoShrink(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.ExtendInfoShrink = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetSn(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.Sn = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetSupplierId(v int64) *CheckThirdRightSendPlanShrinkRequest {
	s.SupplierId = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
