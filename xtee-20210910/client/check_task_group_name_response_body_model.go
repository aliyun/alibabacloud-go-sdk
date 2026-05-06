// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTaskGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckTaskGroupNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CheckTaskGroupNameResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CheckTaskGroupNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckTaskGroupNameResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CheckTaskGroupNameResponseBody
	GetResultObject() *bool
}

type CheckTaskGroupNameResponseBody struct {
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

func (s CheckTaskGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTaskGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTaskGroupNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckTaskGroupNameResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CheckTaskGroupNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckTaskGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTaskGroupNameResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CheckTaskGroupNameResponseBody) SetCode(v string) *CheckTaskGroupNameResponseBody {
	s.Code = &v
	return s
}

func (s *CheckTaskGroupNameResponseBody) SetHttpStatusCode(v string) *CheckTaskGroupNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckTaskGroupNameResponseBody) SetMessage(v string) *CheckTaskGroupNameResponseBody {
	s.Message = &v
	return s
}

func (s *CheckTaskGroupNameResponseBody) SetRequestId(v string) *CheckTaskGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTaskGroupNameResponseBody) SetResultObject(v bool) *CheckTaskGroupNameResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CheckTaskGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}
