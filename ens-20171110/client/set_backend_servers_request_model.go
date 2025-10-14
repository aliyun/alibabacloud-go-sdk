// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v []*SetBackendServersRequestBackendServers) *SetBackendServersRequest
	GetBackendServers() []*SetBackendServersRequestBackendServers
	SetLoadBalancerId(v string) *SetBackendServersRequest
	GetLoadBalancerId() *string
}

type SetBackendServersRequest struct {
	// The list of backend servers that you added. You can modify the weights of up to 20 backend servers in each request.
	//
	// This parameter is required.
	BackendServers []*SetBackendServersRequestBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Repeated"`
	// The ID of the Edge Load Balancer (ELB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5s7crik3yo3bp03gqrbp5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s SetBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersRequest) GoString() string {
	return s.String()
}

func (s *SetBackendServersRequest) GetBackendServers() []*SetBackendServersRequestBackendServers {
	return s.BackendServers
}

func (s *SetBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetBackendServersRequest) SetBackendServers(v []*SetBackendServersRequestBackendServers) *SetBackendServersRequest {
	s.BackendServers = v
	return s
}

func (s *SetBackendServersRequest) SetLoadBalancerId(v string) *SetBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersRequest) Validate() error {
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

type SetBackendServersRequestBackendServers struct {
	// The ID of the instance that you use as the backend server.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5ze0o05xccvbljtn****
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
	// This parameter is required.
	//
	// example:
	//
	// 20
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SetBackendServersRequestBackendServers) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersRequestBackendServers) GoString() string {
	return s.String()
}

func (s *SetBackendServersRequestBackendServers) GetServerId() *string {
	return s.ServerId
}

func (s *SetBackendServersRequestBackendServers) GetType() *string {
	return s.Type
}

func (s *SetBackendServersRequestBackendServers) GetWeight() *int32 {
	return s.Weight
}

func (s *SetBackendServersRequestBackendServers) SetServerId(v string) *SetBackendServersRequestBackendServers {
	s.ServerId = &v
	return s
}

func (s *SetBackendServersRequestBackendServers) SetType(v string) *SetBackendServersRequestBackendServers {
	s.Type = &v
	return s
}

func (s *SetBackendServersRequestBackendServers) SetWeight(v int32) *SetBackendServersRequestBackendServers {
	s.Weight = &v
	return s
}

func (s *SetBackendServersRequestBackendServers) Validate() error {
	return dara.Validate(s)
}
