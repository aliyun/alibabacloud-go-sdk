// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetNetworkType(v string) *UpdateInstanceRequest
	GetNetworkType() *string
	SetZones(v []*UpdateInstanceRequestZones) *UpdateInstanceRequest
	GetZones() []*UpdateInstanceRequestZones
}

type UpdateInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentteams-abc123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 新的实例名称
	InstanceName *string                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	NetworkType  *string                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Zones        []*UpdateInstanceRequestZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateInstanceRequest) GetZones() []*UpdateInstanceRequestZones {
	return s.Zones
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkType(v string) *UpdateInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *UpdateInstanceRequest) SetZones(v []*UpdateInstanceRequestZones) *UpdateInstanceRequest {
	s.Zones = v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateInstanceRequestZones struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateInstanceRequestZones) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestZones) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateInstanceRequestZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateInstanceRequestZones) SetVSwitchId(v string) *UpdateInstanceRequestZones {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceRequestZones) SetZoneId(v string) *UpdateInstanceRequestZones {
	s.ZoneId = &v
	return s
}

func (s *UpdateInstanceRequestZones) Validate() error {
	return dara.Validate(s)
}
