// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAIDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartAIDBClusterRequest
	GetDBClusterId() *string
}

type RestartAIDBClusterRequest struct {
	// The instance ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s RestartAIDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartAIDBClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartAIDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartAIDBClusterRequest) SetDBClusterId(v string) *RestartAIDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartAIDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
