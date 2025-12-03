// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVServerGroupBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *AddVServerGroupBackendServersResponseBodyBackendServers) *AddVServerGroupBackendServersResponseBody
	GetBackendServers() *AddVServerGroupBackendServersResponseBodyBackendServers
	SetRequestId(v string) *AddVServerGroupBackendServersResponseBody
	GetRequestId() *string
	SetVServerGroupId(v string) *AddVServerGroupBackendServersResponseBody
	GetVServerGroupId() *string
}

type AddVServerGroupBackendServersResponseBody struct {
	// The backend servers.
	BackendServers *AddVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s AddVServerGroupBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBody) GetBackendServers() *AddVServerGroupBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *AddVServerGroupBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVServerGroupBackendServersResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AddVServerGroupBackendServersResponseBody) SetBackendServers(v *AddVServerGroupBackendServersResponseBodyBackendServers) *AddVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *AddVServerGroupBackendServersResponseBody) SetRequestId(v string) *AddVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *AddVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*AddVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s AddVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServers) GetBackendServer() []*AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) *AddVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServers) Validate() error {
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

type AddVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
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
	// 	- **eni**: ENI
	//
	// 	- **eci**: elastic container instance
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
