// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEpnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceName(v string) *CreateEpnInstanceRequest
	GetEPNInstanceName() *string
	SetEPNInstanceType(v string) *CreateEpnInstanceRequest
	GetEPNInstanceType() *string
	SetInternetChargeType(v string) *CreateEpnInstanceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *CreateEpnInstanceRequest
	GetInternetMaxBandwidthOut() *int32
	SetNetworkingModel(v string) *CreateEpnInstanceRequest
	GetNetworkingModel() *string
}

type CreateEpnInstanceRequest struct {
	// The name of the Edge-to-Edge network instance.
	//
	// example:
	//
	// test EPNInstanceName
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	// The type of the Edge-to-Edge communication instance. Valid value: **EdgeToEdge**.
	//
	// This parameter is required.
	//
	// example:
	//
	// EdgeToEdge
	EPNInstanceType *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	// The network billing type. Valid values:
	//
	// - **BandwidthByDay**: Daily peak bandwidth.
	//
	// - **95BandwidthByMonth**: Monthly 95th percentile peak bandwidth.
	//
	// - **PayByBandwidth4thMonth**: Monthly fourth peak bandwidth.
	//
	// - **PayByBandwidth**: Bandwidth with defined specifications.
	//
	// Only one of the Billable methods is supported. If a billing method already exists, the new value is invalid by default, and the existing one prevails.
	//
	// This parameter is required.
	//
	// example:
	//
	// BandwidthByDay
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound bandwidth to the Internet. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The networking mode. Valid values:
	//
	// - **SpeedUp**: Intelligent acceleration (Internet).
	//
	// - **Connection**: Internal network connectivity (internal network mode).
	//
	// - **SpeedUpAndConnection**: Intelligent acceleration and internal network connectivity (internal network mode).
	//
	// This parameter is required.
	//
	// example:
	//
	// SpeedUp
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
}

func (s CreateEpnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEpnInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceRequest) GetEPNInstanceName() *string {
	return s.EPNInstanceName
}

func (s *CreateEpnInstanceRequest) GetEPNInstanceType() *string {
	return s.EPNInstanceType
}

func (s *CreateEpnInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateEpnInstanceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateEpnInstanceRequest) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceName(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetEPNInstanceType(v string) *CreateEpnInstanceRequest {
	s.EPNInstanceType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetChargeType(v string) *CreateEpnInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateEpnInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateEpnInstanceRequest) SetNetworkingModel(v string) *CreateEpnInstanceRequest {
	s.NetworkingModel = &v
	return s
}

func (s *CreateEpnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
