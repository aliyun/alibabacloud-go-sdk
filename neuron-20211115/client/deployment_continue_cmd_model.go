// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentContinueCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v int64) *DeploymentContinueCmd
	GetDeploymentId() *int64
}

type DeploymentContinueCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeploymentId *int64 `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
}

func (s DeploymentContinueCmd) String() string {
	return dara.Prettify(s)
}

func (s DeploymentContinueCmd) GoString() string {
	return s.String()
}

func (s *DeploymentContinueCmd) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *DeploymentContinueCmd) SetDeploymentId(v int64) *DeploymentContinueCmd {
	s.DeploymentId = &v
	return s
}

func (s *DeploymentContinueCmd) Validate() error {
	return dara.Validate(s)
}
