// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *CreateVerifySettingResponseBody
	GetBizName() *string
	SetBizType(v string) *CreateVerifySettingResponseBody
	GetBizType() *string
	SetRequestId(v string) *CreateVerifySettingResponseBody
	GetRequestId() *string
	SetSolution(v string) *CreateVerifySettingResponseBody
	GetSolution() *string
	SetStepList(v []*string) *CreateVerifySettingResponseBody
	GetStepList() []*string
}

type CreateVerifySettingResponseBody struct {
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// UserRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// RPBasic
	Solution *string   `json:"Solution,omitempty" xml:"Solution,omitempty"`
	StepList []*string `json:"StepList,omitempty" xml:"StepList,omitempty" type:"Repeated"`
}

func (s CreateVerifySettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySettingResponseBody) GetBizName() *string {
	return s.BizName
}

func (s *CreateVerifySettingResponseBody) GetBizType() *string {
	return s.BizType
}

func (s *CreateVerifySettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVerifySettingResponseBody) GetSolution() *string {
	return s.Solution
}

func (s *CreateVerifySettingResponseBody) GetStepList() []*string {
	return s.StepList
}

func (s *CreateVerifySettingResponseBody) SetBizName(v string) *CreateVerifySettingResponseBody {
	s.BizName = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetBizType(v string) *CreateVerifySettingResponseBody {
	s.BizType = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetRequestId(v string) *CreateVerifySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetSolution(v string) *CreateVerifySettingResponseBody {
	s.Solution = &v
	return s
}

func (s *CreateVerifySettingResponseBody) SetStepList(v []*string) *CreateVerifySettingResponseBody {
	s.StepList = v
	return s
}

func (s *CreateVerifySettingResponseBody) Validate() error {
	return dara.Validate(s)
}
