// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateModuleBasicInfoResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *UpdateModuleBasicInfoResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *UpdateModuleBasicInfoResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *UpdateModuleBasicInfoResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *UpdateModuleBasicInfoResponseBody
	GetSuccess() *bool
}

type UpdateModuleBasicInfoResponseBody struct {
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
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModuleBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModuleBasicInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateModuleBasicInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateModuleBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModuleBasicInfoResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *UpdateModuleBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModuleBasicInfoResponseBody) SetCode(v int64) *UpdateModuleBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModuleBasicInfoResponseBody) SetHttpStatusCode(v int64) *UpdateModuleBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModuleBasicInfoResponseBody) SetRequestId(v string) *UpdateModuleBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModuleBasicInfoResponseBody) SetResultObject(v bool) *UpdateModuleBasicInfoResponseBody {
	s.ResultObject = &v
	return s
}

func (s *UpdateModuleBasicInfoResponseBody) SetSuccess(v bool) *UpdateModuleBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModuleBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
