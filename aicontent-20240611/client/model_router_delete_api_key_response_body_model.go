// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteApiKeyResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteApiKeyResponseBody struct {
	// example:
	//
	// true
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

func (s ModelRouterDeleteApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetData(v bool) *ModelRouterDeleteApiKeyResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetErrCode(v string) *ModelRouterDeleteApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetErrMessage(v string) *ModelRouterDeleteApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetRequestId(v string) *ModelRouterDeleteApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) SetSuccess(v bool) *ModelRouterDeleteApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
