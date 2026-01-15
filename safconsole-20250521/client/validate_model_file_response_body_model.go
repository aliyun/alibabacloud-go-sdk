// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModelFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ValidateModelFileResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *ValidateModelFileResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ValidateModelFileResponseBody
	GetRequestId() *string
	SetResultObject(v *ValidateModelFileResponseBodyResultObject) *ValidateModelFileResponseBody
	GetResultObject() *ValidateModelFileResponseBodyResultObject
	SetSuccess(v bool) *ValidateModelFileResponseBody
	GetSuccess() *bool
}

type ValidateModelFileResponseBody struct {
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
	// Returned result.
	ResultObject *ValidateModelFileResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
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

func (s ValidateModelFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateModelFileResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateModelFileResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ValidateModelFileResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ValidateModelFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateModelFileResponseBody) GetResultObject() *ValidateModelFileResponseBodyResultObject {
	return s.ResultObject
}

func (s *ValidateModelFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateModelFileResponseBody) SetCode(v int64) *ValidateModelFileResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateModelFileResponseBody) SetHttpStatusCode(v int64) *ValidateModelFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ValidateModelFileResponseBody) SetRequestId(v string) *ValidateModelFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateModelFileResponseBody) SetResultObject(v *ValidateModelFileResponseBodyResultObject) *ValidateModelFileResponseBody {
	s.ResultObject = v
	return s
}

func (s *ValidateModelFileResponseBody) SetSuccess(v bool) *ValidateModelFileResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateModelFileResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateModelFileResponseBodyResultObject struct {
	// Validation result.
	//
	// example:
	//
	// true
	FileValid *bool `json:"FileValid,omitempty" xml:"FileValid,omitempty"`
}

func (s ValidateModelFileResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ValidateModelFileResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ValidateModelFileResponseBodyResultObject) GetFileValid() *bool {
	return s.FileValid
}

func (s *ValidateModelFileResponseBodyResultObject) SetFileValid(v bool) *ValidateModelFileResponseBodyResultObject {
	s.FileValid = &v
	return s
}

func (s *ValidateModelFileResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
