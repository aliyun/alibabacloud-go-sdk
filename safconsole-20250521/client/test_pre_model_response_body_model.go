// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestPreModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TestPreModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *TestPreModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *TestPreModelResponseBody
	GetRequestId() *string
	SetResultObject(v []*bool) *TestPreModelResponseBody
	GetResultObject() []*bool
	SetSuccess(v bool) *TestPreModelResponseBody
	GetSuccess() *bool
}

type TestPreModelResponseBody struct {
	// Status code. A return of 200 indicates success.
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
	ResultObject []*bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Whether the call was successful.
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

func (s TestPreModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestPreModelResponseBody) GoString() string {
	return s.String()
}

func (s *TestPreModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TestPreModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *TestPreModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestPreModelResponseBody) GetResultObject() []*bool {
	return s.ResultObject
}

func (s *TestPreModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestPreModelResponseBody) SetCode(v int64) *TestPreModelResponseBody {
	s.Code = &v
	return s
}

func (s *TestPreModelResponseBody) SetHttpStatusCode(v int64) *TestPreModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TestPreModelResponseBody) SetRequestId(v string) *TestPreModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestPreModelResponseBody) SetResultObject(v []*bool) *TestPreModelResponseBody {
	s.ResultObject = v
	return s
}

func (s *TestPreModelResponseBody) SetSuccess(v bool) *TestPreModelResponseBody {
	s.Success = &v
	return s
}

func (s *TestPreModelResponseBody) Validate() error {
	return dara.Validate(s)
}
