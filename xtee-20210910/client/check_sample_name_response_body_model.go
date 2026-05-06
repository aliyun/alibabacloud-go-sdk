// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckSampleNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CheckSampleNameResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CheckSampleNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckSampleNameResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckSampleNameResponseBody
	GetResultObject() *bool
}

type CheckSampleNameResponseBody struct {
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

func (s CheckSampleNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSampleNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckSampleNameResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CheckSampleNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckSampleNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSampleNameResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckSampleNameResponseBody) SetCode(v string) *CheckSampleNameResponseBody {
	s.Code = &v
	return s
}

func (s *CheckSampleNameResponseBody) SetHttpStatusCode(v string) *CheckSampleNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckSampleNameResponseBody) SetMessage(v string) *CheckSampleNameResponseBody {
	s.Message = &v
	return s
}

func (s *CheckSampleNameResponseBody) SetRequestId(v string) *CheckSampleNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSampleNameResponseBody) SetResultObject(v bool) *CheckSampleNameResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckSampleNameResponseBody) Validate() error {
	return dara.Validate(s)
}
