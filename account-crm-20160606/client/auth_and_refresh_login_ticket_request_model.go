// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndRefreshLoginTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AuthAndRefreshLoginTicketRequest
	GetAppName() *string
	SetHavanaId(v string) *AuthAndRefreshLoginTicketRequest
	GetHavanaId() *string
	SetSessionId(v string) *AuthAndRefreshLoginTicketRequest
	GetSessionId() *string
}

type AuthAndRefreshLoginTicketRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaId  *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthAndRefreshLoginTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthAndRefreshLoginTicketRequest) GoString() string {
	return s.String()
}

func (s *AuthAndRefreshLoginTicketRequest) GetAppName() *string {
	return s.AppName
}

func (s *AuthAndRefreshLoginTicketRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *AuthAndRefreshLoginTicketRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthAndRefreshLoginTicketRequest) SetAppName(v string) *AuthAndRefreshLoginTicketRequest {
	s.AppName = &v
	return s
}

func (s *AuthAndRefreshLoginTicketRequest) SetHavanaId(v string) *AuthAndRefreshLoginTicketRequest {
	s.HavanaId = &v
	return s
}

func (s *AuthAndRefreshLoginTicketRequest) SetSessionId(v string) *AuthAndRefreshLoginTicketRequest {
	s.SessionId = &v
	return s
}

func (s *AuthAndRefreshLoginTicketRequest) Validate() error {
	return dara.Validate(s)
}
