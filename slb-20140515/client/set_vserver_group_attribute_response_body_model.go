// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVServerGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *SetVServerGroupAttributeResponseBodyBackendServers) *SetVServerGroupAttributeResponseBody
	GetBackendServers() *SetVServerGroupAttributeResponseBodyBackendServers
	SetRequestId(v string) *SetVServerGroupAttributeResponseBody
	GetRequestId() *string
	SetVServerGroupId(v string) *SetVServerGroupAttributeResponseBody
	GetVServerGroupId() *string
	SetVServerGroupName(v string) *SetVServerGroupAttributeResponseBody
	GetVServerGroupName() *string
}

type SetVServerGroupAttributeResponseBody struct {
	// The backend servers.
	BackendServers *SetVServerGroupAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The name of the vServer group.
	//
	// example:
	//
	// Group1
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s SetVServerGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBody) GetBackendServers() *SetVServerGroupAttributeResponseBodyBackendServers {
	return s.BackendServers
}

func (s *SetVServerGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVServerGroupAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetVServerGroupAttributeResponseBody) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *SetVServerGroupAttributeResponseBody) SetBackendServers(v *SetVServerGroupAttributeResponseBodyBackendServers) *SetVServerGroupAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetRequestId(v string) *SetVServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetVServerGroupId(v string) *SetVServerGroupAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetVServerGroupName(v string) *SetVServerGroupAttributeResponseBody {
	s.VServerGroupName = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetVServerGroupAttributeResponseBodyBackendServers struct {
	BackendServer []*SetVServerGroupAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s SetVServerGroupAttributeResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBodyBackendServers) GetBackendServer() []*SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *SetVServerGroupAttributeResponseBodyBackendServers) SetBackendServer(v []*SetVServerGroupAttributeResponseBodyBackendServersBackendServer) *SetVServerGroupAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServers) Validate() error {
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

type SetVServerGroupAttributeResponseBodyBackendServersBackendServer struct {
	// The description of the server group.
	//
	// example:
	//
	// Backend server group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server.
	//
	// example:
	//
	// 70
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// i-bp1ek6yd7jvkx****
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

func (s SetVServerGroupAttributeResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetPort(v int32) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetType(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
