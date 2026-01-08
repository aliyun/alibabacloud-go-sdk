// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointZones(v []*CreateEndpointRequestEndpointZones) *CreateEndpointRequest
	GetEndpointZones() []*CreateEndpointRequestEndpointZones
	SetName(v string) *CreateEndpointRequest
	GetName() *string
	SetVpcId(v string) *CreateEndpointRequest
	GetVpcId() *string
	SetType(v string) *CreateEndpointRequest
	GetType() *string
}

type CreateEndpointRequest struct {
	// This parameter is required.
	EndpointZones []*CreateEndpointRequestEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// testendpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf664nyle5khp5***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) GetEndpointZones() []*CreateEndpointRequestEndpointZones {
	return s.EndpointZones
}

func (s *CreateEndpointRequest) GetName() *string {
	return s.Name
}

func (s *CreateEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEndpointRequest) GetType() *string {
	return s.Type
}

func (s *CreateEndpointRequest) SetEndpointZones(v []*CreateEndpointRequestEndpointZones) *CreateEndpointRequest {
	s.EndpointZones = v
	return s
}

func (s *CreateEndpointRequest) SetName(v string) *CreateEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointRequest) SetVpcId(v string) *CreateEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointRequest) SetType(v string) *CreateEndpointRequest {
	s.Type = &v
	return s
}

func (s *CreateEndpointRequest) Validate() error {
	if s.EndpointZones != nil {
		for _, item := range s.EndpointZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEndpointRequestEndpointZones struct {
	// example:
	//
	// vsw-uf6qmfkqdcw*****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateEndpointRequestEndpointZones) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointRequestEndpointZones) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequestEndpointZones) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateEndpointRequestEndpointZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateEndpointRequestEndpointZones) SetVswitchId(v string) *CreateEndpointRequestEndpointZones {
	s.VswitchId = &v
	return s
}

func (s *CreateEndpointRequestEndpointZones) SetZoneId(v string) *CreateEndpointRequestEndpointZones {
	s.ZoneId = &v
	return s
}

func (s *CreateEndpointRequestEndpointZones) Validate() error {
	return dara.Validate(s)
}
