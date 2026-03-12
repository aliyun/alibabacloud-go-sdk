// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v int64) *QueryFailReasonListRequest
	GetContactTemplateId() *int64
	SetDomainName(v string) *QueryFailReasonListRequest
	GetDomainName() *string
	SetLang(v string) *QueryFailReasonListRequest
	GetLang() *string
	SetSaleId(v string) *QueryFailReasonListRequest
	GetSaleId() *string
	SetUserClientIp(v string) *QueryFailReasonListRequest
	GetUserClientIp() *string
}

type QueryFailReasonListRequest struct {
	ContactTemplateId *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SaleId            *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailReasonListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonListRequest) GoString() string {
	return s.String()
}

func (s *QueryFailReasonListRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *QueryFailReasonListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryFailReasonListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryFailReasonListRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryFailReasonListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryFailReasonListRequest) SetContactTemplateId(v int64) *QueryFailReasonListRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *QueryFailReasonListRequest) SetDomainName(v string) *QueryFailReasonListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryFailReasonListRequest) SetLang(v string) *QueryFailReasonListRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailReasonListRequest) SetSaleId(v string) *QueryFailReasonListRequest {
	s.SaleId = &v
	return s
}

func (s *QueryFailReasonListRequest) SetUserClientIp(v string) *QueryFailReasonListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryFailReasonListRequest) Validate() error {
	return dara.Validate(s)
}
