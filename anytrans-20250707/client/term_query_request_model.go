// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v *TermQueryRequestExt) *TermQueryRequest
	GetExt() *TermQueryRequestExt
	SetScene(v string) *TermQueryRequest
	GetScene() *string
	SetSourceLanguage(v string) *TermQueryRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TermQueryRequest
	GetTargetLanguage() *string
	SetText(v string) *TermQueryRequest
	GetText() *string
	SetWorkspaceId(v string) *TermQueryRequest
	GetWorkspaceId() *string
}

type TermQueryRequest struct {
	Ext *TermQueryRequestExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
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
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	//
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

func (s TermQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TermQueryRequest) GoString() string {
	return s.String()
}

func (s *TermQueryRequest) GetExt() *TermQueryRequestExt {
	return s.Ext
}

func (s *TermQueryRequest) GetScene() *string {
	return s.Scene
}

func (s *TermQueryRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TermQueryRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TermQueryRequest) GetText() *string {
	return s.Text
}

func (s *TermQueryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *TermQueryRequest) SetExt(v *TermQueryRequestExt) *TermQueryRequest {
	s.Ext = v
	return s
}

func (s *TermQueryRequest) SetScene(v string) *TermQueryRequest {
	s.Scene = &v
	return s
}

func (s *TermQueryRequest) SetSourceLanguage(v string) *TermQueryRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TermQueryRequest) SetTargetLanguage(v string) *TermQueryRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TermQueryRequest) SetText(v string) *TermQueryRequest {
	s.Text = &v
	return s
}

func (s *TermQueryRequest) SetWorkspaceId(v string) *TermQueryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *TermQueryRequest) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TermQueryRequestExt struct {
	ParamMap interface{} `json:"paramMap,omitempty" xml:"paramMap,omitempty"`
}

func (s TermQueryRequestExt) String() string {
	return dara.Prettify(s)
}

func (s TermQueryRequestExt) GoString() string {
	return s.String()
}

func (s *TermQueryRequestExt) GetParamMap() interface{} {
	return s.ParamMap
}

func (s *TermQueryRequestExt) SetParamMap(v interface{}) *TermQueryRequestExt {
	s.ParamMap = v
	return s
}

func (s *TermQueryRequestExt) Validate() error {
	return dara.Validate(s)
}
