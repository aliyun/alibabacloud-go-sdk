// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageAddInferenceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNumber(v int32) *PersonalizedTextToImageAddInferenceJobRequest
	GetImageNumber() *int32
	SetImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobRequest
	GetImageUrl() []*string
	SetPrompt(v string) *PersonalizedTextToImageAddInferenceJobRequest
	GetPrompt() *string
	SetSeed(v int64) *PersonalizedTextToImageAddInferenceJobRequest
	GetSeed() *int64
	SetStrength(v float64) *PersonalizedTextToImageAddInferenceJobRequest
	GetStrength() *float64
	SetTrainSteps(v int32) *PersonalizedTextToImageAddInferenceJobRequest
	GetTrainSteps() *int32
}

type PersonalizedTextToImageAddInferenceJobRequest struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// example:
	//
	// 1
	Strength *float64 `json:"strength,omitempty" xml:"strength,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetImageNumber() *int32 {
	return s.ImageNumber
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetSeed() *int64 {
	return s.Seed
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetStrength() *float64 {
	return s.Strength
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) GetTrainSteps() *int32 {
	return s.TrainSteps
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageNumber(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageUrl = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetPrompt(v string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetSeed(v int64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Seed = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetStrength(v float64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Strength = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetTrainSteps(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.TrainSteps = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) Validate() error {
	return dara.Validate(s)
}
