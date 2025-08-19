// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *CreateVerifySettingRequest
	GetBizName() *string
	SetBizType(v string) *CreateVerifySettingRequest
	GetBizType() *string
	SetGuideStep(v bool) *CreateVerifySettingRequest
	GetGuideStep() *bool
	SetPrivacyStep(v bool) *CreateVerifySettingRequest
	GetPrivacyStep() *bool
	SetResultStep(v bool) *CreateVerifySettingRequest
	GetResultStep() *bool
	SetSolution(v string) *CreateVerifySettingRequest
	GetSolution() *string
}

type CreateVerifySettingRequest struct {
	// This parameter is required.
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UserRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// false
	GuideStep *bool `json:"GuideStep,omitempty" xml:"GuideStep,omitempty"`
	// example:
	//
	// true
	PrivacyStep *bool `json:"PrivacyStep,omitempty" xml:"PrivacyStep,omitempty"`
	// example:
	//
	// false
	ResultStep *bool `json:"ResultStep,omitempty" xml:"ResultStep,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RPBasic
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s CreateVerifySettingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySettingRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateVerifySettingRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateVerifySettingRequest) GetGuideStep() *bool {
	return s.GuideStep
}

func (s *CreateVerifySettingRequest) GetPrivacyStep() *bool {
	return s.PrivacyStep
}

func (s *CreateVerifySettingRequest) GetResultStep() *bool {
	return s.ResultStep
}

func (s *CreateVerifySettingRequest) GetSolution() *string {
	return s.Solution
}

func (s *CreateVerifySettingRequest) SetBizName(v string) *CreateVerifySettingRequest {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingRequest) SetBizType(v string) *CreateVerifySettingRequest {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingRequest) SetGuideStep(v bool) *CreateVerifySettingRequest {
	s.GuideStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetPrivacyStep(v bool) *CreateVerifySettingRequest {
	s.PrivacyStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetResultStep(v bool) *CreateVerifySettingRequest {
	s.ResultStep = &v
	return s
}

func (s *CreateVerifySettingRequest) SetSolution(v string) *CreateVerifySettingRequest {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingRequest) Validate() error {
	return dara.Validate(s)
}
