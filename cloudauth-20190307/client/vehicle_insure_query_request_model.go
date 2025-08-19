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
	// Parameter type:
	//
	// - **normal**: Unencrypted.
	//
	// - **md5**: MD5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// License plate number
	//
	// >
	//
	// > - When `paramType` is set to `normal`, enter the plain text.
	//
	// > - When `paramType` is set to `md5`, enter the plain text of all but the last two characters of the license plate + the MD5 encryption (32 lowercase characters) of the last two characters of the license plate.
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// Driver\\"s license vehicle type.
	//
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// Vehicle identification code, i.e., the vehicle VIN
	//
	//
	// >
	//
	// > - When `paramType` is set to `normal`, enter the plain text.
	//
	// > - When `paramType` is set to `md5`, enter the plain text of all but the last four characters of the VIN + the MD5 encryption (32 lowercase characters) of the last four characters of the VIN.
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
