// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputGroupConfig interface {
	dara.Model
	String() string
	GoString() string
	SetManifestName(v string) *MediaConvertOutputGroupConfig
	GetManifestName() *string
	SetOutputFileBase(v *MediaObject) *MediaConvertOutputGroupConfig
	GetOutputFileBase() *MediaObject
	SetType(v string) *MediaConvertOutputGroupConfig
	GetType() *string
}

type MediaConvertOutputGroupConfig struct {
	ManifestName   *string      `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	OutputFileBase *MediaObject `json:"OutputFileBase,omitempty" xml:"OutputFileBase,omitempty"`
	Type           *string      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertOutputGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputGroupConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputGroupConfig) GetManifestName() *string {
	return s.ManifestName
}

func (s *MediaConvertOutputGroupConfig) GetOutputFileBase() *MediaObject {
	return s.OutputFileBase
}

func (s *MediaConvertOutputGroupConfig) GetType() *string {
	return s.Type
}

func (s *MediaConvertOutputGroupConfig) SetManifestName(v string) *MediaConvertOutputGroupConfig {
	s.ManifestName = &v
	return s
}

func (s *MediaConvertOutputGroupConfig) SetOutputFileBase(v *MediaObject) *MediaConvertOutputGroupConfig {
	s.OutputFileBase = v
	return s
}

func (s *MediaConvertOutputGroupConfig) SetType(v string) *MediaConvertOutputGroupConfig {
	s.Type = &v
	return s
}

func (s *MediaConvertOutputGroupConfig) Validate() error {
	if s.OutputFileBase != nil {
		if err := s.OutputFileBase.Validate(); err != nil {
			return err
		}
	}
	return nil
}
