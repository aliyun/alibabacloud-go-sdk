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
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type AddBackendServersResponseBodyBackendServersBackendServer struct {
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
