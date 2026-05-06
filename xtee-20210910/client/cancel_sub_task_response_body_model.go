// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelSubTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CancelSubTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CancelSubTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelSubTaskResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CancelSubTaskResponseBody
	GetResultObject() *bool
}

type CancelSubTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CancelSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSubTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelSubTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CancelSubTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSubTaskResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CancelSubTaskResponseBody) SetCode(v string) *CancelSubTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelSubTaskResponseBody) SetHttpStatusCode(v string) *CancelSubTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelSubTaskResponseBody) SetMessage(v string) *CancelSubTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelSubTaskResponseBody) SetRequestId(v string) *CancelSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSubTaskResponseBody) SetResultObject(v bool) *CancelSubTaskResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CancelSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
