// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScalingConfigStatusOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListScalingConfigStatusOutput
	GetNextToken() *string
	SetResult(v []*ScalingConfigStatus) *ListScalingConfigStatusOutput
	GetResult() []*ScalingConfigStatus
}

type ListScalingConfigStatusOutput struct {
	NextToken *string                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*ScalingConfigStatus `json:"result" xml:"result" type:"Repeated"`
}

func (s ListScalingConfigStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s ListScalingConfigStatusOutput) GoString() string {
	return s.String()
}

func (s *ListScalingConfigStatusOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScalingConfigStatusOutput) GetResult() []*ScalingConfigStatus {
	return s.Result
}

func (s *ListScalingConfigStatusOutput) SetNextToken(v string) *ListScalingConfigStatusOutput {
	s.NextToken = &v
	return s
}

func (s *ListScalingConfigStatusOutput) SetResult(v []*ScalingConfigStatus) *ListScalingConfigStatusOutput {
	s.Result = v
	return s
}

func (s *ListScalingConfigStatusOutput) Validate() error {
	return dara.Validate(s)
}
