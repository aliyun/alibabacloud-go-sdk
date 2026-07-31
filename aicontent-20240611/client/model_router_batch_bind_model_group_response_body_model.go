// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchBindModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterBatchBindModelGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterBatchBindModelGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchBindModelGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchBindModelGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchBindModelGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchBindModelGroupResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchBindModelGroupResponseBody struct {
	// The data object.
	//
	// example:
	//
	// []
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The error message code.
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

func (s ModelRouterBatchBindModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchBindModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchBindModelGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetData(v bool) *ModelRouterBatchBindModelGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetErrCode(v string) *ModelRouterBatchBindModelGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetErrMessage(v string) *ModelRouterBatchBindModelGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchBindModelGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetRequestId(v string) *ModelRouterBatchBindModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) SetSuccess(v bool) *ModelRouterBatchBindModelGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchBindModelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
