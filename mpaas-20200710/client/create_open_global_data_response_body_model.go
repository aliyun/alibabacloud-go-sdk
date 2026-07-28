// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenGlobalDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateOpenGlobalDataResponseBody
	GetAccessDeniedDetail() *string
	SetHttpStatusCode(v int32) *CreateOpenGlobalDataResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateOpenGlobalDataResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateOpenGlobalDataResponseBody
	GetResult() *bool
	SetResultCode(v string) *CreateOpenGlobalDataResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateOpenGlobalDataResponseBody
	GetResultMessage() *string
}

type CreateOpenGlobalDataResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result             *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateOpenGlobalDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenGlobalDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenGlobalDataResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateOpenGlobalDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateOpenGlobalDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenGlobalDataResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateOpenGlobalDataResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateOpenGlobalDataResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateOpenGlobalDataResponseBody) SetAccessDeniedDetail(v string) *CreateOpenGlobalDataResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetHttpStatusCode(v int32) *CreateOpenGlobalDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetRequestId(v string) *CreateOpenGlobalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResult(v bool) *CreateOpenGlobalDataResponseBody {
	s.Result = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResultCode(v string) *CreateOpenGlobalDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) SetResultMessage(v string) *CreateOpenGlobalDataResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateOpenGlobalDataResponseBody) Validate() error {
	return dara.Validate(s)
}
