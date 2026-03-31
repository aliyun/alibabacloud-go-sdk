// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceByConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceType(v string) *GetResourceComplianceByConfigRuleRequest
	GetComplianceType() *string
	SetConfigRuleId(v string) *GetResourceComplianceByConfigRuleRequest
	GetConfigRuleId() *string
}

type GetResourceComplianceByConfigRuleRequest struct {
	// The compliance evaluation result. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-d369626622af008e****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s GetResourceComplianceByConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetResourceComplianceByConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetResourceComplianceByConfigRuleRequest) SetComplianceType(v string) *GetResourceComplianceByConfigRuleRequest {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleRequest) SetConfigRuleId(v string) *GetResourceComplianceByConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleRequest) Validate() error {
	return dara.Validate(s)
}
