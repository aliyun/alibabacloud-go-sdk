// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputGroupOutput interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v string) *MediaConvertOutputGroupOutput
	GetFeatures() *string
	SetName(v string) *MediaConvertOutputGroupOutput
	GetName() *string
	SetOutputFileName(v string) *MediaConvertOutputGroupOutput
	GetOutputFileName() *string
	SetOverrideParams(v string) *MediaConvertOutputGroupOutput
	GetOverrideParams() *string
	SetPriority(v int32) *MediaConvertOutputGroupOutput
	GetPriority() *int32
	SetTemplateId(v string) *MediaConvertOutputGroupOutput
	GetTemplateId() *string
}

type MediaConvertOutputGroupOutput struct {
	Features       *string `json:"Features,omitempty" xml:"Features,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputFileName *string `json:"OutputFileName,omitempty" xml:"OutputFileName,omitempty"`
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MediaConvertOutputGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputGroupOutput) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputGroupOutput) GetFeatures() *string {
	return s.Features
}

func (s *MediaConvertOutputGroupOutput) GetName() *string {
	return s.Name
}

func (s *MediaConvertOutputGroupOutput) GetOutputFileName() *string {
	return s.OutputFileName
}

func (s *MediaConvertOutputGroupOutput) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *MediaConvertOutputGroupOutput) GetPriority() *int32 {
	return s.Priority
}

func (s *MediaConvertOutputGroupOutput) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MediaConvertOutputGroupOutput) SetFeatures(v string) *MediaConvertOutputGroupOutput {
	s.Features = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) SetName(v string) *MediaConvertOutputGroupOutput {
	s.Name = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) SetOutputFileName(v string) *MediaConvertOutputGroupOutput {
	s.OutputFileName = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) SetOverrideParams(v string) *MediaConvertOutputGroupOutput {
	s.OverrideParams = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) SetPriority(v int32) *MediaConvertOutputGroupOutput {
	s.Priority = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) SetTemplateId(v string) *MediaConvertOutputGroupOutput {
	s.TemplateId = &v
	return s
}

func (s *MediaConvertOutputGroupOutput) Validate() error {
	return dara.Validate(s)
}
