// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberByRequestNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetViberByRequestNoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetViberByRequestNoResponseBody
	GetCode() *string
	SetData(v *GetViberByRequestNoResponseBodyData) *GetViberByRequestNoResponseBody
	GetData() *GetViberByRequestNoResponseBodyData
	SetMessage(v string) *GetViberByRequestNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetViberByRequestNoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetViberByRequestNoResponseBody
	GetSuccess() *bool
}

type GetViberByRequestNoResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetViberByRequestNoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8**9-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetViberByRequestNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetViberByRequestNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetViberByRequestNoResponseBody) GetData() *GetViberByRequestNoResponseBodyData {
	return s.Data
}

func (s *GetViberByRequestNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetViberByRequestNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetViberByRequestNoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetViberByRequestNoResponseBody) SetAccessDeniedDetail(v string) *GetViberByRequestNoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetViberByRequestNoResponseBody) SetCode(v string) *GetViberByRequestNoResponseBody {
	s.Code = &v
	return s
}

func (s *GetViberByRequestNoResponseBody) SetData(v *GetViberByRequestNoResponseBodyData) *GetViberByRequestNoResponseBody {
	s.Data = v
	return s
}

func (s *GetViberByRequestNoResponseBody) SetMessage(v string) *GetViberByRequestNoResponseBody {
	s.Message = &v
	return s
}

func (s *GetViberByRequestNoResponseBody) SetRequestId(v string) *GetViberByRequestNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViberByRequestNoResponseBody) SetSuccess(v bool) *GetViberByRequestNoResponseBody {
	s.Success = &v
	return s
}

