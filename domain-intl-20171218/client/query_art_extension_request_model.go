// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArtExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryArtExtensionRequest
	GetDomainName() *string
	SetLang(v string) *QueryArtExtensionRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryArtExtensionRequest
	GetUserClientIp() *string
}

type QueryArtExtensionRequest struct {
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryArtExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryArtExtensionRequest) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryArtExtensionRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryArtExtensionRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryArtExtensionRequest) SetDomainName(v string) *QueryArtExtensionRequest {
	s.DomainName = &v
	return s
}

func (s *QueryArtExtensionRequest) SetLang(v string) *QueryArtExtensionRequest {
	s.Lang = &v
	return s
}

func (s *QueryArtExtensionRequest) SetUserClientIp(v string) *QueryArtExtensionRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryArtExtensionRequest) Validate() error {
	return dara.Validate(s)
}
