// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddInferenceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobRequest
	GetImageNumber() *int32
	SetModelId(v string) *Personalizedtxt2imgAddInferenceJobRequest
	GetModelId() *string
	SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobRequest
	GetPrompt() *string
	SetSeed(v int64) *Personalizedtxt2imgAddInferenceJobRequest
	GetSeed() *int64
}

type Personalizedtxt2imgAddInferenceJobRequest struct {
	// The number of images to generate. Note: Due to resource limits in the test environment, you can generate up to 10 images per request. The system automatically sets values greater than 10 to 10.
	//
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// The model ID to use for the inference job.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// An English prompt describing the image to generate. Replace the subject with . For example, change "a man in the snow" to "a in the snow", and "a photo of a girl" to "a photo of a ".
	//
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// The seed for the random number generator. Using the same seed ensures reproducible results. The value must be between -1 and 2,147,483,647. If the value is outside this range or is not specified, the system automatically generates a suitable seed.
	//
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) GetImageNumber() *int32 {
	return s.ImageNumber
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) GetSeed() *int64 {
	return s.Seed
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetSeed(v int64) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Seed = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) Validate() error {
	return dara.Validate(s)
}
