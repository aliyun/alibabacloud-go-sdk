// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetProjectRequest
	GetEnv() *string
	SetId(v int64) *GetProjectRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetProjectRequest
	GetOpTenantId() *int64
}

type GetProjectRequest struct {
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetEnv() *string {
	return s.Env
}

func (s *GetProjectRequest) GetId() *int64 {
	return s.Id
}

func (s *GetProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetProjectRequest) SetEnv(v string) *GetProjectRequest {
	s.Env = &v
	return s
}

func (s *GetProjectRequest) SetId(v int64) *GetProjectRequest {
	s.Id = &v
	return s
}

func (s *GetProjectRequest) SetOpTenantId(v int64) *GetProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
