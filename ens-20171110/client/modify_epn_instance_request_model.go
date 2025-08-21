// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *ModifyEpnInstanceRequest
	GetEPNInstanceId() *string
	SetEPNInstanceName(v string) *ModifyEpnInstanceRequest
	GetEPNInstanceName() *string
	SetInternetMaxBandwidthOut(v int32) *ModifyEpnInstanceRequest
	GetInternetMaxBandwidthOut() *int32
	SetNetworkingModel(v string) *ModifyEpnInstanceRequest
	GetNetworkingModel() *string
}

type ModifyEpnInstanceRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-****
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The name of the EPN instance.
	//
	// example:
	//
	// ens_test_epn
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 1 Mbit/s to 100 Mbit/s.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The networking mode. Valid values:
	//
	// 	- **SpeedUp**: Intelligent acceleration network (Internet).
	//
	// 	- **Connection**: Internal network.
	//
	// 	- **SpeedUpAndConnection**: Intelligent acceleration network and internal network.
	//
	// >  The internal network supports only **Connection*	- and **SpeedUpAndConnection**.
	//
	// example:
	//
	// SpeedUp
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
}

func (s ModifyEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *ModifyEpnInstanceRequest) GetEPNInstanceName() *string {
	return s.EPNInstanceName
}

func (s *ModifyEpnInstanceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyEpnInstanceRequest) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceId(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetEPNInstanceName(v string) *ModifyEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetInternetMaxBandwidthOut(v int32) *ModifyEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyEpnInstanceRequest) SetNetworkingModel(v string) *ModifyEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *ModifyEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
