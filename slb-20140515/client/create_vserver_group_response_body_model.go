// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *CreateVServerGroupResponseBodyBackendServers) *CreateVServerGroupResponseBody
	GetBackendServers() *CreateVServerGroupResponseBodyBackendServers
	SetRequestId(v string) *CreateVServerGroupResponseBody
	GetRequestId() *string
	SetVServerGroupId(v string) *CreateVServerGroupResponseBody
	GetVServerGroupId() *string
}

type CreateVServerGroupResponseBody struct {
	// The list of backend servers.
	BackendServers *CreateVServerGroupResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s CreateVServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBody) GetBackendServers() *CreateVServerGroupResponseBodyBackendServers {
	return s.BackendServers
}

func (s *CreateVServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVServerGroupResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateVServerGroupResponseBody) SetBackendServers(v *CreateVServerGroupResponseBodyBackendServers) *CreateVServerGroupResponseBody {
	s.BackendServers = v
	return s
}

func (s *CreateVServerGroupResponseBody) SetRequestId(v string) *CreateVServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVServerGroupResponseBody) SetVServerGroupId(v string) *CreateVServerGroupResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *CreateVServerGroupResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVServerGroupResponseBodyBackendServers struct {
	BackendServer []*CreateVServerGroupResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s CreateVServerGroupResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBodyBackendServers) GetBackendServer() []*CreateVServerGroupResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *CreateVServerGroupResponseBodyBackendServers) SetBackendServer(v []*CreateVServerGroupResponseBodyBackendServersBackendServer) *CreateVServerGroupResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServers) Validate() error {
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

type CreateVServerGroupResponseBodyBackendServersBackendServer struct {
	// The description of the vServer group.
	//
	// example:
	//
	// backend server
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server.
	//
	// example:
	//
	// 70
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the ECS instance or ENI.
	//
	// example:
	//
	// eni-hhshhs****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server. Valid values:
	//
	// 	- **ecs*	- (default): ECS instance
	//
	// 	- **eni**: elastic network interface (ENI)
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
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVServerGroupResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetDescription(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetPort(v int32) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetServerId(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetType(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetWeight(v int32) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
