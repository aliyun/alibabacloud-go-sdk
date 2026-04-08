// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTrafficControlTaskCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *DeployTrafficControlTaskCodeRequest
	GetEnvironment() *string
	SetInstanceId(v string) *DeployTrafficControlTaskCodeRequest
	GetInstanceId() *string
	SetRetryDeploy(v bool) *DeployTrafficControlTaskCodeRequest
	GetRetryDeploy() *bool
}

type DeployTrafficControlTaskCodeRequest struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-test1
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RetryDeploy *bool   `json:"RetryDeploy,omitempty" xml:"RetryDeploy,omitempty"`
}

func (s DeployTrafficControlTaskCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployTrafficControlTaskCodeRequest) GoString() string {
	return s.String()
}

func (s *DeployTrafficControlTaskCodeRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *DeployTrafficControlTaskCodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeployTrafficControlTaskCodeRequest) GetRetryDeploy() *bool {
	return s.RetryDeploy
}

func (s *DeployTrafficControlTaskCodeRequest) SetEnvironment(v string) *DeployTrafficControlTaskCodeRequest {
	s.Environment = &v
	return s
}

func (s *DeployTrafficControlTaskCodeRequest) SetInstanceId(v string) *DeployTrafficControlTaskCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeployTrafficControlTaskCodeRequest) SetRetryDeploy(v bool) *DeployTrafficControlTaskCodeRequest {
	s.RetryDeploy = &v
	return s
}

func (s *DeployTrafficControlTaskCodeRequest) Validate() error {
	return dara.Validate(s)
}
