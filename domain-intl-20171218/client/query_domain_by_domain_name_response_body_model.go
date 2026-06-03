// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsList(v *QueryDomainByDomainNameResponseBodyDnsList) *QueryDomainByDomainNameResponseBody
	GetDnsList() *QueryDomainByDomainNameResponseBodyDnsList
	SetDomainName(v string) *QueryDomainByDomainNameResponseBody
	GetDomainName() *string
	SetDomainNameProxyService(v bool) *QueryDomainByDomainNameResponseBody
	GetDomainNameProxyService() *bool
	SetDomainNameVerificationStatus(v string) *QueryDomainByDomainNameResponseBody
	GetDomainNameVerificationStatus() *string
	SetEmail(v string) *QueryDomainByDomainNameResponseBody
	GetEmail() *string
	SetEmailVerificationClientHold(v bool) *QueryDomainByDomainNameResponseBody
	GetEmailVerificationClientHold() *bool
	SetEmailVerificationStatus(v int32) *QueryDomainByDomainNameResponseBody
	GetEmailVerificationStatus() *int32
	SetExpirationDate(v string) *QueryDomainByDomainNameResponseBody
	GetExpirationDate() *string
	SetExpirationDateLong(v int64) *QueryDomainByDomainNameResponseBody
	GetExpirationDateLong() *int64
	SetInstanceId(v string) *QueryDomainByDomainNameResponseBody
	GetInstanceId() *string
	SetPremium(v bool) *QueryDomainByDomainNameResponseBody
	GetPremium() *bool
	SetRealNameStatus(v string) *QueryDomainByDomainNameResponseBody
	GetRealNameStatus() *string
	SetRegistrantName(v string) *QueryDomainByDomainNameResponseBody
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody
	GetRegistrantOrganization() *string
	SetRegistrantType(v string) *QueryDomainByDomainNameResponseBody
	GetRegistrantType() *string
	SetRegistrantUpdatingStatus(v string) *QueryDomainByDomainNameResponseBody
	GetRegistrantUpdatingStatus() *string
	SetRegistrationDate(v string) *QueryDomainByDomainNameResponseBody
	GetRegistrationDate() *string
	SetRegistrationDateLong(v int64) *QueryDomainByDomainNameResponseBody
	GetRegistrationDateLong() *int64
	SetRequestId(v string) *QueryDomainByDomainNameResponseBody
	GetRequestId() *string
	SetTransferOutStatus(v string) *QueryDomainByDomainNameResponseBody
	GetTransferOutStatus() *string
	SetTransferProhibitionLock(v string) *QueryDomainByDomainNameResponseBody
	GetTransferProhibitionLock() *string
	SetUpdateProhibitionLock(v string) *QueryDomainByDomainNameResponseBody
	GetUpdateProhibitionLock() *string
	SetUserId(v string) *QueryDomainByDomainNameResponseBody
	GetUserId() *string
}

type QueryDomainByDomainNameResponseBody struct {
	DnsList                      *QueryDomainByDomainNameResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainName                   *string                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameProxyService       *bool                                       `json:"DomainNameProxyService,omitempty" xml:"DomainNameProxyService,omitempty"`
	DomainNameVerificationStatus *string                                     `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	Email                        *string                                     `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationClientHold  *bool                                       `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	EmailVerificationStatus      *int32                                      `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	ExpirationDate               *string                                     `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong           *int64                                      `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	InstanceId                   *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                      *bool                                       `json:"Premium,omitempty" xml:"Premium,omitempty"`
	RealNameStatus               *string                                     `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantName               *string                                     `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization       *string                                     `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType               *string                                     `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrantUpdatingStatus     *string                                     `json:"RegistrantUpdatingStatus,omitempty" xml:"RegistrantUpdatingStatus,omitempty"`
	RegistrationDate             *string                                     `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong         *int64                                      `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	RequestId                    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransferOutStatus            *string                                     `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	TransferProhibitionLock      *string                                     `json:"TransferProhibitionLock,omitempty" xml:"TransferProhibitionLock,omitempty"`
	UpdateProhibitionLock        *string                                     `json:"UpdateProhibitionLock,omitempty" xml:"UpdateProhibitionLock,omitempty"`
	UserId                       *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDomainByDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBody) GetDnsList() *QueryDomainByDomainNameResponseBodyDnsList {
	return s.DnsList
}

func (s *QueryDomainByDomainNameResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainByDomainNameResponseBody) GetDomainNameProxyService() *bool {
	return s.DomainNameProxyService
}

func (s *QueryDomainByDomainNameResponseBody) GetDomainNameVerificationStatus() *string {
	return s.DomainNameVerificationStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryDomainByDomainNameResponseBody) GetEmailVerificationClientHold() *bool {
	return s.EmailVerificationClientHold
}

func (s *QueryDomainByDomainNameResponseBody) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainByDomainNameResponseBody) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryDomainByDomainNameResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainByDomainNameResponseBody) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainByDomainNameResponseBody) GetRealNameStatus() *string {
	return s.RealNameStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrantUpdatingStatus() *string {
	return s.RegistrantUpdatingStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *QueryDomainByDomainNameResponseBody) GetRegistrationDateLong() *int64 {
	return s.RegistrationDateLong
}

func (s *QueryDomainByDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainByDomainNameResponseBody) GetTransferOutStatus() *string {
	return s.TransferOutStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetTransferProhibitionLock() *string {
	return s.TransferProhibitionLock
}

func (s *QueryDomainByDomainNameResponseBody) GetUpdateProhibitionLock() *string {
	return s.UpdateProhibitionLock
}

func (s *QueryDomainByDomainNameResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryDomainByDomainNameResponseBody) SetDnsList(v *QueryDomainByDomainNameResponseBodyDnsList) *QueryDomainByDomainNameResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainName(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainNameProxyService(v bool) *QueryDomainByDomainNameResponseBody {
	s.DomainNameProxyService = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainNameVerificationStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainNameVerificationStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmail(v string) *QueryDomainByDomainNameResponseBody {
	s.Email = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmailVerificationClientHold(v bool) *QueryDomainByDomainNameResponseBody {
	s.EmailVerificationClientHold = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmailVerificationStatus(v int32) *QueryDomainByDomainNameResponseBody {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDate(v string) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDateLong(v int64) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetInstanceId(v string) *QueryDomainByDomainNameResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetPremium(v bool) *QueryDomainByDomainNameResponseBody {
	s.Premium = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRealNameStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.RealNameStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantName(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantType(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantUpdatingStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantUpdatingStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrationDate(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrationDateLong(v int64) *QueryDomainByDomainNameResponseBody {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRequestId(v string) *QueryDomainByDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetTransferOutStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.TransferOutStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetTransferProhibitionLock(v string) *QueryDomainByDomainNameResponseBody {
	s.TransferProhibitionLock = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetUpdateProhibitionLock(v string) *QueryDomainByDomainNameResponseBody {
	s.UpdateProhibitionLock = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetUserId(v string) *QueryDomainByDomainNameResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) Validate() error {
	if s.DnsList != nil {
		if err := s.DnsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainByDomainNameResponseBodyDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryDomainByDomainNameResponseBodyDnsList) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBodyDnsList) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBodyDnsList) GetDns() []*string {
	return s.Dns
}

func (s *QueryDomainByDomainNameResponseBodyDnsList) SetDns(v []*string) *QueryDomainByDomainNameResponseBodyDnsList {
	s.Dns = v
	return s
}

func (s *QueryDomainByDomainNameResponseBodyDnsList) Validate() error {
	return dara.Validate(s)
}
