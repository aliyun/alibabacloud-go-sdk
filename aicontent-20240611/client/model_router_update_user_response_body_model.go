// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterUpdateUserResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterUpdateUserResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateUserResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateUserResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateUserResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateUserResponseBody struct {
	// The data object. This parameter contains no business data and is always empty.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterUpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterUpdateUserResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateUserResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateUserResponseBody) SetData(v bool) *ModelRouterUpdateUserResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) SetErrCode(v string) *ModelRouterUpdateUserResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) SetErrMessage(v string) *ModelRouterUpdateUserResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) SetRequestId(v string) *ModelRouterUpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) SetSuccess(v bool) *ModelRouterUpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
