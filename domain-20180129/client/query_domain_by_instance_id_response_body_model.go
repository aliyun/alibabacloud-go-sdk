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
	SetDomainGroupId(v int64) *QueryDomainByInstanceIdResponseBody
	GetDomainGroupId() *int64
	SetDomainGroupName(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainGroupName() *string
	SetDomainLifecycleStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainLifecycleStatus() *string
	SetDomainName(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainName() *string
	SetDomainNameProxyService(v bool) *QueryDomainByInstanceIdResponseBody
	GetDomainNameProxyService() *bool
	SetDomainNameVerificationStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainNameVerificationStatus() *string
	SetDomainStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainStatus() *string
	SetDomainType(v string) *QueryDomainByInstanceIdResponseBody
	GetDomainType() *string
	SetEmail(v string) *QueryDomainByInstanceIdResponseBody
	GetEmail() *string
	SetEmailVerificationClientHold(v bool) *QueryDomainByInstanceIdResponseBody
	GetEmailVerificationClientHold() *bool
	SetEmailVerificationStatus(v int32) *QueryDomainByInstanceIdResponseBody
	GetEmailVerificationStatus() *int32
	SetExpirationCurrDateDiff(v int32) *QueryDomainByInstanceIdResponseBody
	GetExpirationCurrDateDiff() *int32
	SetExpirationDate(v string) *QueryDomainByInstanceIdResponseBody
	GetExpirationDate() *string
	SetExpirationDateLong(v int64) *QueryDomainByInstanceIdResponseBody
	GetExpirationDateLong() *int64
	SetExpirationDateStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetExpirationDateStatus() *string
	SetInstanceId(v string) *QueryDomainByInstanceIdResponseBody
	GetInstanceId() *string
	SetPremium(v bool) *QueryDomainByInstanceIdResponseBody
	GetPremium() *bool
	SetPrivacyServiceStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetPrivacyServiceStatus() *string
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
	SetRemark(v string) *QueryDomainByInstanceIdResponseBody
	GetRemark() *string
	SetRequestId(v string) *QueryDomainByInstanceIdResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *QueryDomainByInstanceIdResponseBody
	GetResourceGroupId() *string
	SetTag(v *QueryDomainByInstanceIdResponseBodyTag) *QueryDomainByInstanceIdResponseBody
	GetTag() *QueryDomainByInstanceIdResponseBodyTag
	SetTransferOutStatus(v string) *QueryDomainByInstanceIdResponseBody
	GetTransferOutStatus() *string
	SetTransferProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody
	GetTransferProhibitionLock() *string
	SetUpdateProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody
	GetUpdateProhibitionLock() *string
	SetUserId(v string) *QueryDomainByInstanceIdResponseBody
	GetUserId() *string
	SetZhRegistrantName(v string) *QueryDomainByInstanceIdResponseBody
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody
	GetZhRegistrantOrganization() *string
}

type QueryDomainByInstanceIdResponseBody struct {
	DnsList *QueryDomainByInstanceIdResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	DomainGroupId         *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName       *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainLifecycleStatus *string `json:"DomainLifecycleStatus,omitempty" xml:"DomainLifecycleStatus,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// false
	DomainNameProxyService *bool `json:"DomainNameProxyService,omitempty" xml:"DomainNameProxyService,omitempty"`
	// example:
	//
	// NONAUDIT
	DomainNameVerificationStatus *string `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	// example:
	//
	// 1
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// example:
	//
	// gTLD
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// false
	EmailVerificationClientHold *bool `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	// example:
	//
	// 1
	EmailVerificationStatus *int32 `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	// example:
	//
	// 356
	ExpirationCurrDateDiff *int32 `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	// example:
	//
	// 2019-12-07 17:02:13
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// example:
	//
	// 1625111915000
	ExpirationDateLong *int64 `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	// example:
	//
	// 1
	ExpirationDateStatus *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	// example:
	//
	// S20179H1BBI9test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	Premium              *bool   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	PrivacyServiceStatus *string `json:"PrivacyServiceStatus,omitempty" xml:"PrivacyServiceStatus,omitempty"`
	// example:
	//
	// NONAUDIT
	RealNameStatus *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	// example:
	//
	// Test litm
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// Test litm
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// example:
	//
	// NORMAL
	RegistrantUpdatingStatus *string `json:"RegistrantUpdatingStatus,omitempty" xml:"RegistrantUpdatingStatus,omitempty"`
	// example:
	//
	// 2017-12-07 17:02:13
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// example:
	//
	// 1625111915000
	RegistrationDateLong *int64  `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 23C9B3C4-9E2C-4405-A88D-BD33E459D140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmw6bpc6n7zai
	ResourceGroupId *string                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             *QueryDomainByInstanceIdResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// example:
	//
	// NORMAL
	TransferOutStatus *string `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	// example:
	//
	// CLOSE
	TransferProhibitionLock *string `json:"TransferProhibitionLock,omitempty" xml:"TransferProhibitionLock,omitempty"`
	// example:
	//
	// CLOSE
	UpdateProhibitionLock *string `json:"UpdateProhibitionLock,omitempty" xml:"UpdateProhibitionLock,omitempty"`
	// example:
	//
	// 121000000****
	UserId                   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
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

func (s *QueryDomainByInstanceIdResponseBody) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainLifecycleStatus() *string {
	return s.DomainLifecycleStatus
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

func (s *QueryDomainByInstanceIdResponseBody) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetDomainType() *string {
	return s.DomainType
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

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationCurrDateDiff() *int32 {
	return s.ExpirationCurrDateDiff
}

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryDomainByInstanceIdResponseBody) GetExpirationDateStatus() *string {
	return s.ExpirationDateStatus
}

func (s *QueryDomainByInstanceIdResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainByInstanceIdResponseBody) GetPremium() *bool {
	return s.Premium
}

func (s *QueryDomainByInstanceIdResponseBody) GetPrivacyServiceStatus() *string {
	return s.PrivacyServiceStatus
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

func (s *QueryDomainByInstanceIdResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *QueryDomainByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainByInstanceIdResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryDomainByInstanceIdResponseBody) GetTag() *QueryDomainByInstanceIdResponseBodyTag {
	return s.Tag
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

func (s *QueryDomainByInstanceIdResponseBody) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *QueryDomainByInstanceIdResponseBody) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryDomainByInstanceIdResponseBody) SetDnsList(v *QueryDomainByInstanceIdResponseBodyDnsList) *QueryDomainByInstanceIdResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainGroupId(v int64) *QueryDomainByInstanceIdResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainGroupName(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainLifecycleStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainLifecycleStatus = &v
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

func (s *QueryDomainByInstanceIdResponseBody) SetDomainStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainType(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainType = &v
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

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationCurrDateDiff(v int32) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationCurrDateDiff = &v
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

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDateStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDateStatus = &v
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

func (s *QueryDomainByInstanceIdResponseBody) SetPrivacyServiceStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.PrivacyServiceStatus = &v
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

func (s *QueryDomainByInstanceIdResponseBody) SetRemark(v string) *QueryDomainByInstanceIdResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRequestId(v string) *QueryDomainByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetResourceGroupId(v string) *QueryDomainByInstanceIdResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetTag(v *QueryDomainByInstanceIdResponseBodyTag) *QueryDomainByInstanceIdResponseBody {
	s.Tag = v
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

func (s *QueryDomainByInstanceIdResponseBody) SetZhRegistrantName(v string) *QueryDomainByInstanceIdResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetZhRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) Validate() error {
	if s.DnsList != nil {
		if err := s.DnsList.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
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

type QueryDomainByInstanceIdResponseBodyTag struct {
	Tag []*QueryDomainByInstanceIdResponseBodyTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryDomainByInstanceIdResponseBodyTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBodyTag) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBodyTag) GetTag() []*QueryDomainByInstanceIdResponseBodyTagTag {
	return s.Tag
}

func (s *QueryDomainByInstanceIdResponseBodyTag) SetTag(v []*QueryDomainByInstanceIdResponseBodyTagTag) *QueryDomainByInstanceIdResponseBodyTag {
	s.Tag = v
	return s
}

func (s *QueryDomainByInstanceIdResponseBodyTag) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainByInstanceIdResponseBodyTagTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDomainByInstanceIdResponseBodyTagTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBodyTagTag) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBodyTagTag) GetKey() *string {
	return s.Key
}

func (s *QueryDomainByInstanceIdResponseBodyTagTag) GetValue() *string {
	return s.Value
}

func (s *QueryDomainByInstanceIdResponseBodyTagTag) SetKey(v string) *QueryDomainByInstanceIdResponseBodyTagTag {
	s.Key = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBodyTagTag) SetValue(v string) *QueryDomainByInstanceIdResponseBodyTagTag {
	s.Value = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBodyTagTag) Validate() error {
	return dara.Validate(s)
}
