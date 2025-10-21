// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentTarget interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeploymentTarget
	GetName() *string
	SetNamespace(v string) *DeploymentTarget
	GetNamespace() *string
	SetQuota(v *ResourceQuota) *DeploymentTarget
	GetQuota() *ResourceQuota
}

type DeploymentTarget struct {
	// example:
	//
	// deployment target
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// namespace
	Namespace *string        `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Quota     *ResourceQuota `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s DeploymentTarget) String() string {
	return dara.Prettify(s)
}

func (s DeploymentTarget) GoString() string {
	return s.String()
}

func (s *DeploymentTarget) GetName() *string {
	return s.Name
}

func (s *DeploymentTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *DeploymentTarget) GetQuota() *ResourceQuota {
	return s.Quota
}

func (s *DeploymentTarget) SetName(v string) *DeploymentTarget {
	s.Name = &v
	return s
}

func (s *DeploymentTarget) SetNamespace(v string) *DeploymentTarget {
	s.Namespace = &v
	return s
}

func (s *DeploymentTarget) SetQuota(v *ResourceQuota) *DeploymentTarget {
	s.Quota = v
	return s
}

func (s *DeploymentTarget) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}
