// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest
	GetAggregatorId() *string
	SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest
	GetConfigRuleIds() *string
}

type GetAggregateResourceComplianceGroupByResourceTypeRequest struct {
	// This parameter is required.
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
