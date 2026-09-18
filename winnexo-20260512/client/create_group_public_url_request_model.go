// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupPublicUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupPublicUrlRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupPublicUrlRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupPublicUrlRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupPublicUrlRequest
	GetName() *string
	SetNotes(v string) *CreateGroupPublicUrlRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupPublicUrlRequest
	GetOperatingObjectName() *string
	SetOriginalUrl(v string) *CreateGroupPublicUrlRequest
	GetOriginalUrl() *string
	SetSourceTags(v string) *CreateGroupPublicUrlRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupPublicUrlRequest
	GetTenantId() *string
}

type CreateGroupPublicUrlRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// Group collaboration document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_tenant_kb_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The name.
	//
	// example:
	//
	// Enterprise Policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content. This parameter is optional. The notes participate in auxiliary analysis.
	//
	// example:
	//
	// Extract applicable scope and key clauses
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (monitored object name). This parameter is optional.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The URL of the web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://mp.weixin.qq.com/s/iHqLKhkJcOyHNCOGejO32A
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The resource labels. This parameter is optional. Specify a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Important","Meeting"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupPublicUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupPublicUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupPublicUrlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupPublicUrlRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupPublicUrlRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupPublicUrlRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupPublicUrlRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupPublicUrlRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupPublicUrlRequest) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreateGroupPublicUrlRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupPublicUrlRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupPublicUrlRequest) SetDescription(v string) *CreateGroupPublicUrlRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetDirectoryId(v string) *CreateGroupPublicUrlRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetGroupId(v string) *CreateGroupPublicUrlRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetName(v string) *CreateGroupPublicUrlRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetNotes(v string) *CreateGroupPublicUrlRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetOperatingObjectName(v string) *CreateGroupPublicUrlRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetOriginalUrl(v string) *CreateGroupPublicUrlRequest {
	s.OriginalUrl = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetSourceTags(v string) *CreateGroupPublicUrlRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) SetTenantId(v string) *CreateGroupPublicUrlRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupPublicUrlRequest) Validate() error {
	return dara.Validate(s)
}
