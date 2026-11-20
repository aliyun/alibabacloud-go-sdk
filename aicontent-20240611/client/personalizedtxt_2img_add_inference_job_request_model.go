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
	// The number of images to generate. Note: Due to resource limits in the test environment, a maximum of 10 images can be generated at a time. If this value is greater than 10, it is treated as 10.
	//
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// The ID of the model training task. Specifies the model to use for inference.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The English prompt that describes the image to be generated. Replace the character to be generated with <special-token>. For example, change "a man in the snow" to "a <special-token> in the snow" and "a photo of a girl" to "a photo of a <special-token>".
	//
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// The random seed. A fixed seed generates the same image and ensures the reproducibility of the results. Note: The input value must be in the range of [-1, 2147483647]. If the input value is out of this range or not provided, a suitable seed is automatically generated.
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
