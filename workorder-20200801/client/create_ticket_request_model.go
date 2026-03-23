// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonQuestionId(v int64) *CreateTicketRequest
	GetCommonQuestionId() *int64
	SetContent(v string) *CreateTicketRequest
	GetContent() *string
	SetEmail(v string) *CreateTicketRequest
	GetEmail() *string
	SetFileNameList(v string) *CreateTicketRequest
	GetFileNameList() *string
	SetLanguage(v string) *CreateTicketRequest
	GetLanguage() *string
	SetNotifyTimeRange(v string) *CreateTicketRequest
	GetNotifyTimeRange() *string
	SetPhone(v string) *CreateTicketRequest
	GetPhone() *string
	SetProductCode(v string) *CreateTicketRequest
	GetProductCode() *string
	SetSecretContent(v string) *CreateTicketRequest
	GetSecretContent() *string
	SetTitle(v string) *CreateTicketRequest
	GetTitle() *string
}

type CreateTicketRequest struct {
	CommonQuestionId *int64 `json:"CommonQuestionId,omitempty" xml:"CommonQuestionId,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FileNameList    *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	NotifyTimeRange *string `json:"NotifyTimeRange,omitempty" xml:"NotifyTimeRange,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// This parameter is required.
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SecretContent *string `json:"SecretContent,omitempty" xml:"SecretContent,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetCommonQuestionId() *int64 {
	return s.CommonQuestionId
}

func (s *CreateTicketRequest) GetContent() *string {
	return s.Content
}

func (s *CreateTicketRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateTicketRequest) GetFileNameList() *string {
	return s.FileNameList
}

func (s *CreateTicketRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateTicketRequest) GetNotifyTimeRange() *string {
	return s.NotifyTimeRange
}

func (s *CreateTicketRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateTicketRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateTicketRequest) GetSecretContent() *string {
	return s.SecretContent
}

func (s *CreateTicketRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketRequest) SetCommonQuestionId(v int64) *CreateTicketRequest {
	s.CommonQuestionId = &v
	return s
}

func (s *CreateTicketRequest) SetContent(v string) *CreateTicketRequest {
	s.Content = &v
	return s
}

func (s *CreateTicketRequest) SetEmail(v string) *CreateTicketRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketRequest) SetFileNameList(v string) *CreateTicketRequest {
	s.FileNameList = &v
	return s
}

func (s *CreateTicketRequest) SetLanguage(v string) *CreateTicketRequest {
	s.Language = &v
	return s
}

func (s *CreateTicketRequest) SetNotifyTimeRange(v string) *CreateTicketRequest {
	s.NotifyTimeRange = &v
	return s
}

func (s *CreateTicketRequest) SetPhone(v string) *CreateTicketRequest {
	s.Phone = &v
	return s
}

func (s *CreateTicketRequest) SetProductCode(v string) *CreateTicketRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateTicketRequest) SetSecretContent(v string) *CreateTicketRequest {
	s.SecretContent = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	return dara.Validate(s)
}
