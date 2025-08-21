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
	SetRequestId(v string) *AddBackendServersResponseBody
	GetRequestId() *string
}

type AddBackendServersResponseBody struct {
	// The list of backend servers that you want to add. You can add at most 20 backend servers.
	//
	// >  Only ENS instances that are in the running state can be attached to the ELB instance as backend servers.
	BackendServers *AddBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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

func (s *AddBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBackendServersResponseBody) SetBackendServers(v *AddBackendServersResponseBodyBackendServers) *AddBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *AddBackendServersResponseBody) SetRequestId(v string) *AddBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBackendServersResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type AddBackendServersResponseBodyBackendServersBackendServer struct {
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The backend port that is used by the ELB instance.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the instance that is used as the backend server.
	//
	// example:
	//
	// i-5uf6hj58zvml4ali8****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ens**: ENS instance.
	//
	// 	- **eni**: ENI.
	//
	// example:
	//
	// ens
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server. Default value: 100. Valid values: **0*	- to **100**.
	//
	// >  The value 0 indicates that requests are not forwarded to the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetIp() *string {
	return s.Ip
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetIp(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Ip = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
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

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
