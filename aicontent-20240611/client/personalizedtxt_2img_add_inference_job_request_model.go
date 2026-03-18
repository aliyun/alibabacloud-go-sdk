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
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
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
