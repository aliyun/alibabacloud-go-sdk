// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnsAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryEnsAssociationRequest
	GetDomainName() *string
	SetLang(v string) *QueryEnsAssociationRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryEnsAssociationRequest
	GetUserClientIp() *string
}

type QueryEnsAssociationRequest struct {
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

func (s QueryEnsAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEnsAssociationRequest) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryEnsAssociationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryEnsAssociationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryEnsAssociationRequest) SetDomainName(v string) *QueryEnsAssociationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryEnsAssociationRequest) SetLang(v string) *QueryEnsAssociationRequest {
	s.Lang = &v
	return s
}

func (s *QueryEnsAssociationRequest) SetUserClientIp(v string) *QueryEnsAssociationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryEnsAssociationRequest) Validate() error {
	return dara.Validate(s)
}
