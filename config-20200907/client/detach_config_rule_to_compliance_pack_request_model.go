// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachConfigRuleToCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *DetachConfigRuleToCompliancePackRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *DetachConfigRuleToCompliancePackRequest
	GetConfigRuleIds() *string
}

type DetachConfigRuleToCompliancePackRequest struct {
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

func (s DetachConfigRuleToCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DetachConfigRuleToCompliancePackRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DetachConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *DetachConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *DetachConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackRequest) Validate() error {
	return dara.Validate(s)
}
