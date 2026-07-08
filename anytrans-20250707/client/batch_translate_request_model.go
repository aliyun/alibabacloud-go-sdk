// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *BatchTranslateRequest
	GetAppName() *string
	SetExt(v *BatchTranslateRequestExt) *BatchTranslateRequest
	GetExt() *BatchTranslateRequestExt
	SetFormat(v string) *BatchTranslateRequest
	GetFormat() *string
	SetScene(v string) *BatchTranslateRequest
	GetScene() *string
	SetSourceLanguage(v string) *BatchTranslateRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *BatchTranslateRequest
	GetTargetLanguage() *string
	SetText(v map[string]interface{}) *BatchTranslateRequest
	GetText() map[string]interface{}
	SetWorkspaceId(v string) *BatchTranslateRequest
	GetWorkspaceId() *string
}

type BatchTranslateRequest struct {
	// The name of the calling application.
	//
	// example:
	//
	// baidufanyi
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// The extended parameters that control translation features.
	Ext *BatchTranslateRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// The translation format.
	//
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The translation model.
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The source language.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// The target language.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// A map of texts to translate, in which the key is a custom identifier and the value is the source text.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"0":"明天天气怎么样？","1":"你中午吃饭了吗"}
	Text map[string]interface{} `json:"text,omitempty" xml:"text,omitempty"`
	// The ID of the Model Studio workspace used for this request.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequest) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequest) GetAppName() *string {
	return s.AppName
}

func (s *BatchTranslateRequest) GetExt() *BatchTranslateRequestExt {
	return s.Ext
}

func (s *BatchTranslateRequest) GetFormat() *string {
	return s.Format
}

func (s *BatchTranslateRequest) GetScene() *string {
	return s.Scene
}

func (s *BatchTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *BatchTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *BatchTranslateRequest) GetText() map[string]interface{} {
	return s.Text
}

func (s *BatchTranslateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchTranslateRequest) SetAppName(v string) *BatchTranslateRequest {
	s.AppName = &v
	return s
}

func (s *BatchTranslateRequest) SetExt(v *BatchTranslateRequestExt) *BatchTranslateRequest {
	s.Ext = v
	return s
}

func (s *BatchTranslateRequest) SetFormat(v string) *BatchTranslateRequest {
	s.Format = &v
	return s
}

func (s *BatchTranslateRequest) SetScene(v string) *BatchTranslateRequest {
	s.Scene = &v
	return s
}

func (s *BatchTranslateRequest) SetSourceLanguage(v string) *BatchTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *BatchTranslateRequest) SetTargetLanguage(v string) *BatchTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *BatchTranslateRequest) SetText(v map[string]interface{}) *BatchTranslateRequest {
	s.Text = v
	return s
}

