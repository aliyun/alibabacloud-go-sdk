// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleIds(v string) *GetResourceComplianceGroupByRegionRequest
	GetConfigRuleIds() *string
}

type GetResourceComplianceGroupByRegionRequest struct {
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// example:
	//
	// cr-2541626622af0000****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetResourceComplianceGroupByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GetResourceComplianceGroupByRegionRequest) SetConfigRuleIds(v string) *GetResourceComplianceGroupByRegionRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionRequest) Validate() error {
	return dara.Validate(s)
}
