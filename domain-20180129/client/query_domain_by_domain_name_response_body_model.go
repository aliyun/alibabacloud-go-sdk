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
	SetDomainGroupId(v int64) *QueryDomainByDomainNameResponseBody
	GetDomainGroupId() *int64
	SetDomainGroupName(v string) *QueryDomainByDomainNameResponseBody
	GetDomainGroupName() *string
	SetDomainName(v string) *QueryDomainByDomainNameResponseBody
	GetDomainName() *string
	SetDomainNameProxyService(v bool) *QueryDomainByDomainNameResponseBody
	GetDomainNameProxyService() *bool
	SetDomainNameVerificationStatus(v string) *QueryDomainByDomainNameResponseBody
	GetDomainNameVerificationStatus() *string
	SetDomainStatus(v string) *QueryDomainByDomainNameResponseBody
	GetDomainStatus() *string
	SetDomainType(v string) *QueryDomainByDomainNameResponseBody
	GetDomainType() *string
	SetEmail(v string) *QueryDomainByDomainNameResponseBody
	GetEmail() *string
	SetEmailVerificationClientHold(v bool) *QueryDomainByDomainNameResponseBody
	GetEmailVerificationClientHold() *bool
	SetEmailVerificationStatus(v int32) *QueryDomainByDomainNameResponseBody
	GetEmailVerificationStatus() *int32
	SetExpirationCurrDateDiff(v int32) *QueryDomainByDomainNameResponseBody
	GetExpirationCurrDateDiff() *int32
	SetExpirationDate(v string) *QueryDomainByDomainNameResponseBody
	GetExpirationDate() *string
	SetExpirationDateLong(v int64) *QueryDomainByDomainNameResponseBody
	GetExpirationDateLong() *int64
	SetExpirationDateStatus(v string) *QueryDomainByDomainNameResponseBody
	GetExpirationDateStatus() *string
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
	SetRemark(v string) *QueryDomainByDomainNameResponseBody
	GetRemark() *string
	SetRequestId(v string) *QueryDomainByDomainNameResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *QueryDomainByDomainNameResponseBody
	GetResourceGroupId() *string
	SetTag(v *QueryDomainByDomainNameResponseBodyTag) *QueryDomainByDomainNameResponseBody
	GetTag() *QueryDomainByDomainNameResponseBodyTag
	SetTransferOutStatus(v string) *QueryDomainByDomainNameResponseBody
	GetTransferOutStatus() *string
	SetTransferProhibitionLock(v string) *QueryDomainByDomainNameResponseBody
	GetTransferProhibitionLock() *string
	SetUpdateProhibitionLock(v string) *QueryDomainByDomainNameResponseBody
	GetUpdateProhibitionLock() *string
	SetUserId(v string) *QueryDomainByDomainNameResponseBody
	GetUserId() *string
	SetZhRegistrantName(v string) *QueryDomainByDomainNameResponseBody
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody
	GetZhRegistrantOrganization() *string
}

