// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AuthLoginTicketRequest
	GetAppName() *string
	SetAuthCode(v string) *AuthLoginTicketRequest
	GetAuthCode() *string
	SetMinorAuthCode(v string) *AuthLoginTicketRequest
	GetMinorAuthCode() *string
	SetScene(v string) *AuthLoginTicketRequest
	GetScene() *string
}

type AuthLoginTicketRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	MinorAuthCode *string `json:"MinorAuthCode,omitempty" xml:"MinorAuthCode,omitempty"`
	// This parameter is required.
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s AuthLoginTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginTicketRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginTicketRequest) GetAppName() *string {
	return s.AppName
}

func (s *AuthLoginTicketRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *AuthLoginTicketRequest) GetMinorAuthCode() *string {
	return s.MinorAuthCode
}

func (s *AuthLoginTicketRequest) GetScene() *string {
	return s.Scene
}

func (s *AuthLoginTicketRequest) SetAppName(v string) *AuthLoginTicketRequest {
	s.AppName = &v
	return s
}

func (s *AuthLoginTicketRequest) SetAuthCode(v string) *AuthLoginTicketRequest {
	s.AuthCode = &v
	return s
}

func (s *AuthLoginTicketRequest) SetMinorAuthCode(v string) *AuthLoginTicketRequest {
	s.MinorAuthCode = &v
	return s
}

func (s *AuthLoginTicketRequest) SetScene(v string) *AuthLoginTicketRequest {
	s.Scene = &v
	return s
}

func (s *AuthLoginTicketRequest) Validate() error {
	return dara.Validate(s)
}
