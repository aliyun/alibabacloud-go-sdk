// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLongTextTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *SubmitLongTextTranslateTaskRequestExt) *SubmitLongTextTranslateTaskRequest
	GetExt() *SubmitLongTextTranslateTaskRequestExt
	SetFormat(v string) *SubmitLongTextTranslateTaskRequest
	GetFormat() *string
	SetScene(v string) *SubmitLongTextTranslateTaskRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitLongTextTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitLongTextTranslateTaskRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitLongTextTranslateTaskRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitLongTextTranslateTaskRequest
	GetWorkspaceId() *string
}

type SubmitLongTextTranslateTaskRequest struct {
	Ext *SubmitLongTextTranslateTaskRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
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

func (s SubmitLongTextTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequest) GetExt() *SubmitLongTextTranslateTaskRequestExt {
	return s.Ext
}

func (s *SubmitLongTextTranslateTaskRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitLongTextTranslateTaskRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitLongTextTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitLongTextTranslateTaskRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitLongTextTranslateTaskRequest) GetText() *string {
	return s.Text
}

func (s *SubmitLongTextTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitLongTextTranslateTaskRequest) SetExt(v *SubmitLongTextTranslateTaskRequestExt) *SubmitLongTextTranslateTaskRequest {
	s.Ext = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetFormat(v string) *SubmitLongTextTranslateTaskRequest {
	s.Format = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetScene(v string) *SubmitLongTextTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetSourceLanguage(v string) *SubmitLongTextTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetTargetLanguage(v string) *SubmitLongTextTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetText(v string) *SubmitLongTextTranslateTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) SetWorkspaceId(v string) *SubmitLongTextTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitLongTextTranslateTaskRequestExt struct {
	Config *SubmitLongTextTranslateTaskRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// technology
	DomainHint    *string                                               `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*SubmitLongTextTranslateTaskRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	ParamMap      interface{}                                           `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
	Sensitives    []*string                                             `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*SubmitLongTextTranslateTaskRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *SubmitLongTextTranslateTaskRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s SubmitLongTextTranslateTaskRequestExt) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequestExt) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetConfig() *SubmitLongTextTranslateTaskRequestExtConfig {
	return s.Config
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetExamples() []*SubmitLongTextTranslateTaskRequestExtExamples {
	return s.Examples
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetParamMap() interface{} {
	return s.ParamMap
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetTerminologies() []*SubmitLongTextTranslateTaskRequestExtTerminologies {
	return s.Terminologies
}

func (s *SubmitLongTextTranslateTaskRequestExt) GetTextTransform() *SubmitLongTextTranslateTaskRequestExtTextTransform {
	return s.TextTransform
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetConfig(v *SubmitLongTextTranslateTaskRequestExtConfig) *SubmitLongTextTranslateTaskRequestExt {
	s.Config = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetDomainHint(v string) *SubmitLongTextTranslateTaskRequestExt {
	s.DomainHint = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetExamples(v []*SubmitLongTextTranslateTaskRequestExtExamples) *SubmitLongTextTranslateTaskRequestExt {
	s.Examples = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetParamMap(v interface{}) *SubmitLongTextTranslateTaskRequestExt {
	s.ParamMap = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetSensitives(v []*string) *SubmitLongTextTranslateTaskRequestExt {
	s.Sensitives = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetTerminologies(v []*SubmitLongTextTranslateTaskRequestExtTerminologies) *SubmitLongTextTranslateTaskRequestExt {
	s.Terminologies = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) SetTextTransform(v *SubmitLongTextTranslateTaskRequestExtTextTransform) *SubmitLongTextTranslateTaskRequestExt {
	s.TextTransform = v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExt) Validate() error {
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

type SubmitLongTextTranslateTaskRequestExtConfig struct {
	SkipCsiCheck *bool `json:"skipCsiCheck,omitempty" xml:"skipCsiCheck,omitempty"`
}

func (s SubmitLongTextTranslateTaskRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequestExtConfig) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequestExtConfig) GetSkipCsiCheck() *bool {
	return s.SkipCsiCheck
}

func (s *SubmitLongTextTranslateTaskRequestExtConfig) SetSkipCsiCheck(v bool) *SubmitLongTextTranslateTaskRequestExtConfig {
	s.SkipCsiCheck = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitLongTextTranslateTaskRequestExtExamples struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitLongTextTranslateTaskRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequestExtExamples) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *SubmitLongTextTranslateTaskRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitLongTextTranslateTaskRequestExtExamples) SetSrc(v string) *SubmitLongTextTranslateTaskRequestExtExamples {
	s.Src = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtExamples) SetTgt(v string) *SubmitLongTextTranslateTaskRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type SubmitLongTextTranslateTaskRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// ML
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitLongTextTranslateTaskRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *SubmitLongTextTranslateTaskRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitLongTextTranslateTaskRequestExtTerminologies) SetSrc(v string) *SubmitLongTextTranslateTaskRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtTerminologies) SetTgt(v string) *SubmitLongTextTranslateTaskRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type SubmitLongTextTranslateTaskRequestExtTextTransform struct {
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

func (s SubmitLongTextTranslateTaskRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) SetToLower(v bool) *SubmitLongTextTranslateTaskRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) SetToTitle(v bool) *SubmitLongTextTranslateTaskRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) SetToUpper(v bool) *SubmitLongTextTranslateTaskRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *SubmitLongTextTranslateTaskRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
