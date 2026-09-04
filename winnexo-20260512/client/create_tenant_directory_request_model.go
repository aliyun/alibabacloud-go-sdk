// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTenantDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreateTenantDirectoryRequest
	GetName() *string
	SetParentId(v int64) *CreateTenantDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *CreateTenantDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *CreateTenantDirectoryRequest
	GetTenantId() *string
}

type CreateTenantDirectoryRequest struct {
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the parent node.
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The ID of the tenant for which the operation takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateTenantDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTenantDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *CreateTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTenantDirectoryRequest) SetDescription(v string) *CreateTenantDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetName(v string) *CreateTenantDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetParentId(v int64) *CreateTenantDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetPath(v string) *CreateTenantDirectoryRequest {
	s.Path = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetTenantId(v string) *CreateTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
