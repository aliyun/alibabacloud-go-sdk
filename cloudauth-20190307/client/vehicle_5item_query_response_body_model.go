// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicle5ItemQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Vehicle5ItemQueryResponseBody
	GetCode() *string
	SetMessage(v string) *Vehicle5ItemQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *Vehicle5ItemQueryResponseBody
	GetRequestId() *string
	SetResultObject(v *Vehicle5ItemQueryResponseBodyResultObject) *Vehicle5ItemQueryResponseBody
	GetResultObject() *Vehicle5ItemQueryResponseBodyResultObject
}

type Vehicle5ItemQueryResponseBody struct {
	// Return code
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
	// Request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *Vehicle5ItemQueryResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Vehicle5ItemQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Vehicle5ItemQueryResponseBody) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *Vehicle5ItemQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Vehicle5ItemQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Vehicle5ItemQueryResponseBody) GetResultObject() *Vehicle5ItemQueryResponseBodyResultObject {
	return s.ResultObject
}

func (s *Vehicle5ItemQueryResponseBody) SetCode(v string) *Vehicle5ItemQueryResponseBody {
	s.Code = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetMessage(v string) *Vehicle5ItemQueryResponseBody {
	s.Message = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetRequestId(v string) *Vehicle5ItemQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) SetResultObject(v *Vehicle5ItemQueryResponseBodyResultObject) *Vehicle5ItemQueryResponseBody {
	s.ResultObject = v
	return s
}

func (s *Vehicle5ItemQueryResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Vehicle5ItemQueryResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: Found (charged)
	//
	// - **3**: No record found (not charged)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Vehicle information
	//
	// example:
	//
	// {
	//
	//         "engineNo": "F0******",
	//
	//         "useProperty": 0,
	//
	//         "registrationDate": "****-**-*	- **:**:**",
	//
	//         "engineType": "B******",
	//
	//         "vin": "A********",
	//
	//         "state": "0",
	//
	//         "modelNo": "B********",
	//
	//         "type": "小型轿车",
	//
	//         "brand": "宝马",
	//
	//         "usePropertyDesc": "非营运",
	//
	//         "stateDesc": "正常"
	//
	//     }
	VehicleInfo *string `json:"VehicleInfo,omitempty" xml:"VehicleInfo,omitempty"`
}

func (s Vehicle5ItemQueryResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Vehicle5ItemQueryResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) GetVehicleInfo() *string {
	return s.VehicleInfo
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) SetBizCode(v string) *Vehicle5ItemQueryResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) SetVehicleInfo(v string) *Vehicle5ItemQueryResponseBodyResultObject {
	s.VehicleInfo = &v
	return s
}

func (s *Vehicle5ItemQueryResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
