// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceGroupByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByRegionRequest
	GetAggregatorId() *string
	SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByRegionRequest
	GetConfigRuleIds() *string
}

type GetAggregateResourceComplianceGroupByRegionRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-d6c9626622af0052****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// example:
	//
	// cr-2652626622af005e****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByRegionRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByRegionRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) Validate() error {
	return dara.Validate(s)
}
