// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainBySaleIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryDomainBySaleIdRequest
	GetLang() *string
	SetSaleId(v string) *QueryDomainBySaleIdRequest
	GetSaleId() *string
	SetUserClientIp(v string) *QueryDomainBySaleIdRequest
	GetUserClientIp() *string
}

type QueryDomainBySaleIdRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	SaleId       *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainBySaleIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainBySaleIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainBySaleIdRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainBySaleIdRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainBySaleIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainBySaleIdRequest) SetLang(v string) *QueryDomainBySaleIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainBySaleIdRequest) SetSaleId(v string) *QueryDomainBySaleIdRequest {
	s.SaleId = &v
	return s
}

func (s *QueryDomainBySaleIdRequest) SetUserClientIp(v string) *QueryDomainBySaleIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainBySaleIdRequest) Validate() error {
	return dara.Validate(s)
}
