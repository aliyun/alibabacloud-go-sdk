// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutput interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v string) *MediaConvertOutput
	GetFeatures() *string
	SetName(v string) *MediaConvertOutput
	GetName() *string
	SetOutputFile(v *MediaObject) *MediaConvertOutput
	GetOutputFile() *MediaObject
	SetOverrideParams(v string) *MediaConvertOutput
	GetOverrideParams() *string
	SetPriority(v int32) *MediaConvertOutput
	GetPriority() *int32
	SetTemplateId(v string) *MediaConvertOutput
	GetTemplateId() *string
}

type MediaConvertOutput struct {
	// The feature parameters.
	//
	// example:
	//
	// {}
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The name of the output.
	//
	// example:
	//
	// output-video
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output file.
	OutputFile *MediaObject `json:"OutputFile,omitempty" xml:"OutputFile,omitempty"`
	// A JSON string containing parameters to overwrite the corresponding settings of the template.
	//
	// example:
	//
	// {}
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	// The priority. Valid values: 1 to 10. A larger value indicates a higher priority. Default value: 6.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MediaConvertOutput) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutput) GoString() string {
	return s.String()
}

func (s *MediaConvertOutput) GetFeatures() *string {
	return s.Features
}

func (s *MediaConvertOutput) GetName() *string {
	return s.Name
}

func (s *MediaConvertOutput) GetOutputFile() *MediaObject {
	return s.OutputFile
}

func (s *MediaConvertOutput) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *MediaConvertOutput) GetPriority() *int32 {
	return s.Priority
}

func (s *MediaConvertOutput) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MediaConvertOutput) SetFeatures(v string) *MediaConvertOutput {
	s.Features = &v
	return s
}

func (s *MediaConvertOutput) SetName(v string) *MediaConvertOutput {
	s.Name = &v
	return s
}

func (s *MediaConvertOutput) SetOutputFile(v *MediaObject) *MediaConvertOutput {
	s.OutputFile = v
	return s
}

func (s *MediaConvertOutput) SetOverrideParams(v string) *MediaConvertOutput {
	s.OverrideParams = &v
	return s
}

func (s *MediaConvertOutput) SetPriority(v int32) *MediaConvertOutput {
	s.Priority = &v
	return s
}

func (s *MediaConvertOutput) SetTemplateId(v string) *MediaConvertOutput {
	s.TemplateId = &v
	return s
}

func (s *MediaConvertOutput) Validate() error {
	if s.OutputFile != nil {
		if err := s.OutputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