func (s *BatchTranslateRequest) SetWorkspaceId(v string) *BatchTranslateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchTranslateRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchTranslateRequestExt struct {
	// Controls the translation behavior.
	Config *BatchTranslateRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// A natural language instruction in English that guides the model\\"s translation style.
	//
	// example:
	//
	// this sentence from an e-commerce product image, please provide a translation that is both highly concise and no more than 1.2 times the length of the original.
	DomainHint *string `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	// A list of translation examples.
	Examples []*BatchTranslateRequestExtExamples `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	// Specifies whether to enable automatic detection of the source language. If set to true, the `sourceLanguage` parameter is ignored.
	LangDetect *bool `json:"langDetect,omitempty" xml:"langDetect,omitempty"`
	// Extended parameters for applying custom terminology that is isolated by user or business scenario.
	//
	// example:
	//
	// {"bizUserld":"123456","bizType":session"}
	ParamMap interface{} `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
	// A list of sensitive terms.
	Sensitives []*string `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	// A list of custom terminology for overriding translations.
	Terminologies []*BatchTranslateRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	// Specifies case transformations for the translated text.
	TextTransform *BatchTranslateRequestExtTextTransform `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
}

func (s BatchTranslateRequestExt) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequestExt) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequestExt) GetConfig() *BatchTranslateRequestExtConfig {
	return s.Config
}

func (s *BatchTranslateRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *BatchTranslateRequestExt) GetExamples() []*BatchTranslateRequestExtExamples {
	return s.Examples
}

func (s *BatchTranslateRequestExt) GetLangDetect() *bool {
	return s.LangDetect
}

func (s *BatchTranslateRequestExt) GetParamMap() interface{} {
	return s.ParamMap
}

func (s *BatchTranslateRequestExt) GetSensitives() []*string {
	return s.Sensitives
}

func (s *BatchTranslateRequestExt) GetTerminologies() []*BatchTranslateRequestExtTerminologies {
	return s.Terminologies
}

func (s *BatchTranslateRequestExt) GetTextTransform() *BatchTranslateRequestExtTextTransform {
	return s.TextTransform
}

func (s *BatchTranslateRequestExt) SetConfig(v *BatchTranslateRequestExtConfig) *BatchTranslateRequestExt {
	s.Config = v
	return s
}

func (s *BatchTranslateRequestExt) SetDomainHint(v string) *BatchTranslateRequestExt {
	s.DomainHint = &v
	return s
}

func (s *BatchTranslateRequestExt) SetExamples(v []*BatchTranslateRequestExtExamples) *BatchTranslateRequestExt {
	s.Examples = v
	return s
}

func (s *BatchTranslateRequestExt) SetLangDetect(v bool) *BatchTranslateRequestExt {
	s.LangDetect = &v
	return s
}

func (s *BatchTranslateRequestExt) SetParamMap(v interface{}) *BatchTranslateRequestExt {
	s.ParamMap = v
	return s
}

func (s *BatchTranslateRequestExt) SetSensitives(v []*string) *BatchTranslateRequestExt {
	s.Sensitives = v
	return s
}

func (s *BatchTranslateRequestExt) SetTerminologies(v []*BatchTranslateRequestExtTerminologies) *BatchTranslateRequestExt {
	s.Terminologies = v
	return s
}

func (s *BatchTranslateRequestExt) SetTextTransform(v *BatchTranslateRequestExtTextTransform) *BatchTranslateRequestExt {
	s.TextTransform = v
	return s
}

func (s *BatchTranslateRequestExt) Validate() error {
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

type BatchTranslateRequestExtConfig struct {
	// Specifies whether to skip the Content Moderation check. To set this to true, you must first complete the required process to disable Content Moderation.
	//
	// example:
	//
	// false
	SkipCsiCheck *bool `json:"skipCsiCheck,omitempty" xml:"skipCsiCheck,omitempty"`
}

func (s BatchTranslateRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequestExtConfig) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequestExtConfig) GetSkipCsiCheck() *bool {
	return s.SkipCsiCheck
}

func (s *BatchTranslateRequestExtConfig) SetSkipCsiCheck(v bool) *BatchTranslateRequestExtConfig {
	s.SkipCsiCheck = &v
	return s
}

func (s *BatchTranslateRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateRequestExtExamples struct {
	// The source text.
	//
	// example:
	//
	// 你好
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// The target text.
	//
	// example:
	//
	// hello
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s BatchTranslateRequestExtExamples) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequestExtExamples) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequestExtExamples) GetSrc() *string {
	return s.Src
}

func (s *BatchTranslateRequestExtExamples) GetTgt() *string {
	return s.Tgt
}

func (s *BatchTranslateRequestExtExamples) SetSrc(v string) *BatchTranslateRequestExtExamples {
	s.Src = &v
	return s
}

func (s *BatchTranslateRequestExtExamples) SetTgt(v string) *BatchTranslateRequestExtExamples {
	s.Tgt = &v
	return s
}

func (s *BatchTranslateRequestExtExamples) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateRequestExtTerminologies struct {
	// The source text to be overridden.
	//
	// example:
	//
	// 应用接口
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// The target text to use for the override.
	//
	// example:
	//
	// API
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s BatchTranslateRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *BatchTranslateRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *BatchTranslateRequestExtTerminologies) SetSrc(v string) *BatchTranslateRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *BatchTranslateRequestExtTerminologies) SetTgt(v string) *BatchTranslateRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *BatchTranslateRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}

type BatchTranslateRequestExtTextTransform struct {
	// Specifies whether to convert the entire translated text to lowercase.
	//
	// example:
	//
	// false
	ToLower *bool `json:"toLower,omitempty" xml:"toLower,omitempty"`
	// Specifies whether to convert the entire translated text to title case.
	//
	// example:
	//
	// false
	ToTitle *bool `json:"toTitle,omitempty" xml:"toTitle,omitempty"`
	// Specifies whether to convert the entire translated text to uppercase.
	//
	// example:
	//
	// false
	ToUpper *bool `json:"toUpper,omitempty" xml:"toUpper,omitempty"`
}

func (s BatchTranslateRequestExtTextTransform) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateRequestExtTextTransform) GoString() string {
	return s.String()
}

func (s *BatchTranslateRequestExtTextTransform) GetToLower() *bool {
	return s.ToLower
}

func (s *BatchTranslateRequestExtTextTransform) GetToTitle() *bool {
	return s.ToTitle
}

func (s *BatchTranslateRequestExtTextTransform) GetToUpper() *bool {
	return s.ToUpper
}

func (s *BatchTranslateRequestExtTextTransform) SetToLower(v bool) *BatchTranslateRequestExtTextTransform {
	s.ToLower = &v
	return s
}

func (s *BatchTranslateRequestExtTextTransform) SetToTitle(v bool) *BatchTranslateRequestExtTextTransform {
	s.ToTitle = &v
	return s
}

func (s *BatchTranslateRequestExtTextTransform) SetToUpper(v bool) *BatchTranslateRequestExtTextTransform {
	s.ToUpper = &v
	return s
}

func (s *BatchTranslateRequestExtTextTransform) Validate() error {
	return dara.Validate(s)
}
