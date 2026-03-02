// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentRollbackCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v int64) *DeploymentRollbackCmd
	GetDeploymentId() *int64
}

type DeploymentRollbackCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
}

func (s DeploymentRollbackCmd) String() string {
	return dara.Prettify(s)
}

func (s DeploymentRollbackCmd) GoString() string {
	return s.String()
}

func (s *DeploymentRollbackCmd) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *DeploymentRollbackCmd) SetDeploymentId(v int64) *DeploymentRollbackCmd {
	s.DeploymentId = &v
	return s
}

func (s *DeploymentRollbackCmd) Validate() error {
	return dara.Validate(s)
}