type QueryDomainByDomainNameResponseBody struct {
	// The Domain Name System (DNS) servers of the domain name.
	DnsList *QueryDomainByDomainNameResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	// The ID of the domain name group. You can call the [QueryDomainGroupList](https://help.aliyun.com/document_detail/69362.html) operation to query the ID of the domain name group.
	//
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// The name of the domain name group.
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether privacy protection is enabled for the domain name.
	//
	// example:
	//
	// false
	DomainNameProxyService *bool `json:"DomainNameProxyService,omitempty" xml:"DomainNameProxyService,omitempty"`
	// The status of name auditing for the domain name. Valid values:
	//
	// 	- **NONAUDIT**: The name auditing for the domain name is not performed.
	//
	// 	- **SUCCEED**: The name auditing for the domain name is successful.
	//
	// 	- **FAILED**: The name auditing for the domain name fails.
	//
	// 	- **AUDITING**: The name auditing for the domain name is in progress.
	//
	// example:
	//
	// SUCCEED
	DomainNameVerificationStatus *string `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- 1: The domain name needs to be renewed.
	//
	// 	- 2: The domain name needs to be redeemed.
	//
	// 	- 3: The domain name is normal.
	//
	// example:
	//
	// 3
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- New gTLD
	//
	// 	- gTLD
	//
	// 	- ccTLD
	//
	// example:
	//
	// gTLD
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The email address of the domain name registrant.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the domain name is in the ClientHold state.
	//
	// example:
	//
	// false
	EmailVerificationClientHold *bool `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	// Indicates whether the email address passes verification. Valid values:
	//
	// 	- **0**: The email address fails the verification.
	//
	// 	- **1**: The email address passes the verification.
	//
	// example:
	//
	// 1
	EmailVerificationStatus *int32 `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	// The number of days from the expiration date of the domain name to the current date.
	//
	// example:
	//
	// 356
	ExpirationCurrDateDiff *int32 `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	// The expiration date.
	//
	// example:
	//
	// 2019-12-07 17:02:13
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The timestamp generated when the domain name expired.
	//
	// example:
	//
	// 1625111915000
	ExpirationDateLong *int64 `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	// Indicates whether the domain name expires. Valid values:
	//
	// 	- **1**: The domain name does not expire.
	//
	// 	- **2**: The domain name expires.
	//
	// example:
	//
	// 1
	ExpirationDateStatus *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	// The instance ID of the domain name.
	//
	// example:
	//
	// S20179H1BBI9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the domain name is a premium domain name.
	//
	// example:
	//
	// false
	Premium *bool `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// The status of real-name verification for the domain name. Valid values:
	//
	// 	- **NONAUDIT**: The real-name verification is not performed.
	//
	// 	- **SUCCEED**: The real-name verification is successful.
	//
	// 	- **FAILED**: The real-name verification fails.
	//
	// 	- **AUDITING**: The real-name verification is in progress.
	//
	// example:
	//
	// NONAUDIT
	RealNameStatus *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// Test litm
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// The registrant of the domain name.
	//
	// example:
	//
	// Test litm
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// The type of contact who registers the domain name. Valid values:
	//
	// 	- **1**: individual.
	//
	// 	- **2**: enterprise.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// The status of the information about the domain name registrant. Valid values:
	//
	// 	- **PENDING**: The information about the domain name registrant is being modified.
	//
	// 	- **NORMAL**: normal.
	//
	// example:
	//
	// NORMAL
	RegistrantUpdatingStatus *string `json:"RegistrantUpdatingStatus,omitempty" xml:"RegistrantUpdatingStatus,omitempty"`
	// The time when the domain name was registered.
	//
	// example:
	//
	// 2017-12-07 17:02:13
	RegistrationDate *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	// The timestamp generated when the domain name was registered.
	//
	// example:
	//
	// 1584675448000
	RegistrationDateLong *int64 `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	// The remarks on the domain name.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 44101664-3E70-4F0E-89E5-CCB74BF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw6bpc6n7zai
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag *QueryDomainByDomainNameResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	// The transfer status of the domain name. Valid values:
	//
	// 	- **NORMAL**: The domain name is normal.
	//
	// 	- **PENDING**: The domain name is being transferred out from Alibaba Cloud.
	//
	// example:
	//
	// NORMAL
	TransferOutStatus *string `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	// The status of the transfer lock for the domain name. Valid values:
	//
	// 	- **NONE_SETTING**: No transfer lock is configured.
	//
	// 	- **OPEN**: The transfer lock is enabled.
	//
	// 	- **CLOSE**: The transfer lock is disabled.
	//
	// example:
	//
	// CLOSE
	TransferProhibitionLock *string `json:"TransferProhibitionLock,omitempty" xml:"TransferProhibitionLock,omitempty"`
	// The status of the security lock for the domain name. Valid values:
	//
	// 	- **NONE_SETTING**: No security lock is configured.
	//
	// 	- **OPEN**: The security lock is enabled.
	//
	// 	- **CLOSE**: The security lock is disabled.
	//
	// example:
	//
	// CLOSE
	UpdateProhibitionLock *string `json:"UpdateProhibitionLock,omitempty" xml:"UpdateProhibitionLock,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 121000000****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The Chinese name of the domain name contact.
	ZhRegistrantName *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	// The Chinese name of the domain name registrant.
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
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

func (s *QueryDomainByDomainNameResponseBody) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *QueryDomainByDomainNameResponseBody) GetDomainGroupName() *string {
	return s.DomainGroupName
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

func (s *QueryDomainByDomainNameResponseBody) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *QueryDomainByDomainNameResponseBody) GetDomainType() *string {
	return s.DomainType
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

func (s *QueryDomainByDomainNameResponseBody) GetExpirationCurrDateDiff() *int32 {
	return s.ExpirationCurrDateDiff
}

func (s *QueryDomainByDomainNameResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryDomainByDomainNameResponseBody) GetExpirationDateLong() *int64 {
	return s.ExpirationDateLong
}

func (s *QueryDomainByDomainNameResponseBody) GetExpirationDateStatus() *string {
	return s.ExpirationDateStatus
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

func (s *QueryDomainByDomainNameResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *QueryDomainByDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainByDomainNameResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *QueryDomainByDomainNameResponseBody) GetTag() *QueryDomainByDomainNameResponseBodyTag {
	return s.Tag
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

func (s *QueryDomainByDomainNameResponseBody) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *QueryDomainByDomainNameResponseBody) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryDomainByDomainNameResponseBody) SetDnsList(v *QueryDomainByDomainNameResponseBodyDnsList) *QueryDomainByDomainNameResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainGroupId(v int64) *QueryDomainByDomainNameResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainGroupName(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainGroupName = &v
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

func (s *QueryDomainByDomainNameResponseBody) SetDomainStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainType(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainType = &v
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

func (s *QueryDomainByDomainNameResponseBody) SetExpirationCurrDateDiff(v int32) *QueryDomainByDomainNameResponseBody {
	s.ExpirationCurrDateDiff = &v
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

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDateStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDateStatus = &v
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

func (s *QueryDomainByDomainNameResponseBody) SetRemark(v string) *QueryDomainByDomainNameResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRequestId(v string) *QueryDomainByDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetResourceGroupId(v string) *QueryDomainByDomainNameResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetTag(v *QueryDomainByDomainNameResponseBodyTag) *QueryDomainByDomainNameResponseBody {
	s.Tag = v
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

func (s *QueryDomainByDomainNameResponseBody) SetZhRegistrantName(v string) *QueryDomainByDomainNameResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetZhRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) Validate() error {
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

type QueryDomainByDomainNameResponseBodyTag struct {
	Tag []*QueryDomainByDomainNameResponseBodyTagTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryDomainByDomainNameResponseBodyTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBodyTag) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBodyTag) GetTag() []*QueryDomainByDomainNameResponseBodyTagTag {
	return s.Tag
}

func (s *QueryDomainByDomainNameResponseBodyTag) SetTag(v []*QueryDomainByDomainNameResponseBodyTagTag) *QueryDomainByDomainNameResponseBodyTag {
	s.Tag = v
	return s
}

func (s *QueryDomainByDomainNameResponseBodyTag) Validate() error {
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

type QueryDomainByDomainNameResponseBodyTagTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Vaue *string `json:"Vaue,omitempty" xml:"Vaue,omitempty"`
}

func (s QueryDomainByDomainNameResponseBodyTagTag) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBodyTagTag) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBodyTagTag) GetKey() *string {
	return s.Key
}

func (s *QueryDomainByDomainNameResponseBodyTagTag) GetVaue() *string {
	return s.Vaue
}

func (s *QueryDomainByDomainNameResponseBodyTagTag) SetKey(v string) *QueryDomainByDomainNameResponseBodyTagTag {
	s.Key = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBodyTagTag) SetVaue(v string) *QueryDomainByDomainNameResponseBodyTagTag {
	s.Vaue = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBodyTagTag) Validate() error {
	return dara.Validate(s)
}
