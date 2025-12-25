// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *TextTranslateRequestExt) *TextTranslateRequest
	GetExt() *TextTranslateRequestExt
	SetFormat(v string) *TextTranslateRequest
	GetFormat() *string
	SetScene(v string) *TextTranslateRequest
	GetScene() *string
	SetSourceLanguage(v string) *TextTranslateRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TextTranslateRequest
	GetTargetLanguage() *string
	SetText(v string) *TextTranslateRequest
	GetText() *string
	SetWorkspaceId(v string) *TextTranslateRequest
	GetWorkspaceId() *string
}

type TextTranslateRequest struct {
	Ext *TextTranslateRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s TextTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateRequest) GetExt() *TextTranslateRequestExt {
	return s.Ext
}

func (s *TextTranslateRequest) GetFormat() *string {
	return s.Format
}

func (s *TextTranslateRequest) GetScene() *string {
	return s.Scene
}

func (s *TextTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateRequest) GetText() *string {
	return s.Text
}

func (s *TextTranslateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *TextTranslateRequest) SetExt(v *TextTranslateRequestExt) *TextTranslateRequest {
	s.Ext = v
	return s
}

func (s *TextTranslateRequest) SetFormat(v string) *TextTranslateRequest {
	s.Format = &v
	return s
}

func (s *TextTranslateRequest) SetScene(v string) *TextTranslateRequest {
	s.Scene = &v
	return s
}

func (s *TextTranslateRequest) SetSourceLanguage(v string) *TextTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateRequest) SetTargetLanguage(v string) *TextTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateRequest) SetText(v string) *TextTranslateRequest {
	s.Text = &v
	return s
}

func (s *TextTranslateRequest) SetWorkspaceId(v string) *TextTranslateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *TextTranslateRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextTranslateRequestExt struct {
	Agent  *string                        `json:"agent,omitempty" xml:"agent,omitempty"`
	Config *TextTranslateRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// technology
	DomainHint    *string                                 `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*TextTranslateRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	ParamMap      interface{}                             `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
	Prefix        *string                                 `json:"prefix,omitempty" xml:"prefix,omitempty"`
	Sensitives    []*string                               `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*TextTranslateRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *TextTranslateRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s TextTranslateRequestExt) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequestExt) GoString() string {
	return s.String()
}

func (s *TextTranslateRequestExt) GetAgent() *string {
	return s.Agent
}

func (s *TextTranslateRequestExt) GetConfig() *TextTranslateRequestExtConfig {
	return s.Config
}

func (s *TextTranslateRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *TextTranslateRequestExt) GetExamples() []*TextTranslateRequestExtExamples {
	return s.Examples
}

func (s *TextTranslateRequestExt) GetParamMap() interface{} {
	return s.ParamMap
}

func (s *TextTranslateRequestExt) GetPrefix() *string {
	return s.Prefix
}

func (s *TextTranslateRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *TextTranslateRequestExt) GetTerminologies() []*TextTranslateRequestExtTerminologies {
	return s.Terminologies
}

func (s *TextTranslateRequestExt) GetTextTransform() *TextTranslateRequestExtTextTransform {
	return s.TextTransform
}

func (s *TextTranslateRequestExt) SetAgent(v string) *TextTranslateRequestExt {
	s.Agent = &v
	return s
}

func (s *TextTranslateRequestExt) SetConfig(v *TextTranslateRequestExtConfig) *TextTranslateRequestExt {
	s.Config = v
	return s
}

func (s *TextTranslateRequestExt) SetDomainHint(v string) *TextTranslateRequestExt {
	s.DomainHint = &v
	return s
}

func (s *TextTranslateRequestExt) SetExamples(v []*TextTranslateRequestExtExamples) *TextTranslateRequestExt {
	s.Examples = v
	return s
}

func (s *TextTranslateRequestExt) SetParamMap(v interface{}) *TextTranslateRequestExt {
	s.ParamMap = v
	return s
}

func (s *TextTranslateRequestExt) SetPrefix(v string) *TextTranslateRequestExt {
	s.Prefix = &v
	return s
}

func (s *TextTranslateRequestExt) SetSensitives(v []*string) *TextTranslateRequestExt {
	s.Sensitives = v
	return s
}

func (s *TextTranslateRequestExt) SetTerminologies(v []*TextTranslateRequestExtTerminologies) *TextTranslateRequestExt {
	s.Terminologies = v
	return s
}

func (s *TextTranslateRequestExt) SetTextTransform(v *TextTranslateRequestExtTextTransform) *TextTranslateRequestExt {
	s.TextTransform = v
	return s
}

func (s *TextTranslateRequestExt) Validate() error {
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

type TextTranslateRequestExtConfig struct {
	SkipCsiCheck *bool `json:"skipCsiCheck,omitempty" xml:"skipCsiCheck,omitempty"`
}

func (s TextTranslateRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequestExtConfig) GoString() string {
	return s.String()
}

func (s *TextTranslateRequestExtConfig) GetSkipCsiCheck() *bool {
	return s.SkipCsiCheck
}

func (s *TextTranslateRequestExtConfig) SetSkipCsiCheck(v bool) *TextTranslateRequestExtConfig {
	s.SkipCsiCheck = &v
	return s
}

func (s *TextTranslateRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type TextTranslateRequestExtExamples struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s TextTranslateRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequestExtExamples) GoString() string {
	return s.String()
}

func (s *TextTranslateRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *TextTranslateRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *TextTranslateRequestExtExamples) SetSrc(v string) *TextTranslateRequestExtExamples {
	s.Src = &v
	return s
}

func (s *TextTranslateRequestExtExamples) SetTgt(v string) *TextTranslateRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *TextTranslateRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type TextTranslateRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// ML
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s TextTranslateRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *TextTranslateRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *TextTranslateRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *TextTranslateRequestExtTerminologies) SetSrc(v string) *TextTranslateRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *TextTranslateRequestExtTerminologies) SetTgt(v string) *TextTranslateRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *TextTranslateRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type TextTranslateRequestExtTextTransform struct {
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

func (s TextTranslateRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *TextTranslateRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *TextTranslateRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *TextTranslateRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *TextTranslateRequestExtTextTransform) SetToLower(v bool) *TextTranslateRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *TextTranslateRequestExtTextTransform) SetToTitle(v bool) *TextTranslateRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *TextTranslateRequestExtTextTransform) SetToUpper(v bool) *TextTranslateRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *TextTranslateRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
