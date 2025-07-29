// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWfInstanceSuccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SetWfInstanceSuccessResponseBody
	GetCode() *int32
	SetMessage(v string) *SetWfInstanceSuccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetWfInstanceSuccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetWfInstanceSuccessResponseBody
	GetSuccess() *bool
}

type SetWfInstanceSuccessResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// wofkflowId is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
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

func (s SetWfInstanceSuccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWfInstanceSuccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SetWfInstanceSuccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWfInstanceSuccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWfInstanceSuccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetWfInstanceSuccessResponseBody) SetCode(v int32) *SetWfInstanceSuccessResponseBody {
	s.Code = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetMessage(v string) *SetWfInstanceSuccessResponseBody {
	s.Message = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetRequestId(v string) *SetWfInstanceSuccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetSuccess(v bool) *SetWfInstanceSuccessResponseBody {
	s.Success = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) Validate() error {
	return dara.Validate(s)
}
