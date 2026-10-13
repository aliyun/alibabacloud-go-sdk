// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *CreateTicketShrinkRequest
	GetCategoryId() *string
	SetCreatorId(v string) *CreateTicketShrinkRequest
	GetCreatorId() *string
	SetDescription(v string) *CreateTicketShrinkRequest
	GetDescription() *string
	SetEmail(v string) *CreateTicketShrinkRequest
	GetEmail() *string
	SetFileNameListShrink(v string) *CreateTicketShrinkRequest
	GetFileNameListShrink() *string
	SetSecretInfoShrink(v string) *CreateTicketShrinkRequest
	GetSecretInfoShrink() *string
	SetSeverity(v int32) *CreateTicketShrinkRequest
	GetSeverity() *int32
	SetTitle(v string) *CreateTicketShrinkRequest
	GetTitle() *string
}

type CreateTicketShrinkRequest struct {
	// The ID of the issue category. Obtain this ID from the ListCategories operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53812
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The Alibaba Cloud UID of the account that submits the ticket. This parameter is required for Marketplace reseller scenarios.
	//
	// example:
	//
	// 1289427240739141
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the issue in the ticket. Only plain text is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS backup failed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The contact email address for the ticket.
	//
	// example:
	//
	// sdahkjdshga@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// A list of attachment names. Use the ObjectKey value returned by the GetAttachmentUploadUrl operation.
	FileNameListShrink *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	// Sensitive information.
	SecretInfoShrink *string `json:"SecretInfo,omitempty" xml:"SecretInfo,omitempty"`
	// An enumeration value. A value of 1 indicates a normal issue. A value of 2 indicates an urgent issue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The title of the ticket.
	//
	// example:
	//
	// ECS backup failed
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *CreateTicketShrinkRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateTicketShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTicketShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateTicketShrinkRequest) GetFileNameListShrink() *string {
	return s.FileNameListShrink
}

func (s *CreateTicketShrinkRequest) GetSecretInfoShrink() *string {
	return s.SecretInfoShrink
}

func (s *CreateTicketShrinkRequest) GetSeverity() *int32 {
	return s.Severity
}

func (s *CreateTicketShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketShrinkRequest) SetCategoryId(v string) *CreateTicketShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetCreatorId(v string) *CreateTicketShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetDescription(v string) *CreateTicketShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetEmail(v string) *CreateTicketShrinkRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetFileNameListShrink(v string) *CreateTicketShrinkRequest {
	s.FileNameListShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetSecretInfoShrink(v string) *CreateTicketShrinkRequest {
	s.SecretInfoShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetSeverity(v int32) *CreateTicketShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetTitle(v string) *CreateTicketShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
