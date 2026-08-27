// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalTextRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalTextRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalTextRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreatePersonalTextRequest
	GetTextContent() *string
}

type CreatePersonalTextRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// PublicApplication
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_mysql_10_34_4_255_6306_password
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
	// 3668
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The message content for text messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreatePersonalTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalTextRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalTextRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreatePersonalTextRequest) SetDescription(v string) *CreatePersonalTextRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTextRequest) SetDirectoryId(v string) *CreatePersonalTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetName(v string) *CreatePersonalTextRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalTextRequest) SetOperatingObjectName(v string) *CreatePersonalTextRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTenantId(v string) *CreatePersonalTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTextContent(v string) *CreatePersonalTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreatePersonalTextRequest) Validate() error {
	return dara.Validate(s)
}
