// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsList(v *QueryDomainByInstanceIdResponseBodyDnsList) *QueryDomainByInstanceIdResponseBody
	GetDnsList() *QueryDomainByInstanceIdResponseBodyDnsList
	SetDomainName(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainName() *string
	SetDomainNameProxyService(v bool) *QueryDomainByInstanceIdResponseBody
	GetDomainNameProxyService() *bool
	SetDomainNameVerificationStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainNameVerificationStatus() *string
	SetEmail(v string) *QueryDomainByInstanceIdResponseBody
	GetEmail() *string
	SetEmailVerificationClientHold(v bool) *QueryDomainByInstanceIdResponseBody
	GetEmailVerificationClientHold() *bool
	SetEmailVerificationStatus(v int32) *QueryDomainByInstanceIdResponseBody
	GetEmailVerificationStatus() *int32
	SetExpirationDate(v string) *QueryDomainByInstanceIdResponseBody
	GetExpirationDate() *string
	SetExpirationDateLong(v int64) *QueryDomainByInstanceIdResponseBody
	GetExpirationDateLong() *int64
	SetInstanceId(v string) *QueryDomainByInstanceIdResponseBody
	GetInstanceId() *string
	SetPremium(v bool) *QueryDomainByInstanceIdResponseBody
	GetPremium() *bool
	SetRealNameStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetRealNameStatus() *string
	SetRegistrantName(v string) *QueryDomainByInstanceIdResponseBody
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody
	GetRegistrantOrganization() *string
	SetRegistrantType(v string) *QueryDomainByInstanceIdResponseBody
	GetRegistrantType() *string
	SetRegistrantUpdatingStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetRegistrantUpdatingStatus() *string
	SetRegistrationDate(v string) *QueryDomainByInstanceIdResponseBody
	GetRegistrationDate() *string
	SetRegistrationDateLong(v int64) *QueryDomainByInstanceIdResponseBody
	GetRegistrationDateLong() *int64
	SetRequestId(v string) *QueryDomainByInstanceIdResponseBody
	GetRequestId() *string
	SetTransferOutStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetTransferOutStatus() *string
	SetTransferProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody
	GetTransferProhibitionLock() *string
	SetUpdateProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody
	GetUpdateProhibitionLock() *string
	SetUserId(v string) *QueryDomainByInstanceIdResponseBody
	GetUserId() *string
}

type QueryDomainByInstanceIdResponseBody struct {
	DnsList                      *QueryDomainByInstanceIdResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
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

func (s QueryDomainByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBody) GetDnsList() *QueryDomainByInstanceIdResponseBodyDnsList {
	return s.DnsList
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainNameProxyService() *bool {
	return s.DomainNameProxyService
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainNameVerificationStatus() *string {
	return s.DomainNameVerificationStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryDomainByInstanceIdResponseBody) GetEmailVerificationClientHold() *bool {
	return s.EmailVerificationClientHold
}

func (s *QueryDomainByInstanceIdResponseBody) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryDomainByInstanceIdResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainByInstanceIdResponseBody) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainByInstanceIdResponseBody) GetRealNameStatus() *string {
	return s.RealNameStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrantUpdatingStatus() *string {
	return s.RegistrantUpdatingStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrationDate() *string {
	return s.RegistrationDate
}

func (s *QueryDomainByInstanceIdResponseBody) GetRegistrationDateLong() *int64 {
	return s.RegistrationDateLong
}

func (s *QueryDomainByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainByInstanceIdResponseBody) GetTransferOutStatus() *string {
	return s.TransferOutStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetTransferProhibitionLock() *string {
	return s.TransferProhibitionLock
}

func (s *QueryDomainByInstanceIdResponseBody) GetUpdateProhibitionLock() *string {
	return s.UpdateProhibitionLock
}

func (s *QueryDomainByInstanceIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryDomainByInstanceIdResponseBody) SetDnsList(v *QueryDomainByInstanceIdResponseBodyDnsList) *QueryDomainByInstanceIdResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainName(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainNameProxyService(v bool) *QueryDomainByInstanceIdResponseBody {
	s.DomainNameProxyService = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainNameVerificationStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainNameVerificationStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmail(v string) *QueryDomainByInstanceIdResponseBody {
	s.Email = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmailVerificationClientHold(v bool) *QueryDomainByInstanceIdResponseBody {
	s.EmailVerificationClientHold = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmailVerificationStatus(v int32) *QueryDomainByInstanceIdResponseBody {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDate(v string) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDateLong(v int64) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetInstanceId(v string) *QueryDomainByInstanceIdResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetPremium(v bool) *QueryDomainByInstanceIdResponseBody {
	s.Premium = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRealNameStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.RealNameStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantName(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantType(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantUpdatingStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantUpdatingStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrationDate(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrationDateLong(v int64) *QueryDomainByInstanceIdResponseBody {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRequestId(v string) *QueryDomainByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetTransferOutStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.TransferOutStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetTransferProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody {
	s.TransferProhibitionLock = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetUpdateProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody {
	s.UpdateProhibitionLock = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetUserId(v string) *QueryDomainByInstanceIdResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) Validate() error {
	if s.DnsList != nil {
		if err := s.DnsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainByInstanceIdResponseBodyDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryDomainByInstanceIdResponseBodyDnsList) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBodyDnsList) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBodyDnsList) GetDns() []*string {
	return s.Dns
}

func (s *QueryDomainByInstanceIdResponseBodyDnsList) SetDns(v []*string) *QueryDomainByInstanceIdResponseBodyDnsList {
	s.Dns = v
	return s
}

func (s *QueryDomainByInstanceIdResponseBodyDnsList) Validate() error {
	return dara.Validate(s)
}
