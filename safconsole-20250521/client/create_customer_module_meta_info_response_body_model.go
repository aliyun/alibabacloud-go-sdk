// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleMetaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateCustomerModuleMetaInfoResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *CreateCustomerModuleMetaInfoResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *CreateCustomerModuleMetaInfoResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateCustomerModuleMetaInfoResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *CreateCustomerModuleMetaInfoResponseBody
	GetSuccess() *bool
}

type CreateCustomerModuleMetaInfoResponseBody struct {
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
	// The result indicating if the creation was successful:
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Indicates whether the request was successful, with values as follows:
	//
	// - true, request succeeded
	//
	// - false, request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerModuleMetaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleMetaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleMetaInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateCustomerModuleMetaInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateCustomerModuleMetaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomerModuleMetaInfoResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateCustomerModuleMetaInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomerModuleMetaInfoResponseBody) SetCode(v int64) *CreateCustomerModuleMetaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponseBody) SetHttpStatusCode(v int64) *CreateCustomerModuleMetaInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponseBody) SetRequestId(v string) *CreateCustomerModuleMetaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponseBody) SetResultObject(v bool) *CreateCustomerModuleMetaInfoResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponseBody) SetSuccess(v bool) *CreateCustomerModuleMetaInfoResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
