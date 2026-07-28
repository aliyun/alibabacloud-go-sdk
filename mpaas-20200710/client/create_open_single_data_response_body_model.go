// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSingleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateOpenSingleDataResponseBody
	GetAccessDeniedDetail() *string
	SetHttpStatusCode(v int32) *CreateOpenSingleDataResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateOpenSingleDataResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateOpenSingleDataResponseBody
	GetResult() *bool
	SetResultCode(v string) *CreateOpenSingleDataResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateOpenSingleDataResponseBody
	GetResultMessage() *string
}

type CreateOpenSingleDataResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result             *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultCode         *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateOpenSingleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSingleDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSingleDataResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateOpenSingleDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateOpenSingleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSingleDataResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateOpenSingleDataResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateOpenSingleDataResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateOpenSingleDataResponseBody) SetAccessDeniedDetail(v string) *CreateOpenSingleDataResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) SetHttpStatusCode(v int32) *CreateOpenSingleDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) SetRequestId(v string) *CreateOpenSingleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) SetResult(v bool) *CreateOpenSingleDataResponseBody {
	s.Result = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) SetResultCode(v string) *CreateOpenSingleDataResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) SetResultMessage(v string) *CreateOpenSingleDataResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateOpenSingleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
