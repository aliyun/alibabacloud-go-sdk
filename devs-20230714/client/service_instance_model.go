// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceInstance interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *ServiceConfig) *ServiceInstance
	GetConfig() *ServiceConfig
	SetLatestDeployment(v *ServiceInstanceLatestDeployment) *ServiceInstance
	GetLatestDeployment() *ServiceInstanceLatestDeployment
	SetOutputs(v map[string]interface{}) *ServiceInstance
	GetOutputs() map[string]interface{}
	SetVariables(v map[string]*Variable) *ServiceInstance
	GetVariables() map[string]*Variable
}

type ServiceInstance struct {
	Config           *ServiceConfig                   `json:"config,omitempty" xml:"config,omitempty"`
	LatestDeployment *ServiceInstanceLatestDeployment `json:"latestDeployment,omitempty" xml:"latestDeployment,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Outputs   map[string]interface{} `json:"outputs,omitempty" xml:"outputs,omitempty"`
	Variables map[string]*Variable   `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ServiceInstance) String() string {
	return dara.Prettify(s)
}

func (s ServiceInstance) GoString() string {
	return s.String()
}

func (s *ServiceInstance) GetConfig() *ServiceConfig {
	return s.Config
}

func (s *ServiceInstance) GetLatestDeployment() *ServiceInstanceLatestDeployment {
	return s.LatestDeployment
}

func (s *ServiceInstance) GetOutputs() map[string]interface{} {
	return s.Outputs
}

func (s *ServiceInstance) GetVariables() map[string]*Variable {
	return s.Variables
}

func (s *ServiceInstance) SetConfig(v *ServiceConfig) *ServiceInstance {
	s.Config = v
	return s
}

func (s *ServiceInstance) SetLatestDeployment(v *ServiceInstanceLatestDeployment) *ServiceInstance {
	s.LatestDeployment = v
	return s
}

func (s *ServiceInstance) SetOutputs(v map[string]interface{}) *ServiceInstance {
	s.Outputs = v
	return s
}

func (s *ServiceInstance) SetVariables(v map[string]*Variable) *ServiceInstance {
	s.Variables = v
	return s
}

func (s *ServiceInstance) Validate() error {
	return dara.Validate(s)
}

type ServiceInstanceLatestDeployment struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Running
	Phase     *string `json:"phase,omitempty" xml:"phase,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ServiceInstanceLatestDeployment) String() string {
	return dara.Prettify(s)
}

func (s ServiceInstanceLatestDeployment) GoString() string {
	return s.String()
}

func (s *ServiceInstanceLatestDeployment) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ServiceInstanceLatestDeployment) GetName() *string {
	return s.Name
}

func (s *ServiceInstanceLatestDeployment) GetPhase() *string {
	return s.Phase
}

func (s *ServiceInstanceLatestDeployment) GetStartTime() *string {
	return s.StartTime
}

func (s *ServiceInstanceLatestDeployment) SetFinishedTime(v string) *ServiceInstanceLatestDeployment {
	s.FinishedTime = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetName(v string) *ServiceInstanceLatestDeployment {
	s.Name = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetPhase(v string) *ServiceInstanceLatestDeployment {
	s.Phase = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) SetStartTime(v string) *ServiceInstanceLatestDeployment {
	s.StartTime = &v
	return s
}

func (s *ServiceInstanceLatestDeployment) Validate() error {
	return dara.Validate(s)
}
