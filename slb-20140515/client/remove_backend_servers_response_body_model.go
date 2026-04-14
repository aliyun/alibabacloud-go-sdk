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
	SetLoadBalancerId(v string) *RemoveBackendServersResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *RemoveBackendServersResponseBody
	GetRequestId() *string
}

type RemoveBackendServersResponseBody struct {
	BackendServers *RemoveBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp15lbk8uja8rvm4a****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
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

func (s *RemoveBackendServersResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveBackendServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveBackendServersResponseBody) SetBackendServers(v *RemoveBackendServersResponseBodyBackendServers) *RemoveBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *RemoveBackendServersResponseBody) SetLoadBalancerId(v string) *RemoveBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersResponseBody) SetRequestId(v string) *RemoveBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBackendServersResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type RemoveBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
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

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
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
