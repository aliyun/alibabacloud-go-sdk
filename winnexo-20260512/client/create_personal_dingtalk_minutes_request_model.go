// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalDingtalkMinutesRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMinutesRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalDingtalkMinutesRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalDingtalkMinutesRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkMinutesRequest
	GetOperatingObjectName() *string
	SetShanjiUrl(v string) *CreatePersonalDingtalkMinutesRequest
	GetShanjiUrl() *string
	SetTenantId(v string) *CreatePersonalDingtalkMinutesRequest
	GetTenantId() *string
}

type CreatePersonalDingtalkMinutesRequest struct {
	// The description of the pipeline.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the worksheet.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
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
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDingtalkMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMinutesRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMinutesRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkMinutesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMinutesRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMinutesRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkMinutesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkMinutesRequest) GetShanjiUrl() *string {
	return s.ShanjiUrl
}

func (s *CreatePersonalDingtalkMinutesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkMinutesRequest) SetDescription(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetDirectoryId(v string) *CreatePersonalDingtalkMinutesRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetName(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetNotes(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkMinutesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetShanjiUrl(v string) *CreatePersonalDingtalkMinutesRequest {
	s.ShanjiUrl = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetTenantId(v string) *CreatePersonalDingtalkMinutesRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) Validate() error {
	return dara.Validate(s)
}
