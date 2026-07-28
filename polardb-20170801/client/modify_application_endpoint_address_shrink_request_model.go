// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationEndpointAddressShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationEndpointAddressShrinkRequest
	GetApplicationId() *string
	SetEndpointId(v string) *ModifyApplicationEndpointAddressShrinkRequest
	GetEndpointId() *string
	SetNetType(v string) *ModifyApplicationEndpointAddressShrinkRequest
	GetNetType() *string
	SetNewConnectionStringPrefix(v string) *ModifyApplicationEndpointAddressShrinkRequest
	GetNewConnectionStringPrefix() *string
	SetNewPortsShrink(v string) *ModifyApplicationEndpointAddressShrinkRequest
	GetNewPortsShrink() *string
}

type ModifyApplicationEndpointAddressShrinkRequest struct {
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
	NewPortsShrink *string `json:"NewPorts,omitempty" xml:"NewPorts,omitempty"`
}

func (s ModifyApplicationEndpointAddressShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationEndpointAddressShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) GetNetType() *string {
	return s.NetType
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) GetNewConnectionStringPrefix() *string {
	return s.NewConnectionStringPrefix
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) GetNewPortsShrink() *string {
	return s.NewPortsShrink
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) SetApplicationId(v string) *ModifyApplicationEndpointAddressShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) SetEndpointId(v string) *ModifyApplicationEndpointAddressShrinkRequest {
	s.EndpointId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) SetNetType(v string) *ModifyApplicationEndpointAddressShrinkRequest {
	s.NetType = &v
	return s
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) SetNewConnectionStringPrefix(v string) *ModifyApplicationEndpointAddressShrinkRequest {
	s.NewConnectionStringPrefix = &v
	return s
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) SetNewPortsShrink(v string) *ModifyApplicationEndpointAddressShrinkRequest {
	s.NewPortsShrink = &v
	return s
}

func (s *ModifyApplicationEndpointAddressShrinkRequest) Validate() error {
	return dara.Validate(s)
}
