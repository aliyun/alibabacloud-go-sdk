// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermEditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TermEditRequest
	GetAction() *string
	SetExt(v *TermEditRequestExt) *TermEditRequest
	GetExt() *TermEditRequestExt
	SetScene(v string) *TermEditRequest
	GetScene() *string
	SetSourceLanguage(v string) *TermEditRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TermEditRequest
	GetTargetLanguage() *string
	SetWorkspaceId(v string) *TermEditRequest
	GetWorkspaceId() *string
}

type TermEditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ADD
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	Ext *TermEditRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s TermEditRequest) String() string {
	return dara.Prettify(s)
}

func (s TermEditRequest) GoString() string {
	return s.String()
}

func (s *TermEditRequest) GetAction() *string {
	return s.Action
}

func (s *TermEditRequest) GetExt() *TermEditRequestExt {
	return s.Ext
}

func (s *TermEditRequest) GetScene() *string {
	return s.Scene
}

func (s *TermEditRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TermEditRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TermEditRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *TermEditRequest) SetAction(v string) *TermEditRequest {
	s.Action = &v
	return s
}

func (s *TermEditRequest) SetExt(v *TermEditRequestExt) *TermEditRequest {
	s.Ext = v
	return s
}

func (s *TermEditRequest) SetScene(v string) *TermEditRequest {
	s.Scene = &v
	return s
}

func (s *TermEditRequest) SetSourceLanguage(v string) *TermEditRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TermEditRequest) SetTargetLanguage(v string) *TermEditRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TermEditRequest) SetWorkspaceId(v string) *TermEditRequest {
	s.WorkspaceId = &v
	return s
}

func (s *TermEditRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TermEditRequestExt struct {
	ParamMap interface{} `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
	// This parameter is required.
	Terms []*TermEditRequestExtTerms `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
}

func (s TermEditRequestExt) String() string {
	return dara.Prettify(s)
}

func (s TermEditRequestExt) GoString() string {
	return s.String()
}

func (s *TermEditRequestExt) GetParamMap() interface{} {
	return s.ParamMap
}

func (s *TermEditRequestExt) GetTerms() []*TermEditRequestExtTerms {
	return s.Terms
}

func (s *TermEditRequestExt) SetParamMap(v interface{}) *TermEditRequestExt {
	s.ParamMap = v
	return s
}

func (s *TermEditRequestExt) SetTerms(v []*TermEditRequestExtTerms) *TermEditRequestExt {
	s.Terms = v
	return s
}

func (s *TermEditRequestExt) Validate() error {
	if s.Terms != nil {
		for _, item := range s.Terms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TermEditRequestExtTerms struct {
	// This parameter is required.
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// 92669964
	TermId *string `json:"termId,omitempty" xml:"termId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LLM
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s TermEditRequestExtTerms) String() string {
	return dara.Prettify(s)
}

func (s TermEditRequestExtTerms) GoString() string {
	return s.String()
}

func (s *TermEditRequestExtTerms) GetSrc() *string {
	return s.Src
}

func (s *TermEditRequestExtTerms) GetTermId() *string {
	return s.TermId
}

func (s *TermEditRequestExtTerms) GetTgt() *string {
	return s.Tgt
}

func (s *TermEditRequestExtTerms) SetSrc(v string) *TermEditRequestExtTerms {
	s.Src = &v
	return s
}

func (s *TermEditRequestExtTerms) SetTermId(v string) *TermEditRequestExtTerms {
	s.TermId = &v
	return s
}

func (s *TermEditRequestExtTerms) SetTgt(v string) *TermEditRequestExtTerms {
	s.Tgt = &v
	return s
}

func (s *TermEditRequestExtTerms) Validate() error {
	return dara.Validate(s)
}
