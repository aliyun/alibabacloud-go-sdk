// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceType(v string) *DryRunConfigRuleResponseBody
	GetComplianceType() *string
	SetRequestId(v string) *DryRunConfigRuleResponseBody
	GetRequestId() *string
	SetRuleConditionContext(v string) *DryRunConfigRuleResponseBody
	GetRuleConditionContext() *string
}

type DryRunConfigRuleResponseBody struct {
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C32EEAD7-BF64-5927-977A-AFF9342B7275
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {
	//
	// 			"result":"NON_COMPLIANT",
	//
	// 			"children":[
	//
	// 				{
	//
	// 					"featureValue":"1",
	//
	// 					"featureSource":"CONFIGURATION",
	//
	// 					"result":"NON_COMPLIANT",
	//
	// 					"desired":"4",
	//
	// 					"children":[],
	//
	// 					"operator":"StringEquals",
	//
	// 					"featurePath":"$.Cpu"
	//
	// 				}
	//
	// 			],
	//
	// 			"operator":"and"
	//
	// 		}
	//
	// 	}
	RuleConditionContext *string `json:"RuleConditionContext,omitempty" xml:"RuleConditionContext,omitempty"`
}

func (s DryRunConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DryRunConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DryRunConfigRuleResponseBody) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *DryRunConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DryRunConfigRuleResponseBody) GetRuleConditionContext() *string {
	return s.RuleConditionContext
}

func (s *DryRunConfigRuleResponseBody) SetComplianceType(v string) *DryRunConfigRuleResponseBody {
	s.ComplianceType = &v
	return s
}

func (s *DryRunConfigRuleResponseBody) SetRequestId(v string) *DryRunConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DryRunConfigRuleResponseBody) SetRuleConditionContext(v string) *DryRunConfigRuleResponseBody {
	s.RuleConditionContext = &v
	return s
}

func (s *DryRunConfigRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
