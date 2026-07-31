// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteModelGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteModelGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteModelGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteModelGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteModelGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteModelGroupResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteModelGroupResponseBody struct {
	// The response struct.
	//
	// example:
	//
	// []
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The fault error message encoding.
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
	// 0A5E9849-A2F0-551D-A7D8-1A8118557BAB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterDeleteModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteModelGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetData(v bool) *ModelRouterDeleteModelGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetErrCode(v string) *ModelRouterDeleteModelGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetErrMessage(v string) *ModelRouterDeleteModelGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteModelGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetRequestId(v string) *ModelRouterDeleteModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) SetSuccess(v bool) *ModelRouterDeleteModelGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteModelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
