// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVServerGroupBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *RemoveVServerGroupBackendServersResponseBodyBackendServers) *RemoveVServerGroupBackendServersResponseBody
	GetBackendServers() *RemoveVServerGroupBackendServersResponseBodyBackendServers
	SetRequestId(v string) *RemoveVServerGroupBackendServersResponseBody
	GetRequestId() *string
	SetVServerGroupId(v string) *RemoveVServerGroupBackendServersResponseBody
	GetVServerGroupId() *string
}

type RemoveVServerGroupBackendServersResponseBody struct {
	// The backend servers.
	BackendServers *RemoveVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
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

func (s RemoveVServerGroupBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBody) GetBackendServers() *RemoveVServerGroupBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *RemoveVServerGroupBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveVServerGroupBackendServersResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetBackendServers(v *RemoveVServerGroupBackendServersResponseBodyBackendServers) *RemoveVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetRequestId(v string) *RemoveVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *RemoveVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServers) GetBackendServer() []*RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) *RemoveVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServers) Validate() error {
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

type RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
	// The port that is used by the backend server.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// vm-230
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
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

func (s RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
