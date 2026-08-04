// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndActiveWithHidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AuthAndActiveWithHidRequest
	GetAppName() *string
	SetHavanaId(v string) *AuthAndActiveWithHidRequest
	GetHavanaId() *string
	SetSessionId(v string) *AuthAndActiveWithHidRequest
	GetSessionId() *string
}

type AuthAndActiveWithHidRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaId  *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AuthAndActiveWithHidRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidRequest) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidRequest) GetAppName() *string {
	return s.AppName
}

func (s *AuthAndActiveWithHidRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *AuthAndActiveWithHidRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AuthAndActiveWithHidRequest) SetAppName(v string) *AuthAndActiveWithHidRequest {
	s.AppName = &v
	return s
}

func (s *AuthAndActiveWithHidRequest) SetHavanaId(v string) *AuthAndActiveWithHidRequest {
	s.HavanaId = &v
	return s
}

func (s *AuthAndActiveWithHidRequest) SetSessionId(v string) *AuthAndActiveWithHidRequest {
	s.SessionId = &v
	return s
}

func (s *AuthAndActiveWithHidRequest) Validate() error {
	return dara.Validate(s)
}
