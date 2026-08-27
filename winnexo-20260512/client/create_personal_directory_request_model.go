// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreatePersonalDirectoryRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest
	GetOperatingObjectName() *string
	SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreatePersonalDirectoryRequest
	GetTenantId() *string
}

type CreatePersonalDirectoryRequest struct {
	// The workspace description.
	//
	// example:
	//
	// hangzhou-release-version-3-eventbridge-numeric-queue-fix-20260529
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// sandbox-conversation-webpage-github-default-p32JG2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital human (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 235454102432001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDirectoryRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreatePersonalDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDirectoryRequest) SetDescription(v string) *CreatePersonalDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetName(v string) *CreatePersonalDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetTenantId(v string) *CreatePersonalDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
