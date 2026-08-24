// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v string) *DeleteVulScanScheduledStrategyRequest
	GetStrategyId() *string
}

type DeleteVulScanScheduledStrategyRequest struct {
	// The ID of the scheduled vulnerability scanning policy to delete. You can obtain the value from the following operations:
	//
	// - [ListVulScanScheduledStrategies](~~ListVulScanScheduledStrategies~~): Lists scheduled vulnerability scanning policies.
	//
	// - [CreateVulScanScheduledStrategy](~~CreateVulScanScheduledStrategy~~): Creates a scheduled vulnerability scanning policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul-scan-scheduled-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteVulScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVulScanScheduledStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DeleteVulScanScheduledStrategyRequest) SetStrategyId(v string) *DeleteVulScanScheduledStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DeleteVulScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
