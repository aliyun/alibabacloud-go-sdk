// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerSpec interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *ContainerSpec
	GetArgs() []*string
	SetCommand(v []*string) *ContainerSpec
	GetCommand() []*string
	SetEnv(v []*EnvVar) *ContainerSpec
	GetEnv() []*EnvVar
	SetImage(v string) *ContainerSpec
	GetImage() *string
	SetName(v string) *ContainerSpec
	GetName() *string
	SetResources(v *ResourceRequirements) *ContainerSpec
	GetResources() *ResourceRequirements
	SetWorkingDir(v string) *ContainerSpec
	GetWorkingDir() *string
}

type ContainerSpec struct {
	// The command parameters.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The user command.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables.
	Env []*EnvVar `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
	// The endpoint of the container image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-dlc/curl:v1.0.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// data-init
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The container resources.
	Resources *ResourceRequirements `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The working directory in the container.
	//
	// example:
	//
	// /root
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ContainerSpec) String() string {
	return dara.Prettify(s)
}

func (s ContainerSpec) GoString() string {
	return s.String()
}

func (s *ContainerSpec) GetArgs() []*string {
	return s.Args
}

func (s *ContainerSpec) GetCommand() []*string {
	return s.Command
}

func (s *ContainerSpec) GetEnv() []*EnvVar {
	return s.Env
}

func (s *ContainerSpec) GetImage() *string {
	return s.Image
}

func (s *ContainerSpec) GetName() *string {
	return s.Name
}

func (s *ContainerSpec) GetResources() *ResourceRequirements {
	return s.Resources
}

func (s *ContainerSpec) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *ContainerSpec) SetArgs(v []*string) *ContainerSpec {
	s.Args = v
	return s
}

func (s *ContainerSpec) SetCommand(v []*string) *ContainerSpec {
	s.Command = v
	return s
}

func (s *ContainerSpec) SetEnv(v []*EnvVar) *ContainerSpec {
	s.Env = v
	return s
}

func (s *ContainerSpec) SetImage(v string) *ContainerSpec {
	s.Image = &v
	return s
}

func (s *ContainerSpec) SetName(v string) *ContainerSpec {
	s.Name = &v
	return s
}

func (s *ContainerSpec) SetResources(v *ResourceRequirements) *ContainerSpec {
	s.Resources = v
	return s
}

func (s *ContainerSpec) SetWorkingDir(v string) *ContainerSpec {
	s.WorkingDir = &v
	return s
}

func (s *ContainerSpec) Validate() error {
	if s.Env != nil {
		for _, item := range s.Env {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}
