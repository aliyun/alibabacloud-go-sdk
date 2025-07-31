// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceByVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetResourceByVersionRequest
	GetName() *string
	SetOpTenantId(v int64) *GetResourceByVersionRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetResourceByVersionRequest
	GetProjectId() *int64
	SetVersionId(v int64) *GetResourceByVersionRequest
	GetVersionId() *int64
}

type GetResourceByVersionRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetResourceByVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceByVersionRequest) GoString() string {
	return s.String()
}

func (s *GetResourceByVersionRequest) GetName() *string {
	return s.Name
}

func (s *GetResourceByVersionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetResourceByVersionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceByVersionRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetResourceByVersionRequest) SetName(v string) *GetResourceByVersionRequest {
	s.Name = &v
	return s
}

func (s *GetResourceByVersionRequest) SetOpTenantId(v int64) *GetResourceByVersionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetResourceByVersionRequest) SetProjectId(v int64) *GetResourceByVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetResourceByVersionRequest) SetVersionId(v int64) *GetResourceByVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetResourceByVersionRequest) Validate() error {
	return dara.Validate(s)
}
