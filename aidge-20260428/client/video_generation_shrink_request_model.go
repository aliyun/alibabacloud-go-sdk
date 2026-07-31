// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *VideoGenerationShrinkRequest
	GetInputShrink() *string
	SetIntentShrink(v string) *VideoGenerationShrinkRequest
	GetIntentShrink() *string
	SetOutputShrink(v string) *VideoGenerationShrinkRequest
	GetOutputShrink() *string
}

type VideoGenerationShrinkRequest struct {
	// The product input.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The intent parameters. Currently unavailable.
	IntentShrink *string `json:"Intent,omitempty" xml:"Intent,omitempty"`
	// The output parameters.
	//
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s VideoGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *VideoGenerationShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *VideoGenerationShrinkRequest) GetIntentShrink() *string {
	return s.IntentShrink
}

func (s *VideoGenerationShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *VideoGenerationShrinkRequest) SetInputShrink(v string) *VideoGenerationShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *VideoGenerationShrinkRequest) SetIntentShrink(v string) *VideoGenerationShrinkRequest {
	s.IntentShrink = &v
	return s
}

func (s *VideoGenerationShrinkRequest) SetOutputShrink(v string) *VideoGenerationShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *VideoGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
