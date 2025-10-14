// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v []*AddBackendServersRequestBackendServers) *AddBackendServersRequest
	GetBackendServers() []*AddBackendServersRequestBackendServers
	SetLoadBalancerId(v string) *AddBackendServersRequest
	GetLoadBalancerId() *string
}

type AddBackendServersRequest struct {
	// The list of backend servers that you want to add to the Edge Load Balancer (ELB) instance. You can add up to 20 backend servers at a time.
	//
	// >  Only Edge Node Service (ENS) instances that are in the running state can be added to the ELB instance as backend servers.
	//
	// This parameter is required.
	BackendServers []*AddBackendServersRequestBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Repeated"`
	// The frontend port that is used by the Edge Load Balance (ELB) instance. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5qzdmxefgrpxd7oz2mefonvtx
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s AddBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddBackendServersRequest) GetBackendServers() []*AddBackendServersRequestBackendServers {
	return s.BackendServers
}

func (s *AddBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddBackendServersRequest) SetBackendServers(v []*AddBackendServersRequestBackendServers) *AddBackendServersRequest {
	s.BackendServers = v
	return s
}

func (s *AddBackendServersRequest) SetLoadBalancerId(v string) *AddBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersRequest) Validate() error {
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

type AddBackendServersRequestBackendServers struct {
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.X.X
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The backend port that is used by the ELB instance.
	//
	// example:
	//
	// 3309
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the ENS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5uf6dwyzch3wly790****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ens**: ENS instance
	//
	// 	- **eni**: elastic network interface (ENI)
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

func (s AddBackendServersRequestBackendServers) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersRequestBackendServers) GoString() string {
	return s.String()
}

func (s *AddBackendServersRequestBackendServers) GetIp() *string {
	return s.Ip
}

func (s *AddBackendServersRequestBackendServers) GetPort() *int32 {
	return s.Port
}

func (s *AddBackendServersRequestBackendServers) GetServerId() *string {
	return s.ServerId
}

func (s *AddBackendServersRequestBackendServers) GetType() *string {
	return s.Type
}

func (s *AddBackendServersRequestBackendServers) GetWeight() *int32 {
	return s.Weight
}

func (s *AddBackendServersRequestBackendServers) SetIp(v string) *AddBackendServersRequestBackendServers {
	s.Ip = &v
	return s
}

func (s *AddBackendServersRequestBackendServers) SetPort(v int32) *AddBackendServersRequestBackendServers {
	s.Port = &v
	return s
}

func (s *AddBackendServersRequestBackendServers) SetServerId(v string) *AddBackendServersRequestBackendServers {
	s.ServerId = &v
	return s
}

func (s *AddBackendServersRequestBackendServers) SetType(v string) *AddBackendServersRequestBackendServers {
	s.Type = &v
	return s
}

func (s *AddBackendServersRequestBackendServers) SetWeight(v int32) *AddBackendServersRequestBackendServers {
	s.Weight = &v
	return s
}

func (s *AddBackendServersRequestBackendServers) Validate() error {
	return dara.Validate(s)
}
