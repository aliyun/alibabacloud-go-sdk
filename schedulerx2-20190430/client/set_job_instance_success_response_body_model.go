// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobInstanceSuccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SetJobInstanceSuccessResponseBody
	GetCode() *int32
	SetMessage(v string) *SetJobInstanceSuccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetJobInstanceSuccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetJobInstanceSuccessResponseBody
	GetSuccess() *bool
}

type SetJobInstanceSuccessResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetJobInstanceSuccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetJobInstanceSuccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SetJobInstanceSuccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetJobInstanceSuccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetJobInstanceSuccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetJobInstanceSuccessResponseBody) SetCode(v int32) *SetJobInstanceSuccessResponseBody {
	s.Code = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetMessage(v string) *SetJobInstanceSuccessResponseBody {
	s.Message = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetRequestId(v string) *SetJobInstanceSuccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetSuccess(v bool) *SetJobInstanceSuccessResponseBody {
	s.Success = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) Validate() error {
	return dara.Validate(s)
}
