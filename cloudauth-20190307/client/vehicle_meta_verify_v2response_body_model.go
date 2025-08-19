// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VehicleMetaVerifyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *VehicleMetaVerifyV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *VehicleMetaVerifyV2ResponseBody
	GetRequestId() *string
	SetResultObject(v *VehicleMetaVerifyV2ResponseBodyResultObject) *VehicleMetaVerifyV2ResponseBody
	GetResultObject() *VehicleMetaVerifyV2ResponseBodyResultObject
}

type VehicleMetaVerifyV2ResponseBody struct {
	// Return code, **200*	- indicates successful API response.
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
	// 5A6229C0-E156-48E4-B6EC-0F528B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result
	ResultObject *VehicleMetaVerifyV2ResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleMetaVerifyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *VehicleMetaVerifyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VehicleMetaVerifyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VehicleMetaVerifyV2ResponseBody) GetResultObject() *VehicleMetaVerifyV2ResponseBodyResultObject {
	return s.ResultObject
}

func (s *VehicleMetaVerifyV2ResponseBody) SetCode(v string) *VehicleMetaVerifyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetMessage(v string) *VehicleMetaVerifyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetRequestId(v string) *VehicleMetaVerifyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) SetResultObject(v *VehicleMetaVerifyV2ResponseBodyResultObject) *VehicleMetaVerifyV2ResponseBody {
	s.ResultObject = v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type VehicleMetaVerifyV2ResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: Verification consistent.
	//
	// - **2**: Verification inconsistent.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Detailed vehicle information.
	//
	// example:
	//
	// {
	//
	// 	"approvedCount": "7",
	//
	// 	"approvedLoad": "0",
	//
	// 	"axleCount": "2",
	//
	// 	"backWheelDistance": "1585",
	//
	// 	"brand": "雷克萨斯",
	//
	// 	"displacement": "4608",
	//
	// 	"engineNum": "1*********",
	//
	// 	"engineType": "1**",
	//
	// 	"frontWheelDistance": "1585",
	//
	// 	"fuelType": "汽油",
	//
	// 	"inspectionDate": "****-**-*	- **:**:**",
	//
	// 	"modelNum": "UR*****-******",
	//
	// 	"power": "228",
	//
	// 	"registrationDate": "****-**-*	- **:**:**",
	//
	// 	"retirementDate": "****-**-*	- **:**:**",
	//
	// 	"totalMass": "2940",
	//
	// 	"unladenMass": "2350",
	//
	// 	"useProperty": "非营运",
	//
	// 	"vehicleState": "正常",
	//
	// 	"vin": "JT***************",
	//
	// 	"wheelBase": "2790"
	//
	// }
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleMetaVerifyV2ResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyV2ResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) GetVehicleInfo() *string {
	return s.VehicleInfo
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) SetBizCode(v string) *VehicleMetaVerifyV2ResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) SetVehicleInfo(v string) *VehicleMetaVerifyV2ResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

func (s *VehicleMetaVerifyV2ResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
