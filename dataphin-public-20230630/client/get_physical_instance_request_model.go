// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalInstanceRequest
	GetEnv() *string
	SetInstanceId(v string) *GetPhysicalInstanceRequest
	GetInstanceId() *string
	SetOpTenantId(v int64) *GetPhysicalInstanceRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetPhysicalInstanceRequest
	GetProjectId() *int64
}

type GetPhysicalInstanceRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_23231
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 2323131
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPhysicalInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPhysicalInstanceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPhysicalInstanceRequest) SetEnv(v string) *GetPhysicalInstanceRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetInstanceId(v string) *GetPhysicalInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetOpTenantId(v int64) *GetPhysicalInstanceRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetProjectId(v int64) *GetPhysicalInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPhysicalInstanceRequest) Validate() error {
	return dara.Validate(s)
}
