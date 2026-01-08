// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecord(v *UpdateAuditRequestRequestAuditRecord) *UpdateAuditRequestRequest
	GetAuditRecord() *UpdateAuditRequestRequestAuditRecord
	SetAuditResult(v string) *UpdateAuditRequestRequest
	GetAuditResult() *string
	SetCustSpaceId(v string) *UpdateAuditRequestRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *UpdateAuditRequestRequest
	GetOwnerId() *int64
	SetRequestNo(v string) *UpdateAuditRequestRequest
	GetRequestNo() *string
	SetResourceOwnerAccount(v string) *UpdateAuditRequestRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAuditRequestRequest
	GetResourceOwnerId() *int64
}

type UpdateAuditRequestRequest struct {
	// This parameter is required.
	AuditRecord *UpdateAuditRequestRequestAuditRecord `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty" type:"Struct"`
	// example:
	//
	// unAudit
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-8pi**urn5s
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 114624518**5956096
	RequestNo            *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateAuditRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestRequest) GetAuditRecord() *UpdateAuditRequestRequestAuditRecord {
	return s.AuditRecord
}

func (s *UpdateAuditRequestRequest) GetAuditResult() *string {
	return s.AuditResult
}

func (s *UpdateAuditRequestRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateAuditRequestRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAuditRequestRequest) GetRequestNo() *string {
	return s.RequestNo
}

func (s *UpdateAuditRequestRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAuditRequestRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAuditRequestRequest) SetAuditRecord(v *UpdateAuditRequestRequestAuditRecord) *UpdateAuditRequestRequest {
	s.AuditRecord = v
	return s
}

func (s *UpdateAuditRequestRequest) SetAuditResult(v string) *UpdateAuditRequestRequest {
	s.AuditResult = &v
	return s
}

func (s *UpdateAuditRequestRequest) SetCustSpaceId(v string) *UpdateAuditRequestRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateAuditRequestRequest) SetOwnerId(v int64) *UpdateAuditRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAuditRequestRequest) SetRequestNo(v string) *UpdateAuditRequestRequest {
	s.RequestNo = &v
	return s
}

func (s *UpdateAuditRequestRequest) SetResourceOwnerAccount(v string) *UpdateAuditRequestRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAuditRequestRequest) SetResourceOwnerId(v int64) *UpdateAuditRequestRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAuditRequestRequest) Validate() error {
	if s.AuditRecord != nil {
		if err := s.AuditRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAuditRequestRequestAuditRecord struct {
	// example:
	//
	// unlimited
	AgeLimit *string `json:"AgeLimit,omitempty" xml:"AgeLimit,omitempty"`
	// example:
	//
	// reason
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// example:
	//
	// testaccount
	BusinessAccountName *string `json:"BusinessAccountName,omitempty" xml:"BusinessAccountName,omitempty"`
	// example:
	//
	// 9211515**345
	BusinessLicenseRegistrationNumber *string                                               `json:"BusinessLicenseRegistrationNumber,omitempty" xml:"BusinessLicenseRegistrationNumber,omitempty"`
	CompanyAddress                    []*UpdateAuditRequestRequestAuditRecordCompanyAddress `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty" type:"Repeated"`
	// example:
	//
	// baidu
	CompanyEnglishName *string `json:"CompanyEnglishName,omitempty" xml:"CompanyEnglishName,omitempty"`
	// example:
	//
	// 示例值示例值
	CompanyLegalName *string `json:"CompanyLegalName,omitempty" xml:"CompanyLegalName,omitempty"`
	// example:
	//
	// 示例值示例值
	CompanyLegalPerson *string `json:"CompanyLegalPerson,omitempty" xml:"CompanyLegalPerson,omitempty"`
	// example:
	//
	// xxx
	CompanyProfile *string `json:"CompanyProfile,omitempty" xml:"CompanyProfile,omitempty"`
	// example:
	//
	// 示例值示例值
	CompanyRegisteredCountry *string                                           `json:"CompanyRegisteredCountry,omitempty" xml:"CompanyRegisteredCountry,omitempty"`
	CompanyTel               []*UpdateAuditRequestRequestAuditRecordCompanyTel `json:"CompanyTel,omitempty" xml:"CompanyTel,omitempty" type:"Repeated"`
	// example:
	//
	// xxxstreet
	CompleteAddressOfHeadquarters *string `json:"CompleteAddressOfHeadquarters,omitempty" xml:"CompleteAddressOfHeadquarters,omitempty"`
	// example:
	//
	// xxxx@gamil.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// N/A‌
	EnableAutoReply *string `json:"EnableAutoReply,omitempty" xml:"EnableAutoReply,omitempty"`
	// example:
	//
	// description
	IndustryDescription *string `json:"IndustryDescription,omitempty" xml:"IndustryDescription,omitempty"`
	// example:
	//
	// telecommunicationsAndIT
	IndustryInvolved *string `json:"IndustryInvolved,omitempty" xml:"IndustryInvolved,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	LetterGuarantee *string `json:"LetterGuarantee,omitempty" xml:"LetterGuarantee,omitempty"`
	// logo
	Logo                                   []*string `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
	MessageDestinationCountry              []*string `json:"MessageDestinationCountry,omitempty" xml:"MessageDestinationCountry,omitempty" type:"Repeated"`
	MessageDestinationInternationalCountry []*string `json:"MessageDestinationInternationalCountry,omitempty" xml:"MessageDestinationInternationalCountry,omitempty" type:"Repeated"`
	// example:
	//
	// hello!
	MessageDialogueIntroduction *string `json:"MessageDialogueIntroduction,omitempty" xml:"MessageDialogueIntroduction,omitempty"`
	// example:
	//
	// 2025-11-07
	MessageEnableDate *string `json:"MessageEnableDate,omitempty" xml:"MessageEnableDate,omitempty"`
	// example:
	//
	// test
	MessageSessionName *string `json:"MessageSessionName,omitempty" xml:"MessageSessionName,omitempty"`
	// example:
	//
	// https://xxx.xxx.com
	OtherLetterGuarantee *string `json:"OtherLetterGuarantee,omitempty" xml:"OtherLetterGuarantee,omitempty"`
	// example:
	//
	// 2025-11-07
	RecoveryDate *string `json:"RecoveryDate,omitempty" xml:"RecoveryDate,omitempty"`
	// example:
	//
	// N/A
	ReplyContent *string `json:"ReplyContent,omitempty" xml:"ReplyContent,omitempty"`
	// example:
	//
	// N/A
	SuspensionDate *string `json:"SuspensionDate,omitempty" xml:"SuspensionDate,omitempty"`
	// example:
	//
	// https://www.xxx.xxx.com
	WebAddress *string `json:"WebAddress,omitempty" xml:"WebAddress,omitempty"`
}

func (s UpdateAuditRequestRequestAuditRecord) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestRequestAuditRecord) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestRequestAuditRecord) GetAgeLimit() *string {
	return s.AgeLimit
}

