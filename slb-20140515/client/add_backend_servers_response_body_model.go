// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *AddBackendServersResponseBodyBackendServers) *AddBackendServersResponseBody
	GetBackendServers() *AddBackendServersResponseBodyBackendServers
	SetLoadBalancerId(v string) *AddBackendServersResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *AddBackendServersResponseBody
	GetRequestId() *string
}

type AddBackendServersResponseBody struct {
	// The list of backend servers.
	BackendServers *AddBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-2ze7o5h52g02kkzz****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34B82C81-F13B-4EEB-99F6-A048C67CC830
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBody) GetBackendServers() *AddBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *AddBackendServersResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBackendServersResponseBody) SetBackendServers(v *AddBackendServersResponseBodyBackendServers) *AddBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *AddBackendServersResponseBody) SetLoadBalancerId(v string) *AddBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersResponseBody) SetRequestId(v string) *AddBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddBackendServersResponseBodyBackendServers struct {
	BackendServer []*AddBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s AddBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBodyBackendServers) GetBackendServer() []*AddBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *AddBackendServersResponseBodyBackendServers) SetBackendServer(v []*AddBackendServersResponseBodyBackendServersBackendServer) *AddBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *AddBackendServersResponseBodyBackendServers) Validate() error {
	if s.BackendServer != nil {
		for _, item := range s.BackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddBackendServersResponseBodyBackendServersBackendServer struct {
	// The description of the backend server.
	//
	// example:
	//
	// backend server
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ECS instance, ENI, or elastic container instance.
	//
	// example:
	//
	// i-2zej4lxhjoq1icu*****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ecs*	- (default): an ECS instance
	//
	// 	- **eni**: an ENI
	//
	// 	- **eci**: an elastic container instance
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// Valid values: **0 to 100**. Default value: **100**.
	//
	// If the value is set to **0**, no requests are forwarded to the backend server.
	//
	// example:
	//
	// 100
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetWeight() *string {
	return s.Weight
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetWeight(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
