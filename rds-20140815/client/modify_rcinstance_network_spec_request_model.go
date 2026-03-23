// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceNetworkSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyRCInstanceNetworkSpecRequest
	GetInstanceId() *string
	SetInternetMaxBandwidthOut(v string) *ModifyRCInstanceNetworkSpecRequest
	GetInternetMaxBandwidthOut() *string
	SetNetworkChargeType(v string) *ModifyRCInstanceNetworkSpecRequest
	GetNetworkChargeType() *string
	SetRegionId(v string) *ModifyRCInstanceNetworkSpecRequest
	GetRegionId() *string
}

type ModifyRCInstanceNetworkSpecRequest struct {
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetMaxBandwidthOut *string `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	NetworkChargeType       *string `json:"NetworkChargeType,omitempty" xml:"NetworkChargeType,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCInstanceNetworkSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceNetworkSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceNetworkSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() *string {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyRCInstanceNetworkSpecRequest) GetNetworkChargeType() *string {
	return s.NetworkChargeType
}

func (s *ModifyRCInstanceNetworkSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceNetworkSpecRequest) SetInstanceId(v string) *ModifyRCInstanceNetworkSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(v string) *ModifyRCInstanceNetworkSpecRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecRequest) SetNetworkChargeType(v string) *ModifyRCInstanceNetworkSpecRequest {
	s.NetworkChargeType = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecRequest) SetRegionId(v string) *ModifyRCInstanceNetworkSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceNetworkSpecRequest) Validate() error {
	return dara.Validate(s)
}
