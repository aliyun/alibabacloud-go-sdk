// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLearningSamples(v []*string) *RunStyleWritingRequest
	GetLearningSamples() []*string
	SetProcessStage(v string) *RunStyleWritingRequest
	GetProcessStage() *string
	SetReferenceMaterials(v []*string) *RunStyleWritingRequest
	GetReferenceMaterials() []*string
	SetStyleFeature(v string) *RunStyleWritingRequest
	GetStyleFeature() *string
	SetUseSearch(v bool) *RunStyleWritingRequest
	GetUseSearch() *bool
	SetWritingTheme(v string) *RunStyleWritingRequest
	GetWritingTheme() *string
}

type RunStyleWritingRequest struct {
	LearningSamples    []*string `json:"learningSamples,omitempty" xml:"learningSamples,omitempty" type:"Repeated"`
	ProcessStage       *string   `json:"processStage,omitempty" xml:"processStage,omitempty"`
	ReferenceMaterials []*string `json:"referenceMaterials,omitempty" xml:"referenceMaterials,omitempty" type:"Repeated"`
	StyleFeature       *string   `json:"styleFeature,omitempty" xml:"styleFeature,omitempty"`
	UseSearch          *bool     `json:"useSearch,omitempty" xml:"useSearch,omitempty"`
	WritingTheme       *string   `json:"writingTheme,omitempty" xml:"writingTheme,omitempty"`
}

func (s RunStyleWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingRequest) GoString() string {
	return s.String()
}

func (s *RunStyleWritingRequest) GetLearningSamples() []*string {
	return s.LearningSamples
}

func (s *RunStyleWritingRequest) GetProcessStage() *string {
	return s.ProcessStage
}

func (s *RunStyleWritingRequest) GetReferenceMaterials() []*string {
	return s.ReferenceMaterials
}

func (s *RunStyleWritingRequest) GetStyleFeature() *string {
	return s.StyleFeature
}

func (s *RunStyleWritingRequest) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunStyleWritingRequest) GetWritingTheme() *string {
	return s.WritingTheme
}

func (s *RunStyleWritingRequest) SetLearningSamples(v []*string) *RunStyleWritingRequest {
	s.LearningSamples = v
	return s
}

func (s *RunStyleWritingRequest) SetProcessStage(v string) *RunStyleWritingRequest {
	s.ProcessStage = &v
	return s
}

func (s *RunStyleWritingRequest) SetReferenceMaterials(v []*string) *RunStyleWritingRequest {
	s.ReferenceMaterials = v
	return s
}

func (s *RunStyleWritingRequest) SetStyleFeature(v string) *RunStyleWritingRequest {
	s.StyleFeature = &v
	return s
}

func (s *RunStyleWritingRequest) SetUseSearch(v bool) *RunStyleWritingRequest {
	s.UseSearch = &v
	return s
}

func (s *RunStyleWritingRequest) SetWritingTheme(v string) *RunStyleWritingRequest {
	s.WritingTheme = &v
	return s
}

func (s *RunStyleWritingRequest) Validate() error {
	return dara.Validate(s)
}
