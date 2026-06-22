// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditRequestByTypeUnAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAuditRequestByTypeUnAuditResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetAuditRequestByTypeUnAuditResponseBody
	GetCode() *string
	SetData(v *GetAuditRequestByTypeUnAuditResponseBodyData) *GetAuditRequestByTypeUnAuditResponseBody
	GetData() *GetAuditRequestByTypeUnAuditResponseBodyData
	SetMessage(v string) *GetAuditRequestByTypeUnAuditResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuditRequestByTypeUnAuditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuditRequestByTypeUnAuditResponseBody
	GetSuccess() *bool
}

type GetAuditRequestByTypeUnAuditResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAuditRequestByTypeUnAuditResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetData() *GetAuditRequestByTypeUnAuditResponseBodyData {
	return s.Data
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetAccessDeniedDetail(v string) *GetAuditRequestByTypeUnAuditResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetCode(v string) *GetAuditRequestByTypeUnAuditResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetData(v *GetAuditRequestByTypeUnAuditResponseBodyData) *GetAuditRequestByTypeUnAuditResponseBody {
	s.Data = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetMessage(v string) *GetAuditRequestByTypeUnAuditResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetRequestId(v string) *GetAuditRequestByTypeUnAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) SetSuccess(v bool) *GetAuditRequestByTypeUnAuditResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuditRequestByTypeUnAuditResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	AliUid      *string                                                  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AuditRecord *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// example:
	//
	// 示例值
	AuditTime *string `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 示例值示例值
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	GmtModifier *string `json:"GmtModifier,omitempty" xml:"GmtModifier,omitempty"`
	// example:
	//
	// 示例值示例值
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 示例值示例值
	MaapServiceNo *string `json:"MaapServiceNo,omitempty" xml:"MaapServiceNo,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestNo *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestType *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	// example:
	//
	// 示例值示例值
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 示例值示例值
	SubscriberCode *string `json:"SubscriberCode,omitempty" xml:"SubscriberCode,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetAliUid() *string {
	return s.AliUid
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetAuditRecord() *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	return s.AuditRecord
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetAuditResult() *string {
	return s.AuditResult
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetAuditTime() *string {
	return s.AuditTime
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetGmtModifier() *string {
	return s.GmtModifier
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetMaapServiceNo() *string {
	return s.MaapServiceNo
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetRequestNo() *string {
	return s.RequestNo
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetRequestType() *string {
	return s.RequestType
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) GetSubscriberCode() *string {
	return s.SubscriberCode
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetAliUid(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetAuditRecord(v *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.AuditRecord = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetAuditResult(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.AuditResult = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetAuditTime(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.AuditTime = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetCreator(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetGmtCreate(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetGmtModifier(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.GmtModifier = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetId(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetMaapServiceNo(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.MaapServiceNo = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetModifier(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetRequestNo(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.RequestNo = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetRequestType(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.RequestType = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetState(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.State = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) SetSubscriberCode(v string) *GetAuditRequestByTypeUnAuditResponseBodyData {
	s.SubscriberCode = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyData) Validate() error {
	if s.AuditRecord != nil {
		if err := s.AuditRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord struct {
	// example:
	//
	// 示例值示例值示例值
	AgeLimit *string `json:"AgeLimit,omitempty" xml:"AgeLimit,omitempty"`
	// example:
	//
	// 示例值示例值
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// example:
	//
	// 示例值
	BusinessAccountName *string                                                                  `json:"BusinessAccountName,omitempty" xml:"BusinessAccountName,omitempty"`
	CompanyAddress      []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
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
	// 示例值示例值示例值
	CompanyProfile           *string                                                              `json:"CompanyProfile,omitempty" xml:"CompanyProfile,omitempty"`
	CompanyRegisteredCountry []*string                                                            `json:"CompanyRegisteredCountry,omitempty" xml:"CompanyRegisteredCountry,omitempty" type:"Repeated"`
	CompanyTel               []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel `json:"CompanyTel,omitempty" xml:"CompanyTel,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// example:
	//
	// 示例值示例值
	EnableAutoReply *string `json:"EnableAutoReply,omitempty" xml:"EnableAutoReply,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	IndustryDescription *string `json:"IndustryDescription,omitempty" xml:"IndustryDescription,omitempty"`
	// example:
	//
	// 示例值
	IndustryInvolved *string `json:"IndustryInvolved,omitempty" xml:"IndustryInvolved,omitempty"`
	// example:
	//
	// 示例值示例值
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
	// 示例值示例值
	MessageSessionName *string `json:"MessageSessionName,omitempty" xml:"MessageSessionName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	NowRecovery *string `json:"NowRecovery,omitempty" xml:"NowRecovery,omitempty"`
	// example:
	//
	// 示例值示例值
	RecoveryDate *string `json:"RecoveryDate,omitempty" xml:"RecoveryDate,omitempty"`
	// example:
	//
	// 示例值示例值
	ReplyContent *string `json:"ReplyContent,omitempty" xml:"ReplyContent,omitempty"`
	// example:
	//
	// 示例值示例值
	SuspensionDate *string `json:"SuspensionDate,omitempty" xml:"SuspensionDate,omitempty"`
	// example:
	//
	// 示例值
	WebAddress *string `json:"WebAddress,omitempty" xml:"WebAddress,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetAgeLimit() *string {
	return s.AgeLimit
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetBusinessAccountName() *string {
	return s.BusinessAccountName
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyAddress() []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress {
	return s.CompanyAddress
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyEnglishName() *string {
	return s.CompanyEnglishName
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyLegalName() *string {
	return s.CompanyLegalName
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyLegalPerson() *string {
	return s.CompanyLegalPerson
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyProfile() *string {
	return s.CompanyProfile
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyRegisteredCountry() []*string {
	return s.CompanyRegisteredCountry
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetCompanyTel() []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel {
	return s.CompanyTel
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetContactMail() *string {
	return s.ContactMail
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetEnableAutoReply() *string {
	return s.EnableAutoReply
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetIndustryDescription() *string {
	return s.IndustryDescription
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetIndustryInvolved() *string {
	return s.IndustryInvolved
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetLetterGuarantee() *string {
	return s.LetterGuarantee
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetLogo() []*string {
	return s.Logo
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetMessageDestinationCountry() []*string {
	return s.MessageDestinationCountry
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetMessageDestinationInternationalCountry() []*string {
	return s.MessageDestinationInternationalCountry
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetMessageDialogueIntroduction() *string {
	return s.MessageDialogueIntroduction
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetMessageEnableDate() *string {
	return s.MessageEnableDate
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetMessageSessionName() *string {
	return s.MessageSessionName
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetNowRecovery() *string {
	return s.NowRecovery
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetRecoveryDate() *string {
	return s.RecoveryDate
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetReplyContent() *string {
	return s.ReplyContent
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetSuspensionDate() *string {
	return s.SuspensionDate
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) GetWebAddress() *string {
	return s.WebAddress
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetAgeLimit(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.AgeLimit = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetApplyReason(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.ApplyReason = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetBusinessAccountName(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.BusinessAccountName = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyAddress(v []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyAddress = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyEnglishName(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyEnglishName = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyLegalName(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyLegalName = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyLegalPerson(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyLegalPerson = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyProfile(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyProfile = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyRegisteredCountry(v []*string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyRegisteredCountry = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetCompanyTel(v []*GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.CompanyTel = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetContactMail(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.ContactMail = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetEnableAutoReply(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.EnableAutoReply = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetIndustryDescription(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.IndustryDescription = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetIndustryInvolved(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.IndustryInvolved = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetLetterGuarantee(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.LetterGuarantee = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetLogo(v []*string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.Logo = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetMessageDestinationCountry(v []*string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.MessageDestinationCountry = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetMessageDestinationInternationalCountry(v []*string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.MessageDestinationInternationalCountry = v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetMessageDialogueIntroduction(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.MessageDialogueIntroduction = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetMessageEnableDate(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.MessageEnableDate = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetMessageSessionName(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.MessageSessionName = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetNowRecovery(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.NowRecovery = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetRecoveryDate(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.RecoveryDate = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetReplyContent(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.ReplyContent = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetSuspensionDate(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.SuspensionDate = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) SetWebAddress(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord {
	s.WebAddress = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecord) Validate() error {
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

type GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress struct {
	// example:
	//
	// 示例值示例值示例值
	CompanyAddress *string `json:"CompanyAddress,omitempty" xml:"CompanyAddress,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CompanyAddressTitle *string `json:"CompanyAddressTitle,omitempty" xml:"CompanyAddressTitle,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) GetCompanyAddress() *string {
	return s.CompanyAddress
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) GetCompanyAddressTitle() *string {
	return s.CompanyAddressTitle
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) SetCompanyAddress(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress {
	s.CompanyAddress = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) SetCompanyAddressTitle(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress {
	s.CompanyAddressTitle = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyAddress) Validate() error {
	return dara.Validate(s)
}

type GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel struct {
	// example:
	//
	// 示例值
	CompanyTelNumber *string `json:"CompanyTelNumber,omitempty" xml:"CompanyTelNumber,omitempty"`
	// example:
	//
	// 示例值
	CompanyTelTitle *string `json:"CompanyTelTitle,omitempty" xml:"CompanyTelTitle,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) GetCompanyTelNumber() *string {
	return s.CompanyTelNumber
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) GetCompanyTelTitle() *string {
	return s.CompanyTelTitle
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) SetCompanyTelNumber(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel {
	s.CompanyTelNumber = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) SetCompanyTelTitle(v string) *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel {
	s.CompanyTelTitle = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditResponseBodyDataAuditRecordCompanyTel) Validate() error {
	return dara.Validate(s)
}
