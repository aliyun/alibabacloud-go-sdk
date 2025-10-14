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
	SetRequestId(v string) *SetBackendServersResponseBody
	GetRequestId() *string
}

type SetBackendServersResponseBody struct {
	// The backend servers.
	BackendServers *SetBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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

func (s *SetBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetBackendServersResponseBody) SetBackendServers(v *SetBackendServersResponseBodyBackendServers) *SetBackendServersResponseBody {
	s.BackendServers = v
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
	// The ID of the instance that you want to use as the backend server.
	//
	// example:
	//
	// i-5uf6hj58zvml4ali8****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ens**: ENS instance.
	//
	// 	- **eni**: ENI instance.
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

func (s SetBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetIp() *string {
	return s.Ip
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetIp(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Ip = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
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

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
