// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *SubmitImageTranslateTaskRequestExt) *SubmitImageTranslateTaskRequest
	GetExt() *SubmitImageTranslateTaskRequestExt
	SetFormat(v string) *SubmitImageTranslateTaskRequest
	GetFormat() *string
	SetScene(v string) *SubmitImageTranslateTaskRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitImageTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v []*string) *SubmitImageTranslateTaskRequest
	GetTargetLanguage() []*string
	SetText(v string) *SubmitImageTranslateTaskRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitImageTranslateTaskRequest
	GetWorkspaceId() *string
}

type SubmitImageTranslateTaskRequest struct {
	Ext *SubmitImageTranslateTaskRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// example:
	//
	// image
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flash
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	TargetLanguage []*string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty" type:"Repeated"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitImageTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskRequest) GetExt() *SubmitImageTranslateTaskRequestExt {
	return s.Ext
}

func (s *SubmitImageTranslateTaskRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitImageTranslateTaskRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitImageTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitImageTranslateTaskRequest) GetTargetLanguage() []*string {
	return s.TargetLanguage
}

func (s *SubmitImageTranslateTaskRequest) GetText() *string {
	return s.Text
}

func (s *SubmitImageTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitImageTranslateTaskRequest) SetExt(v *SubmitImageTranslateTaskRequestExt) *SubmitImageTranslateTaskRequest {
	s.Ext = v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetFormat(v string) *SubmitImageTranslateTaskRequest {
	s.Format = &v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetScene(v string) *SubmitImageTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetSourceLanguage(v string) *SubmitImageTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetTargetLanguage(v []*string) *SubmitImageTranslateTaskRequest {
	s.TargetLanguage = v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetText(v string) *SubmitImageTranslateTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitImageTranslateTaskRequest) SetWorkspaceId(v string) *SubmitImageTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitImageTranslateTaskRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImageTranslateTaskRequestExt struct {
	// example:
	//
	// technology
	DomainHint    *string                                            `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*SubmitImageTranslateTaskRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Sensitives    []*string                                          `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*SubmitImageTranslateTaskRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *SubmitImageTranslateTaskRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s SubmitImageTranslateTaskRequestExt) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskRequestExt) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *SubmitImageTranslateTaskRequestExt) GetExamples() []*SubmitImageTranslateTaskRequestExtExamples {
	return s.Examples
}

func (s *SubmitImageTranslateTaskRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *SubmitImageTranslateTaskRequestExt) GetTerminologies() []*SubmitImageTranslateTaskRequestExtTerminologies {
	return s.Terminologies
}

func (s *SubmitImageTranslateTaskRequestExt) GetTextTransform() *SubmitImageTranslateTaskRequestExtTextTransform {
	return s.TextTransform
}

func (s *SubmitImageTranslateTaskRequestExt) SetDomainHint(v string) *SubmitImageTranslateTaskRequestExt {
	s.DomainHint = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExt) SetExamples(v []*SubmitImageTranslateTaskRequestExtExamples) *SubmitImageTranslateTaskRequestExt {
	s.Examples = v
	return s
}

func (s *SubmitImageTranslateTaskRequestExt) SetSensitives(v []*string) *SubmitImageTranslateTaskRequestExt {
	s.Sensitives = v
	return s
}

func (s *SubmitImageTranslateTaskRequestExt) SetTerminologies(v []*SubmitImageTranslateTaskRequestExtTerminologies) *SubmitImageTranslateTaskRequestExt {
	s.Terminologies = v
	return s
}

func (s *SubmitImageTranslateTaskRequestExt) SetTextTransform(v *SubmitImageTranslateTaskRequestExtTextTransform) *SubmitImageTranslateTaskRequestExt {
	s.TextTransform = v
	return s
}

func (s *SubmitImageTranslateTaskRequestExt) Validate() error {
	if s.Examples != nil {
		for _, item := range s.Examples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Terminologies != nil {
		for _, item := range s.Terminologies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextTransform != nil {
		if err := s.TextTransform.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImageTranslateTaskRequestExtExamples struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitImageTranslateTaskRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskRequestExtExamples) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *SubmitImageTranslateTaskRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitImageTranslateTaskRequestExtExamples) SetSrc(v string) *SubmitImageTranslateTaskRequestExtExamples {
	s.Src = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtExamples) SetTgt(v string) *SubmitImageTranslateTaskRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type SubmitImageTranslateTaskRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// ML
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitImageTranslateTaskRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *SubmitImageTranslateTaskRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitImageTranslateTaskRequestExtTerminologies) SetSrc(v string) *SubmitImageTranslateTaskRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtTerminologies) SetTgt(v string) *SubmitImageTranslateTaskRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type SubmitImageTranslateTaskRequestExtTextTransform struct {
	// example:
	//
	// false
	ToLower *bool `json:"toLower,omitempty" xml:"toLower,omitempty"`
	// example:
	//
	// false
	ToTitle *bool `json:"toTitle,omitempty" xml:"toTitle,omitempty"`
	// example:
	//
	// false
	ToUpper *bool `json:"toUpper,omitempty" xml:"toUpper,omitempty"`
}

func (s SubmitImageTranslateTaskRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) SetToLower(v bool) *SubmitImageTranslateTaskRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) SetToTitle(v bool) *SubmitImageTranslateTaskRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) SetToUpper(v bool) *SubmitImageTranslateTaskRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *SubmitImageTranslateTaskRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
