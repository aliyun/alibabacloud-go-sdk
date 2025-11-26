// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentLayout interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *DeploymentLayout
	GetApplicationName() *string
	SetComponentName(v string) *DeploymentLayout
	GetComponentName() *string
	SetNodeSelector(v *NodeSelector) *DeploymentLayout
	GetNodeSelector() *NodeSelector
}

type DeploymentLayout struct {
	ApplicationName *string       `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ComponentName   *string       `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	NodeSelector    *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
}

func (s DeploymentLayout) String() string {
	return dara.Prettify(s)
}

func (s DeploymentLayout) GoString() string {
	return s.String()
}

func (s *DeploymentLayout) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DeploymentLayout) GetComponentName() *string {
	return s.ComponentName
}

func (s *DeploymentLayout) GetNodeSelector() *NodeSelector {
	return s.NodeSelector
}

func (s *DeploymentLayout) SetApplicationName(v string) *DeploymentLayout {
	s.ApplicationName = &v
	return s
}

func (s *DeploymentLayout) SetComponentName(v string) *DeploymentLayout {
	s.ComponentName = &v
	return s
}

func (s *DeploymentLayout) SetNodeSelector(v *NodeSelector) *DeploymentLayout {
	s.NodeSelector = v
	return s
}

func (s *DeploymentLayout) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
