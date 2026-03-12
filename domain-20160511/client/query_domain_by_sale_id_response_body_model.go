// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainBySaleIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChineseContactPerson(v string) *QueryDomainBySaleIdResponseBody
	GetChineseContactPerson() *string
	SetChineseHolder(v string) *QueryDomainBySaleIdResponseBody
	GetChineseHolder() *string
	SetCreationDate(v string) *QueryDomainBySaleIdResponseBody
	GetCreationDate() *string
	SetDnsList(v *QueryDomainBySaleIdResponseBodyDnsList) *QueryDomainBySaleIdResponseBody
	GetDnsList() *QueryDomainBySaleIdResponseBodyDnsList
	SetDomainName(v string) *QueryDomainBySaleIdResponseBody
	GetDomainName() *string
	SetDomainRegType(v string) *QueryDomainBySaleIdResponseBody
	GetDomainRegType() *string
	SetEmailVerificationClientHold(v bool) *QueryDomainBySaleIdResponseBody
	GetEmailVerificationClientHold() *bool
	SetEmailVerificationStatus(v int32) *QueryDomainBySaleIdResponseBody
	GetEmailVerificationStatus() *int32
	SetEnglishContactPerson(v string) *QueryDomainBySaleIdResponseBody
	GetEnglishContactPerson() *string
	SetEnglishHolder(v string) *QueryDomainBySaleIdResponseBody
	GetEnglishHolder() *string
	SetExpirationDate(v string) *QueryDomainBySaleIdResponseBody
	GetExpirationDate() *string
	SetHolderEmail(v string) *QueryDomainBySaleIdResponseBody
	GetHolderEmail() *string
	SetPremium(v bool) *QueryDomainBySaleIdResponseBody
	GetPremium() *bool
	SetRemark(v string) *QueryDomainBySaleIdResponseBody
	GetRemark() *string
	SetSafetyLock(v string) *QueryDomainBySaleIdResponseBody
	GetSafetyLock() *string
	SetSaleId(v string) *QueryDomainBySaleIdResponseBody
	GetSaleId() *string
	SetTransferLock(v string) *QueryDomainBySaleIdResponseBody
	GetTransferLock() *string
	SetTransferOutStatus(v string) *QueryDomainBySaleIdResponseBody
	GetTransferOutStatus() *string
	SetUserId(v string) *QueryDomainBySaleIdResponseBody
	GetUserId() *string
	SetWhoisProtected(v bool) *QueryDomainBySaleIdResponseBody
	GetWhoisProtected() *bool
}

