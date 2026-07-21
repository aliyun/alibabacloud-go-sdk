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
	// The number of images to generate. Note: The maximum is 10 images per request in the test environment. If the value exceeds 10, it is treated as 10.
	//
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// An array containing one or more image URLs. For example, `["url_1", "url_2", ...]`.
	//
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// The English prompt for image generation. Use the placeholder for the subject. For example, change "a man in the snow" to "a in the snow".
	//
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// A random seed to ensure reproducible image generation. The value must be within `[-1, 2147483647]`. If the value is outside this range or omitted, the system automatically generates a seed.
	//
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// Determines the influence of the reference image.
	//
	// Valid values: `0.3`, `0.4`, `0.5`, `0.6`, `0.7`, and `0.8`.
	//
	// A lower value decreases the influence of the reference image and increases the influence of the text prompt.
	//
	// The default is `0.5`, and you typically do not need to change this value.
	//
	// example:
	//
	// 1
	Strength *float64 `json:"strength,omitempty" xml:"strength,omitempty"`
	// The number of training steps for the model.
	//
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
