// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationEndpointAddressRequest
	GetApplicationId() *string
	SetEndpointId(v string) *ModifyApplicationEndpointAddressRequest
	GetEndpointId() *string
	SetNetType(v string) *ModifyApplicationEndpointAddressRequest
	GetNetType() *string
	SetNewConnectionStringPrefix(v string) *ModifyApplicationEndpointAddressRequest
	GetNewConnectionStringPrefix() *string
	SetNewPorts(v []*ModifyApplicationEndpointAddressRequestNewPorts) *ModifyApplicationEndpointAddressRequest
	GetNewPorts() []*ModifyApplicationEndpointAddressRequestNewPorts
}

type ModifyApplicationEndpointAddressRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The network type of the endpoint address. Valid values:
	//
	// 	- **Public**: public network.
	//
	// 	- **Private**: private network.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The new endpoint prefix.
	//
	// example:
	//
	// xg06iror0l
	NewConnectionStringPrefix *string `json:"NewConnectionStringPrefix,omitempty" xml:"NewConnectionStringPrefix,omitempty"`
	// The list of new ports.
	NewPorts []*ModifyApplicationEndpointAddressRequestNewPorts `json:"NewPorts,omitempty" xml:"NewPorts,omitempty" type:"Repeated"`
}

func (s ModifyApplicationEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationEndpointAddressRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationEndpointAddressRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ModifyApplicationEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *ModifyApplicationEndpointAddressRequest) GetNewConnectionStringPrefix() *string {
	return s.NewConnectionStringPrefix
}

func (s *ModifyApplicationEndpointAddressRequest) GetNewPorts() []*ModifyApplicationEndpointAddressRequestNewPorts {
	return s.NewPorts
}

func (s *ModifyApplicationEndpointAddressRequest) SetApplicationId(v string) *ModifyApplicationEndpointAddressRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequest) SetEndpointId(v string) *ModifyApplicationEndpointAddressRequest {
	s.EndpointId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequest) SetNetType(v string) *ModifyApplicationEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequest) SetNewConnectionStringPrefix(v string) *ModifyApplicationEndpointAddressRequest {
	s.NewConnectionStringPrefix = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequest) SetNewPorts(v []*ModifyApplicationEndpointAddressRequestNewPorts) *ModifyApplicationEndpointAddressRequest {
	s.NewPorts = v
	return s
}

func (s *ModifyApplicationEndpointAddressRequest) Validate() error {
	if s.NewPorts != nil {
		for _, item := range s.NewPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyApplicationEndpointAddressRequestNewPorts struct {
	// The new port value.
	//
	// example:
	//
	// 443
	NewPort *int32 `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// The old port value.
	//
	// example:
	//
	// 18789
	OldPort *int32 `json:"OldPort,omitempty" xml:"OldPort,omitempty"`
	// The port name.
	//
	// example:
	//
	// polarclaw
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
}

func (s ModifyApplicationEndpointAddressRequestNewPorts) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationEndpointAddressRequestNewPorts) GoString() string {
	return s.String()
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) GetNewPort() *int32 {
	return s.NewPort
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) GetOldPort() *int32 {
	return s.OldPort
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) GetPortName() *string {
	return s.PortName
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) SetNewPort(v int32) *ModifyApplicationEndpointAddressRequestNewPorts {
	s.NewPort = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) SetOldPort(v int32) *ModifyApplicationEndpointAddressRequestNewPorts {
	s.OldPort = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) SetPortName(v string) *ModifyApplicationEndpointAddressRequestNewPorts {
	s.PortName = &v
	return s
}

func (s *ModifyApplicationEndpointAddressRequestNewPorts) Validate() error {
	return dara.Validate(s)
}
