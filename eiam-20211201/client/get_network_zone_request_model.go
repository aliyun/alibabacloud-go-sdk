// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNetworkZoneRequest
	GetInstanceId() *string
	SetNetworkZoneId(v string) *GetNetworkZoneRequest
	GetNetworkZoneId() *string
}

type GetNetworkZoneRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS的网络区域主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// network_11111
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
}

func (s GetNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkZoneRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *GetNetworkZoneRequest) SetInstanceId(v string) *GetNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkZoneRequest) SetNetworkZoneId(v string) *GetNetworkZoneRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *GetNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
