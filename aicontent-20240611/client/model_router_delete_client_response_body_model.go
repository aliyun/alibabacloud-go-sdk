// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteClientResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteClientResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteClientResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteClientResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteClientResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteClientResponseBody struct {
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

func (s ModelRouterDeleteClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteClientResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteClientResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteClientResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteClientResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteClientResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteClientResponseBody) SetData(v bool) *ModelRouterDeleteClientResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) SetErrCode(v string) *ModelRouterDeleteClientResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) SetErrMessage(v string) *ModelRouterDeleteClientResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteClientResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) SetRequestId(v string) *ModelRouterDeleteClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) SetSuccess(v bool) *ModelRouterDeleteClientResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteClientResponseBody) Validate() error {
	return dara.Validate(s)
}