type QueryDomainBySaleIdResponseBody struct {
	ChineseContactPerson        *string                                 `json:"ChineseContactPerson,omitempty" xml:"ChineseContactPerson,omitempty"`
	ChineseHolder               *string                                 `json:"ChineseHolder,omitempty" xml:"ChineseHolder,omitempty"`
	CreationDate                *string                                 `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	DnsList                     *QueryDomainBySaleIdResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainName                  *string                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRegType               *string                                 `json:"DomainRegType,omitempty" xml:"DomainRegType,omitempty"`
	EmailVerificationClientHold *bool                                   `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	EmailVerificationStatus     *int32                                  `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	EnglishContactPerson        *string                                 `json:"EnglishContactPerson,omitempty" xml:"EnglishContactPerson,omitempty"`
	EnglishHolder               *string                                 `json:"EnglishHolder,omitempty" xml:"EnglishHolder,omitempty"`
	ExpirationDate              *string                                 `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	HolderEmail                 *string                                 `json:"HolderEmail,omitempty" xml:"HolderEmail,omitempty"`
	Premium                     *bool                                   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	Remark                      *string                                 `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SafetyLock                  *string                                 `json:"SafetyLock,omitempty" xml:"SafetyLock,omitempty"`
	SaleId                      *string                                 `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	TransferLock                *string                                 `json:"TransferLock,omitempty" xml:"TransferLock,omitempty"`
	TransferOutStatus           *string                                 `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	UserId                      *string                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WhoisProtected              *bool                                   `json:"WhoisProtected,omitempty" xml:"WhoisProtected,omitempty"`
}

func (s QueryDomainBySaleIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainBySaleIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainBySaleIdResponseBody) GetChineseContactPerson() *string {
	return s.ChineseContactPerson
}

func (s *QueryDomainBySaleIdResponseBody) GetChineseHolder() *string {
	return s.ChineseHolder
}

func (s *QueryDomainBySaleIdResponseBody) GetCreationDate() *string {
	return s.CreationDate
}

func (s *QueryDomainBySaleIdResponseBody) GetDnsList() *QueryDomainBySaleIdResponseBodyDnsList {
	return s.DnsList
}

func (s *QueryDomainBySaleIdResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainBySaleIdResponseBody) GetDomainRegType() *string {
	return s.DomainRegType
}

func (s *QueryDomainBySaleIdResponseBody) GetEmailVerificationClientHold() *bool {
	return s.EmailVerificationClientHold
}

func (s *QueryDomainBySaleIdResponseBody) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryDomainBySaleIdResponseBody) GetEnglishContactPerson() *string {
	return s.EnglishContactPerson
}

func (s *QueryDomainBySaleIdResponseBody) GetEnglishHolder() *string {
	return s.EnglishHolder
}

func (s *QueryDomainBySaleIdResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainBySaleIdResponseBody) GetHolderEmail() *string {
	return s.HolderEmail
}

func (s *QueryDomainBySaleIdResponseBody) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainBySaleIdResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *QueryDomainBySaleIdResponseBody) GetSafetyLock() *string {
	return s.SafetyLock
}

func (s *QueryDomainBySaleIdResponseBody) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainBySaleIdResponseBody) GetTransferLock() *string {
	return s.TransferLock
}

func (s *QueryDomainBySaleIdResponseBody) GetTransferOutStatus() *string {
	return s.TransferOutStatus
}

func (s *QueryDomainBySaleIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryDomainBySaleIdResponseBody) GetWhoisProtected() *bool {
	return s.WhoisProtected
}

func (s *QueryDomainBySaleIdResponseBody) SetChineseContactPerson(v string) *QueryDomainBySaleIdResponseBody {
	s.ChineseContactPerson = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetChineseHolder(v string) *QueryDomainBySaleIdResponseBody {
	s.ChineseHolder = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetCreationDate(v string) *QueryDomainBySaleIdResponseBody {
	s.CreationDate = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetDnsList(v *QueryDomainBySaleIdResponseBodyDnsList) *QueryDomainBySaleIdResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetDomainName(v string) *QueryDomainBySaleIdResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetDomainRegType(v string) *QueryDomainBySaleIdResponseBody {
	s.DomainRegType = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetEmailVerificationClientHold(v bool) *QueryDomainBySaleIdResponseBody {
	s.EmailVerificationClientHold = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetEmailVerificationStatus(v int32) *QueryDomainBySaleIdResponseBody {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetEnglishContactPerson(v string) *QueryDomainBySaleIdResponseBody {
	s.EnglishContactPerson = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetEnglishHolder(v string) *QueryDomainBySaleIdResponseBody {
	s.EnglishHolder = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetExpirationDate(v string) *QueryDomainBySaleIdResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetHolderEmail(v string) *QueryDomainBySaleIdResponseBody {
	s.HolderEmail = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetPremium(v bool) *QueryDomainBySaleIdResponseBody {
	s.Premium = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetRemark(v string) *QueryDomainBySaleIdResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetSafetyLock(v string) *QueryDomainBySaleIdResponseBody {
	s.SafetyLock = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetSaleId(v string) *QueryDomainBySaleIdResponseBody {
	s.SaleId = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetTransferLock(v string) *QueryDomainBySaleIdResponseBody {
	s.TransferLock = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetTransferOutStatus(v string) *QueryDomainBySaleIdResponseBody {
	s.TransferOutStatus = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetUserId(v string) *QueryDomainBySaleIdResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) SetWhoisProtected(v bool) *QueryDomainBySaleIdResponseBody {
	s.WhoisProtected = &v
	return s
}

func (s *QueryDomainBySaleIdResponseBody) Validate() error {
	if s.DnsList != nil {
		if err := s.DnsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainBySaleIdResponseBodyDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryDomainBySaleIdResponseBodyDnsList) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainBySaleIdResponseBodyDnsList) GoString() string {
	return s.String()
}

func (s *QueryDomainBySaleIdResponseBodyDnsList) GetDns() []*string {
	return s.Dns
}

func (s *QueryDomainBySaleIdResponseBodyDnsList) SetDns(v []*string) *QueryDomainBySaleIdResponseBodyDnsList {
	s.Dns = v
	return s
}

func (s *QueryDomainBySaleIdResponseBodyDnsList) Validate() error {
	return dara.Validate(s)
}
