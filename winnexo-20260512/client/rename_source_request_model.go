// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewName(v string) *RenameSourceRequest
	GetNewName() *string
	SetSourceId(v string) *RenameSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *RenameSourceRequest
	GetTenantId() *string
}

type RenameSourceRequest struct {
	// The new name of the data source.
	//
	// example:
	//
	// string_value
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// The data source ID, which is unique within the tenant.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. You can pass this parameter explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RenameSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameSourceRequest) GoString() string {
	return s.String()
}

func (s *RenameSourceRequest) GetNewName() *string {
	return s.NewName
}

func (s *RenameSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RenameSourceRequest) SetNewName(v string) *RenameSourceRequest {
	s.NewName = &v
	return s
}

func (s *RenameSourceRequest) SetSourceId(v string) *RenameSourceRequest {
	s.SourceId = &v
	return s
}

func (s *RenameSourceRequest) SetTenantId(v string) *RenameSourceRequest {
	s.TenantId = &v
	return s
}

func (s *RenameSourceRequest) Validate() error {
	return dara.Validate(s)
}
