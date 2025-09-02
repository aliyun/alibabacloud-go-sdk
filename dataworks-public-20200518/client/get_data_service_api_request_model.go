// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *GetDataServiceApiRequest
	GetApiId() *int64
	SetProjectId(v int64) *GetDataServiceApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *GetDataServiceApiRequest
	GetTenantId() *int64
}

type GetDataServiceApiRequest struct {
	// The ID of the DataService Studio API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetDataServiceApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceApiRequest) SetApiId(v int64) *GetDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *GetDataServiceApiRequest) SetProjectId(v int64) *GetDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiRequest) SetTenantId(v int64) *GetDataServiceApiRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
