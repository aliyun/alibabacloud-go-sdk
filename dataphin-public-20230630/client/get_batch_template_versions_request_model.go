// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetBatchTemplateVersionsRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetBatchTemplateVersionsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetBatchTemplateVersionsRequest
	GetOpUserId() *string
	SetProjectId(v int64) *GetBatchTemplateVersionsRequest
	GetProjectId() *int64
	SetTemplateId(v int64) *GetBatchTemplateVersionsRequest
	GetTemplateId() *int64
}

type GetBatchTemplateVersionsRequest struct {
	// The environment. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD: production environment.
	//
	// Default value for dev_prod projects: DEV.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetBatchTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsRequest) GetEnv() *string {
	return s.Env
}

func (s *GetBatchTemplateVersionsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBatchTemplateVersionsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetBatchTemplateVersionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTemplateVersionsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetBatchTemplateVersionsRequest) SetEnv(v string) *GetBatchTemplateVersionsRequest {
	s.Env = &v
	return s
}

func (s *GetBatchTemplateVersionsRequest) SetOpTenantId(v int64) *GetBatchTemplateVersionsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBatchTemplateVersionsRequest) SetOpUserId(v string) *GetBatchTemplateVersionsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetBatchTemplateVersionsRequest) SetProjectId(v int64) *GetBatchTemplateVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTemplateVersionsRequest) SetTemplateId(v int64) *GetBatchTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetBatchTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
