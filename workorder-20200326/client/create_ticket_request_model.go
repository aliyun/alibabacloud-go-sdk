// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateTicketRequest
	GetCategory() *string
	SetContent(v string) *CreateTicketRequest
	GetContent() *string
	SetEmail(v string) *CreateTicketRequest
	GetEmail() *string
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
	// This parameter is required.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// support@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// all
	NotifyTimeRange *string `json:"NotifyTimeRange,omitempty" xml:"NotifyTimeRange,omitempty"`
	// example:
	//
	// 13288888888
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
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

func (s *CreateTicketRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateTicketRequest) GetContent() *string {
	return s.Content
}

func (s *CreateTicketRequest) GetEmail() *string {
	return s.Email
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

func (s *CreateTicketRequest) SetCategory(v string) *CreateTicketRequest {
	s.Category = &v
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
