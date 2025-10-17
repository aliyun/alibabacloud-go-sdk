// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtraPodSpec interface {
  dara.Model
  String() string
  GoString() string
  SetInitContainers(v []*ContainerSpec) *ExtraPodSpec
  GetInitContainers() []*ContainerSpec 
  SetLifecycle(v *Lifecycle) *ExtraPodSpec
  GetLifecycle() *Lifecycle 
  SetPodAnnotations(v map[string]*string) *ExtraPodSpec
  GetPodAnnotations() map[string]*string 
  SetPodLabels(v map[string]*string) *ExtraPodSpec
  GetPodLabels() map[string]*string 
  SetSharedVolumeMountPaths(v []*string) *ExtraPodSpec
  GetSharedVolumeMountPaths() []*string 
  SetSideCarContainers(v []*ContainerSpec) *ExtraPodSpec
  GetSideCarContainers() []*ContainerSpec 
}

type ExtraPodSpec struct {
  InitContainers []*ContainerSpec `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
  Lifecycle *Lifecycle `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
  // Deprecated
  PodAnnotations map[string]*string `json:"PodAnnotations,omitempty" xml:"PodAnnotations,omitempty"`
  // Deprecated
  PodLabels map[string]*string `json:"PodLabels,omitempty" xml:"PodLabels,omitempty"`
  SharedVolumeMountPaths []*string `json:"SharedVolumeMountPaths,omitempty" xml:"SharedVolumeMountPaths,omitempty" type:"Repeated"`
  SideCarContainers []*ContainerSpec `json:"SideCarContainers,omitempty" xml:"SideCarContainers,omitempty" type:"Repeated"`
}

func (s ExtraPodSpec) String() string {
  return dara.Prettify(s)
}

func (s ExtraPodSpec) GoString() string {
  return s.String()
}

func (s *ExtraPodSpec) GetInitContainers() []*ContainerSpec  {
  return s.InitContainers
}

func (s *ExtraPodSpec) GetLifecycle() *Lifecycle  {
  return s.Lifecycle
}

func (s *ExtraPodSpec) GetPodAnnotations() map[string]*string  {
  return s.PodAnnotations
}

func (s *ExtraPodSpec) GetPodLabels() map[string]*string  {
  return s.PodLabels
}

func (s *ExtraPodSpec) GetSharedVolumeMountPaths() []*string  {
  return s.SharedVolumeMountPaths
}

func (s *ExtraPodSpec) GetSideCarContainers() []*ContainerSpec  {
  return s.SideCarContainers
}

func (s *ExtraPodSpec) SetInitContainers(v []*ContainerSpec) *ExtraPodSpec {
  s.InitContainers = v
  return s
}

func (s *ExtraPodSpec) SetLifecycle(v *Lifecycle) *ExtraPodSpec {
  s.Lifecycle = v
  return s
}

func (s *ExtraPodSpec) SetPodAnnotations(v map[string]*string) *ExtraPodSpec {
  s.PodAnnotations = v
  return s
}

func (s *ExtraPodSpec) SetPodLabels(v map[string]*string) *ExtraPodSpec {
  s.PodLabels = v
  return s
}

func (s *ExtraPodSpec) SetSharedVolumeMountPaths(v []*string) *ExtraPodSpec {
  s.SharedVolumeMountPaths = v
  return s
}

func (s *ExtraPodSpec) SetSideCarContainers(v []*ContainerSpec) *ExtraPodSpec {
  s.SideCarContainers = v
  return s
}

func (s *ExtraPodSpec) Validate() error {
  if s.InitContainers != nil {
    for _, item := range s.InitContainers {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Lifecycle != nil {
    if err := s.Lifecycle.Validate(); err != nil {
      return err
    }
  }
  if s.SideCarContainers != nil {
    for _, item := range s.SideCarContainers {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

