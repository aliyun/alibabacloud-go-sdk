// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *VideoGenerationRequestInput) *VideoGenerationRequest
	GetInput() *VideoGenerationRequestInput
	SetIntent(v *VideoGenerationRequestIntent) *VideoGenerationRequest
	GetIntent() *VideoGenerationRequestIntent
	SetOutput(v *VideoGenerationRequestOutput) *VideoGenerationRequest
	GetOutput() *VideoGenerationRequestOutput
}

type VideoGenerationRequest struct {
	// This parameter is required.
	Input  *VideoGenerationRequestInput  `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	Intent *VideoGenerationRequestIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Struct"`
	// This parameter is required.
	Output *VideoGenerationRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
}

func (s VideoGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequest) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequest) GetInput() *VideoGenerationRequestInput {
	return s.Input
}

func (s *VideoGenerationRequest) GetIntent() *VideoGenerationRequestIntent {
	return s.Intent
}

func (s *VideoGenerationRequest) GetOutput() *VideoGenerationRequestOutput {
	return s.Output
}

func (s *VideoGenerationRequest) SetInput(v *VideoGenerationRequestInput) *VideoGenerationRequest {
	s.Input = v
	return s
}

func (s *VideoGenerationRequest) SetIntent(v *VideoGenerationRequestIntent) *VideoGenerationRequest {
	s.Intent = v
	return s
}

func (s *VideoGenerationRequest) SetOutput(v *VideoGenerationRequestOutput) *VideoGenerationRequest {
	s.Output = v
	return s
}

func (s *VideoGenerationRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Intent != nil {
		if err := s.Intent.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoGenerationRequestInput struct {
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2026 New Slimming Women\\"s Summer Dress with Mid-Length Design, High-Quality Waist Definition for a Slender Look
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s VideoGenerationRequestInput) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestInput) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestInput) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *VideoGenerationRequestInput) GetImages() []*string {
	return s.Images
}

func (s *VideoGenerationRequestInput) GetTitle() *string {
	return s.Title
}

func (s *VideoGenerationRequestInput) SetExtra(v map[string]interface{}) *VideoGenerationRequestInput {
	s.Extra = v
	return s
}

func (s *VideoGenerationRequestInput) SetImages(v []*string) *VideoGenerationRequestInput {
	s.Images = v
	return s
}

func (s *VideoGenerationRequestInput) SetTitle(v string) *VideoGenerationRequestInput {
	s.Title = &v
	return s
}

func (s *VideoGenerationRequestInput) Validate() error {
	return dara.Validate(s)
}

type VideoGenerationRequestIntent struct {
	// example:
	//
	// -
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// -
	Goal *string `json:"Goal,omitempty" xml:"Goal,omitempty"`
}

func (s VideoGenerationRequestIntent) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestIntent) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestIntent) GetChannel() *string {
	return s.Channel
}

func (s *VideoGenerationRequestIntent) GetGoal() *string {
	return s.Goal
}

func (s *VideoGenerationRequestIntent) SetChannel(v string) *VideoGenerationRequestIntent {
	s.Channel = &v
	return s
}

func (s *VideoGenerationRequestIntent) SetGoal(v string) *VideoGenerationRequestIntent {
	s.Goal = &v
	return s
}

func (s *VideoGenerationRequestIntent) Validate() error {
	return dara.Validate(s)
}

type VideoGenerationRequestOutput struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1080p
	Quality *string `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// example:
	//
	// 9:16
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s VideoGenerationRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationRequestOutput) GoString() string {
	return s.String()
}

func (s *VideoGenerationRequestOutput) GetDuration() *int64 {
	return s.Duration
}

func (s *VideoGenerationRequestOutput) GetQuality() *string {
	return s.Quality
}

func (s *VideoGenerationRequestOutput) GetRatio() *string {
	return s.Ratio
}

func (s *VideoGenerationRequestOutput) SetDuration(v int64) *VideoGenerationRequestOutput {
	s.Duration = &v
	return s
}

func (s *VideoGenerationRequestOutput) SetQuality(v string) *VideoGenerationRequestOutput {
	s.Quality = &v
	return s
}

func (s *VideoGenerationRequestOutput) SetRatio(v string) *VideoGenerationRequestOutput {
	s.Ratio = &v
	return s
}

func (s *VideoGenerationRequestOutput) Validate() error {
	return dara.Validate(s)
}
