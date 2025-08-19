// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *VehicleMetaVerifyV2Request
	GetIdentifyNum() *string
	SetParamType(v string) *VehicleMetaVerifyV2Request
	GetParamType() *string
	SetUserName(v string) *VehicleMetaVerifyV2Request
	GetUserName() *string
	SetVehicleNum(v string) *VehicleMetaVerifyV2Request
	GetVehicleNum() *string
	SetVehicleType(v string) *VehicleMetaVerifyV2Request
	GetVehicleType() *string
	SetVerifyMetaType(v string) *VehicleMetaVerifyV2Request
	GetVerifyMetaType() *string
}

type VehicleMetaVerifyV2Request struct {
	// ID number.
	//
	// This is a required field when VerifyMetaType is VEHICLE_3_META.
	//
	// >
	//
	// > - When paramType is normal, enter plain text.
	//
	// > - When paramType is md5, enter the first 6 digits in plain text + MD5 (32 lowercase) of the birth date + the last 4 digits in plain text.
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: Md5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name
	//
	// >
	//
	// > - When paramType is normal, enter plain text.
	//
	// > - When paramType is md5, enter the first character of the name as MD5 (32 lowercase) + the rest of the name in plain text.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// License plate number
	//
	// >
	//
	// > - When paramType is normal, enter plain text.
	//
	// > - When paramType is md5, enter all but the last two characters in plain text + the last two characters as MD5 (32 lowercase).
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
	// > - VEHICLE_2_META: Two-element verification, name + license plate number verification;
	//
	// > - VEHICLE_3_META: Three-element verification, name + license plate number + ID number verification;
	//
	// example:
	//
	// VEHICLE_3_META
	VerifyMetaType *string `json:"VerifyMetaType,omitempty" xml:"VerifyMetaType,omitempty"`
}

func (s VehicleMetaVerifyV2Request) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyV2Request) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2Request) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *VehicleMetaVerifyV2Request) GetParamType() *string {
	return s.ParamType
}

func (s *VehicleMetaVerifyV2Request) GetUserName() *string {
	return s.UserName
}

func (s *VehicleMetaVerifyV2Request) GetVehicleNum() *string {
	return s.VehicleNum
}

func (s *VehicleMetaVerifyV2Request) GetVehicleType() *string {
	return s.VehicleType
}

func (s *VehicleMetaVerifyV2Request) GetVerifyMetaType() *string {
	return s.VerifyMetaType
}

func (s *VehicleMetaVerifyV2Request) SetIdentifyNum(v string) *VehicleMetaVerifyV2Request {
	s.IdentifyNum = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetParamType(v string) *VehicleMetaVerifyV2Request {
	s.ParamType = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetUserName(v string) *VehicleMetaVerifyV2Request {
	s.UserName = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVehicleNum(v string) *VehicleMetaVerifyV2Request {
	s.VehicleNum = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVehicleType(v string) *VehicleMetaVerifyV2Request {
	s.VehicleType = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) SetVerifyMetaType(v string) *VehicleMetaVerifyV2Request {
	s.VerifyMetaType = &v
	return s
}

func (s *VehicleMetaVerifyV2Request) Validate() error {
	return dara.Validate(s)
}
