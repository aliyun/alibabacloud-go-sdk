// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDisassociatingEnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForDisassociatingEnsRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForDisassociatingEnsRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForDisassociatingEnsRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForDisassociatingEnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test.luxe
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDisassociatingEnsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetDomainName(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetLang(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetUserClientIp(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) Validate() error {
	return dara.Validate(s)
}
