// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterUpdateModelGroupResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterUpdateModelGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateModelGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateModelGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateModelGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateModelGroupResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateModelGroupResponseBody struct {
	// The response struct.
	//
	// example:
	//
	// {}
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The error message code.
	//
	// example:
	//
	// null
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

func (s ModelRouterUpdateModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateModelGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetData(v bool) *ModelRouterUpdateModelGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetErrCode(v string) *ModelRouterUpdateModelGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetErrMessage(v string) *ModelRouterUpdateModelGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateModelGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetRequestId(v string) *ModelRouterUpdateModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) SetSuccess(v bool) *ModelRouterUpdateModelGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateModelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
