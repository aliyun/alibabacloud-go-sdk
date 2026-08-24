// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanScheduledStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v string) *GetVirusScanScheduledStrategyRequest
	GetStrategyId() *string
}

type GetVirusScanScheduledStrategyRequest struct {
	// The ID of the scheduled virus scan policy. You can obtain the value from the following operations:
	//
	// - [ListVirusScanScheduledStrategies](~~ListVirusScanScheduledStrategies~~): Lists scheduled virus scan policies.
	//
	// - [CreateVirusScanScheduledStrategy](~~CreateVirusScanScheduledStrategy~~): Creates a scheduled virus scan policy.
	//
	// example:
	//
	// vc-strategy-8a3f6c2e91b7****
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s GetVirusScanScheduledStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanScheduledStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetVirusScanScheduledStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetVirusScanScheduledStrategyRequest) SetStrategyId(v string) *GetVirusScanScheduledStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *GetVirusScanScheduledStrategyRequest) Validate() error {
	return dara.Validate(s)
}
