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
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
