// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalPublicUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalPublicUrlRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalPublicUrlRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalPublicUrlRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalPublicUrlRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalPublicUrlRequest
	GetOperatingObjectName() *string
	SetOriginalUrl(v string) *CreatePersonalPublicUrlRequest
	GetOriginalUrl() *string
	SetSourceTags(v string) *CreatePersonalPublicUrlRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalPublicUrlRequest
	GetTenantId() *string
}

type CreatePersonalPublicUrlRequest struct {
	// The resource description.
	//
	// example:
	//
	// Project design document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the target personal directory. If not specified, the current user\\"s default personal root directory is used.
	//
	// example:
	//
	// dir_personal_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The resource name. If omitted, the URL is used.
	//
	// example:
	//
	// Project Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis instruction.
	//
	// example:
	//
	// Extract decisions and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the operating object.
	//
	// example:
	//
	// R&D Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The URL of the public HTTP/HTTPS web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The list of resource tag JSON strings.
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalPublicUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalPublicUrlRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalPublicUrlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalPublicUrlRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalPublicUrlRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalPublicUrlRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalPublicUrlRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalPublicUrlRequest) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreatePersonalPublicUrlRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalPublicUrlRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalPublicUrlRequest) SetDescription(v string) *CreatePersonalPublicUrlRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetDirectoryId(v string) *CreatePersonalPublicUrlRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetName(v string) *CreatePersonalPublicUrlRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetNotes(v string) *CreatePersonalPublicUrlRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetOperatingObjectName(v string) *CreatePersonalPublicUrlRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetOriginalUrl(v string) *CreatePersonalPublicUrlRequest {
	s.OriginalUrl = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetSourceTags(v string) *CreatePersonalPublicUrlRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) SetTenantId(v string) *CreatePersonalPublicUrlRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalPublicUrlRequest) Validate() error {
	return dara.Validate(s)
}
