// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetPhysicalInstanceLogRequest
	GetEnv() *string
	SetInstanceId(v string) *GetPhysicalInstanceLogRequest
	GetInstanceId() *string
	SetOpTenantId(v int64) *GetPhysicalInstanceLogRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetPhysicalInstanceLogRequest
	GetProjectId() *int64
}

type GetPhysicalInstanceLogRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_5929472_20210411_9577721
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
	// 123131
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPhysicalInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogRequest) GetEnv() *string {
	return s.Env
}

func (s *GetPhysicalInstanceLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPhysicalInstanceLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPhysicalInstanceLogRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPhysicalInstanceLogRequest) SetEnv(v string) *GetPhysicalInstanceLogRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetInstanceId(v string) *GetPhysicalInstanceLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetOpTenantId(v int64) *GetPhysicalInstanceLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetProjectId(v int64) *GetPhysicalInstanceLogRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}
