// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyAIDBClusterModelRequest
	GetDBClusterId() *string
	SetDisplayModelName(v string) *ModifyAIDBClusterModelRequest
	GetDisplayModelName() *string
	SetDryRun(v bool) *ModifyAIDBClusterModelRequest
	GetDryRun() *bool
	SetModelName(v string) *ModifyAIDBClusterModelRequest
	GetModelName() *string
	SetRegionId(v string) *ModifyAIDBClusterModelRequest
	GetRegionId() *string
	SetRestartMode(v string) *ModifyAIDBClusterModelRequest
	GetRestartMode() *string
	SetWorkerBatchSize(v int64) *ModifyAIDBClusterModelRequest
	GetWorkerBatchSize() *int64
}

type ModifyAIDBClusterModelRequest struct {
	// The ID of the PolarDB AI 3.0 logical instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-2ze4x2mwo81knj08a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The new client-facing invocation name. If this parameter is not specified, the existing invocation name is retained.
	//
	// example:
	//
	// my-flagship-chat
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// Specifies whether to only preview the change without actually performing it.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the target model. Select a value from the ModelName values returned by the DescribeAvailableModels operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-32B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The restart mode for workers. Valid values:
	//
	// - inPlace
	//
	// - recreate
	//
	// example:
	//
	// inPlace
	RestartMode *string `json:"RestartMode,omitempty" xml:"RestartMode,omitempty"`
	// The maximum number of workers to restart per batch within a single MSD. Valid values: 1 to 30. This parameter takes effect only when RestartMode is set to inPlace.
	//
	// example:
	//
	// 8
	WorkerBatchSize *int64 `json:"WorkerBatchSize,omitempty" xml:"WorkerBatchSize,omitempty"`
}

func (s ModifyAIDBClusterModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAIDBClusterModelRequest) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *ModifyAIDBClusterModelRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAIDBClusterModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModifyAIDBClusterModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAIDBClusterModelRequest) GetRestartMode() *string {
	return s.RestartMode
}

func (s *ModifyAIDBClusterModelRequest) GetWorkerBatchSize() *int64 {
	return s.WorkerBatchSize
}

func (s *ModifyAIDBClusterModelRequest) SetDBClusterId(v string) *ModifyAIDBClusterModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetDisplayModelName(v string) *ModifyAIDBClusterModelRequest {
	s.DisplayModelName = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetDryRun(v bool) *ModifyAIDBClusterModelRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetModelName(v string) *ModifyAIDBClusterModelRequest {
	s.ModelName = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetRegionId(v string) *ModifyAIDBClusterModelRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetRestartMode(v string) *ModifyAIDBClusterModelRequest {
	s.RestartMode = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) SetWorkerBatchSize(v int64) *ModifyAIDBClusterModelRequest {
	s.WorkerBatchSize = &v
	return s
}

func (s *ModifyAIDBClusterModelRequest) Validate() error {
	return dara.Validate(s)
}
