// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLearningSamplesShrink(v string) *RunStyleWritingShrinkRequest
	GetLearningSamplesShrink() *string
	SetProcessStage(v string) *RunStyleWritingShrinkRequest
	GetProcessStage() *string
	SetReferenceMaterialsShrink(v string) *RunStyleWritingShrinkRequest
	GetReferenceMaterialsShrink() *string
	SetStyleFeature(v string) *RunStyleWritingShrinkRequest
	GetStyleFeature() *string
	SetUseSearch(v bool) *RunStyleWritingShrinkRequest
	GetUseSearch() *bool
	SetWritingTheme(v string) *RunStyleWritingShrinkRequest
	GetWritingTheme() *string
}

type RunStyleWritingShrinkRequest struct {
	LearningSamplesShrink    *string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty"`
	ProcessStage             *string `json:"processStage,omitempty" xml:"processStage,omitempty"`
	ReferenceMaterialsShrink *string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty"`
	StyleFeature             *string `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	UseSearch                *bool   `json:"useSearch,omitempty" xml:"useSearch,omitempty"`
	WritingTheme             *string `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
}

func (s RunStyleWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStyleWritingShrinkRequest) GetLearningSamplesShrink() *string {
	return s.LearningSamplesShrink
}

func (s *RunStyleWritingShrinkRequest) GetProcessStage() *string {
	return s.ProcessStage
}

func (s *RunStyleWritingShrinkRequest) GetReferenceMaterialsShrink() *string {
	return s.ReferenceMaterialsShrink
}

func (s *RunStyleWritingShrinkRequest) GetStyleFeature() *string {
	return s.StyleFeature
}

func (s *RunStyleWritingShrinkRequest) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunStyleWritingShrinkRequest) GetWritingTheme() *string {
	return s.WritingTheme
}

func (s *RunStyleWritingShrinkRequest) SetLearningSamplesShrink(v string) *RunStyleWritingShrinkRequest {
	s.LearningSamplesShrink = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetProcessStage(v string) *RunStyleWritingShrinkRequest {
	s.ProcessStage = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetReferenceMaterialsShrink(v string) *RunStyleWritingShrinkRequest {
	s.ReferenceMaterialsShrink = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetStyleFeature(v string) *RunStyleWritingShrinkRequest {
	s.StyleFeature = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetUseSearch(v bool) *RunStyleWritingShrinkRequest {
	s.UseSearch = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) SetWritingTheme(v string) *RunStyleWritingShrinkRequest {
	s.WritingTheme = &v
	return s
}

func (s *RunStyleWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
