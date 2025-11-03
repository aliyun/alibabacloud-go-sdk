// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaiLianAgentTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *BaiLianAgentTransformParameters
	GetApiKey() *string
	SetApplicationId(v string) *BaiLianAgentTransformParameters
	GetApplicationId() *string
	SetPrompt(v *BaiLianAgentTransformParametersPrompt) *BaiLianAgentTransformParameters
	GetPrompt() *BaiLianAgentTransformParametersPrompt
	SetRequestPerMinute(v int64) *BaiLianAgentTransformParameters
	GetRequestPerMinute() *int64
	SetTokenPerMinute(v int64) *BaiLianAgentTransformParameters
	GetTokenPerMinute() *int64
}

type BaiLianAgentTransformParameters struct {
	ApiKey           *string                                `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ApplicationId    *string                                `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Prompt           *BaiLianAgentTransformParametersPrompt `json:"Prompt,omitempty" xml:"Prompt,omitempty" type:"Struct"`
	RequestPerMinute *int64                                 `json:"RequestPerMinute,omitempty" xml:"RequestPerMinute,omitempty"`
	TokenPerMinute   *int64                                 `json:"TokenPerMinute,omitempty" xml:"TokenPerMinute,omitempty"`
}

func (s BaiLianAgentTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s BaiLianAgentTransformParameters) GoString() string {
	return s.String()
}

func (s *BaiLianAgentTransformParameters) GetApiKey() *string {
	return s.ApiKey
}

func (s *BaiLianAgentTransformParameters) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *BaiLianAgentTransformParameters) GetPrompt() *BaiLianAgentTransformParametersPrompt {
	return s.Prompt
}

func (s *BaiLianAgentTransformParameters) GetRequestPerMinute() *int64 {
	return s.RequestPerMinute
}

func (s *BaiLianAgentTransformParameters) GetTokenPerMinute() *int64 {
	return s.TokenPerMinute
}

func (s *BaiLianAgentTransformParameters) SetApiKey(v string) *BaiLianAgentTransformParameters {
	s.ApiKey = &v
	return s
}

func (s *BaiLianAgentTransformParameters) SetApplicationId(v string) *BaiLianAgentTransformParameters {
	s.ApplicationId = &v
	return s
}

func (s *BaiLianAgentTransformParameters) SetPrompt(v *BaiLianAgentTransformParametersPrompt) *BaiLianAgentTransformParameters {
	s.Prompt = v
	return s
}

func (s *BaiLianAgentTransformParameters) SetRequestPerMinute(v int64) *BaiLianAgentTransformParameters {
	s.RequestPerMinute = &v
	return s
}

func (s *BaiLianAgentTransformParameters) SetTokenPerMinute(v int64) *BaiLianAgentTransformParameters {
	s.TokenPerMinute = &v
	return s
}

func (s *BaiLianAgentTransformParameters) Validate() error {
	if s.Prompt != nil {
		if err := s.Prompt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BaiLianAgentTransformParametersPrompt struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BaiLianAgentTransformParametersPrompt) String() string {
	return dara.Prettify(s)
}

func (s BaiLianAgentTransformParametersPrompt) GoString() string {
	return s.String()
}

func (s *BaiLianAgentTransformParametersPrompt) GetForm() *string {
	return s.Form
}

func (s *BaiLianAgentTransformParametersPrompt) GetTemplate() *string {
	return s.Template
}

func (s *BaiLianAgentTransformParametersPrompt) GetValue() *string {
	return s.Value
}

func (s *BaiLianAgentTransformParametersPrompt) SetForm(v string) *BaiLianAgentTransformParametersPrompt {
	s.Form = &v
	return s
}

func (s *BaiLianAgentTransformParametersPrompt) SetTemplate(v string) *BaiLianAgentTransformParametersPrompt {
	s.Template = &v
	return s
}

func (s *BaiLianAgentTransformParametersPrompt) SetValue(v string) *BaiLianAgentTransformParametersPrompt {
	s.Value = &v
	return s
}

func (s *BaiLianAgentTransformParametersPrompt) Validate() error {
	return dara.Validate(s)
}
