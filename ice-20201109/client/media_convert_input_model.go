// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertInput interface {
	dara.Model
	String() string
	GoString() string
	SetInputFile(v *MediaObject) *MediaConvertInput
	GetInputFile() *MediaObject
	SetName(v string) *MediaConvertInput
	GetName() *string
}

type MediaConvertInput struct {
	InputFile *MediaObject `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	Name      *string      `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MediaConvertInput) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertInput) GoString() string {
	return s.String()
}

func (s *MediaConvertInput) GetInputFile() *MediaObject {
	return s.InputFile
}

func (s *MediaConvertInput) GetName() *string {
	return s.Name
}

func (s *MediaConvertInput) SetInputFile(v *MediaObject) *MediaConvertInput {
	s.InputFile = v
	return s
}

func (s *MediaConvertInput) SetName(v string) *MediaConvertInput {
	s.Name = &v
	return s
}

func (s *MediaConvertInput) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
