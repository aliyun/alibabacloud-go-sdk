// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *RemoveBackendServersResponseBodyBackendServers) *RemoveBackendServersResponseBody
	GetBackendServers() *RemoveBackendServersResponseBodyBackendServers
	SetRequestId(v string) *RemoveBackendServersResponseBody
	GetRequestId() *string
}

type RemoveBackendServersResponseBody struct {
	// The list of backend servers that you want to add to the SLB instance.
	BackendServers *RemoveBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBody) GetBackendServers() *RemoveBackendServersResponseBodyBackendServers {
	return s.BackendServers
}

func (s *RemoveBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveBackendServersResponseBody) SetBackendServers(v *RemoveBackendServersResponseBodyBackendServers) *RemoveBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *RemoveBackendServersResponseBody) SetRequestId(v string) *RemoveBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBackendServersResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveBackendServersResponseBodyBackendServers struct {
	BackendServer []*RemoveBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s RemoveBackendServersResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBodyBackendServers) GetBackendServer() []*RemoveBackendServersResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *RemoveBackendServersResponseBodyBackendServers) SetBackendServer(v []*RemoveBackendServersResponseBodyBackendServersBackendServer) *RemoveBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServers) Validate() error {
	return dara.Validate(s)
}

type RemoveBackendServersResponseBodyBackendServersBackendServer struct {
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.1XX.X.X
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The backend port that is used by the ELB instance.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The instance ID of the backend server.
	//
	// example:
	//
	// i-5vb5h5njxiuhn48a****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ens**: an ENS instance.
	//
	// 	- **eni**: an ENI.
	//
	// example:
	//
	// ens
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// >  The value 0 indicates that requests are not forwarded to the backend server.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetIp() *string {
	return s.Ip
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetIp(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Ip = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
