// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *CreateTicketRequest
	GetCategoryId() *string
	SetCreatorId(v string) *CreateTicketRequest
	GetCreatorId() *string
	SetDescription(v string) *CreateTicketRequest
	GetDescription() *string
	SetEmail(v string) *CreateTicketRequest
	GetEmail() *string
	SetFileNameList(v []*string) *CreateTicketRequest
	GetFileNameList() []*string
	SetSecretInfo(v *CreateTicketRequestSecretInfo) *CreateTicketRequest
	GetSecretInfo() *CreateTicketRequestSecretInfo
	SetSeverity(v int32) *CreateTicketRequest
	GetSeverity() *int32
	SetTitle(v string) *CreateTicketRequest
	GetTitle() *string
}

type CreateTicketRequest struct {
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
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
	// Sensitive information
	SecretInfo *CreateTicketRequestSecretInfo `json:"SecretInfo,omitempty" xml:"SecretInfo,omitempty" type:"Struct"`
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

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *CreateTicketRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateTicketRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTicketRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateTicketRequest) GetFileNameList() []*string {
	return s.FileNameList
}

func (s *CreateTicketRequest) GetSecretInfo() *CreateTicketRequestSecretInfo {
	return s.SecretInfo
}

func (s *CreateTicketRequest) GetSeverity() *int32 {
	return s.Severity
}

func (s *CreateTicketRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketRequest) SetCategoryId(v string) *CreateTicketRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorId(v string) *CreateTicketRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketRequest) SetDescription(v string) *CreateTicketRequest {
	s.Description = &v
	return s
}

func (s *CreateTicketRequest) SetEmail(v string) *CreateTicketRequest {
	s.Email = &v
	return s
}

func (s *CreateTicketRequest) SetFileNameList(v []*string) *CreateTicketRequest {
	s.FileNameList = v
	return s
}

func (s *CreateTicketRequest) SetSecretInfo(v *CreateTicketRequestSecretInfo) *CreateTicketRequest {
	s.SecretInfo = v
	return s
}

func (s *CreateTicketRequest) SetSeverity(v int32) *CreateTicketRequest {
	s.Severity = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	if s.SecretInfo != nil {
		if err := s.SecretInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTicketRequestSecretInfo struct {
	// Sensitive information-text content
	//
	// example:
	//
	// ID:330102xxxxxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Sensitive Information-Attachment Name List
	FileNameList []*string `json:"FileNameList,omitempty" xml:"FileNameList,omitempty" type:"Repeated"`
}

func (s CreateTicketRequestSecretInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequestSecretInfo) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSecretInfo) GetContent() *string {
	return s.Content
}

func (s *CreateTicketRequestSecretInfo) GetFileNameList() []*string {
	return s.FileNameList
}

func (s *CreateTicketRequestSecretInfo) SetContent(v string) *CreateTicketRequestSecretInfo {
	s.Content = &v
	return s
}

func (s *CreateTicketRequestSecretInfo) SetFileNameList(v []*string) *CreateTicketRequestSecretInfo {
	s.FileNameList = v
	return s
}

func (s *CreateTicketRequestSecretInfo) Validate() error {
	return dara.Validate(s)
}
