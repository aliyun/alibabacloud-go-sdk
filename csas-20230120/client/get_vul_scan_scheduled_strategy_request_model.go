// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v string) *GetVulScanScheduledStrategyRequest
	GetStrategyId() *string
}

type GetVulScanScheduledStrategyRequest struct {
	// The vulnerability scheduled scan policy ID. You can obtain the value from the following operations:
	//
	// - [ListVulScanScheduledStrategies](~~ListVulScanScheduledStrategies~~): Lists vulnerability scheduled scan policies.
	//
	// - [CreateVulScanScheduledStrategy](~~CreateVulScanScheduledStrategy~~): Creates a vulnerability scheduled scan policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s GetVulScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVulScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetVulScanScheduledStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetVulScanScheduledStrategyRequest) SetStrategyId(v string) *GetVulScanScheduledStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *GetVulScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
