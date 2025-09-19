// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfSetupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyConfig(v string) *ModifyClientConfSetupRequest
	GetStrategyConfig() *string
	SetStrategyTag(v string) *ModifyClientConfSetupRequest
	GetStrategyTag() *string
	SetStrategyTagValue(v string) *ModifyClientConfSetupRequest
	GetStrategyTagValue() *string
}

type ModifyClientConfSetupRequest struct {
	// The configurations of the Security Center agent.
	//
	// 	- cpu: the maximum CPU utilization that can be occupied by the Security Center agent on the server
	//
	// 	- mem: the maximum memory usage that can be occupied by the Security Center agent on the server
	//
	// example:
	//
	// {
	//
	//       "cpu": "20"
	//
	// }
	StrategyConfig *string `json:"StrategyConfig,omitempty" xml:"StrategyConfig,omitempty"`
	// The type of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// machineResource
	StrategyTag *string `json:"StrategyTag,omitempty" xml:"StrategyTag,omitempty"`
	// The value of the tag. Valid values:
	//
	// 	- major
	//
	// 	- advanced
	//
	// 	- basic
	//
	// This parameter is required.
	//
	// example:
	//
	// major
	StrategyTagValue *string `json:"StrategyTagValue,omitempty" xml:"StrategyTagValue,omitempty"`
}

func (s ModifyClientConfSetupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfSetupRequest) GoString() string {
	return s.String()
}

func (s *ModifyClientConfSetupRequest) GetStrategyConfig() *string {
	return s.StrategyConfig
}

func (s *ModifyClientConfSetupRequest) GetStrategyTag() *string {
	return s.StrategyTag
}

func (s *ModifyClientConfSetupRequest) GetStrategyTagValue() *string {
	return s.StrategyTagValue
}

func (s *ModifyClientConfSetupRequest) SetStrategyConfig(v string) *ModifyClientConfSetupRequest {
	s.StrategyConfig = &v
	return s
}

func (s *ModifyClientConfSetupRequest) SetStrategyTag(v string) *ModifyClientConfSetupRequest {
	s.StrategyTag = &v
	return s
}

func (s *ModifyClientConfSetupRequest) SetStrategyTagValue(v string) *ModifyClientConfSetupRequest {
	s.StrategyTagValue = &v
	return s
}

func (s *ModifyClientConfSetupRequest) Validate() error {
	return dara.Validate(s)
}
