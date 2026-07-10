// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamType(v string) *VehicleQueryRequest
	GetParamType() *string
	SetVehicleNum(v string) *VehicleQueryRequest
	GetVehicleNum() *string
	SetVehicleType(v string) *VehicleQueryRequest
	GetVehicleType() *string
}

type VehicleQueryRequest struct {
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
	// > - If ParamType is set to normal, enter the plaintext license plate number.
	//
	// > - If ParamType is set to md5, enter the plaintext of the license plate number excluding the last two characters, concatenated with the MD5-encrypted value of the last two characters (32-bit lowercase MD5).
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// The vehicle type.
	//
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s VehicleQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s VehicleQueryRequest) GoString() string {
	return s.String()
}

func (s *VehicleQueryRequest) GetParamType() *string {
	return s.ParamType
}

func (s *VehicleQueryRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *VehicleQueryRequest) GetVehicleType() *string {
	return s.VehicleType
}

func (s *VehicleQueryRequest) SetParamType(v string) *VehicleQueryRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleQueryRequest) SetVehicleNum(v string) *VehicleQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleQueryRequest) SetVehicleType(v string) *VehicleQueryRequest {
	s.VehicleType = &v
	return s
}

func (s *VehicleQueryRequest) Validate() error {
	return dara.Validate(s)
}
