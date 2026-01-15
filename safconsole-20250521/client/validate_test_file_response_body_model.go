// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTestFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ValidateTestFileResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *ValidateTestFileResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ValidateTestFileResponseBody
	GetRequestId() *string
	SetResultObject(v *ValidateTestFileResponseBodyResultObject) *ValidateTestFileResponseBody
	GetResultObject() *ValidateTestFileResponseBodyResultObject
	SetSuccess(v bool) *ValidateTestFileResponseBody
	GetSuccess() *bool
}

type ValidateTestFileResponseBody struct {
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
	ResultObject *ValidateTestFileResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
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

func (s ValidateTestFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateTestFileResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTestFileResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ValidateTestFileResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ValidateTestFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateTestFileResponseBody) GetResultObject() *ValidateTestFileResponseBodyResultObject {
	return s.ResultObject
}

func (s *ValidateTestFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateTestFileResponseBody) SetCode(v int64) *ValidateTestFileResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateTestFileResponseBody) SetHttpStatusCode(v int64) *ValidateTestFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ValidateTestFileResponseBody) SetRequestId(v string) *ValidateTestFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTestFileResponseBody) SetResultObject(v *ValidateTestFileResponseBodyResultObject) *ValidateTestFileResponseBody {
	s.ResultObject = v
	return s
}

func (s *ValidateTestFileResponseBody) SetSuccess(v bool) *ValidateTestFileResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateTestFileResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateTestFileResponseBodyResultObject struct {
	// Whether the test file is valid.
	FileValid *bool `json:"FileValid,omitempty" xml:"FileValid,omitempty"`
}

func (s ValidateTestFileResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ValidateTestFileResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ValidateTestFileResponseBodyResultObject) GetFileValid() *bool {
	return s.FileValid
}

func (s *ValidateTestFileResponseBodyResultObject) SetFileValid(v bool) *ValidateTestFileResponseBodyResultObject {
	s.FileValid = &v
	return s
}

func (s *ValidateTestFileResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
