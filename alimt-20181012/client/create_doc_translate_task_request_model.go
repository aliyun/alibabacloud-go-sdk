// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *CreateDocTranslateTaskRequest
	GetCallbackUrl() *string
	SetClientToken(v string) *CreateDocTranslateTaskRequest
	GetClientToken() *string
	SetFileUrl(v string) *CreateDocTranslateTaskRequest
	GetFileUrl() *string
	SetScene(v string) *CreateDocTranslateTaskRequest
	GetScene() *string
	SetSourceLanguage(v string) *CreateDocTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *CreateDocTranslateTaskRequest
	GetTargetLanguage() *string
}

type CreateDocTranslateTaskRequest struct {
	// example:
	//
	// http://callbackUrl
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://fileUrl
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s CreateDocTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateDocTranslateTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDocTranslateTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateDocTranslateTaskRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateDocTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *CreateDocTranslateTaskRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *CreateDocTranslateTaskRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetClientToken(v string) *CreateDocTranslateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetFileUrl(v string) *CreateDocTranslateTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetScene(v string) *CreateDocTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetSourceLanguage(v string) *CreateDocTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetTargetLanguage(v string) *CreateDocTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}
