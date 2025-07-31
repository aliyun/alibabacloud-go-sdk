// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetUdfRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetUdfRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetUdfRequest
	GetProjectId() *int64
}

type GetUdfRequest struct {
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
}

func (s GetUdfRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUdfRequest) GoString() string {
	return s.String()
}

func (s *GetUdfRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUdfRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetUdfRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetUdfRequest) SetId(v int64) *GetUdfRequest {
	s.Id = &v
	return s
}

func (s *GetUdfRequest) SetOpTenantId(v int64) *GetUdfRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUdfRequest) SetProjectId(v int64) *GetUdfRequest {
	s.ProjectId = &v
	return s
}

func (s *GetUdfRequest) Validate() error {
	return dara.Validate(s)
}
