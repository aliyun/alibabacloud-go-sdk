// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VehicleMetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *VehicleMetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *VehicleMetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *VehicleMetaVerifyResponseBodyResultObject) *VehicleMetaVerifyResponseBody
	GetResultObject() *VehicleMetaVerifyResponseBodyResultObject
}

type VehicleMetaVerifyResponseBody struct {
	// Response code, **200*	- indicates that the API response was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF4*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *VehicleMetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s VehicleMetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *VehicleMetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VehicleMetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VehicleMetaVerifyResponseBody) GetResultObject() *VehicleMetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *VehicleMetaVerifyResponseBody) SetCode(v string) *VehicleMetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetMessage(v string) *VehicleMetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetRequestId(v string) *VehicleMetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *VehicleMetaVerifyResponseBody) SetResultObject(v *VehicleMetaVerifyResponseBodyResultObject) *VehicleMetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *VehicleMetaVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type VehicleMetaVerifyResponseBodyResultObject struct {
	// Verification result.
	//
	// - 1: Consistent (billable)
	//
	// - 2: Inconsistent (billable)
	//
	// - 3: No record found (non-billable)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s VehicleMetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *VehicleMetaVerifyResponseBodyResultObject) SetBizCode(v string) *VehicleMetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *VehicleMetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
