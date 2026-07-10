// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicle5ItemQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamType(v string) *Vehicle5ItemQueryRequest
	GetParamType() *string
	SetVehicleNum(v string) *Vehicle5ItemQueryRequest
	GetVehicleNum() *string
	SetVehicleType(v string) *Vehicle5ItemQueryRequest
	GetVehicleType() *string
}

type Vehicle5ItemQueryRequest struct {
	// The parameter type. Valid values:
	//
	// - **normal**: Not encrypted.
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
	// > - If ParamType is set to md5, enter the plaintext of the license plate number excluding the last two characters, concatenated with the MD5-encrypted last two characters (32-bit lowercase MD5).
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// The vehicle type.
	//
	// >
	//
	// > - 02: standard passenger car
	//
	// > - 52: new energy passenger car.
	//
	// example:
	//
	// 02
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s Vehicle5ItemQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s Vehicle5ItemQueryRequest) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Vehicle5ItemQueryRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *Vehicle5ItemQueryRequest) GetVehicleType() *string {
	return s.VehicleType
}

func (s *Vehicle5ItemQueryRequest) SetParamType(v string) *Vehicle5ItemQueryRequest {
	s.ParamType = &v
	return s
}

func (s *Vehicle5ItemQueryRequest) SetVehicleNum(v string) *Vehicle5ItemQueryRequest {
	s.VehicleNum = &v
	return s
}

func (s *Vehicle5ItemQueryRequest) SetVehicleType(v string) *Vehicle5ItemQueryRequest {
	s.VehicleType = &v
	return s
}

func (s *Vehicle5ItemQueryRequest) Validate() error {
	return dara.Validate(s)
}
