// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalAlidingDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAlidingDocRequest
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreatePersonalAlidingDocRequest
	GetFilePublicUrl() *string
	SetName(v string) *CreatePersonalAlidingDocRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalAlidingDocRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalAlidingDocRequest
	GetTenantId() *string
}

type CreatePersonalAlidingDocRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// controll service user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The publicly accessible URL of the AliDing online document.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The customer group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// issue_research
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1729094555111072
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAlidingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingDocRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAlidingDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingDocRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreatePersonalAlidingDocRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAlidingDocRequest) SetDescription(v string) *CreatePersonalAlidingDocRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetDirectoryId(v string) *CreatePersonalAlidingDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetFilePublicUrl(v string) *CreatePersonalAlidingDocRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetName(v string) *CreatePersonalAlidingDocRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetOperatingObjectName(v string) *CreatePersonalAlidingDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetTenantId(v string) *CreatePersonalAlidingDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) Validate() error {
	return dara.Validate(s)
}
