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
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-5bb1626622af00bd****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The rule ID. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-6cc4626622af00e7****
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
