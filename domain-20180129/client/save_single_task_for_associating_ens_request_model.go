// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAssociatingEnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveSingleTaskForAssociatingEnsRequest
	GetAddress() *string
	SetDomainName(v string) *SaveSingleTaskForAssociatingEnsRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForAssociatingEnsRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForAssociatingEnsRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForAssociatingEnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0x1234567890123456789012345678901234567890
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
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

func (s SaveSingleTaskForAssociatingEnsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveSingleTaskForAssociatingEnsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForAssociatingEnsRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForAssociatingEnsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetAddress(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.Address = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetDomainName(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetLang(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetUserClientIp(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) Validate() error {
	return dara.Validate(s)
}
