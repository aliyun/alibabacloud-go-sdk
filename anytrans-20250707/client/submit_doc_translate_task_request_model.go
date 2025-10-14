// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *SubmitDocTranslateTaskRequestExt) *SubmitDocTranslateTaskRequest
	GetExt() *SubmitDocTranslateTaskRequestExt
	SetFormat(v string) *SubmitDocTranslateTaskRequest
	GetFormat() *string
	SetScene(v string) *SubmitDocTranslateTaskRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitDocTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitDocTranslateTaskRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitDocTranslateTaskRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitDocTranslateTaskRequest
	GetWorkspaceId() *string
}

type SubmitDocTranslateTaskRequest struct {
	Ext *SubmitDocTranslateTaskRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spoke-llm
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx-hangzhou.aliyuncs.com/docs/tmp/%E6%A0%B7%E4%BE%8B_%E6%97%A0%E5%9B%BE.pdf
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xc
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitDocTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskRequest) GetExt() *SubmitDocTranslateTaskRequestExt {
	return s.Ext
}

func (s *SubmitDocTranslateTaskRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitDocTranslateTaskRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitDocTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitDocTranslateTaskRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitDocTranslateTaskRequest) GetText() *string {
	return s.Text
}

func (s *SubmitDocTranslateTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitDocTranslateTaskRequest) SetExt(v *SubmitDocTranslateTaskRequestExt) *SubmitDocTranslateTaskRequest {
	s.Ext = v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetFormat(v string) *SubmitDocTranslateTaskRequest {
	s.Format = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetScene(v string) *SubmitDocTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetSourceLanguage(v string) *SubmitDocTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetTargetLanguage(v string) *SubmitDocTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetText(v string) *SubmitDocTranslateTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) SetWorkspaceId(v string) *SubmitDocTranslateTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitDocTranslateTaskRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDocTranslateTaskRequestExt struct {
	Config *SubmitDocTranslateTaskRequestExtConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// This text comes from a rigorous academic paper. Please provide a translation that complies with academic standards.
	DomainHint    *string                                          `json:"domainHint,omitempty" xml:"domainHint,omitempty"`
	Terminologies []*SubmitDocTranslateTaskRequestExtTerminologies `json:"terminologies,omitempty" xml:"terminologies,omitempty" type:"Repeated"`
}

func (s SubmitDocTranslateTaskRequestExt) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskRequestExt) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskRequestExt) GetConfig() *SubmitDocTranslateTaskRequestExtConfig {
	return s.Config
}

func (s *SubmitDocTranslateTaskRequestExt) GetDomainHint() *string {
	return s.DomainHint
}

func (s *SubmitDocTranslateTaskRequestExt) GetTerminologies() []*SubmitDocTranslateTaskRequestExtTerminologies {
	return s.Terminologies
}

func (s *SubmitDocTranslateTaskRequestExt) SetConfig(v *SubmitDocTranslateTaskRequestExtConfig) *SubmitDocTranslateTaskRequestExt {
	s.Config = v
	return s
}

func (s *SubmitDocTranslateTaskRequestExt) SetDomainHint(v string) *SubmitDocTranslateTaskRequestExt {
	s.DomainHint = &v
	return s
}

func (s *SubmitDocTranslateTaskRequestExt) SetTerminologies(v []*SubmitDocTranslateTaskRequestExtTerminologies) *SubmitDocTranslateTaskRequestExt {
	s.Terminologies = v
	return s
}

func (s *SubmitDocTranslateTaskRequestExt) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
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
	return nil
}

type SubmitDocTranslateTaskRequestExtConfig struct {
	SkipImgTrans *bool `json:"skipImgTrans,omitempty" xml:"skipImgTrans,omitempty"`
}

func (s SubmitDocTranslateTaskRequestExtConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskRequestExtConfig) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskRequestExtConfig) GetSkipImgTrans() *bool {
	return s.SkipImgTrans
}

func (s *SubmitDocTranslateTaskRequestExtConfig) SetSkipImgTrans(v bool) *SubmitDocTranslateTaskRequestExtConfig {
	s.SkipImgTrans = &v
	return s
}

func (s *SubmitDocTranslateTaskRequestExtConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitDocTranslateTaskRequestExtTerminologies struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// llm
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s SubmitDocTranslateTaskRequestExtTerminologies) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskRequestExtTerminologies) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskRequestExtTerminologies) GetSrc() *string {
	return s.Src
}

func (s *SubmitDocTranslateTaskRequestExtTerminologies) GetTgt() *string {
	return s.Tgt
}

func (s *SubmitDocTranslateTaskRequestExtTerminologies) SetSrc(v string) *SubmitDocTranslateTaskRequestExtTerminologies {
	s.Src = &v
	return s
}

func (s *SubmitDocTranslateTaskRequestExtTerminologies) SetTgt(v string) *SubmitDocTranslateTaskRequestExtTerminologies {
	s.Tgt = &v
	return s
}

func (s *SubmitDocTranslateTaskRequestExtTerminologies) Validate() error {
	return dara.Validate(s)
}
