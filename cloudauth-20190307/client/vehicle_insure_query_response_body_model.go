// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleInsureQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VehicleInsureQueryResponseBody
	GetCode() *string
	SetMessage(v string) *VehicleInsureQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *VehicleInsureQueryResponseBody
	GetRequestId() *string
	SetResultObject(v *VehicleInsureQueryResponseBodyResultObject) *VehicleInsureQueryResponseBody
	GetResultObject() *VehicleInsureQueryResponseBodyResultObject
}

type VehicleInsureQueryResponseBody struct {
	// Return code: 200 indicates success, others indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message for the request information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F52********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *VehicleInsureQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleInsureQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VehicleInsureQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *VehicleInsureQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VehicleInsureQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VehicleInsureQueryResponseBody) GetResultObject() *VehicleInsureQueryResponseBodyResultObject {
	return s.ResultObject
}

func (s *VehicleInsureQueryResponseBody) SetCode(v string) *VehicleInsureQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetMessage(v string) *VehicleInsureQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetRequestId(v string) *VehicleInsureQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleInsureQueryResponseBody) SetResultObject(v *VehicleInsureQueryResponseBodyResultObject) *VehicleInsureQueryResponseBody {
	s.ResultObject = v
	return s
}

func (s *VehicleInsureQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type VehicleInsureQueryResponseBodyResultObject struct {
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
	// Insurance date information
	//
	// example:
	//
	// {
	//
	//     "firstInsuranceDate": "****-**-**",
	//
	//     "lastInsuranceDate": "****-**",
	//
	//     "latestInsuranceDate": "****-**",
	//
	//     "latestInsuranceDateStart": "****-**"
	//
	//   }
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s VehicleInsureQueryResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s VehicleInsureQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *VehicleInsureQueryResponseBodyResultObject) GetVehicleInfo() *string {
	return s.VehicleInfo
}

func (s *VehicleInsureQueryResponseBodyResultObject) SetBizCode(v string) *VehicleInsureQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleInsureQueryResponseBodyResultObject) SetVehicleInfo(v string) *VehicleInsureQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

func (s *VehicleInsureQueryResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
