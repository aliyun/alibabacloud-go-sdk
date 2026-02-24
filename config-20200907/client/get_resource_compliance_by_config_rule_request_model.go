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
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// This parameter is required.
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
