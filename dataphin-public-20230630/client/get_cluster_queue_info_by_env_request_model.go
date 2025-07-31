// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterQueueInfoByEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetClusterQueueInfoByEnvRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetClusterQueueInfoByEnvRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetClusterQueueInfoByEnvRequest
	GetProjectId() *int64
	SetStreamBatchMode(v string) *GetClusterQueueInfoByEnvRequest
	GetStreamBatchMode() *string
}

type GetClusterQueueInfoByEnvRequest struct {
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
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BOTH
	StreamBatchMode *string `json:"StreamBatchMode,omitempty" xml:"StreamBatchMode,omitempty"`
}

func (s GetClusterQueueInfoByEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterQueueInfoByEnvRequest) GoString() string {
	return s.String()
}

func (s *GetClusterQueueInfoByEnvRequest) GetEnv() *string {
	return s.Env
}

func (s *GetClusterQueueInfoByEnvRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetClusterQueueInfoByEnvRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetClusterQueueInfoByEnvRequest) GetStreamBatchMode() *string {
	return s.StreamBatchMode
}

func (s *GetClusterQueueInfoByEnvRequest) SetEnv(v string) *GetClusterQueueInfoByEnvRequest {
	s.Env = &v
	return s
}

func (s *GetClusterQueueInfoByEnvRequest) SetOpTenantId(v int64) *GetClusterQueueInfoByEnvRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetClusterQueueInfoByEnvRequest) SetProjectId(v int64) *GetClusterQueueInfoByEnvRequest {
	s.ProjectId = &v
	return s
}

func (s *GetClusterQueueInfoByEnvRequest) SetStreamBatchMode(v string) *GetClusterQueueInfoByEnvRequest {
	s.StreamBatchMode = &v
	return s
}

func (s *GetClusterQueueInfoByEnvRequest) Validate() error {
	return dara.Validate(s)
}
