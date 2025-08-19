// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointZones(v []*UpdateEndpointRequestEndpointZones) *UpdateEndpointRequest
	GetEndpointZones() []*UpdateEndpointRequestEndpointZones
	SetName(v string) *UpdateEndpointRequest
	GetName() *string
}

type UpdateEndpointRequest struct {
	// This parameter is required.
	EndpointZones []*UpdateEndpointRequestEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointRequest) GetEndpointZones() []*UpdateEndpointRequestEndpointZones {
	return s.EndpointZones
}

func (s *UpdateEndpointRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEndpointRequest) SetEndpointZones(v []*UpdateEndpointRequestEndpointZones) *UpdateEndpointRequest {
	s.EndpointZones = v
	return s
}

func (s *UpdateEndpointRequest) SetName(v string) *UpdateEndpointRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateEndpointRequestEndpointZones struct {
	// example:
	//
	// vsw-bp18r8uwnukv3rvi9****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s UpdateEndpointRequestEndpointZones) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointRequestEndpointZones) GoString() string {
	return s.String()
}

func (s *UpdateEndpointRequestEndpointZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateEndpointRequestEndpointZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateEndpointRequestEndpointZones) SetVSwitchId(v string) *UpdateEndpointRequestEndpointZones {
	s.VSwitchId = &v
	return s
}

func (s *UpdateEndpointRequestEndpointZones) SetZoneId(v string) *UpdateEndpointRequestEndpointZones {
	s.ZoneId = &v
	return s
}

func (s *UpdateEndpointRequestEndpointZones) Validate() error {
	return dara.Validate(s)
}
