// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateCustomerModuleBasicInfoResponseBody
	GetCode() *int64
	SetRequestId(v string) *CreateCustomerModuleBasicInfoResponseBody
	GetRequestId() *string
	SetResultObject(v string) *CreateCustomerModuleBasicInfoResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *CreateCustomerModuleBasicInfoResponseBody
	GetSuccess() *bool
}

type CreateCustomerModuleBasicInfoResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	//
	// example:
	//
	// 1
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerModuleBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleBasicInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateCustomerModuleBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomerModuleBasicInfoResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *CreateCustomerModuleBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomerModuleBasicInfoResponseBody) SetCode(v int64) *CreateCustomerModuleBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponseBody) SetRequestId(v string) *CreateCustomerModuleBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponseBody) SetResultObject(v string) *CreateCustomerModuleBasicInfoResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponseBody) SetSuccess(v bool) *CreateCustomerModuleBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
