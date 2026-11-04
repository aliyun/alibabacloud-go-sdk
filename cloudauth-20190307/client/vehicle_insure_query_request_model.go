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
	// - **normal**: No encryption.
	//
	// - **md5**: MD5 encryption.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The license plate number.
	//
	// >
	//
	// > - If paramType is set to normal, enter the plaintext.
	//
	// > - If paramType is set to md5, enter the plaintext of the license plate number excluding the last two characters, followed by the MD5 hash (32-character lowercase) of the last two characters.
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// The permitted vehicle type of the driver\\"s license.
	//
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// The vehicle identification number (VIN).
	//
	// >
	//
	// > - If paramType is set to normal, enter the plaintext.
	//
	// > - If paramType is set to md5, enter the plaintext of the VIN excluding the last four characters, followed by the MD5 hash (32-character lowercase) of the last four characters.
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
