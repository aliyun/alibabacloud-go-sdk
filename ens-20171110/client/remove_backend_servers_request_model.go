// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v []*RemoveBackendServersRequestBackendServers) *RemoveBackendServersRequest
	GetBackendServers() []*RemoveBackendServersRequestBackendServers
	SetLoadBalancerId(v string) *RemoveBackendServersRequest
	GetLoadBalancerId() *string
}

type RemoveBackendServersRequest struct {
	// The list of backend servers that you want to remove. You can remove up to 20 backend servers at a time.
	//
	// This parameter is required.
	BackendServers []*RemoveBackendServersRequestBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Repeated"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5ovkn1piwqmoqrfjdyhq4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s RemoveBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersRequest) GetBackendServers() []*RemoveBackendServersRequestBackendServers {
	return s.BackendServers
}

func (s *RemoveBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveBackendServersRequest) SetBackendServers(v []*RemoveBackendServersRequestBackendServers) *RemoveBackendServersRequest {
	s.BackendServers = v
	return s
}

func (s *RemoveBackendServersRequest) SetLoadBalancerId(v string) *RemoveBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersRequest) Validate() error {
	if s.BackendServers != nil {
		for _, item := range s.BackendServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveBackendServersRequestBackendServers struct {
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.XXX.X.X
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The backend port that is used by the Edge Load Balancer (ELB) instance.
	//
	// example:
	//
	// 0
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The instance ID of the backend server.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5uf68ts0fqexe1a4n****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server. Valid values:
	//
	// 	- **ens**: an Edge Node Service (ENS) instance.
	//
	// 	- **eni**: an Elastic Network Interface (ENI).
	//
	// example:
	//
	// ens
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RemoveBackendServersRequestBackendServers) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersRequestBackendServers) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersRequestBackendServers) GetIp() *string {
	return s.Ip
}

func (s *RemoveBackendServersRequestBackendServers) GetPort() *int32 {
	return s.Port
}

func (s *RemoveBackendServersRequestBackendServers) GetServerId() *string {
	return s.ServerId
}

func (s *RemoveBackendServersRequestBackendServers) GetType() *string {
	return s.Type
}

func (s *RemoveBackendServersRequestBackendServers) GetWeight() *int32 {
	return s.Weight
}

func (s *RemoveBackendServersRequestBackendServers) SetIp(v string) *RemoveBackendServersRequestBackendServers {
	s.Ip = &v
	return s
}

func (s *RemoveBackendServersRequestBackendServers) SetPort(v int32) *RemoveBackendServersRequestBackendServers {
	s.Port = &v
	return s
}

func (s *RemoveBackendServersRequestBackendServers) SetServerId(v string) *RemoveBackendServersRequestBackendServers {
	s.ServerId = &v
	return s
}

func (s *RemoveBackendServersRequestBackendServers) SetType(v string) *RemoveBackendServersRequestBackendServers {
	s.Type = &v
	return s
}

func (s *RemoveBackendServersRequestBackendServers) SetWeight(v int32) *RemoveBackendServersRequestBackendServers {
	s.Weight = &v
	return s
}

func (s *RemoveBackendServersRequestBackendServers) Validate() error {
	return dara.Validate(s)
}
