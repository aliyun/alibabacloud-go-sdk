// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetStreamJobsRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetStreamJobsRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetStreamJobsRequest
	GetProjectId() *int64
}

type GetStreamJobsRequest struct {
	// This parameter is required.
	//
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
	// 7162269257990111
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetStreamJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStreamJobsRequest) GoString() string {
	return s.String()
}

func (s *GetStreamJobsRequest) GetEnv() *string {
	return s.Env
}

func (s *GetStreamJobsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStreamJobsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetStreamJobsRequest) SetEnv(v string) *GetStreamJobsRequest {
	s.Env = &v
	return s
}

func (s *GetStreamJobsRequest) SetOpTenantId(v int64) *GetStreamJobsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStreamJobsRequest) SetProjectId(v int64) *GetStreamJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetStreamJobsRequest) Validate() error {
	return dara.Validate(s)
}
