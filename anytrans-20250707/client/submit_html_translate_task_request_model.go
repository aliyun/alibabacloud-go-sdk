// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHtmlTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *SubmitHtmlTranslateTaskRequestExt) *SubmitHtmlTranslateTaskRequest
	GetExt() *SubmitHtmlTranslateTaskRequestExt
	SetFormat(v string) *SubmitHtmlTranslateTaskRequest
	GetFormat() *string
	SetScene(v string) *SubmitHtmlTranslateTaskRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitHtmlTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitHtmlTranslateTaskRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitHtmlTranslateTaskRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitHtmlTranslateTaskRequest
	GetWorkspaceId() *string
}

type SubmitHtmlTranslateTaskRequest struct {
	Ext *SubmitHtmlTranslateTaskRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	Text           *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitHtmlTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequest) GetExt() *SubmitHtmlTranslateTaskRequestExt {
	return s.Ext
}

func (s *SubmitHtmlTranslateTaskRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitHtmlTranslateTaskRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitHtmlTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitHtmlTranslateTaskRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitHtmlTranslateTaskRequest) GetText() *string {
	return s.Text
}

func (s *SubmitHtmlTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitHtmlTranslateTaskRequest) SetExt(v *SubmitHtmlTranslateTaskRequestExt) *SubmitHtmlTranslateTaskRequest {
	s.Ext = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetFormat(v string) *SubmitHtmlTranslateTaskRequest {
	s.Format = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetScene(v string) *SubmitHtmlTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetSourceLanguage(v string) *SubmitHtmlTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetTargetLanguage(v string) *SubmitHtmlTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetText(v string) *SubmitHtmlTranslateTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) SetWorkspaceId(v string) *SubmitHtmlTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitHtmlTranslateTaskRequestExt struct {
	Config *SubmitHtmlTranslateTaskRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// technology
	DomainHint    *string                                           `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*SubmitHtmlTranslateTaskRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Sensitives    []*string                                         `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*SubmitHtmlTranslateTaskRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *SubmitHtmlTranslateTaskRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s SubmitHtmlTranslateTaskRequestExt) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequestExt) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetConfig() *SubmitHtmlTranslateTaskRequestExtConfig {
	return s.Config
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetExamples() []*SubmitHtmlTranslateTaskRequestExtExamples {
	return s.Examples
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetTerminologies() []*SubmitHtmlTranslateTaskRequestExtTerminologies {
	return s.Terminologies
}

func (s *SubmitHtmlTranslateTaskRequestExt) GetTextTransform() *SubmitHtmlTranslateTaskRequestExtTextTransform {
	return s.TextTransform
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetConfig(v *SubmitHtmlTranslateTaskRequestExtConfig) *SubmitHtmlTranslateTaskRequestExt {
	s.Config = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetDomainHint(v string) *SubmitHtmlTranslateTaskRequestExt {
	s.DomainHint = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetExamples(v []*SubmitHtmlTranslateTaskRequestExtExamples) *SubmitHtmlTranslateTaskRequestExt {
	s.Examples = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetSensitives(v []*string) *SubmitHtmlTranslateTaskRequestExt {
	s.Sensitives = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetTerminologies(v []*SubmitHtmlTranslateTaskRequestExtTerminologies) *SubmitHtmlTranslateTaskRequestExt {
	s.Terminologies = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) SetTextTransform(v *SubmitHtmlTranslateTaskRequestExtTextTransform) *SubmitHtmlTranslateTaskRequestExt {
	s.TextTransform = v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExt) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
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

type SubmitHtmlTranslateTaskRequestExtConfig struct {
	CallbackUrl  *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	SkipCsiCheck *bool   `json:"skipCsiCheck,omitempty" xml:"skipCsiCheck,omitempty"`
}

func (s SubmitHtmlTranslateTaskRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequestExtConfig) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequestExtConfig) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SubmitHtmlTranslateTaskRequestExtConfig) GetSkipCsiCheck() *bool {
	return s.SkipCsiCheck
}

func (s *SubmitHtmlTranslateTaskRequestExtConfig) SetCallbackUrl(v string) *SubmitHtmlTranslateTaskRequestExtConfig {
	s.CallbackUrl = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtConfig) SetSkipCsiCheck(v bool) *SubmitHtmlTranslateTaskRequestExtConfig {
	s.SkipCsiCheck = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitHtmlTranslateTaskRequestExtExamples struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitHtmlTranslateTaskRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequestExtExamples) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *SubmitHtmlTranslateTaskRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitHtmlTranslateTaskRequestExtExamples) SetSrc(v string) *SubmitHtmlTranslateTaskRequestExtExamples {
	s.Src = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtExamples) SetTgt(v string) *SubmitHtmlTranslateTaskRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type SubmitHtmlTranslateTaskRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// ML
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitHtmlTranslateTaskRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *SubmitHtmlTranslateTaskRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitHtmlTranslateTaskRequestExtTerminologies) SetSrc(v string) *SubmitHtmlTranslateTaskRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtTerminologies) SetTgt(v string) *SubmitHtmlTranslateTaskRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type SubmitHtmlTranslateTaskRequestExtTextTransform struct {
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

func (s SubmitHtmlTranslateTaskRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) SetToLower(v bool) *SubmitHtmlTranslateTaskRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) SetToTitle(v bool) *SubmitHtmlTranslateTaskRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) SetToUpper(v bool) *SubmitHtmlTranslateTaskRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *SubmitHtmlTranslateTaskRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
