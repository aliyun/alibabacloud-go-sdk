// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteModelResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DeleteModelResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DeleteModelResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteModelResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *DeleteModelResponseBody
	GetSuccess() *bool
}

type DeleteModelResponseBody struct {
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
	// Result of the operation, indicating whether it was successful:
	//
	// - true: Success
	//
	// - false: Failure
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

func (s DeleteModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteModelResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DeleteModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelResponseBody) SetCode(v int64) *DeleteModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelResponseBody) SetHttpStatusCode(v int64) *DeleteModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetResultObject(v bool) *DeleteModelResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteModelResponseBody) SetSuccess(v bool) *DeleteModelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelResponseBody) Validate() error {
	return dara.Validate(s)
}
