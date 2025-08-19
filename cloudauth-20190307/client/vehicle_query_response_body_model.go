// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VehicleQueryResponseBody
	GetCode() *string
	SetMessage(v string) *VehicleQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *VehicleQueryResponseBody
	GetRequestId() *string
	SetResultObject(v *VehicleQueryResponseBodyResultObject) *VehicleQueryResponseBody
	GetResultObject() *VehicleQueryResponseBodyResultObject
}

type VehicleQueryResponseBody struct {
	// Return code: 200 for success, others for failure
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result
	ResultObject *VehicleQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VehicleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *VehicleQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VehicleQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VehicleQueryResponseBody) GetResultObject() *VehicleQueryResponseBodyResultObject {
	return s.ResultObject
}

func (s *VehicleQueryResponseBody) SetCode(v string) *VehicleQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleQueryResponseBody) SetMessage(v string) *VehicleQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleQueryResponseBody) SetRequestId(v string) *VehicleQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleQueryResponseBody) SetResultObject(v *VehicleQueryResponseBodyResultObject) *VehicleQueryResponseBody {
	s.ResultObject = v
	return s
}

func (s *VehicleQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type VehicleQueryResponseBodyResultObject struct {
	// Verification result code:
	//
	// >
	//
	// > - 1: Found (charged)
	//
	// > - 3: No record found (not charged)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Vehicle information.
	//
	// example:
	//
	// {
	//
	// 	"approvedCount": 7,
	//
	// 	"approvedLoad": 0,
	//
	// 	"axleCount": 2,
	//
	// 	"backWheelDistance": 1585,
	//
	// 	"brand": "雷克萨斯",
	//
	// 	"compulsoryScrapTo": "****-**-*	- **:**:**",
	//
	// 	"displacement": 4608,
	//
	// 	"engineNo": "1UR0******",
	//
	// 	"engineType": "1**",
	//
	// 	"frontWheelDistance": 1585,
	//
	// 	"fuelType": "汽油",
	//
	// 	"high": 1860,
	//
	// 	"inspectionResultEffectiveTo": "****-**-*	- **:**:**",
	//
	// 	"modelNo": "UR*****-******",
	//
	// 	"power": 228,
	//
	// 	"registrationDate": "****-**-*	- **:**:**",
	//
	// 	"releaseDate": "",
	//
	// 	"state": "0",
	//
	// 	"stateDesc": "正常",
	//
	// 	"totalMass": 2940,
	//
	// 	"unladenMass": 2350,
	//
	// 	"useProperty": 0,
	//
	// 	"usePropertyDesc": "非营运",
	//
	// 	"long": 4890,
	//
	// 	"vin": "A***********",
	//
	// 	"wheelBase": 2790,
	//
	// 	"wide": 1910
	//
	// }
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleQueryResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s VehicleQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *VehicleQueryResponseBodyResultObject) GetVehicleInfo() *string {
	return s.VehicleInfo
}

func (s *VehicleQueryResponseBodyResultObject) SetBizCode(v string) *VehicleQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleQueryResponseBodyResultObject) SetVehicleInfo(v string) *VehicleQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

func (s *VehicleQueryResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
