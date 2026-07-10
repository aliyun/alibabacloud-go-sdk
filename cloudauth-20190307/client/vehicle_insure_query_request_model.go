// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleInsureQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamType(v string) *VehicleInsureQueryRequest
	GetParamType() *string
	SetVehicleNum(v string) *VehicleInsureQueryRequest
	GetVehicleNum() *string
	SetVehicleType(v string) *VehicleInsureQueryRequest
	GetVehicleType() *string
	SetVin(v string) *VehicleInsureQueryRequest
	GetVin() *string
}

type VehicleInsureQueryRequest struct {
	// The parameter type. Valid values:
	//
	// - **normal**: not encrypted.
	//
	// - **md5**: MD5-encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The license plate number.
	//
	// >
	//
	// > - If ParamType is set to normal, enter the plaintext value.
	//
	// > - If ParamType is set to md5, enter the plaintext of the license plate number excluding the last two characters, concatenated with the MD5-encrypted value of the last two characters (32-bit lowercase MD5).
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// The vehicle type on the driving license.
	//
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// The vehicle identification number (VIN).
	//
	//
	// >
	//
	// > - If ParamType is set to normal, enter the plaintext value.
	//
	// > - If ParamType is set to md5, enter the plaintext of the VIN excluding the last 4 characters, concatenated with the MD5-encrypted value of the last 4 characters (32-bit lowercase MD5).
	//
	// example:
	//
	// LB**************
	Vin *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s VehicleInsureQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s VehicleInsureQueryRequest) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryRequest) GetParamType() *string {
	return s.ParamType
}

func (s *VehicleInsureQueryRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *VehicleInsureQueryRequest) GetVehicleType() *string {
	return s.VehicleType
}

func (s *VehicleInsureQueryRequest) GetVin() *string {
	return s.Vin
}

func (s *VehicleInsureQueryRequest) SetParamType(v string) *VehicleInsureQueryRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVehicleNum(v string) *VehicleInsureQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVehicleType(v string) *VehicleInsureQueryRequest {
	s.VehicleType = &v
	return s
}

func (s *VehicleInsureQueryRequest) SetVin(v string) *VehicleInsureQueryRequest {
	s.Vin = &v
	return s
}

func (s *VehicleInsureQueryRequest) Validate() error {
	return dara.Validate(s)
}
