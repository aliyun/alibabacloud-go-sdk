// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *RollbackModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *RollbackModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *RollbackModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *RollbackModelResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *RollbackModelResponseBody
	GetSuccess() *bool
}

type RollbackModelResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP状态码。
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

func (s RollbackModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackModelResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RollbackModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RollbackModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *RollbackModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackModelResponseBody) SetCode(v int64) *RollbackModelResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackModelResponseBody) SetHttpStatusCode(v int64) *RollbackModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RollbackModelResponseBody) SetRequestId(v string) *RollbackModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackModelResponseBody) SetResultObject(v bool) *RollbackModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RollbackModelResponseBody) SetSuccess(v bool) *RollbackModelResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackModelResponseBody) Validate() error {
	return dara.Validate(s)
}
