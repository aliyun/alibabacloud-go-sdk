// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestModelFileSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *RequestModelFileSyncResponseBody
	GetCode() *int64
	SetRequestId(v string) *RequestModelFileSyncResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *RequestModelFileSyncResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *RequestModelFileSyncResponseBody
	GetSuccess() *bool
}

type RequestModelFileSyncResponseBody struct {
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

func (s RequestModelFileSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RequestModelFileSyncResponseBody) GoString() string {
	return s.String()
}

func (s *RequestModelFileSyncResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RequestModelFileSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RequestModelFileSyncResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *RequestModelFileSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RequestModelFileSyncResponseBody) SetCode(v int64) *RequestModelFileSyncResponseBody {
	s.Code = &v
	return s
}

func (s *RequestModelFileSyncResponseBody) SetRequestId(v string) *RequestModelFileSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestModelFileSyncResponseBody) SetResultObject(v bool) *RequestModelFileSyncResponseBody {
	s.ResultObject = &v
	return s
}

func (s *RequestModelFileSyncResponseBody) SetSuccess(v bool) *RequestModelFileSyncResponseBody {
	s.Success = &v
	return s
}

func (s *RequestModelFileSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
