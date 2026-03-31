// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleIds(v string) *GetResourceComplianceGroupByResourceTypeRequest
	GetConfigRuleIds() *string
}

type GetResourceComplianceGroupByResourceTypeRequest struct {
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// example:
	//
	// cr-a5c6626622af0058****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *GetResourceComplianceGroupByResourceTypeRequest) SetConfigRuleIds(v string) *GetResourceComplianceGroupByResourceTypeRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
