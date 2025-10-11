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
}

type DeleteAIDBClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *DeleteAIDBClusterRequest) SetDBClusterId(v string) *DeleteAIDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
