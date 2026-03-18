// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddInferenceJobCmd interface {
	dara.Model
	String() string
	GoString() string
	SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobCmd
	GetImageNumber() *int32
	SetModelId(v string) *Personalizedtxt2imgAddInferenceJobCmd
	GetModelId() *string
	SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobCmd
	GetPrompt() *string
	SetSeed(v int32) *Personalizedtxt2imgAddInferenceJobCmd
	GetSeed() *int32
}

type Personalizedtxt2imgAddInferenceJobCmd struct {
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
	Seed *int32 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobCmd) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) GetImageNumber() *int32 {
	return s.ImageNumber
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) GetPrompt() *string {
	return s.Prompt
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) GetSeed() *int32 {
	return s.Seed
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetSeed(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Seed = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) Validate() error {
	return dara.Validate(s)
}
