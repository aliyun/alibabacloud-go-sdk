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
	// The name of a container runtime. The following types of runtime are supported by Container Service for Kubernetes (ACK).
	//
	// 	- `containerd`: supports all Kubernetes versions. We recommend that you set the parameter to this value.
	//
	// 	- `Sandboxed-Container.runv`: Sandboxed container provides enhanced isolation and supports Kubernetes 1.24 and earlier.
	//
	// 	- `docker`: supports Kubernetes 1.22 and earlier.
	//
	// Default value: `containerd`.
	//
	// example:
	//
	// docker
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the container runtime. By default, the latest version is used.
	//
	// For more information about the changes to Sandboxed-Container, see [Sandboxed-Container release notes](https://help.aliyun.com/document_detail/160312.html).
	//
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
