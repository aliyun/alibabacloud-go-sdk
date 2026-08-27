// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateDirectoryRequest
	GetName() *string
	SetParentId(v int64) *UpdateDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *UpdateDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *UpdateDirectoryRequest
	GetTenantId() *string
}

type UpdateDirectoryRequest struct {
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name.
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
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateDirectoryRequest) SetDescription(v string) *UpdateDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateDirectoryRequest) SetDirectoryId(v string) *UpdateDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetName(v string) *UpdateDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateDirectoryRequest) SetParentId(v int64) *UpdateDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetPath(v string) *UpdateDirectoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateDirectoryRequest) SetTenantId(v string) *UpdateDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
