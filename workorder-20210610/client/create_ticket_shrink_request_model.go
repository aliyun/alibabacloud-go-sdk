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
	// The ID of the problem category. You can obtain the returned value from the ListCategories operation by using the CategoryId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7161
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Submit the Alibaba Cloud UID of the account, which is required for the MPK virtual market scene.
	//
	// example:
	//
	// 1289427240739141
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the ticket. Only pure text format is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Why ECS backup failed?
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// sdahkjdshga@qq.com
	//
	// example:
	//
	// 163@163.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The list of attachment names, GetAttachmentUploadUrl the ObjectKey field returned by the interface.
	FileNameListShrink *string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty"`
	// Sensitive information
	SecretInfoShrink *string `json:"SecretInfo,omitempty" xml:"SecretInfo,omitempty"`
	// Enumeration value, 1 for general problem, 2 for urgent problem
	//
	// Enumeration values:
	//
	// 	- 1: regular.
	//
	// 	- 2: emergency.
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
	// Why ECS backup failed?
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
