// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetResourceRequest
	GetName() *string
	SetOpTenantId(v int64) *GetResourceRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetResourceRequest
	GetProjectId() *int64
}

type GetResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// udf_sleep.jar
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) GetName() *string {
	return s.Name
}

func (s *GetResourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceRequest) SetName(v string) *GetResourceRequest {
	s.Name = &v
	return s
}

func (s *GetResourceRequest) SetOpTenantId(v int64) *GetResourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetResourceRequest) SetProjectId(v int64) *GetResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *GetResourceRequest) Validate() error {
	return dara.Validate(s)
}
