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
	// > - When paramType is set to normal, enter the plain text.
	//
	// > - When paramType is set to md5, enter the unencrypted part of the license plate number except for the last two characters + the MD5 (32 lowercase) encryption of the last two characters of the license plate.
	//
	// example:
	//
	// 陕A9****
	VehicleNum *string `json:"VehicleNum,omitempty" xml:"VehicleNum,omitempty"`
	// Vehicle type
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
