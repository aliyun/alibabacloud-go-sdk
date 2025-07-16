// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateDocTranslateTaskAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *CreateDocTranslateTaskAdvanceRequest
	GetCallbackUrl() *string
	SetClientToken(v string) *CreateDocTranslateTaskAdvanceRequest
	GetClientToken() *string
	SetFileUrlObject(v io.Reader) *CreateDocTranslateTaskAdvanceRequest
	GetFileUrlObject() io.Reader
	SetScene(v string) *CreateDocTranslateTaskAdvanceRequest
	GetScene() *string
	SetSourceLanguage(v string) *CreateDocTranslateTaskAdvanceRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *CreateDocTranslateTaskAdvanceRequest
	GetTargetLanguage() *string
}

type CreateDocTranslateTaskAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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

func (s CreateDocTranslateTaskAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocTranslateTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *CreateDocTranslateTaskAdvanceRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetClientToken(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocTranslateTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetScene(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.Scene = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetSourceLanguage(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetTargetLanguage(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
