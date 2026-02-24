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
	// This parameter is required.
	AggregatorId   *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// This parameter is required.
	ConfigRuleId      *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
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
