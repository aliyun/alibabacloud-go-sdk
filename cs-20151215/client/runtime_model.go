// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *Runtime
	GetName() *string
	SetVersion(v string) *Runtime
	GetVersion() *string
}

type Runtime struct {
	// example:
	//
	// docker
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 19.03.5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Runtime) String() string {
	return dara.Prettify(s)
}

func (s Runtime) GoString() string {
	return s.String()
}

func (s *Runtime) GetName() *string {
	return s.Name
}

func (s *Runtime) GetVersion() *string {
	return s.Version
}

func (s *Runtime) SetName(v string) *Runtime {
	s.Name = &v
	return s
}

func (s *Runtime) SetVersion(v string) *Runtime {
	s.Version = &v
	return s
}

func (s *Runtime) Validate() error {
	return dara.Validate(s)
}
