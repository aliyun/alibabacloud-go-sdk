// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfByVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetUdfByVersionRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetUdfByVersionRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetUdfByVersionRequest
	GetProjectId() *int64
	SetVersionId(v int64) *GetUdfByVersionRequest
	GetVersionId() *int64
}

type GetUdfByVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s GetUdfByVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUdfByVersionRequest) GoString() string {
	return s.String()
}

func (s *GetUdfByVersionRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUdfByVersionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUdfByVersionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetUdfByVersionRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetUdfByVersionRequest) SetId(v int64) *GetUdfByVersionRequest {
	s.Id = &v
	return s
}

func (s *GetUdfByVersionRequest) SetOpTenantId(v int64) *GetUdfByVersionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUdfByVersionRequest) SetProjectId(v int64) *GetUdfByVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetUdfByVersionRequest) SetVersionId(v int64) *GetUdfByVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetUdfByVersionRequest) Validate() error {
	return dara.Validate(s)
}
