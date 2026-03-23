// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetTicketResponseBody
	GetCode() *int32
	SetCreateTime(v string) *GetTicketResponseBody
	GetCreateTime() *string
	SetEmail(v string) *GetTicketResponseBody
	GetEmail() *string
	SetList(v []*GetTicketResponseBodyList) *GetTicketResponseBody
	GetList() []*GetTicketResponseBodyList
	SetMessage(v string) *GetTicketResponseBody
	GetMessage() *string
	SetNotifyTimeRange(v string) *GetTicketResponseBody
	GetNotifyTimeRange() *string
	SetPhone(v string) *GetTicketResponseBody
	GetPhone() *string
	SetProductCode(v string) *GetTicketResponseBody
	GetProductCode() *string
	SetRequestId(v string) *GetTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTicketResponseBody
	GetSuccess() *bool
	SetTicketId(v string) *GetTicketResponseBody
	GetTicketId() *string
	SetTicketStatus(v string) *GetTicketResponseBody
	GetTicketStatus() *string
	SetTitle(v string) *GetTicketResponseBody
	GetTitle() *string
}

type GetTicketResponseBody struct {
	AccessDeniedDetail *string                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime         *string                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Email              *string                      `json:"Email,omitempty" xml:"Email,omitempty"`
	List               []*GetTicketResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message            *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyTimeRange    *string                      `json:"NotifyTimeRange,omitempty" xml:"NotifyTimeRange,omitempty"`
	Phone              *string                      `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProductCode        *string                      `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RequestId          *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TicketId           *string                      `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TicketStatus       *string                      `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	Title              *string                      `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTicketResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTicketResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetTicketResponseBody) GetList() []*GetTicketResponseBodyList {
	return s.List
}

func (s *GetTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketResponseBody) GetNotifyTimeRange() *string {
	return s.NotifyTimeRange
}

func (s *GetTicketResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetTicketResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTicketResponseBody) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketResponseBody) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *GetTicketResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetTicketResponseBody) SetAccessDeniedDetail(v string) *GetTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTicketResponseBody) SetCode(v int32) *GetTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketResponseBody) SetCreateTime(v string) *GetTicketResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetEmail(v string) *GetTicketResponseBody {
	s.Email = &v
	return s
}

func (s *GetTicketResponseBody) SetList(v []*GetTicketResponseBodyList) *GetTicketResponseBody {
	s.List = v
	return s
}

func (s *GetTicketResponseBody) SetMessage(v string) *GetTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketResponseBody) SetNotifyTimeRange(v string) *GetTicketResponseBody {
	s.NotifyTimeRange = &v
	return s
}

func (s *GetTicketResponseBody) SetPhone(v string) *GetTicketResponseBody {
	s.Phone = &v
	return s
}

func (s *GetTicketResponseBody) SetProductCode(v string) *GetTicketResponseBody {
	s.ProductCode = &v
	return s
}

func (s *GetTicketResponseBody) SetRequestId(v string) *GetTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketResponseBody) SetSuccess(v bool) *GetTicketResponseBody {
	s.Success = &v
	return s
}

func (s *GetTicketResponseBody) SetTicketId(v string) *GetTicketResponseBody {
	s.TicketId = &v
	return s
}

func (s *GetTicketResponseBody) SetTicketStatus(v string) *GetTicketResponseBody {
	s.TicketStatus = &v
	return s
}

func (s *GetTicketResponseBody) SetTitle(v string) *GetTicketResponseBody {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTicketResponseBodyList struct {
	AttachList   []*string `json:"AttachList,omitempty" xml:"AttachList,omitempty" type:"Repeated"`
	Content      *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	FromOfficial *bool     `json:"FromOfficial,omitempty" xml:"FromOfficial,omitempty"`
	GmtCreated   *string   `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	NoteId       *string   `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	Special      *bool     `json:"Special,omitempty" xml:"Special,omitempty"`
}

func (s GetTicketResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyList) GetAttachList() []*string {
	return s.AttachList
}

func (s *GetTicketResponseBodyList) GetContent() *string {
	return s.Content
}

func (s *GetTicketResponseBodyList) GetFromOfficial() *bool {
	return s.FromOfficial
}

func (s *GetTicketResponseBodyList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *GetTicketResponseBodyList) GetNoteId() *string {
	return s.NoteId
}

func (s *GetTicketResponseBodyList) GetSpecial() *bool {
	return s.Special
}

func (s *GetTicketResponseBodyList) SetAttachList(v []*string) *GetTicketResponseBodyList {
	s.AttachList = v
	return s
}

func (s *GetTicketResponseBodyList) SetContent(v string) *GetTicketResponseBodyList {
	s.Content = &v
	return s
}

func (s *GetTicketResponseBodyList) SetFromOfficial(v bool) *GetTicketResponseBodyList {
	s.FromOfficial = &v
	return s
}

func (s *GetTicketResponseBodyList) SetGmtCreated(v string) *GetTicketResponseBodyList {
	s.GmtCreated = &v
	return s
}

func (s *GetTicketResponseBodyList) SetNoteId(v string) *GetTicketResponseBodyList {
	s.NoteId = &v
	return s
}

func (s *GetTicketResponseBodyList) SetSpecial(v bool) *GetTicketResponseBodyList {
	s.Special = &v
	return s
}

func (s *GetTicketResponseBodyList) Validate() error {
	return dara.Validate(s)
}
