// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteUserResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteUserResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteUserResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteUserResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteUserResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteUserResponseBody struct {
	// The data object. This object contains no business data and is always empty.
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

func (s ModelRouterDeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteUserResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteUserResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteUserResponseBody) SetData(v bool) *ModelRouterDeleteUserResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) SetErrCode(v string) *ModelRouterDeleteUserResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) SetErrMessage(v string) *ModelRouterDeleteUserResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) SetRequestId(v string) *ModelRouterDeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) SetSuccess(v bool) *ModelRouterDeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteUserResponseBody) Validate() error {
	return dara.Validate(s)
}
