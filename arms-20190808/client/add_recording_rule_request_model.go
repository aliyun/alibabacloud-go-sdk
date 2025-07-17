// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddRecordingRuleRequest
	GetClusterId() *string
	SetRegionId(v string) *AddRecordingRuleRequest
	GetRegionId() *string
	SetRuleYaml(v string) *AddRecordingRuleRequest
	GetRuleYaml() *string
}

type AddRecordingRuleRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The custom recording rule. The value is in the YAML format.
	//
	// This parameter is required.
	//
	// example:
	//
	// groups: - name: "recording_demo"   rules:   - expr: "sum(jvm_memory_max_bytes)"     record: "rate_coredns_demo"
	RuleYaml *string `json:"RuleYaml,omitempty" xml:"RuleYaml,omitempty"`
}

func (s AddRecordingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddRecordingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddRecordingRuleRequest) GetRuleYaml() *string {
	return s.RuleYaml
}

func (s *AddRecordingRuleRequest) SetClusterId(v string) *AddRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *AddRecordingRuleRequest) SetRegionId(v string) *AddRecordingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddRecordingRuleRequest) SetRuleYaml(v string) *AddRecordingRuleRequest {
	s.RuleYaml = &v
	return s
}

func (s *AddRecordingRuleRequest) Validate() error {
	return dara.Validate(s)
}
