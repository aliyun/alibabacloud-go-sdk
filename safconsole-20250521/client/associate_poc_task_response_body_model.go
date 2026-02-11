// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePocTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *AssociatePocTaskResponseBody
	GetCode() *int64
	SetRequestId(v string) *AssociatePocTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *AssociatePocTaskResponseBody
	GetResultObject() *bool
	SetSuccess(v bool) *AssociatePocTaskResponseBody
	GetSuccess() *bool
}

type AssociatePocTaskResponseBody struct {
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

func (s AssociatePocTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociatePocTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AssociatePocTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AssociatePocTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociatePocTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *AssociatePocTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssociatePocTaskResponseBody) SetCode(v int64) *AssociatePocTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AssociatePocTaskResponseBody) SetRequestId(v string) *AssociatePocTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociatePocTaskResponseBody) SetResultObject(v bool) *AssociatePocTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *AssociatePocTaskResponseBody) SetSuccess(v bool) *AssociatePocTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AssociatePocTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
