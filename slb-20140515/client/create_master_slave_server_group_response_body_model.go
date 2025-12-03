// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMasterSlaveServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterSlaveBackendServers(v *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) *CreateMasterSlaveServerGroupResponseBody
	GetMasterSlaveBackendServers() *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers
	SetMasterSlaveServerGroupId(v string) *CreateMasterSlaveServerGroupResponseBody
	GetMasterSlaveServerGroupId() *string
	SetRequestId(v string) *CreateMasterSlaveServerGroupResponseBody
	GetRequestId() *string
}

type CreateMasterSlaveServerGroupResponseBody struct {
	// The backend servers in the primary/secondary server group.
	MasterSlaveBackendServers *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty" type:"Struct"`
	// The ID of the active/standby server group.
	//
	// example:
	//
	// rsp-bp19au4******
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7CA4DB76-4D32-523B-822E-5C9494613D46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMasterSlaveServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBody) GetMasterSlaveBackendServers() *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers {
	return s.MasterSlaveBackendServers
}

func (s *CreateMasterSlaveServerGroupResponseBody) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *CreateMasterSlaveServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetMasterSlaveBackendServers(v *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) *CreateMasterSlaveServerGroupResponseBody {
	s.MasterSlaveBackendServers = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetMasterSlaveServerGroupId(v string) *CreateMasterSlaveServerGroupResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetRequestId(v string) *CreateMasterSlaveServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBody) Validate() error {
	if s.MasterSlaveBackendServers != nil {
		if err := s.MasterSlaveBackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers struct {
	MasterSlaveBackendServer []*CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer `json:"MasterSlaveBackendServer,omitempty" xml:"MasterSlaveBackendServer,omitempty" type:"Repeated"`
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) GetMasterSlaveBackendServer() []*CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	return s.MasterSlaveBackendServer
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) SetMasterSlaveBackendServer(v []*CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers {
	s.MasterSlaveBackendServer = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) Validate() error {
	if s.MasterSlaveBackendServer != nil {
		for _, item := range s.MasterSlaveBackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer struct {
	// The description of the primary/secondary server group.
	//
	// example:
	//
	// test-112
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server.
	//
	// example:
	//
	// 82
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server that you want to add.
	//
	// example:
	//
	// i-bp1fq61enf4loa5i****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server.
	//
	// Valid values: **Master*	- and **Slave**.
	//
	// example:
	//
	// Master
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The type of backend server. Valid values:
	//
	// 	- **ecs**: ECS instance
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

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetDescription() *string {
	return s.Description
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetServerType() *string {
	return s.ServerType
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetType() *string {
	return s.Type
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetDescription(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Description = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetPort(v int32) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Port = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerId(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerType(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerType = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetType(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Type = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetWeight(v int32) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Weight = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) Validate() error {
	return dara.Validate(s)
}
