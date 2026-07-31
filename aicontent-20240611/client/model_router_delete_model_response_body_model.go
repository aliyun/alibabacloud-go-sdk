// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteModelResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteModelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteModelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteModelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteModelResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteModelResponseBody struct {
	// Indicates whether the model was deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault message code.
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

func (s ModelRouterDeleteModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteModelResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteModelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteModelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteModelResponseBody) SetData(v bool) *ModelRouterDeleteModelResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) SetErrCode(v string) *ModelRouterDeleteModelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) SetErrMessage(v string) *ModelRouterDeleteModelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) SetRequestId(v string) *ModelRouterDeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) SetSuccess(v bool) *ModelRouterDeleteModelResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteModelResponseBody) Validate() error {
	return dara.Validate(s)
}
