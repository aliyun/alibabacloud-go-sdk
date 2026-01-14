// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelUsage interface {
	dara.Model
	String() string
	GoString() string
	SetInputTokens(v int64) *ModelUsage
	GetInputTokens() *int64
	SetOutputTokens(v int64) *ModelUsage
	GetOutputTokens() *int64
	SetTotalTokens(v int64) *ModelUsage
	GetTotalTokens() *int64
}

type ModelUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s ModelUsage) String() string {
	return dara.Prettify(s)
}

func (s ModelUsage) GoString() string {
	return s.String()
}

func (s *ModelUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *ModelUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *ModelUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ModelUsage) SetInputTokens(v int64) *ModelUsage {
	s.InputTokens = &v
	return s
}

func (s *ModelUsage) SetOutputTokens(v int64) *ModelUsage {
	s.OutputTokens = &v
	return s
}

func (s *ModelUsage) SetTotalTokens(v int64) *ModelUsage {
	s.TotalTokens = &v
	return s
}

func (s *ModelUsage) Validate() error {
	return dara.Validate(s)
}
