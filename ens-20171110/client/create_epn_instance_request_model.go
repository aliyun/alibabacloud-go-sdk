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
	// The name of the EPN instance.
	//
	// example:
	//
	// test EPNInstanceName
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	// The type of the EPN instance. Set the value to **EdgeToEdge**.
	//
	// This parameter is required.
	//
	// example:
	//
	// EdgeToEdge
	EPNInstanceType *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- **BandwidthByDay**: Pay by daily peak bandwidth.
	//
	// 	- **95BandwidthByMonth**: Pay by monthly 95th percentile bandwidth.
	//
	// 	- **PayByBandwidth4thMonth**: Pay by monthly fourth peak bandwidth.
	//
	// 	- **PayByBandwidth**: Pay by fixed bandwidth.
	//
	// You can specify only one metering method for network usage and cannot overwrite the existing metering method.
	//
	// This parameter is required.
	//
	// example:
	//
	// BandwidthByDay
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The networking mode. Valid values:
	//
	// 	- **SpeedUp**: intelligent acceleration network (Internet)
	//
	// 	- **Connection**: internal network
	//
	// 	- **SpeedUpAndConnection**: intelligent acceleration network and internal network
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
