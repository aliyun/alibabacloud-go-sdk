// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamOptions interface {
	dara.Model
	String() string
	GoString() string
	SetIncrementalOutput(v bool) *StreamOptions
	GetIncrementalOutput() *bool
	SetNeedReturnFinalResult(v bool) *StreamOptions
	GetNeedReturnFinalResult() *bool
}

type StreamOptions struct {
	IncrementalOutput     *bool `json:"IncrementalOutput,omitempty" xml:"IncrementalOutput,omitempty"`
	NeedReturnFinalResult *bool `json:"NeedReturnFinalResult,omitempty" xml:"NeedReturnFinalResult,omitempty"`
}

func (s StreamOptions) String() string {
	return dara.Prettify(s)
}

func (s StreamOptions) GoString() string {
	return s.String()
}

func (s *StreamOptions) GetIncrementalOutput() *bool {
	return s.IncrementalOutput
}

func (s *StreamOptions) GetNeedReturnFinalResult() *bool {
	return s.NeedReturnFinalResult
}

func (s *StreamOptions) SetIncrementalOutput(v bool) *StreamOptions {
	s.IncrementalOutput = &v
	return s
}

func (s *StreamOptions) SetNeedReturnFinalResult(v bool) *StreamOptions {
	s.NeedReturnFinalResult = &v
	return s
}

func (s *StreamOptions) Validate() error {
	return dara.Validate(s)
}
