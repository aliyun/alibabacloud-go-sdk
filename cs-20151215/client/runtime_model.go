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
	// Specifies the container runtime. Valid values:
	//
	// - `containerd`: Recommended. Supported in version 1.20 and later.
	//
	// - `Sandboxed-Container.runv`: Sandboxed container that provides higher isolation. Supported in version 1.16 and later.
	//
	// - `docker`: Supported in clusters of version 1.22 and earlier.
	//
	// Default value: `containerd`.
	//
	// example:
	//
	// containerd
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The container runtime version. By default, the latest version is used. For more information about the release notes, see [Container runtime release notes](https://help.aliyun.com/document_detail/207829.html).
	//
	// example:
	//
	// 1.6.20
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
