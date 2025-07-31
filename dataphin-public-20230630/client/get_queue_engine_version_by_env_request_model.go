// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueEngineVersionByEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetQueueEngineVersionByEnvRequest
	GetClusterId() *string
	SetEnv(v string) *GetQueueEngineVersionByEnvRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetQueueEngineVersionByEnvRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetQueueEngineVersionByEnvRequest
	GetProjectId() *int64
	SetQueueName(v string) *GetQueueEngineVersionByEnvRequest
	GetQueueName() *string
	SetStreamBatchMode(v string) *GetQueueEngineVersionByEnvRequest
	GetStreamBatchMode() *string
}

type GetQueueEngineVersionByEnvRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 61187014-a3ba-4cdd-8609-1f0aa3df4a3d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROD
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
	// default-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BOTH
	StreamBatchMode *string `json:"StreamBatchMode,omitempty" xml:"StreamBatchMode,omitempty"`
}

func (s GetQueueEngineVersionByEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueEngineVersionByEnvRequest) GoString() string {
	return s.String()
}

func (s *GetQueueEngineVersionByEnvRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetQueueEngineVersionByEnvRequest) GetEnv() *string {
	return s.Env
}

func (s *GetQueueEngineVersionByEnvRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQueueEngineVersionByEnvRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQueueEngineVersionByEnvRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueEngineVersionByEnvRequest) GetStreamBatchMode() *string {
	return s.StreamBatchMode
}

func (s *GetQueueEngineVersionByEnvRequest) SetClusterId(v string) *GetQueueEngineVersionByEnvRequest {
	s.ClusterId = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) SetEnv(v string) *GetQueueEngineVersionByEnvRequest {
	s.Env = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) SetOpTenantId(v int64) *GetQueueEngineVersionByEnvRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) SetProjectId(v int64) *GetQueueEngineVersionByEnvRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) SetQueueName(v string) *GetQueueEngineVersionByEnvRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) SetStreamBatchMode(v string) *GetQueueEngineVersionByEnvRequest {
	s.StreamBatchMode = &v
	return s
}

func (s *GetQueueEngineVersionByEnvRequest) Validate() error {
	return dara.Validate(s)
}
