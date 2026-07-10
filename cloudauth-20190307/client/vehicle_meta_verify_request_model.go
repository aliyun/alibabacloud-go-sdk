// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *VehicleMetaVerifyRequest
	GetIdentifyNum() *string
	SetParamType(v string) *VehicleMetaVerifyRequest
	GetParamType() *string
	SetUserName(v string) *VehicleMetaVerifyRequest
	GetUserName() *string
	SetVehicleNum(v string) *VehicleMetaVerifyRequest
	GetVehicleNum() *string
	SetVehicleType(v string) *VehicleMetaVerifyRequest
	GetVehicleType() *string
	SetVerifyMetaType(v string) *VehicleMetaVerifyRequest
	GetVerifyMetaType() *string
}

type VehicleMetaVerifyRequest struct {
	// The ID card number.
	//
	// This parameter is required when VerifyMetaType is set to VEHICLE_3_META.
	//
	// >
	//
	// > - When paramType is set to normal, enter the plaintext.
	//
	// > - When paramType is set to md5, enter the first 6 digits of the ID card number in plaintext + the MD5-encrypted date of birth (32-bit lowercase MD5) + the last 4 digits of the ID card number.
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The parameter type. Valid values:
	//
	// - normal: not encrypted.
	//
	// - md5: MD5-encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name.
	//
	// > Note
	//
	// > - When paramType is set to normal, enter the plaintext.
	//
	// > - When paramType is set to md5, enter the MD5-encrypted first character of the name (32-bit lowercase MD5) + the remaining characters of the name in plaintext.
	//
	// example:
	//
	// 张**
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The license plate number.
	//
	// >
	//
	// > - When paramType is set to normal, enter the plaintext.
	//
	// > - When paramType is set to md5, enter the characters of the license plate number except the last two in plaintext + the MD5-encrypted last two characters (32-bit lowercase MD5).
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
	// The verification type.
	//
	// >
	//
	// > - VEHICLE_2_META: two-element verification. Verifies the name and license plate number.
	//
	// > - VEHICLE_3_META: three-element verification. Verifies the name, license plate number, and ID card number.
	//
	// example:
	//
	// VEHICLE_2_META
	VerifyMetaType *string `json:"VerifyMetaType,omitempty" xml:"VerifyMetaType,omitempty"`
}

func (s VehicleMetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *VehicleMetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *VehicleMetaVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *VehicleMetaVerifyRequest) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *VehicleMetaVerifyRequest) GetVehicleType() *string {
	return s.VehicleType
}

func (s *VehicleMetaVerifyRequest) GetVerifyMetaType() *string {
	return s.VerifyMetaType
}

func (s *VehicleMetaVerifyRequest) SetIdentifyNum(v string) *VehicleMetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetParamType(v string) *VehicleMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetUserName(v string) *VehicleMetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVehicleNum(v string) *VehicleMetaVerifyRequest {
	s.VehicleNum = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVehicleType(v string) *VehicleMetaVerifyRequest {
	s.VehicleType = &v
	return s
}

func (s *VehicleMetaVerifyRequest) SetVerifyMetaType(v string) *VehicleMetaVerifyRequest {
	s.VerifyMetaType = &v
	return s
}

func (s *VehicleMetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
