// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v *DescribeBackendServersResponseBodyListeners) *DescribeBackendServersResponseBody
	GetListeners() *DescribeBackendServersResponseBodyListeners
	SetRequestId(v string) *DescribeBackendServersResponseBody
	GetRequestId() *string
}

type DescribeBackendServersResponseBody struct {
	Listeners *DescribeBackendServersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackendServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponseBody) GetListeners() *DescribeBackendServersResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackendServersResponseBody) SetListeners(v *DescribeBackendServersResponseBodyListeners) *DescribeBackendServersResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeBackendServersResponseBody) SetRequestId(v string) *DescribeBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackendServersResponseBody) Validate() error {
	if s.Listeners != nil {
		if err := s.Listeners.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackendServersResponseBodyListeners struct {
	Listener []*DescribeBackendServersResponseBodyListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeBackendServersResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponseBodyListeners) GetListener() []*DescribeBackendServersResponseBodyListenersListener {
	return s.Listener
}

func (s *DescribeBackendServersResponseBodyListeners) SetListener(v []*DescribeBackendServersResponseBodyListenersListener) *DescribeBackendServersResponseBodyListeners {
	s.Listener = v
	return s
}

func (s *DescribeBackendServersResponseBodyListeners) Validate() error {
	if s.Listener != nil {
		for _, item := range s.Listener {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackendServersResponseBodyListenersListener struct {
	BackendServers *DescribeBackendServersResponseBodyListenersListenerBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	ListenerPort   *int32                                                             `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s DescribeBackendServersResponseBodyListenersListener) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponseBodyListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponseBodyListenersListener) GetBackendServers() *DescribeBackendServersResponseBodyListenersListenerBackendServers {
	return s.BackendServers
}

func (s *DescribeBackendServersResponseBodyListenersListener) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeBackendServersResponseBodyListenersListener) SetBackendServers(v *DescribeBackendServersResponseBodyListenersListenerBackendServers) *DescribeBackendServersResponseBodyListenersListener {
	s.BackendServers = v
	return s
}

func (s *DescribeBackendServersResponseBodyListenersListener) SetListenerPort(v int32) *DescribeBackendServersResponseBodyListenersListener {
	s.ListenerPort = &v
	return s
}

func (s *DescribeBackendServersResponseBodyListenersListener) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackendServersResponseBodyListenersListenerBackendServers struct {
	BackendServer []*DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeBackendServersResponseBodyListenersListenerBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponseBodyListenersListenerBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServers) GetBackendServer() []*DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServers) SetBackendServer(v []*DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) *DescribeBackendServersResponseBodyListenersListenerBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServers) Validate() error {
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

type DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer struct {
	ServerHealthStatus *string `json:"ServerHealthStatus,omitempty" xml:"ServerHealthStatus,omitempty"`
	ServerId           *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) GetServerHealthStatus() *string {
	return s.ServerHealthStatus
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) SetServerHealthStatus(v string) *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer {
	s.ServerHealthStatus = &v
	return s
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) SetServerId(v string) *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeBackendServersResponseBodyListenersListenerBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}
