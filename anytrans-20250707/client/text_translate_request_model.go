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
	// Extended parameters to control translation behavior
	Ext *TextTranslateRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// text format
	//
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// Model type
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// Source language code
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// Target Language Code
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// Text to be translated
	//
	// This parameter is required.
	//
	// example:
	//
	// 今天天气怎么样
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// Workspace ID
	//
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
	// Expert agent
	//
	// example:
	//
	// game
	Agent *string `json:"agent,omitempty" xml:"agent,omitempty"`
	// Translation Behavior Control
	Config *TextTranslateRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Domain hint
	//
	// example:
	//
	// technology
	DomainHint *string `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	// List of Translation Examples
	Examples   []*TextTranslateRequestExtExamples `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	LangDetect *bool                              `json:"langDetect,omitempty" xml:"langDetect,omitempty"`
	// Extended parameter configuration (bizUserld: A business-level user ID that distinguishes between different business users. It implements "user-based isolation" for terminology intervention so that interventions for one user do not affect others. bizType: A business scenario type or identifier that distinguishes between different scenarios. It implements "scenario-based isolation" for terminology intervention so that interventions for one scenario do not affect others.)
	//
	// example:
	//
	// {"bizUserld":"123456","bizType":session"}
	ParamMap interface{} `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
	// Prefix Configuration
	//
	// example:
	//
	// Today\\"s weather
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// Sensitive word list
	Sensitives []*string `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	// Translation terminology
	Terminologies []*TextTranslateRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	// Translated Text Conversion
	TextTransform *TextTranslateRequestExtTextTransform `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
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

func (s *TextTranslateRequestExt) GetLangDetect() *bool {
	return s.LangDetect
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

func (s *TextTranslateRequestExt) SetLangDetect(v bool) *TextTranslateRequestExt {
	s.LangDetect = &v
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
	// Whether to skip the Green Web check. (To use this option, you must first complete the process to disable Green Web before making the API call.)
	//
	// example:
	//
	// fasle
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
	// Source text
	//
	// example:
	//
	// 你好
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// Target text
	//
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
	// Source text
	//
	// example:
	//
	// 机器学习
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// Target text
	//
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
	// Convert to lowercase
	//
	// example:
	//
	// false
	ToLower *bool `json:"toLower,omitempty" xml:"toLower,omitempty"`
	// First letter capitalized
	//
	// example:
	//
	// false
	ToTitle *bool `json:"toTitle,omitempty" xml:"toTitle,omitempty"`
	// Convert to uppercase
	//
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
