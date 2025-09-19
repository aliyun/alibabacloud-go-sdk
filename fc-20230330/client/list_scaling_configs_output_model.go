// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScalingConfigsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListScalingConfigsOutput
	GetNextToken() *string
	SetScalingConfigs(v []*ScalingConfigStatus) *ListScalingConfigsOutput
	GetScalingConfigs() []*ScalingConfigStatus
}

type ListScalingConfigsOutput struct {
	NextToken      *string                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ScalingConfigs []*ScalingConfigStatus `json:"scalingConfigs" xml:"scalingConfigs" type:"Repeated"`
}

func (s ListScalingConfigsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListScalingConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListScalingConfigsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScalingConfigsOutput) GetScalingConfigs() []*ScalingConfigStatus {
	return s.ScalingConfigs
}

func (s *ListScalingConfigsOutput) SetNextToken(v string) *ListScalingConfigsOutput {
	s.NextToken = &v
	return s
}

func (s *ListScalingConfigsOutput) SetScalingConfigs(v []*ScalingConfigStatus) *ListScalingConfigsOutput {
	s.ScalingConfigs = v
	return s
}

func (s *ListScalingConfigsOutput) Validate() error {
	return dara.Validate(s)
}
