// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterCustomModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterCustomModelRequest
	GetDBClusterId() *string
	SetModelName(v string) *DeleteAIDBClusterCustomModelRequest
	GetModelName() *string
	SetRegionId(v string) *DeleteAIDBClusterCustomModelRequest
	GetRegionId() *string
}

type DeleteAIDBClusterCustomModelRequest struct {
	// The ID of the PolarDB AI 3.0 logical instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pm-2ze4x2mwo81knj08a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The key of the custom model registration to delete.
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

func (s DeleteAIDBClusterCustomModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterCustomModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterCustomModelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterCustomModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DeleteAIDBClusterCustomModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAIDBClusterCustomModelRequest) SetDBClusterId(v string) *DeleteAIDBClusterCustomModelRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelRequest) SetModelName(v string) *DeleteAIDBClusterCustomModelRequest {
	s.ModelName = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelRequest) SetRegionId(v string) *DeleteAIDBClusterCustomModelRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAIDBClusterCustomModelRequest) Validate() error {
	return dara.Validate(s)
}
