// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterCustomModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelName(v string) *DeleteAIDBClusterCustomModelRequest
	GetModelName() *string
	SetRegionId(v string) *DeleteAIDBClusterCustomModelRequest
	GetRegionId() *string
}

type DeleteAIDBClusterCustomModelRequest struct {
	// The custom model registration key to delete.
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

func (s *DeleteAIDBClusterCustomModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DeleteAIDBClusterCustomModelRequest) GetRegionId() *string {
	return s.RegionId
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
