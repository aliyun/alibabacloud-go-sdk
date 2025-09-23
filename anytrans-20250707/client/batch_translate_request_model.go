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
	AppName *string                   `json:"appName,omitempty" xml:"appName,omitempty"`
	Ext     *BatchTranslateRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type BatchTranslateRequestExt struct {
	Config *BatchTranslateRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// technology
	DomainHint    *string                                  `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Examples      []*BatchTranslateRequestExtExamples      `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Sensitives    []*string                                `json:"sensitives,omitempty" xml:"sensitives,omitempty" type:"Repeated"`
	Terminologies []*BatchTranslateRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
	TextTransform *BatchTranslateRequestExtTextTransform   `json:"textTransform,omitempty" xml:"textTransform,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type BatchTranslateRequestExtConfig struct {
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
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
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
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
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
