// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackToBackCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *StartBackToBackCallRequest
	GetBizId() *string
	SetBizType(v string) *StartBackToBackCallRequest
	GetBizType() *string
	SetCallCenterNumber(v string) *StartBackToBackCallRequest
	GetCallCenterNumber() *string
	SetCaller(v string) *StartBackToBackCallRequest
	GetCaller() *string
	SetMobileKey(v string) *StartBackToBackCallRequest
	GetMobileKey() *string
	SetSkillType(v int64) *StartBackToBackCallRequest
	GetSkillType() *int64
}

type StartBackToBackCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20211203180209000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 02131584184
	CallCenterNumber *string `json:"CallCenterNumber,omitempty" xml:"CallCenterNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13162828888
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// mobile1
	MobileKey *string `json:"MobileKey,omitempty" xml:"MobileKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SkillType *int64 `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s StartBackToBackCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBackToBackCallRequest) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallRequest) GetBizId() *string {
	return s.BizId
}

func (s *StartBackToBackCallRequest) GetBizType() *string {
	return s.BizType
}

func (s *StartBackToBackCallRequest) GetCallCenterNumber() *string {
	return s.CallCenterNumber
}

func (s *StartBackToBackCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *StartBackToBackCallRequest) GetMobileKey() *string {
	return s.MobileKey
}

func (s *StartBackToBackCallRequest) GetSkillType() *int64 {
	return s.SkillType
}

func (s *StartBackToBackCallRequest) SetBizId(v string) *StartBackToBackCallRequest {
	s.BizId = &v
	return s
}

func (s *StartBackToBackCallRequest) SetBizType(v string) *StartBackToBackCallRequest {
	s.BizType = &v
	return s
}

func (s *StartBackToBackCallRequest) SetCallCenterNumber(v string) *StartBackToBackCallRequest {
	s.CallCenterNumber = &v
	return s
}

func (s *StartBackToBackCallRequest) SetCaller(v string) *StartBackToBackCallRequest {
	s.Caller = &v
	return s
}

func (s *StartBackToBackCallRequest) SetMobileKey(v string) *StartBackToBackCallRequest {
	s.MobileKey = &v
	return s
}

func (s *StartBackToBackCallRequest) SetSkillType(v int64) *StartBackToBackCallRequest {
	s.SkillType = &v
	return s
}

func (s *StartBackToBackCallRequest) Validate() error {
	return dara.Validate(s)
}
