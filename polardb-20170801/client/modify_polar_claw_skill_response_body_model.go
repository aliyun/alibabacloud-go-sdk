// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyPolarClawSkillResponseBody
	GetApplicationId() *string
	SetEvaluateStatus(v string) *ModifyPolarClawSkillResponseBody
	GetEvaluateStatus() *string
	SetRequestId(v string) *ModifyPolarClawSkillResponseBody
	GetRequestId() *string
	SetRiskDetail(v string) *ModifyPolarClawSkillResponseBody
	GetRiskDetail() *string
}

type ModifyPolarClawSkillResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The evaluation status.
	//
	// example:
	//
	// 2
	EvaluateStatus *string `json:"EvaluateStatus,omitempty" xml:"EvaluateStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The baseline risk details.
	//
	// example:
	//
	// riskDetail
	RiskDetail *string `json:"RiskDetail,omitempty" xml:"RiskDetail,omitempty"`
}

func (s ModifyPolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarClawSkillResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyPolarClawSkillResponseBody) GetEvaluateStatus() *string {
	return s.EvaluateStatus
}

func (s *ModifyPolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolarClawSkillResponseBody) GetRiskDetail() *string {
	return s.RiskDetail
}

func (s *ModifyPolarClawSkillResponseBody) SetApplicationId(v string) *ModifyPolarClawSkillResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyPolarClawSkillResponseBody) SetEvaluateStatus(v string) *ModifyPolarClawSkillResponseBody {
	s.EvaluateStatus = &v
	return s
}

func (s *ModifyPolarClawSkillResponseBody) SetRequestId(v string) *ModifyPolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarClawSkillResponseBody) SetRiskDetail(v string) *ModifyPolarClawSkillResponseBody {
	s.RiskDetail = &v
	return s
}

func (s *ModifyPolarClawSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
