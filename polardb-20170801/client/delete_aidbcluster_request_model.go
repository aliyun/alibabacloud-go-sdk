// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterRequest
	GetDBClusterId() *string
	SetModelSpace(v string) *DeleteAIDBClusterRequest
	GetModelSpace() *string
}

type DeleteAIDBClusterRequest struct {
	// The AI cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The model operator space.
	//
	// example:
	//
	// pms-xxx
	ModelSpace *string `json:"ModelSpace,omitempty" xml:"ModelSpace,omitempty"`
}

func (s DeleteAIDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterRequest) GetModelSpace() *string {
	return s.ModelSpace
}

func (s *DeleteAIDBClusterRequest) SetDBClusterId(v string) *DeleteAIDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterRequest) SetModelSpace(v string) *DeleteAIDBClusterRequest {
	s.ModelSpace = &v
	return s
}

func (s *DeleteAIDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
