// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalEnsAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryLocalEnsAssociationRequest
	GetDomainName() *string
	SetLang(v string) *QueryLocalEnsAssociationRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryLocalEnsAssociationRequest
	GetUserClientIp() *string
}

type QueryLocalEnsAssociationRequest struct {
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

func (s QueryLocalEnsAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalEnsAssociationRequest) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryLocalEnsAssociationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryLocalEnsAssociationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryLocalEnsAssociationRequest) SetDomainName(v string) *QueryLocalEnsAssociationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryLocalEnsAssociationRequest) SetLang(v string) *QueryLocalEnsAssociationRequest {
	s.Lang = &v
	return s
}

func (s *QueryLocalEnsAssociationRequest) SetUserClientIp(v string) *QueryLocalEnsAssociationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryLocalEnsAssociationRequest) Validate() error {
	return dara.Validate(s)
}
