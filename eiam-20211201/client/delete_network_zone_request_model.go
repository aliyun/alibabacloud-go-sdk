// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNetworkZoneRequest
	GetInstanceId() *string
	SetNetworkZoneId(v string) *DeleteNetworkZoneRequest
	GetNetworkZoneId() *string
}

type DeleteNetworkZoneRequest struct {
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

func (s DeleteNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNetworkZoneRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *DeleteNetworkZoneRequest) SetInstanceId(v string) *DeleteNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNetworkZoneRequest) SetNetworkZoneId(v string) *DeleteNetworkZoneRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *DeleteNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