func (s *UpdateAuditRequestRequestAuditRecord) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *UpdateAuditRequestRequestAuditRecord) GetBusinessAccountName() *string {
	return s.BusinessAccountName
}

func (s *UpdateAuditRequestRequestAuditRecord) GetBusinessLicenseRegistrationNumber() *string {
	return s.BusinessLicenseRegistrationNumber
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyAddress() []*UpdateAuditRequestRequestAuditRecordCompanyAddress {
	return s.CompanyAddress
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyEnglishName() *string {
	return s.CompanyEnglishName
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyLegalName() *string {
	return s.CompanyLegalName
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyLegalPerson() *string {
	return s.CompanyLegalPerson
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyProfile() *string {
	return s.CompanyProfile
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyRegisteredCountry() *string {
	return s.CompanyRegisteredCountry
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompanyTel() []*UpdateAuditRequestRequestAuditRecordCompanyTel {
	return s.CompanyTel
}

func (s *UpdateAuditRequestRequestAuditRecord) GetCompleteAddressOfHeadquarters() *string {
	return s.CompleteAddressOfHeadquarters
}

func (s *UpdateAuditRequestRequestAuditRecord) GetContactMail() *string {
	return s.ContactMail
}

func (s *UpdateAuditRequestRequestAuditRecord) GetEnableAutoReply() *string {
	return s.EnableAutoReply
}

func (s *UpdateAuditRequestRequestAuditRecord) GetIndustryDescription() *string {
	return s.IndustryDescription
}

func (s *UpdateAuditRequestRequestAuditRecord) GetIndustryInvolved() *string {
	return s.IndustryInvolved
}

func (s *UpdateAuditRequestRequestAuditRecord) GetLetterGuarantee() *string {
	return s.LetterGuarantee
}

func (s *UpdateAuditRequestRequestAuditRecord) GetLogo() []*string {
	return s.Logo
}

func (s *UpdateAuditRequestRequestAuditRecord) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *UpdateAuditRequestRequestAuditRecord) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *UpdateAuditRequestRequestAuditRecord) GetMessageDialogueIntroduction() *string {
	return s.MessageDialogueIntroduction
}

func (s *UpdateAuditRequestRequestAuditRecord) GetMessageEnableDate() *string {
	return s.MessageEnableDate
}

func (s *UpdateAuditRequestRequestAuditRecord) GetMessageSessionName() *string {
	return s.MessageSessionName
}

func (s *UpdateAuditRequestRequestAuditRecord) GetOtherLetterGuarantee() *string {
	return s.OtherLetterGuarantee
}

func (s *UpdateAuditRequestRequestAuditRecord) GetRecoveryDate() *string {
	return s.RecoveryDate
}

func (s *UpdateAuditRequestRequestAuditRecord) GetReplyContent() *string {
	return s.ReplyContent
}

func (s *UpdateAuditRequestRequestAuditRecord) GetSuspensionDate() *string {
	return s.SuspensionDate
}

func (s *UpdateAuditRequestRequestAuditRecord) GetWebAddress() *string {
	return s.WebAddress
}

func (s *UpdateAuditRequestRequestAuditRecord) SetAgeLimit(v string) *UpdateAuditRequestRequestAuditRecord {
	s.AgeLimit = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetApplyReason(v string) *UpdateAuditRequestRequestAuditRecord {
	s.ApplyReason = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetBusinessAccountName(v string) *UpdateAuditRequestRequestAuditRecord {
	s.BusinessAccountName = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetBusinessLicenseRegistrationNumber(v string) *UpdateAuditRequestRequestAuditRecord {
	s.BusinessLicenseRegistrationNumber = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyAddress(v []*UpdateAuditRequestRequestAuditRecordCompanyAddress) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyAddress = v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyEnglishName(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyEnglishName = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyLegalName(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyLegalName = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyLegalPerson(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyLegalPerson = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyProfile(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyProfile = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyRegisteredCountry(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyRegisteredCountry = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompanyTel(v []*UpdateAuditRequestRequestAuditRecordCompanyTel) *UpdateAuditRequestRequestAuditRecord {
	s.CompanyTel = v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetCompleteAddressOfHeadquarters(v string) *UpdateAuditRequestRequestAuditRecord {
	s.CompleteAddressOfHeadquarters = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetContactMail(v string) *UpdateAuditRequestRequestAuditRecord {
	s.ContactMail = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetEnableAutoReply(v string) *UpdateAuditRequestRequestAuditRecord {
	s.EnableAutoReply = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetIndustryDescription(v string) *UpdateAuditRequestRequestAuditRecord {
	s.IndustryDescription = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetIndustryInvolved(v string) *UpdateAuditRequestRequestAuditRecord {
	s.IndustryInvolved = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetLetterGuarantee(v string) *UpdateAuditRequestRequestAuditRecord {
	s.LetterGuarantee = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetLogo(v []*string) *UpdateAuditRequestRequestAuditRecord {
	s.Logo = v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetMessageDestinationCountry(v []*string) *UpdateAuditRequestRequestAuditRecord {
	s.MessageDestinationCountry = v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetMessageDestinationInternationalCountry(v []*string) *UpdateAuditRequestRequestAuditRecord {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetMessageDialogueIntroduction(v string) *UpdateAuditRequestRequestAuditRecord {
	s.MessageDialogueIntroduction = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetMessageEnableDate(v string) *UpdateAuditRequestRequestAuditRecord {
	s.MessageEnableDate = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetMessageSessionName(v string) *UpdateAuditRequestRequestAuditRecord {
	s.MessageSessionName = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetOtherLetterGuarantee(v string) *UpdateAuditRequestRequestAuditRecord {
	s.OtherLetterGuarantee = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetRecoveryDate(v string) *UpdateAuditRequestRequestAuditRecord {
	s.RecoveryDate = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetReplyContent(v string) *UpdateAuditRequestRequestAuditRecord {
	s.ReplyContent = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetSuspensionDate(v string) *UpdateAuditRequestRequestAuditRecord {
	s.SuspensionDate = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) SetWebAddress(v string) *UpdateAuditRequestRequestAuditRecord {
	s.WebAddress = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecord) Validate() error {
	if s.CompanyAddress != nil {
		for _, item := range s.CompanyAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CompanyTel != nil {
		for _, item := range s.CompanyTel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAuditRequestRequestAuditRecordCompanyAddress struct {
	// example:
	//
	// 123 Sunshine Street, City, Country
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// example:
	//
	// xxx company
	CompanyAddressTitle *string `json:"CompanyAddressTitle,omitempty" xml:"CompanyAddressTitle,omitempty"`
}

func (s UpdateAuditRequestRequestAuditRecordCompanyAddress) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestRequestAuditRecordCompanyAddress) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyAddress) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyAddress) GetCompanyAddressTitle() *string {
	return s.CompanyAddressTitle
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyAddress) SetCompanyAddress(v string) *UpdateAuditRequestRequestAuditRecordCompanyAddress {
	s.CompanyAddress = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyAddress) SetCompanyAddressTitle(v string) *UpdateAuditRequestRequestAuditRecordCompanyAddress {
	s.CompanyAddressTitle = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyAddress) Validate() error {
	return dara.Validate(s)
}

type UpdateAuditRequestRequestAuditRecordCompanyTel struct {
	// example:
	//
	// 07364245xxxx
	CompanyTelNumber *string `json:"CompanyTelNumber,omitempty" xml:"CompanyTelNumber,omitempty"`
	// example:
	//
	// xxx
	CompanyTelTitle *string `json:"CompanyTelTitle,omitempty" xml:"CompanyTelTitle,omitempty"`
}

func (s UpdateAuditRequestRequestAuditRecordCompanyTel) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestRequestAuditRecordCompanyTel) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyTel) GetCompanyTelNumber() *string {
	return s.CompanyTelNumber
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyTel) GetCompanyTelTitle() *string {
	return s.CompanyTelTitle
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyTel) SetCompanyTelNumber(v string) *UpdateAuditRequestRequestAuditRecordCompanyTel {
	s.CompanyTelNumber = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyTel) SetCompanyTelTitle(v string) *UpdateAuditRequestRequestAuditRecordCompanyTel {
	s.CompanyTelTitle = &v
	return s
}

func (s *UpdateAuditRequestRequestAuditRecordCompanyTel) Validate() error {
	return dara.Validate(s)
}
