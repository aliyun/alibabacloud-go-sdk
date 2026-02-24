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
	// This parameter is required.
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
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
