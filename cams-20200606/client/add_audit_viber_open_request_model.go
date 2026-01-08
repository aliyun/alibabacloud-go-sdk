// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditViberOpenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecord(v *AddAuditViberOpenRequestAuditRecord) *AddAuditViberOpenRequest
	GetAuditRecord() *AddAuditViberOpenRequestAuditRecord
	SetAuditResult(v string) *AddAuditViberOpenRequest
	GetAuditResult() *string
	SetCustSpaceId(v string) *AddAuditViberOpenRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *AddAuditViberOpenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddAuditViberOpenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAuditViberOpenRequest
	GetResourceOwnerId() *int64
}

type AddAuditViberOpenRequest struct {
	AuditRecord *AddAuditViberOpenRequestAuditRecord `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAuditViberOpenRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenRequest) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenRequest) GetAuditRecord() *AddAuditViberOpenRequestAuditRecord {
	return s.AuditRecord
}

func (s *AddAuditViberOpenRequest) GetAuditResult() *string {
	return s.AuditResult
}

func (s *AddAuditViberOpenRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddAuditViberOpenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAuditViberOpenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAuditViberOpenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAuditViberOpenRequest) SetAuditRecord(v *AddAuditViberOpenRequestAuditRecord) *AddAuditViberOpenRequest {
	s.AuditRecord = v
	return s
}

func (s *AddAuditViberOpenRequest) SetAuditResult(v string) *AddAuditViberOpenRequest {
	s.AuditResult = &v
	return s
}

func (s *AddAuditViberOpenRequest) SetCustSpaceId(v string) *AddAuditViberOpenRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddAuditViberOpenRequest) SetOwnerId(v int64) *AddAuditViberOpenRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAuditViberOpenRequest) SetResourceOwnerAccount(v string) *AddAuditViberOpenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAuditViberOpenRequest) SetResourceOwnerId(v int64) *AddAuditViberOpenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAuditViberOpenRequest) Validate() error {
	if s.AuditRecord != nil {
		if err := s.AuditRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAuditViberOpenRequestAuditRecord struct {
	// example:
	//
	// 示例值示例值
	AgeLimit *string `json:"AgeLimit,omitempty" xml:"AgeLimit,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	BusinessAccountName *string `json:"BusinessAccountName,omitempty" xml:"BusinessAccountName,omitempty"`
	// example:
	//
	// 示例值示例值
	BusinessLicenseRegistrationNumber *string                                              `json:"BusinessLicenseRegistrationNumber,omitempty" xml:"BusinessLicenseRegistrationNumber,omitempty"`
	CompanyAddress                    []*AddAuditViberOpenRequestAuditRecordCompanyAddress `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty" type:"Repeated"`
	// example:
	//
	// 7Pj6
	CompanyEnglishName *string `json:"CompanyEnglishName,omitempty" xml:"CompanyEnglishName,omitempty"`
	// example:
	//
	// 示例值
	CompanyLegalName *string `json:"CompanyLegalName,omitempty" xml:"CompanyLegalName,omitempty"`
	// example:
	//
	// jerry
	CompanyLegalPerson *string `json:"CompanyLegalPerson,omitempty" xml:"CompanyLegalPerson,omitempty"`
	// example:
	//
	// CompanyProfile
	CompanyProfile *string `json:"CompanyProfile,omitempty" xml:"CompanyProfile,omitempty"`
	// example:
	//
	// cn
	CompanyRegisteredCountry *string                                          `json:"CompanyRegisteredCountry,omitempty" xml:"CompanyRegisteredCountry,omitempty"`
	CompanyTel               []*AddAuditViberOpenRequestAuditRecordCompanyTel `json:"CompanyTel,omitempty" xml:"CompanyTel,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值示例值
	CompleteAddressOfHeadquarters *string `json:"CompleteAddressOfHeadquarters,omitempty" xml:"CompleteAddressOfHeadquarters,omitempty"`
	// example:
	//
	// 123@123.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// 示例值
	EnableAutoReply *string `json:"EnableAutoReply,omitempty" xml:"EnableAutoReply,omitempty"`
	// example:
	//
	// 示例值示例值
	IndustryDescription *string `json:"IndustryDescription,omitempty" xml:"IndustryDescription,omitempty"`
	// example:
	//
	// NFUwytiBnm11
	IndustryInvolved *string `json:"IndustryInvolved,omitempty" xml:"IndustryInvolved,omitempty"`
	// example:
	//
	// tiCaYNHR8ttt
	LetterGuarantee                        *string   `json:"LetterGuarantee,omitempty" xml:"LetterGuarantee,omitempty"`
	Logo                                   []*string `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
	MessageDestinationCountry              []*string `json:"MessageDestinationCountry,omitempty" xml:"MessageDestinationCountry,omitempty" type:"Repeated"`
	MessageDestinationInternationalCountry []*string `json:"MessageDestinationInternationalCountry,omitempty" xml:"MessageDestinationInternationalCountry,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值示例值
	MessageDialogueIntroduction *string `json:"MessageDialogueIntroduction,omitempty" xml:"MessageDialogueIntroduction,omitempty"`
	// example:
	//
	// 示例值示例值
	MessageEnableDate *string `json:"MessageEnableDate,omitempty" xml:"MessageEnableDate,omitempty"`
	// example:
	//
	// nPHK8uaRo
	MessageSessionName *string `json:"MessageSessionName,omitempty" xml:"MessageSessionName,omitempty"`
	// example:
	//
	// 示例值示例值
	OtherLetterGuarantee *string `json:"OtherLetterGuarantee,omitempty" xml:"OtherLetterGuarantee,omitempty"`
	// example:
	//
	// hello
	ReplyContent *string `json:"ReplyContent,omitempty" xml:"ReplyContent,omitempty"`
	// example:
	//
	// 示例值示例值
	WebAddress *string `json:"WebAddress,omitempty" xml:"WebAddress,omitempty"`
}

func (s AddAuditViberOpenRequestAuditRecord) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenRequestAuditRecord) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenRequestAuditRecord) GetAgeLimit() *string {
	return s.AgeLimit
}

func (s *AddAuditViberOpenRequestAuditRecord) GetBusinessAccountName() *string {
	return s.BusinessAccountName
}

func (s *AddAuditViberOpenRequestAuditRecord) GetBusinessLicenseRegistrationNumber() *string {
	return s.BusinessLicenseRegistrationNumber
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyAddress() []*AddAuditViberOpenRequestAuditRecordCompanyAddress {
	return s.CompanyAddress
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyEnglishName() *string {
	return s.CompanyEnglishName
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyLegalName() *string {
	return s.CompanyLegalName
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyLegalPerson() *string {
	return s.CompanyLegalPerson
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyProfile() *string {
	return s.CompanyProfile
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyRegisteredCountry() *string {
	return s.CompanyRegisteredCountry
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompanyTel() []*AddAuditViberOpenRequestAuditRecordCompanyTel {
	return s.CompanyTel
}

func (s *AddAuditViberOpenRequestAuditRecord) GetCompleteAddressOfHeadquarters() *string {
	return s.CompleteAddressOfHeadquarters
}

func (s *AddAuditViberOpenRequestAuditRecord) GetContactMail() *string {
	return s.ContactMail
}

func (s *AddAuditViberOpenRequestAuditRecord) GetEnableAutoReply() *string {
	return s.EnableAutoReply
}

func (s *AddAuditViberOpenRequestAuditRecord) GetIndustryDescription() *string {
	return s.IndustryDescription
}

func (s *AddAuditViberOpenRequestAuditRecord) GetIndustryInvolved() *string {
	return s.IndustryInvolved
}

func (s *AddAuditViberOpenRequestAuditRecord) GetLetterGuarantee() *string {
	return s.LetterGuarantee
}

func (s *AddAuditViberOpenRequestAuditRecord) GetLogo() []*string {
	return s.Logo
}

func (s *AddAuditViberOpenRequestAuditRecord) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *AddAuditViberOpenRequestAuditRecord) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *AddAuditViberOpenRequestAuditRecord) GetMessageDialogueIntroduction() *string {
	return s.MessageDialogueIntroduction
}

func (s *AddAuditViberOpenRequestAuditRecord) GetMessageEnableDate() *string {
	return s.MessageEnableDate
}

func (s *AddAuditViberOpenRequestAuditRecord) GetMessageSessionName() *string {
	return s.MessageSessionName
}

func (s *AddAuditViberOpenRequestAuditRecord) GetOtherLetterGuarantee() *string {
	return s.OtherLetterGuarantee
}

func (s *AddAuditViberOpenRequestAuditRecord) GetReplyContent() *string {
	return s.ReplyContent
}

func (s *AddAuditViberOpenRequestAuditRecord) GetWebAddress() *string {
	return s.WebAddress
}

func (s *AddAuditViberOpenRequestAuditRecord) SetAgeLimit(v string) *AddAuditViberOpenRequestAuditRecord {
	s.AgeLimit = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetBusinessAccountName(v string) *AddAuditViberOpenRequestAuditRecord {
	s.BusinessAccountName = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetBusinessLicenseRegistrationNumber(v string) *AddAuditViberOpenRequestAuditRecord {
	s.BusinessLicenseRegistrationNumber = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyAddress(v []*AddAuditViberOpenRequestAuditRecordCompanyAddress) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyAddress = v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyEnglishName(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyEnglishName = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyLegalName(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyLegalName = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyLegalPerson(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyLegalPerson = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyProfile(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyProfile = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyRegisteredCountry(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyRegisteredCountry = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompanyTel(v []*AddAuditViberOpenRequestAuditRecordCompanyTel) *AddAuditViberOpenRequestAuditRecord {
	s.CompanyTel = v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetCompleteAddressOfHeadquarters(v string) *AddAuditViberOpenRequestAuditRecord {
	s.CompleteAddressOfHeadquarters = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetContactMail(v string) *AddAuditViberOpenRequestAuditRecord {
	s.ContactMail = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetEnableAutoReply(v string) *AddAuditViberOpenRequestAuditRecord {
	s.EnableAutoReply = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetIndustryDescription(v string) *AddAuditViberOpenRequestAuditRecord {
	s.IndustryDescription = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetIndustryInvolved(v string) *AddAuditViberOpenRequestAuditRecord {
	s.IndustryInvolved = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetLetterGuarantee(v string) *AddAuditViberOpenRequestAuditRecord {
	s.LetterGuarantee = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetLogo(v []*string) *AddAuditViberOpenRequestAuditRecord {
	s.Logo = v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetMessageDestinationCountry(v []*string) *AddAuditViberOpenRequestAuditRecord {
	s.MessageDestinationCountry = v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetMessageDestinationInternationalCountry(v []*string) *AddAuditViberOpenRequestAuditRecord {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetMessageDialogueIntroduction(v string) *AddAuditViberOpenRequestAuditRecord {
	s.MessageDialogueIntroduction = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetMessageEnableDate(v string) *AddAuditViberOpenRequestAuditRecord {
	s.MessageEnableDate = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetMessageSessionName(v string) *AddAuditViberOpenRequestAuditRecord {
	s.MessageSessionName = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetOtherLetterGuarantee(v string) *AddAuditViberOpenRequestAuditRecord {
	s.OtherLetterGuarantee = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetReplyContent(v string) *AddAuditViberOpenRequestAuditRecord {
	s.ReplyContent = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) SetWebAddress(v string) *AddAuditViberOpenRequestAuditRecord {
	s.WebAddress = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecord) Validate() error {
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

type AddAuditViberOpenRequestAuditRecordCompanyAddress struct {
	// example:
	//
	// address
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// example:
	//
	// title
	CompanyAddressTitle *string `json:"CompanyAddressTitle,omitempty" xml:"CompanyAddressTitle,omitempty"`
}

func (s AddAuditViberOpenRequestAuditRecordCompanyAddress) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenRequestAuditRecordCompanyAddress) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyAddress) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyAddress) GetCompanyAddressTitle() *string {
	return s.CompanyAddressTitle
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyAddress) SetCompanyAddress(v string) *AddAuditViberOpenRequestAuditRecordCompanyAddress {
	s.CompanyAddress = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyAddress) SetCompanyAddressTitle(v string) *AddAuditViberOpenRequestAuditRecordCompanyAddress {
	s.CompanyAddressTitle = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyAddress) Validate() error {
	return dara.Validate(s)
}

type AddAuditViberOpenRequestAuditRecordCompanyTel struct {
	// example:
	//
	// 156********
	CompanyTelNumber *string `json:"CompanyTelNumber,omitempty" xml:"CompanyTelNumber,omitempty"`
	// example:
	//
	// title
	CompanyTelTitle *string `json:"CompanyTelTitle,omitempty" xml:"CompanyTelTitle,omitempty"`
}

func (s AddAuditViberOpenRequestAuditRecordCompanyTel) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenRequestAuditRecordCompanyTel) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyTel) GetCompanyTelNumber() *string {
	return s.CompanyTelNumber
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyTel) GetCompanyTelTitle() *string {
	return s.CompanyTelTitle
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyTel) SetCompanyTelNumber(v string) *AddAuditViberOpenRequestAuditRecordCompanyTel {
	s.CompanyTelNumber = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyTel) SetCompanyTelTitle(v string) *AddAuditViberOpenRequestAuditRecordCompanyTel {
	s.CompanyTelTitle = &v
	return s
}

func (s *AddAuditViberOpenRequestAuditRecordCompanyTel) Validate() error {
	return dara.Validate(s)
}
