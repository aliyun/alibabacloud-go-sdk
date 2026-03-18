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
	// example:
	//
	// []
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
