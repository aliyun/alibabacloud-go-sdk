// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnKubeClusterResource interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v string) *OnKubeClusterResource
	GetCpu() *string
	SetMemory(v string) *OnKubeClusterResource
	GetMemory() *string
}

type OnKubeClusterResource struct {
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s OnKubeClusterResource) String() string {
	return dara.Prettify(s)
}

func (s OnKubeClusterResource) GoString() string {
	return s.String()
}

func (s *OnKubeClusterResource) GetCpu() *string {
	return s.Cpu
}

func (s *OnKubeClusterResource) GetMemory() *string {
	return s.Memory
}

func (s *OnKubeClusterResource) SetCpu(v string) *OnKubeClusterResource {
	s.Cpu = &v
	return s
}

func (s *OnKubeClusterResource) SetMemory(v string) *OnKubeClusterResource {
	s.Memory = &v
	return s
}

func (s *OnKubeClusterResource) Validate() error {
	return dara.Validate(s)
}
