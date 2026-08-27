// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalAliDingMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAliDingMeetingRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalAliDingMeetingRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalAliDingMeetingRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalAliDingMeetingRequest
	GetOperatingObjectName() *string
	SetShanjiUrl(v string) *CreatePersonalAliDingMeetingRequest
	GetShanjiUrl() *string
	SetTenantId(v string) *CreatePersonalAliDingMeetingRequest
	GetTenantId() *string
}

type CreatePersonalAliDingMeetingRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// controll service user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the image-trained digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-toolset-3dcef7ca-31b9-4d1c-8692-1ef03099cad3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// string_value
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The original Shanji link (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	ShanjiUrl *string `json:"shanjiUrl,omitempty" xml:"shanjiUrl,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 549003315603714
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAliDingMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAliDingMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAliDingMeetingRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalAliDingMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAliDingMeetingRequest) GetShanjiUrl() *string {
	return s.ShanjiUrl
}

func (s *CreatePersonalAliDingMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAliDingMeetingRequest) SetDescription(v string) *CreatePersonalAliDingMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetDirectoryId(v string) *CreatePersonalAliDingMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetName(v string) *CreatePersonalAliDingMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetNotes(v string) *CreatePersonalAliDingMeetingRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalAliDingMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetShanjiUrl(v string) *CreatePersonalAliDingMeetingRequest {
	s.ShanjiUrl = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetTenantId(v string) *CreatePersonalAliDingMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) Validate() error {
	return dara.Validate(s)
}
