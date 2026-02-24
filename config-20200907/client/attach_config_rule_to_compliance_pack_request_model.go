// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachConfigRuleToCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *AttachConfigRuleToCompliancePackRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *AttachConfigRuleToCompliancePackRequest
	GetConfigRuleIds() *string
}

type AttachConfigRuleToCompliancePackRequest struct {
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// This parameter is required.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s AttachConfigRuleToCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *AttachConfigRuleToCompliancePackRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *AttachConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *AttachConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *AttachConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackRequest) Validate() error {
	return dara.Validate(s)
}
