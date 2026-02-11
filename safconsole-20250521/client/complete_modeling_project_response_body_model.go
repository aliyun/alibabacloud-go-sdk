// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteModelingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CompleteModelingProjectResponseBody
	GetCode() *int64
	SetRequestId(v string) *CompleteModelingProjectResponseBody
	GetRequestId() *string
	SetResultObject(v string) *CompleteModelingProjectResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *CompleteModelingProjectResponseBody
	GetSuccess() *bool
}

type CompleteModelingProjectResponseBody struct {
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
	// 055f1546-f465-4c92-a2da-bfb86abe6f56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result, indicating whether the creation was successful:
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
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

func (s CompleteModelingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteModelingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteModelingProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CompleteModelingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteModelingProjectResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *CompleteModelingProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CompleteModelingProjectResponseBody) SetCode(v int64) *CompleteModelingProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CompleteModelingProjectResponseBody) SetRequestId(v string) *CompleteModelingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteModelingProjectResponseBody) SetResultObject(v string) *CompleteModelingProjectResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CompleteModelingProjectResponseBody) SetSuccess(v bool) *CompleteModelingProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CompleteModelingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
