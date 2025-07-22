// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCallbackMetaResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyCallbackMetaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyCallbackMetaResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCallbackMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCallbackMetaResponseBody
	GetSuccess() *bool
}

type ModifyCallbackMetaResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 0
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCallbackMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCallbackMetaResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCallbackMetaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCallbackMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCallbackMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCallbackMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCallbackMetaResponseBody) SetCode(v string) *ModifyCallbackMetaResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCallbackMetaResponseBody) SetHttpStatusCode(v int32) *ModifyCallbackMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCallbackMetaResponseBody) SetMessage(v string) *ModifyCallbackMetaResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCallbackMetaResponseBody) SetRequestId(v string) *ModifyCallbackMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCallbackMetaResponseBody) SetSuccess(v bool) *ModifyCallbackMetaResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCallbackMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
