// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleOutputInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateCustomerModuleOutputInfoResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *CreateCustomerModuleOutputInfoResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *CreateCustomerModuleOutputInfoResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateCustomerModuleOutputInfoResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *CreateCustomerModuleOutputInfoResponseBody
	GetSuccess() *bool
}

type CreateCustomerModuleOutputInfoResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
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

func (s CreateCustomerModuleOutputInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleOutputInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleOutputInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateCustomerModuleOutputInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateCustomerModuleOutputInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomerModuleOutputInfoResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateCustomerModuleOutputInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomerModuleOutputInfoResponseBody) SetCode(v int64) *CreateCustomerModuleOutputInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponseBody) SetHttpStatusCode(v int64) *CreateCustomerModuleOutputInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponseBody) SetRequestId(v string) *CreateCustomerModuleOutputInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponseBody) SetResultObject(v bool) *CreateCustomerModuleOutputInfoResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponseBody) SetSuccess(v bool) *CreateCustomerModuleOutputInfoResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
