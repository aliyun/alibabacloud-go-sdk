// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginEmail(v string) *CreateAgAccountRequest
	GetLoginEmail() *string
	SetMpk(v string) *CreateAgAccountRequest
	GetMpk() *string
	SetNationCode(v string) *CreateAgAccountRequest
	GetNationCode() *string
	SetOwn(v string) *CreateAgAccountRequest
	GetOwn() *string
	SetRealParentPk(v string) *CreateAgAccountRequest
	GetRealParentPk() *string
	SetSecurityMobile(v string) *CreateAgAccountRequest
	GetSecurityMobile() *string
	SetShowNickName(v string) *CreateAgAccountRequest
	GetShowNickName() *string
	SetSiteNick(v string) *CreateAgAccountRequest
	GetSiteNick() *string
	SetSrcAccountInfo(v string) *CreateAgAccountRequest
	GetSrcAccountInfo() *string
}

type CreateAgAccountRequest struct {
	LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// This parameter is required.
	Mpk            *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	NationCode     *string `json:"NationCode,omitempty" xml:"NationCode,omitempty"`
	Own            *string `json:"Own,omitempty" xml:"Own,omitempty"`
	RealParentPk   *string `json:"RealParentPk,omitempty" xml:"RealParentPk,omitempty"`
	SecurityMobile *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
	ShowNickName   *string `json:"ShowNickName,omitempty" xml:"ShowNickName,omitempty"`
	SiteNick       *string `json:"SiteNick,omitempty" xml:"SiteNick,omitempty"`
	SrcAccountInfo *string `json:"srcAccountInfo,omitempty" xml:"srcAccountInfo,omitempty"`
}

func (s CreateAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAgAccountRequest) GetLoginEmail() *string {
	return s.LoginEmail
}

func (s *CreateAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *CreateAgAccountRequest) GetNationCode() *string {
	return s.NationCode
}

func (s *CreateAgAccountRequest) GetOwn() *string {
	return s.Own
}

func (s *CreateAgAccountRequest) GetRealParentPk() *string {
	return s.RealParentPk
}

func (s *CreateAgAccountRequest) GetSecurityMobile() *string {
	return s.SecurityMobile
}

func (s *CreateAgAccountRequest) GetShowNickName() *string {
	return s.ShowNickName
}

func (s *CreateAgAccountRequest) GetSiteNick() *string {
	return s.SiteNick
}

func (s *CreateAgAccountRequest) GetSrcAccountInfo() *string {
	return s.SrcAccountInfo
}

func (s *CreateAgAccountRequest) SetLoginEmail(v string) *CreateAgAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *CreateAgAccountRequest) SetMpk(v string) *CreateAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *CreateAgAccountRequest) SetNationCode(v string) *CreateAgAccountRequest {
	s.NationCode = &v
	return s
}

func (s *CreateAgAccountRequest) SetOwn(v string) *CreateAgAccountRequest {
	s.Own = &v
	return s
}

func (s *CreateAgAccountRequest) SetRealParentPk(v string) *CreateAgAccountRequest {
	s.RealParentPk = &v
	return s
}

func (s *CreateAgAccountRequest) SetSecurityMobile(v string) *CreateAgAccountRequest {
	s.SecurityMobile = &v
	return s
}

func (s *CreateAgAccountRequest) SetShowNickName(v string) *CreateAgAccountRequest {
	s.ShowNickName = &v
	return s
}

func (s *CreateAgAccountRequest) SetSiteNick(v string) *CreateAgAccountRequest {
	s.SiteNick = &v
	return s
}

func (s *CreateAgAccountRequest) SetSrcAccountInfo(v string) *CreateAgAccountRequest {
	s.SrcAccountInfo = &v
	return s
}

func (s *CreateAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
