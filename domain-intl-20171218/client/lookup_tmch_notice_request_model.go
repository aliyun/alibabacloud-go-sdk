// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupTmchNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClaimKey(v string) *LookupTmchNoticeRequest
	GetClaimKey() *string
	SetLang(v string) *LookupTmchNoticeRequest
	GetLang() *string
	SetUserClientIp(v string) *LookupTmchNoticeRequest
	GetUserClientIp() *string
}

type LookupTmchNoticeRequest struct {
	// This parameter is required.
	ClaimKey     *string `json:"ClaimKey,omitempty" xml:"ClaimKey,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s LookupTmchNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeRequest) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeRequest) GetClaimKey() *string {
	return s.ClaimKey
}

func (s *LookupTmchNoticeRequest) GetLang() *string {
	return s.Lang
}

func (s *LookupTmchNoticeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *LookupTmchNoticeRequest) SetClaimKey(v string) *LookupTmchNoticeRequest {
	s.ClaimKey = &v
	return s
}

func (s *LookupTmchNoticeRequest) SetLang(v string) *LookupTmchNoticeRequest {
	s.Lang = &v
	return s
}

func (s *LookupTmchNoticeRequest) SetUserClientIp(v string) *LookupTmchNoticeRequest {
	s.UserClientIp = &v
	return s
}

func (s *LookupTmchNoticeRequest) Validate() error {
	return dara.Validate(s)
}
