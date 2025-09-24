// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSavingPlanUserDeductRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetSavingPlanUserDeductRuleResponseBody
	GetData() *bool
	SetRequestId(v string) *SetSavingPlanUserDeductRuleResponseBody
	GetRequestId() *string
}

type SetSavingPlanUserDeductRuleResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSavingPlanUserDeductRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetSavingPlanUserDeductRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSavingPlanUserDeductRuleResponseBody) SetData(v bool) *SetSavingPlanUserDeductRuleResponseBody {
	s.Data = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponseBody) SetRequestId(v string) *SetSavingPlanUserDeductRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
