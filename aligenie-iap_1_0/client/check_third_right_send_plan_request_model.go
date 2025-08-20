// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckThirdRightSendPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizGroup(v string) *CheckThirdRightSendPlanRequest
	GetBizGroup() *string
	SetBizType(v string) *CheckThirdRightSendPlanRequest
	GetBizType() *string
	SetExtendInfo(v map[string]interface{}) *CheckThirdRightSendPlanRequest
	GetExtendInfo() map[string]interface{}
	SetSn(v string) *CheckThirdRightSendPlanRequest
	GetSn() *string
	SetSupplierId(v int64) *CheckThirdRightSendPlanRequest
	GetSupplierId() *int64
}

type CheckThirdRightSendPlanRequest struct {
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
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 01000019100307010000600
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 1
	SupplierId *int64 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CheckThirdRightSendPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanRequest) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanRequest) GetBizGroup() *string {
	return s.BizGroup
}

func (s *CheckThirdRightSendPlanRequest) GetBizType() *string {
	return s.BizType
}

func (s *CheckThirdRightSendPlanRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *CheckThirdRightSendPlanRequest) GetSn() *string {
	return s.Sn
}

func (s *CheckThirdRightSendPlanRequest) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *CheckThirdRightSendPlanRequest) SetBizGroup(v string) *CheckThirdRightSendPlanRequest {
	s.BizGroup = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetBizType(v string) *CheckThirdRightSendPlanRequest {
	s.BizType = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetExtendInfo(v map[string]interface{}) *CheckThirdRightSendPlanRequest {
	s.ExtendInfo = v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetSn(v string) *CheckThirdRightSendPlanRequest {
	s.Sn = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetSupplierId(v int64) *CheckThirdRightSendPlanRequest {
	s.SupplierId = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) Validate() error {
	return dara.Validate(s)
}
