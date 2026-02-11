// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestProcessExpressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TestProcessExpressionResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *TestProcessExpressionResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *TestProcessExpressionResponseBody
	GetRequestId() *string
	SetResultObject(v *TestProcessExpressionResponseBodyResultObject) *TestProcessExpressionResponseBody
	GetResultObject() *TestProcessExpressionResponseBodyResultObject
	SetSuccess(v bool) *TestProcessExpressionResponseBody
	GetSuccess() *bool
}

type TestProcessExpressionResponseBody struct {
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
	ResultObject *TestProcessExpressionResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
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

func (s TestProcessExpressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestProcessExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *TestProcessExpressionResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TestProcessExpressionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *TestProcessExpressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestProcessExpressionResponseBody) GetResultObject() *TestProcessExpressionResponseBodyResultObject {
	return s.ResultObject
}

func (s *TestProcessExpressionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestProcessExpressionResponseBody) SetCode(v int64) *TestProcessExpressionResponseBody {
	s.Code = &v
	return s
}

func (s *TestProcessExpressionResponseBody) SetHttpStatusCode(v int64) *TestProcessExpressionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TestProcessExpressionResponseBody) SetRequestId(v string) *TestProcessExpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestProcessExpressionResponseBody) SetResultObject(v *TestProcessExpressionResponseBodyResultObject) *TestProcessExpressionResponseBody {
	s.ResultObject = v
	return s
}

func (s *TestProcessExpressionResponseBody) SetSuccess(v bool) *TestProcessExpressionResponseBody {
	s.Success = &v
	return s
}

func (s *TestProcessExpressionResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestProcessExpressionResponseBodyResultObject struct {
	// Test score.
	//
	// example:
	//
	// 10
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s TestProcessExpressionResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s TestProcessExpressionResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *TestProcessExpressionResponseBodyResultObject) GetScore() *float64 {
	return s.Score
}

func (s *TestProcessExpressionResponseBodyResultObject) SetScore(v float64) *TestProcessExpressionResponseBodyResultObject {
	s.Score = &v
	return s
}

func (s *TestProcessExpressionResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
