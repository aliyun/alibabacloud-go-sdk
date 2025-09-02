// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *DeleteDataServiceApiAuthorityRequest
	GetApiId() *int64
	SetAuthorizedProjectId(v int64) *DeleteDataServiceApiAuthorityRequest
	GetAuthorizedProjectId() *int64
	SetProjectId(v int64) *DeleteDataServiceApiAuthorityRequest
	GetProjectId() *int64
	SetTenantId(v int64) *DeleteDataServiceApiAuthorityRequest
	GetTenantId() *int64
}

type DeleteDataServiceApiAuthorityRequest struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the workspace from which you want to revoke the access permissions on the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	AuthorizedProjectId *int64 `json:"AuthorizedProjectId,omitempty" xml:"AuthorizedProjectId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10003
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. This parameter is deprecated.
	//
	// example:
	//
	// 10004
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteDataServiceApiAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiAuthorityRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiAuthorityRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *DeleteDataServiceApiAuthorityRequest) GetAuthorizedProjectId() *int64 {
	return s.AuthorizedProjectId
}

func (s *DeleteDataServiceApiAuthorityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDataServiceApiAuthorityRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DeleteDataServiceApiAuthorityRequest) SetApiId(v int64) *DeleteDataServiceApiAuthorityRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityRequest) SetAuthorizedProjectId(v int64) *DeleteDataServiceApiAuthorityRequest {
	s.AuthorizedProjectId = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityRequest) SetProjectId(v int64) *DeleteDataServiceApiAuthorityRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityRequest) SetTenantId(v int64) *DeleteDataServiceApiAuthorityRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
