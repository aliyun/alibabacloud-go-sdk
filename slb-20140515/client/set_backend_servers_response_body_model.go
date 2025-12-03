// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *SetBackendServersResponseBodyBackendServers) *SetBackendServersResponseBody
	GetBackendServers() *SetBackendServersResponseBodyBackendServers
	SetLoadBalancerId(v string) *SetBackendServersResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *SetBackendServersResponseBody
	GetRequestId() *string
}

type SetBackendServersResponseBody struct {
	// The backend servers.
	BackendServers *SetBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1qjwo61pqz3a******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBody) GetBackendServers() *SetBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *SetBackendServersResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetBackendServersResponseBody) SetBackendServers(v *SetBackendServersResponseBodyBackendServers) *SetBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *SetBackendServersResponseBody) SetLoadBalancerId(v string) *SetBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersResponseBody) SetRequestId(v string) *SetBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetBackendServersResponseBodyBackendServers struct {
	BackendServer []*SetBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s SetBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBodyBackendServers) GetBackendServer() []*SetBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *SetBackendServersResponseBodyBackendServers) SetBackendServer(v []*SetBackendServersResponseBodyBackendServersBackendServer) *SetBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *SetBackendServersResponseBodyBackendServers) Validate() error {
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

type SetBackendServersResponseBodyBackendServersBackendServer struct {
	// The description of the backend server.
	//
	// example:
	//
	// backend server
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// eni-hhshhs****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server. Valid values:
	//
	// 	- **ecs*	- (default): ECS instance
	//
	// 	- **eni**: ENI
	//
	// 	- **eci**: elastic container instance
	//
	// example:
	//
	// eni
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SetBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetWeight() *string {
	return s.Weight
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetWeight(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
