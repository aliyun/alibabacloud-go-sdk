// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteExporterRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *DeleteExporterRuleRequest
	GetRuleName() *string
}

type DeleteExporterRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the data export rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// myRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteExporterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExporterRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteExporterRuleRequest) SetRegionId(v string) *DeleteExporterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExporterRuleRequest) SetRuleName(v string) *DeleteExporterRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteExporterRuleRequest) Validate() error {
	return dara.Validate(s)
}
