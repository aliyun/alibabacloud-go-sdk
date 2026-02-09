// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackend interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v string) *Backend
	GetScene() *string
	SetServices(v []*BackendServices) *Backend
	GetServices() []*BackendServices
}

type Backend struct {
	// The backend service scenario. Valid values: Single, MultiServiceByRatio, MultiServiceByTag, Mock, and Redirect.
	//
	// example:
	//
	// Single
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The backend services.
	Services []*BackendServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s Backend) String() string {
	return dara.Prettify(s)
}

func (s Backend) GoString() string {
	return s.String()
}

func (s *Backend) GetScene() *string {
	return s.Scene
}

func (s *Backend) GetServices() []*BackendServices {
	return s.Services
}

func (s *Backend) SetScene(v string) *Backend {
	s.Scene = &v
	return s
}

func (s *Backend) SetServices(v []*BackendServices) *Backend {
	s.Services = v
	return s
}

func (s *Backend) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BackendServices struct {
	// The service name.
	//
	// example:
	//
	// item-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service port. If you want to use a dynamic port, do not pass this parameter.
	//
	// example:
	//
	// port
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol. Valid values: HTTP, TCP, and DUBBO.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-cq2bmmdlhtgj***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The percentage value of traffic.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BackendServices) String() string {
	return dara.Prettify(s)
}

func (s BackendServices) GoString() string {
	return s.String()
}

func (s *BackendServices) GetName() *string {
	return s.Name
}

func (s *BackendServices) GetPort() *int32 {
	return s.Port
}

func (s *BackendServices) GetProtocol() *string {
	return s.Protocol
}

func (s *BackendServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *BackendServices) GetVersion() *string {
	return s.Version
}

func (s *BackendServices) GetWeight() *int32 {
	return s.Weight
}

func (s *BackendServices) SetName(v string) *BackendServices {
	s.Name = &v
	return s
}

func (s *BackendServices) SetPort(v int32) *BackendServices {
	s.Port = &v
	return s
}

func (s *BackendServices) SetProtocol(v string) *BackendServices {
	s.Protocol = &v
	return s
}

func (s *BackendServices) SetServiceId(v string) *BackendServices {
	s.ServiceId = &v
	return s
}

func (s *BackendServices) SetVersion(v string) *BackendServices {
	s.Version = &v
	return s
}

func (s *BackendServices) SetWeight(v int32) *BackendServices {
	s.Weight = &v
	return s
}

func (s *BackendServices) Validate() error {
	return dara.Validate(s)
}
