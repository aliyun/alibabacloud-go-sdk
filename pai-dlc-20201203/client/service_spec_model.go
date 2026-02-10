// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultPort(v int32) *ServiceSpec
	GetDefaultPort() *int32
	SetExtraPorts(v []*int32) *ServiceSpec
	GetExtraPorts() []*int32
	SetServiceMode(v string) *ServiceSpec
	GetServiceMode() *string
}

type ServiceSpec struct {
	// example:
	//
	// 8080
	DefaultPort *int32   `json:"DefaultPort,omitempty" xml:"DefaultPort,omitempty"`
	ExtraPorts  []*int32 `json:"ExtraPorts,omitempty" xml:"ExtraPorts,omitempty" type:"Repeated"`
	// example:
	//
	// PerRole
	ServiceMode *string `json:"ServiceMode,omitempty" xml:"ServiceMode,omitempty"`
}

func (s ServiceSpec) String() string {
	return dara.Prettify(s)
}

func (s ServiceSpec) GoString() string {
	return s.String()
}

func (s *ServiceSpec) GetDefaultPort() *int32 {
	return s.DefaultPort
}

func (s *ServiceSpec) GetExtraPorts() []*int32 {
	return s.ExtraPorts
}

func (s *ServiceSpec) GetServiceMode() *string {
	return s.ServiceMode
}

func (s *ServiceSpec) SetDefaultPort(v int32) *ServiceSpec {
	s.DefaultPort = &v
	return s
}

func (s *ServiceSpec) SetExtraPorts(v []*int32) *ServiceSpec {
	s.ExtraPorts = v
	return s
}

func (s *ServiceSpec) SetServiceMode(v string) *ServiceSpec {
	s.ServiceMode = &v
	return s
}

func (s *ServiceSpec) Validate() error {
	return dara.Validate(s)
}