func (s *GetViberByRequestNoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetViberByRequestNoResponseBodyData struct {
	AuditRecord *GetViberByRequestNoResponseBodyDataAuditRecord `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty" type:"Struct"`
	// example:
	//
	// pass
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// example:
	//
	// 1785465
	AuditTime *string `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// example:
	//
	// 11
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1753236426000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 17833636
	GmtModifier *string `json:"GmtModifier,omitempty" xml:"GmtModifier,omitempty"`
	// ID。
	//
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 735221
	MaapServiceNo *string `json:"MaapServiceNo,omitempty" xml:"MaapServiceNo,omitempty"`
	// example:
	//
	// 11
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// aa
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 5435435435
	RequestNo *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	// example:
	//
	// viberOpen
	RequestType *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	// example:
	//
	// A
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 4534**
	SubscriberCode *string `json:"SubscriberCode,omitempty" xml:"SubscriberCode,omitempty"`
}

func (s GetViberByRequestNoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponseBodyData) GetAuditRecord() *GetViberByRequestNoResponseBodyDataAuditRecord {
	return s.AuditRecord
}

func (s *GetViberByRequestNoResponseBodyData) GetAuditResult() *string {
	return s.AuditResult
}

func (s *GetViberByRequestNoResponseBodyData) GetAuditTime() *string {
	return s.AuditTime
}

func (s *GetViberByRequestNoResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetViberByRequestNoResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetViberByRequestNoResponseBodyData) GetGmtModifier() *string {
	return s.GmtModifier
}

func (s *GetViberByRequestNoResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetViberByRequestNoResponseBodyData) GetMaapServiceNo() *string {
	return s.MaapServiceNo
}

func (s *GetViberByRequestNoResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *GetViberByRequestNoResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *GetViberByRequestNoResponseBodyData) GetRequestNo() *string {
	return s.RequestNo
}

func (s *GetViberByRequestNoResponseBodyData) GetRequestType() *string {
	return s.RequestType
}

func (s *GetViberByRequestNoResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetViberByRequestNoResponseBodyData) GetSubscriberCode() *string {
	return s.SubscriberCode
}

func (s *GetViberByRequestNoResponseBodyData) SetAuditRecord(v *GetViberByRequestNoResponseBodyDataAuditRecord) *GetViberByRequestNoResponseBodyData {
	s.AuditRecord = v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetAuditResult(v string) *GetViberByRequestNoResponseBodyData {
	s.AuditResult = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetAuditTime(v string) *GetViberByRequestNoResponseBodyData {
	s.AuditTime = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetCreator(v string) *GetViberByRequestNoResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetGmtCreate(v string) *GetViberByRequestNoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetGmtModifier(v string) *GetViberByRequestNoResponseBodyData {
	s.GmtModifier = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetId(v string) *GetViberByRequestNoResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetMaapServiceNo(v string) *GetViberByRequestNoResponseBodyData {
	s.MaapServiceNo = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetModifier(v string) *GetViberByRequestNoResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetReason(v string) *GetViberByRequestNoResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetRequestNo(v string) *GetViberByRequestNoResponseBodyData {
	s.RequestNo = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetRequestType(v string) *GetViberByRequestNoResponseBodyData {
	s.RequestType = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetState(v string) *GetViberByRequestNoResponseBodyData {
	s.State = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) SetSubscriberCode(v string) *GetViberByRequestNoResponseBodyData {
	s.SubscriberCode = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyData) Validate() error {
	if s.AuditRecord != nil {
		if err := s.AuditRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetViberByRequestNoResponseBodyDataAuditRecord struct {
	// example:
	//
	// eighteenUp
	AgeLimit *string `json:"AgeLimit,omitempty" xml:"AgeLimit,omitempty"`
	// example:
	//
	// aaa
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// example:
	//
	// aaaa
	BusinessAccountName *string `json:"BusinessAccountName,omitempty" xml:"BusinessAccountName,omitempty"`
	// example:
	//
	// 43645465465
	BusinessLicenseRegistrationNumber *string                                                         `json:"BusinessLicenseRegistrationNumber,omitempty" xml:"BusinessLicenseRegistrationNumber,omitempty"`
	CompanyAddress                    []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty" type:"Repeated"`
	// example:
	//
	// english
	CompanyEnglishName *string `json:"CompanyEnglishName,omitempty" xml:"CompanyEnglishName,omitempty"`
	// example:
	//
	// xxxxxxcompany
	CompanyLegalName *string `json:"CompanyLegalName,omitempty" xml:"CompanyLegalName,omitempty"`
	// example:
	//
	// jack
	CompanyLegalPerson *string `json:"CompanyLegalPerson,omitempty" xml:"CompanyLegalPerson,omitempty"`
	// example:
	//
	// new
	CompanyProfile *string `json:"CompanyProfile,omitempty" xml:"CompanyProfile,omitempty"`
	// example:
	//
	// beijing
	CompanyRegisteredCountry *string                                                     `json:"CompanyRegisteredCountry,omitempty" xml:"CompanyRegisteredCountry,omitempty"`
	CompanyTel               []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel `json:"CompanyTel,omitempty" xml:"CompanyTel,omitempty" type:"Repeated"`
	// example:
	//
	// address
	CompleteAddressOfHeadquarters *string `json:"CompleteAddressOfHeadquarters,omitempty" xml:"CompleteAddressOfHeadquarters,omitempty"`
	// example:
	//
	// 示例值示例值
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// N
	EnableAutoReply *string `json:"EnableAutoReply,omitempty" xml:"EnableAutoReply,omitempty"`
	// example:
	//
	// Description
	IndustryDescription *string `json:"IndustryDescription,omitempty" xml:"IndustryDescription,omitempty"`
	// example:
	//
	// cateringServices
	IndustryInvolved *string `json:"IndustryInvolved,omitempty" xml:"IndustryInvolved,omitempty"`
	// example:
	//
	// https://xxxxxxxxxx
	LetterGuarantee                        *string   `json:"LetterGuarantee,omitempty" xml:"LetterGuarantee,omitempty"`
	Logo                                   []*string `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
	MessageDestinationCountry              []*string `json:"MessageDestinationCountry,omitempty" xml:"MessageDestinationCountry,omitempty" type:"Repeated"`
	MessageDestinationInternationalCountry []*string `json:"MessageDestinationInternationalCountry,omitempty" xml:"MessageDestinationInternationalCountry,omitempty" type:"Repeated"`
	// example:
	//
	// aaa
	MessageDialogueIntroduction *string `json:"MessageDialogueIntroduction,omitempty" xml:"MessageDialogueIntroduction,omitempty"`
	// example:
	//
	// 2025-08-05
	MessageEnableDate *string `json:"MessageEnableDate,omitempty" xml:"MessageEnableDate,omitempty"`
	// example:
	//
	// name
	MessageSessionName *string `json:"MessageSessionName,omitempty" xml:"MessageSessionName,omitempty"`
	// example:
	//
	// NowRecovery
	NowRecovery *string `json:"NowRecovery,omitempty" xml:"NowRecovery,omitempty"`
	// example:
	//
	// https://wwwxxxx
	OtherLetterGuarantee *string `json:"OtherLetterGuarantee,omitempty" xml:"OtherLetterGuarantee,omitempty"`
	// example:
	//
	// 2025-11-11
	RecoveryDate *string `json:"RecoveryDate,omitempty" xml:"RecoveryDate,omitempty"`
	// example:
	//
	// aaa
	ReplyContent *string `json:"ReplyContent,omitempty" xml:"ReplyContent,omitempty"`
	// SuspensionDate
	//
	// example:
	//
	// SuspensionDate
	SuspensionDate *string `json:"SuspensionDate,omitempty" xml:"SuspensionDate,omitempty"`
	// example:
	//
	// https://xxxxxx
	WebAddress *string `json:"WebAddress,omitempty" xml:"WebAddress,omitempty"`
}

func (s GetViberByRequestNoResponseBodyDataAuditRecord) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponseBodyDataAuditRecord) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetAgeLimit() *string {
	return s.AgeLimit
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetBusinessAccountName() *string {
	return s.BusinessAccountName
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetBusinessLicenseRegistrationNumber() *string {
	return s.BusinessLicenseRegistrationNumber
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyAddress() []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress {
	return s.CompanyAddress
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyEnglishName() *string {
	return s.CompanyEnglishName
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyLegalName() *string {
	return s.CompanyLegalName
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyLegalPerson() *string {
	return s.CompanyLegalPerson
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyProfile() *string {
	return s.CompanyProfile
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyRegisteredCountry() *string {
	return s.CompanyRegisteredCountry
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompanyTel() []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel {
	return s.CompanyTel
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetCompleteAddressOfHeadquarters() *string {
	return s.CompleteAddressOfHeadquarters
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetContactMail() *string {
	return s.ContactMail
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetEnableAutoReply() *string {
	return s.EnableAutoReply
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetIndustryDescription() *string {
	return s.IndustryDescription
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetIndustryInvolved() *string {
	return s.IndustryInvolved
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetLetterGuarantee() *string {
	return s.LetterGuarantee
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetLogo() []*string {
	return s.Logo
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetMessageDialogueIntroduction() *string {
	return s.MessageDialogueIntroduction
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetMessageEnableDate() *string {
	return s.MessageEnableDate
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetMessageSessionName() *string {
	return s.MessageSessionName
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetNowRecovery() *string {
	return s.NowRecovery
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetOtherLetterGuarantee() *string {
	return s.OtherLetterGuarantee
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetRecoveryDate() *string {
	return s.RecoveryDate
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetReplyContent() *string {
	return s.ReplyContent
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetSuspensionDate() *string {
	return s.SuspensionDate
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) GetWebAddress() *string {
	return s.WebAddress
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetAgeLimit(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.AgeLimit = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetApplyReason(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.ApplyReason = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetBusinessAccountName(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.BusinessAccountName = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetBusinessLicenseRegistrationNumber(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.BusinessLicenseRegistrationNumber = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyAddress(v []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyAddress = v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyEnglishName(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyEnglishName = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyLegalName(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyLegalName = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyLegalPerson(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyLegalPerson = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyProfile(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyProfile = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyRegisteredCountry(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyRegisteredCountry = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompanyTel(v []*GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompanyTel = v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetCompleteAddressOfHeadquarters(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.CompleteAddressOfHeadquarters = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetContactMail(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.ContactMail = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetEnableAutoReply(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.EnableAutoReply = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetIndustryDescription(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.IndustryDescription = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetIndustryInvolved(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.IndustryInvolved = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetLetterGuarantee(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.LetterGuarantee = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetLogo(v []*string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.Logo = v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetMessageDestinationCountry(v []*string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.MessageDestinationCountry = v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetMessageDestinationInternationalCountry(v []*string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetMessageDialogueIntroduction(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.MessageDialogueIntroduction = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetMessageEnableDate(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.MessageEnableDate = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetMessageSessionName(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.MessageSessionName = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetNowRecovery(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.NowRecovery = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetOtherLetterGuarantee(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.OtherLetterGuarantee = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetRecoveryDate(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.RecoveryDate = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetReplyContent(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.ReplyContent = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetSuspensionDate(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.SuspensionDate = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) SetWebAddress(v string) *GetViberByRequestNoResponseBodyDataAuditRecord {
	s.WebAddress = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecord) Validate() error {
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

type GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress struct {
	// example:
	//
	// address
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// example:
	//
	// 111
	CompanyAddressTitle *string `json:"CompanyAddressTitle,omitempty" xml:"CompanyAddressTitle,omitempty"`
}

func (s GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) GetCompanyAddressTitle() *string {
	return s.CompanyAddressTitle
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) SetCompanyAddress(v string) *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress {
	s.CompanyAddress = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) SetCompanyAddressTitle(v string) *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress {
	s.CompanyAddressTitle = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyAddress) Validate() error {
	return dara.Validate(s)
}

type GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel struct {
	// example:
	//
	// 134213213
	CompanyTelNumber *string `json:"CompanyTelNumber,omitempty" xml:"CompanyTelNumber,omitempty"`
	// example:
	//
	// xxx
	CompanyTelTitle *string `json:"CompanyTelTitle,omitempty" xml:"CompanyTelTitle,omitempty"`
}

func (s GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) GetCompanyTelNumber() *string {
	return s.CompanyTelNumber
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) GetCompanyTelTitle() *string {
	return s.CompanyTelTitle
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) SetCompanyTelNumber(v string) *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel {
	s.CompanyTelNumber = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) SetCompanyTelTitle(v string) *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel {
	s.CompanyTelTitle = &v
	return s
}

func (s *GetViberByRequestNoResponseBodyDataAuditRecordCompanyTel) Validate() error {
	return dara.Validate(s)
}
