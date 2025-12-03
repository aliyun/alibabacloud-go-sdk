// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVServerGroupBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *ModifyVServerGroupBackendServersResponseBodyBackendServers) *ModifyVServerGroupBackendServersResponseBody
	GetBackendServers() *ModifyVServerGroupBackendServersResponseBodyBackendServers
	SetRequestId(v string) *ModifyVServerGroupBackendServersResponseBody
	GetRequestId() *string
	SetVServerGroupId(v string) *ModifyVServerGroupBackendServersResponseBody
	GetVServerGroupId() *string
}

type ModifyVServerGroupBackendServersResponseBody struct {
	// The backend servers.
	BackendServers *ModifyVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
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
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ModifyVServerGroupBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBody) GetBackendServers() *ModifyVServerGroupBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *ModifyVServerGroupBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVServerGroupBackendServersResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetBackendServers(v *ModifyVServerGroupBackendServersResponseBodyBackendServers) *ModifyVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetRequestId(v string) *ModifyVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *ModifyVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServers) GetBackendServer() []*ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) *ModifyVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServers) Validate() error {
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

type ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
	// The description of the backend server.
	//
	// example:
	//
	// Backend server description
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
	// i-bp1ge5hrp****
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

func (s ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
