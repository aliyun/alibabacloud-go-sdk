// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BVolumeMount interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *E2BVolumeMount
  GetName() *string 
  SetPath(v string) *E2BVolumeMount
  GetPath() *string 
}

type E2BVolumeMount struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s E2BVolumeMount) String() string {
  return dara.Prettify(s)
}

func (s E2BVolumeMount) GoString() string {
  return s.String()
}

func (s *E2BVolumeMount) GetName() *string  {
  return s.Name
}

func (s *E2BVolumeMount) GetPath() *string  {
  return s.Path
}

func (s *E2BVolumeMount) SetName(v string) *E2BVolumeMount {
  s.Name = &v
  return s
}

func (s *E2BVolumeMount) SetPath(v string) *E2BVolumeMount {
  s.Path = &v
  return s
}

func (s *E2BVolumeMount) Validate() error {
  return dara.Validate(s)
}

