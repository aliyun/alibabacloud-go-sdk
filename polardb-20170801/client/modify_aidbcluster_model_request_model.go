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
	// The new customer-facing invocation name. If this parameter is not specified, the existing invocation name is retained.
	//
	// example:
	//
	// my-flagship-chat
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// Specifies whether to only preview the change.
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

func (s *ModifyAIDBClusterModelRequest) Validate() error {
	return dara.Validate(s)
}
