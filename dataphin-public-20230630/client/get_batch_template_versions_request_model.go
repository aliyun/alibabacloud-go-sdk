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
	SetProjectId(v int64) *GetBatchTemplateVersionsRequest
	GetProjectId() *int64
	SetTemplateId(v int64) *GetBatchTemplateVersionsRequest
	GetTemplateId() *int64
}

type GetBatchTemplateVersionsRequest struct {
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
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
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
