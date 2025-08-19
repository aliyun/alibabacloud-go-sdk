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
	// ID number.
	//
	// This is a required field when VerifyMetaType is set to VEHICLE_3_META.
	//
	// >
	//
	// > - When paramType is set to normal, enter the plain text.
	//
	// > - When paramType is set to md5, enter the first 6 digits in plain text + the birth date encrypted with MD5 (32 lowercase characters) + the last 4 digits in plain text.
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: Encrypted with MD5.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name
	//
	// > This is an explanation
	//
	// > - When paramType is set to normal, enter the plain text.
	//
	// > - When paramType is set to md5, encrypt the first character of the name with MD5 (32 lowercase characters) + the rest of the name in plain text.
	//
	// example:
	//
	// 张**
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Vehicle license plate
	//
	// >
	//
	// > - When paramType is set to normal, enter the plain text.
	//
	// > - When paramType is set to md5, enter the part of the license plate except for the last two characters in plain text + the last two characters of the license plate encrypted with MD5 (32 lowercase characters).
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
	// Verification type
	//
	// >
	//
	// > - VEHICLE_2_META: Two-element verification, name + vehicle license plate verification;
	//
	// > - VEHICLE_3_META: Three-element verification, name + vehicle license plate + ID number verification;
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
