// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CustomerNoteListDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CustomerNoteListDetailResponseBody
	GetCode() *string
	SetData(v *CustomerNoteListDetailResponseBodyData) *CustomerNoteListDetailResponseBody
	GetData() *CustomerNoteListDetailResponseBodyData
	SetHttpStatusCode(v int32) *CustomerNoteListDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CustomerNoteListDetailResponseBody
	GetMessage() *string
	SetMsg(v string) *CustomerNoteListDetailResponseBody
	GetMsg() *string
	SetRequestId(v string) *CustomerNoteListDetailResponseBody
	GetRequestId() *string
}

type CustomerNoteListDetailResponseBody struct {
	// The access denied details returned by the POP API when RAM permissions are missing.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CustomerNoteListDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the POP API.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The prompt message. This is the same as Message.
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomerNoteListDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListDetailResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerNoteListDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CustomerNoteListDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerNoteListDetailResponseBody) GetData() *CustomerNoteListDetailResponseBodyData {
	return s.Data
}

func (s *CustomerNoteListDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CustomerNoteListDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerNoteListDetailResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CustomerNoteListDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerNoteListDetailResponseBody) SetAccessDeniedDetail(v string) *CustomerNoteListDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetCode(v string) *CustomerNoteListDetailResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetData(v *CustomerNoteListDetailResponseBodyData) *CustomerNoteListDetailResponseBody {
	s.Data = v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetHttpStatusCode(v int32) *CustomerNoteListDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetMessage(v string) *CustomerNoteListDetailResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetMsg(v string) *CustomerNoteListDetailResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) SetRequestId(v string) *CustomerNoteListDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerNoteListDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CustomerNoteListDetailResponseBodyData struct {
	// The AI parsing result (JSON string).
	//
	// example:
	//
	// {"TouchDate":"2026-07-17 10:00:01"}
	AiResult *string `json:"AiResult,omitempty" xml:"AiResult,omitempty"`
	// The attachment list.
	Attachment []*CustomerNoteListDetailResponseBodyDataAttachment `json:"Attachment,omitempty" xml:"Attachment,omitempty" type:"Repeated"`
	// The contact information.
	//
	// example:
	//
	// 13833333333
	ContactInformation *string `json:"ContactInformation,omitempty" xml:"ContactInformation,omitempty"`
	// The contact name.
	//
	// example:
	//
	// John
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The UID of the creator.
	//
	// example:
	//
	// 291688841144601701
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The logon name of the creator.
	//
	// example:
	//
	// John
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The customer name.
	//
	// example:
	//
	// XXXX Co., Ltd
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// The customer UID.
	//
	// example:
	//
	// 1271202085096245
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// The creation time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-05-07 10:27:46
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The note content.
	//
	// example:
	//
	// Note content
	NoteContent *string `json:"NoteContent,omitempty" xml:"NoteContent,omitempty"`
	// The note ID.
	//
	// example:
	//
	// 1620737
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// The note type (CUSTOMER).
	//
	// example:
	//
	// CUSTOMER
	NoteType *string `json:"NoteType,omitempty" xml:"NoteType,omitempty"`
	// The note type label.
	//
	// example:
	//
	// Customer
	NoteTypeLabel *string `json:"NoteTypeLabel,omitempty" xml:"NoteTypeLabel,omitempty"`
	// The touch date (timestamp).
	//
	// example:
	//
	// 1784266662000
	TouchDate *string `json:"TouchDate,omitempty" xml:"TouchDate,omitempty"`
}

func (s CustomerNoteListDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomerNoteListDetailResponseBodyData) GetAiResult() *string {
	return s.AiResult
}

func (s *CustomerNoteListDetailResponseBodyData) GetAttachment() []*CustomerNoteListDetailResponseBodyDataAttachment {
	return s.Attachment
}

func (s *CustomerNoteListDetailResponseBodyData) GetContactInformation() *string {
	return s.ContactInformation
}

func (s *CustomerNoteListDetailResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *CustomerNoteListDetailResponseBodyData) GetCreator() *int64 {
	return s.Creator
}

func (s *CustomerNoteListDetailResponseBodyData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *CustomerNoteListDetailResponseBodyData) GetCustomerName() *string {
	return s.CustomerName
}

func (s *CustomerNoteListDetailResponseBodyData) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *CustomerNoteListDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CustomerNoteListDetailResponseBodyData) GetNoteContent() *string {
	return s.NoteContent
}

func (s *CustomerNoteListDetailResponseBodyData) GetNoteId() *int64 {
	return s.NoteId
}

func (s *CustomerNoteListDetailResponseBodyData) GetNoteType() *string {
	return s.NoteType
}

func (s *CustomerNoteListDetailResponseBodyData) GetNoteTypeLabel() *string {
	return s.NoteTypeLabel
}

func (s *CustomerNoteListDetailResponseBodyData) GetTouchDate() *string {
	return s.TouchDate
}

func (s *CustomerNoteListDetailResponseBodyData) SetAiResult(v string) *CustomerNoteListDetailResponseBodyData {
	s.AiResult = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetAttachment(v []*CustomerNoteListDetailResponseBodyDataAttachment) *CustomerNoteListDetailResponseBodyData {
	s.Attachment = v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetContactInformation(v string) *CustomerNoteListDetailResponseBodyData {
	s.ContactInformation = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetContactName(v string) *CustomerNoteListDetailResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetCreator(v int64) *CustomerNoteListDetailResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetCreatorName(v string) *CustomerNoteListDetailResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetCustomerName(v string) *CustomerNoteListDetailResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetCustomerUid(v int64) *CustomerNoteListDetailResponseBodyData {
	s.CustomerUid = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetGmtCreate(v string) *CustomerNoteListDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetNoteContent(v string) *CustomerNoteListDetailResponseBodyData {
	s.NoteContent = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetNoteId(v int64) *CustomerNoteListDetailResponseBodyData {
	s.NoteId = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetNoteType(v string) *CustomerNoteListDetailResponseBodyData {
	s.NoteType = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetNoteTypeLabel(v string) *CustomerNoteListDetailResponseBodyData {
	s.NoteTypeLabel = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) SetTouchDate(v string) *CustomerNoteListDetailResponseBodyData {
	s.TouchDate = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyData) Validate() error {
	if s.Attachment != nil {
		for _, item := range s.Attachment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CustomerNoteListDetailResponseBodyDataAttachment struct {
	// The attachment signature.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The attachment ID.
	//
	// example:
	//
	// 307
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The attachment name.
	//
	// example:
	//
	// Course Training
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attachment signature.
	//
	// example:
	//
	// hF2UrEMc4XWy990sh9LGM0+ScI8=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The attachment size in bytes.
	//
	// example:
	//
	// 111222121
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The attachment type.
	//
	// example:
	//
	// image/png
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CustomerNoteListDetailResponseBodyDataAttachment) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListDetailResponseBodyDataAttachment) GoString() string {
	return s.String()
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetId() *int64 {
	return s.Id
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetName() *string {
	return s.Name
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetSignature() *string {
	return s.Signature
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetSize() *int64 {
	return s.Size
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) GetType() *string {
	return s.Type
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetDownloadUrl(v string) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.DownloadUrl = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetId(v int64) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.Id = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetName(v string) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.Name = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetSignature(v string) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.Signature = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetSize(v int64) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.Size = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) SetType(v string) *CustomerNoteListDetailResponseBodyDataAttachment {
	s.Type = &v
	return s
}

func (s *CustomerNoteListDetailResponseBodyDataAttachment) Validate() error {
	return dara.Validate(s)
}
