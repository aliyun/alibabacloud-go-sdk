// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceByConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceComplianceByConfigRuleRequest
	GetAggregatorId() *string
	SetComplianceType(v string) *GetAggregateResourceComplianceByConfigRuleRequest
	GetComplianceType() *string
	SetConfigRuleId(v string) *GetAggregateResourceComplianceByConfigRuleRequest
	GetConfigRuleId() *string
	SetResourceAccountId(v int64) *GetAggregateResourceComplianceByConfigRuleRequest
	GetResourceAccountId() *int64
	SetResourceOwnerId(v int64) *GetAggregateResourceComplianceByConfigRuleRequest
	GetResourceOwnerId() *int64
}

type GetAggregateResourceComplianceByConfigRuleRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The compliance evaluation result of the resources. Valid values:
	//
	// 	- COMPLIANT: The resource is evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource is evaluated as incompliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to your resources.
	//
	// 	- INSUFFICIENT_DATA: No resource data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-d369626622af008e****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources in the account group belong.
	//
	// > You can use either the ResourceAccountId or ResourceOwnerId parameter. We recommend that you use the ResourceAccountId parameter.
	//
	// example:
	//
	// 100931896542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	// Deprecated
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetComplianceType(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetConfigRuleId(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetResourceAccountId(v int64) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetResourceOwnerId(v int64) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) Validate() error {
	return dara.Validate(s)
}
