// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateJobScriptResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateJobScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJobScriptResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateJobScriptResponseBody
	GetSuccess() *bool
}

type UpdateJobScriptResponseBody struct {
	// The returned code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned only if an error occurs.
	//
	// example:
	//
	// job is not existed, jobId=302
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateJobScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobScriptResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateJobScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJobScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateJobScriptResponseBody) SetCode(v int32) *UpdateJobScriptResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetMessage(v string) *UpdateJobScriptResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetRequestId(v string) *UpdateJobScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobScriptResponseBody) SetSuccess(v bool) *UpdateJobScriptResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateJobScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
