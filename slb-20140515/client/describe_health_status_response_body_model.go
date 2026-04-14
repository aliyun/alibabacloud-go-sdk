// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *DescribeHealthStatusResponseBodyBackendServers) *DescribeHealthStatusResponseBody
	GetBackendServers() *DescribeHealthStatusResponseBodyBackendServers
	SetRequestId(v string) *DescribeHealthStatusResponseBody
	GetRequestId() *string
}

type DescribeHealthStatusResponseBody struct {
	BackendServers *DescribeHealthStatusResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBody) GetBackendServers() *DescribeHealthStatusResponseBodyBackendServers {
	return s.BackendServers
}

func (s *DescribeHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthStatusResponseBody) SetBackendServers(v *DescribeHealthStatusResponseBodyBackendServers) *DescribeHealthStatusResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetRequestId(v string) *DescribeHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHealthStatusResponseBodyBackendServers struct {
	BackendServer []*DescribeHealthStatusResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeHealthStatusResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyBackendServers) GetBackendServer() []*DescribeHealthStatusResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeHealthStatusResponseBodyBackendServers) SetBackendServer(v []*DescribeHealthStatusResponseBodyBackendServersBackendServer) *DescribeHealthStatusResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServers) Validate() error {
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

type DescribeHealthStatusResponseBodyBackendServersBackendServer struct {
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	Port               *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol           *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ServerHealthStatus *string `json:"ServerHealthStatus,omitempty" xml:"ServerHealthStatus,omitempty"`
	ServerId           *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerIp           *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
}

func (s DescribeHealthStatusResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetServerHealthStatus() *string {
	return s.ServerHealthStatus
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) GetServerIp() *string {
	return s.ServerIp
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetListenerPort(v int32) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ListenerPort = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetPort(v int32) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetProtocol(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerHealthStatus(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerHealthStatus = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerIp(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerIp = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
