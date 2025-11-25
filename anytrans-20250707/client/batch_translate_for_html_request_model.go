// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateForHtmlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *BatchTranslateForHtmlRequest
	GetAppName() *string
	SetExt(v *BatchTranslateForHtmlRequestExt) *BatchTranslateForHtmlRequest
	GetExt() *BatchTranslateForHtmlRequestExt
	SetFormat(v string) *BatchTranslateForHtmlRequest
	GetFormat() *string
	SetScene(v string) *BatchTranslateForHtmlRequest
	GetScene() *string
	SetSourceLanguage(v string) *BatchTranslateForHtmlRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *BatchTranslateForHtmlRequest
	GetTargetLanguage() *string
	SetText(v map[string]interface{}) *BatchTranslateForHtmlRequest
	GetText() map[string]interface{}
	SetWorkspaceId(v string) *BatchTranslateForHtmlRequest
	GetWorkspaceId() *string
}

type BatchTranslateForHtmlRequest struct {
	// example:
	//
	// baidufanyi
	AppName *string                          `json:"appName,omitempty" xml:"appName,omitempty"`
	Ext     *BatchTranslateForHtmlRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
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
	Text map[string]interface{} `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchTranslateForHtmlRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequest) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequest) GetAppName() *string {
	return s.AppName
}

func (s *BatchTranslateForHtmlRequest) GetExt() *BatchTranslateForHtmlRequestExt {
	return s.Ext
}

func (s *BatchTranslateForHtmlRequest) GetFormat() *string {
	return s.Format
}

func (s *BatchTranslateForHtmlRequest) GetScene() *string {
	return s.Scene
}

func (s *BatchTranslateForHtmlRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *BatchTranslateForHtmlRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *BatchTranslateForHtmlRequest) GetText() map[string]interface{} {
	return s.Text
}

func (s *BatchTranslateForHtmlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchTranslateForHtmlRequest) SetAppName(v string) *BatchTranslateForHtmlRequest {
	s.AppName = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetExt(v *BatchTranslateForHtmlRequestExt) *BatchTranslateForHtmlRequest {
	s.Ext = v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetFormat(v string) *BatchTranslateForHtmlRequest {
	s.Format = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetScene(v string) *BatchTranslateForHtmlRequest {
	s.Scene = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetSourceLanguage(v string) *BatchTranslateForHtmlRequest {
	s.SourceLanguage = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetTargetLanguage(v string) *BatchTranslateForHtmlRequest {
	s.TargetLanguage = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetText(v map[string]interface{}) *BatchTranslateForHtmlRequest {
	s.Text = v
	return s
}

func (s *BatchTranslateForHtmlRequest) SetWorkspaceId(v string) *BatchTranslateForHtmlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchTranslateForHtmlRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateForHtmlRequestExt struct {
	Config *BatchTranslateForHtmlRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// this sentence from an e-commerce product image, please provide a translation that is both highly concise and no more than 1.2 times the length of the original.
	DomainHint    *string                                         `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*BatchTranslateForHtmlRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Sensitives    []*string                                       `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*BatchTranslateForHtmlRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *BatchTranslateForHtmlRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s BatchTranslateForHtmlRequestExt) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequestExt) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequestExt) GetConfig() *BatchTranslateForHtmlRequestExtConfig {
	return s.Config
}

func (s *BatchTranslateForHtmlRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *BatchTranslateForHtmlRequestExt) GetExamples() []*BatchTranslateForHtmlRequestExtExamples {
	return s.Examples
}

func (s *BatchTranslateForHtmlRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *BatchTranslateForHtmlRequestExt) GetTerminologies() []*BatchTranslateForHtmlRequestExtTerminologies {
	return s.Terminologies
}

func (s *BatchTranslateForHtmlRequestExt) GetTextTransform() *BatchTranslateForHtmlRequestExtTextTransform {
	return s.TextTransform
}

func (s *BatchTranslateForHtmlRequestExt) SetConfig(v *BatchTranslateForHtmlRequestExtConfig) *BatchTranslateForHtmlRequestExt {
	s.Config = v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) SetDomainHint(v string) *BatchTranslateForHtmlRequestExt {
	s.DomainHint = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) SetExamples(v []*BatchTranslateForHtmlRequestExtExamples) *BatchTranslateForHtmlRequestExt {
	s.Examples = v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) SetSensitives(v []*string) *BatchTranslateForHtmlRequestExt {
	s.Sensitives = v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) SetTerminologies(v []*BatchTranslateForHtmlRequestExtTerminologies) *BatchTranslateForHtmlRequestExt {
	s.Terminologies = v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) SetTextTransform(v *BatchTranslateForHtmlRequestExtTextTransform) *BatchTranslateForHtmlRequestExt {
	s.TextTransform = v
	return s
}

func (s *BatchTranslateForHtmlRequestExt) Validate() error {
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

type BatchTranslateForHtmlRequestExtConfig struct {
	// example:
	//
	// fasle
	SkipCsiCheck *bool `json:"skipCsiCheck,omitempty" xml:"skipCsiCheck,omitempty"`
}

func (s BatchTranslateForHtmlRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequestExtConfig) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequestExtConfig) GetSkipCsiCheck() *bool {
	return s.SkipCsiCheck
}

func (s *BatchTranslateForHtmlRequestExtConfig) SetSkipCsiCheck(v bool) *BatchTranslateForHtmlRequestExtConfig {
	s.SkipCsiCheck = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateForHtmlRequestExtExamples struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s BatchTranslateForHtmlRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequestExtExamples) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *BatchTranslateForHtmlRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *BatchTranslateForHtmlRequestExtExamples) SetSrc(v string) *BatchTranslateForHtmlRequestExtExamples {
	s.Src = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtExamples) SetTgt(v string) *BatchTranslateForHtmlRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateForHtmlRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// API
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s BatchTranslateForHtmlRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *BatchTranslateForHtmlRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *BatchTranslateForHtmlRequestExtTerminologies) SetSrc(v string) *BatchTranslateForHtmlRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtTerminologies) SetTgt(v string) *BatchTranslateForHtmlRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateForHtmlRequestExtTextTransform struct {
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

func (s BatchTranslateForHtmlRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) SetToLower(v bool) *BatchTranslateForHtmlRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) SetToTitle(v bool) *BatchTranslateForHtmlRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) SetToUpper(v bool) *BatchTranslateForHtmlRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *BatchTranslateForHtmlRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
