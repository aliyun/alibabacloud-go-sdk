// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAccessTokenResponseBody
	GetAccessDeniedDetail() *string
	SetAccessToken(v string) *GetAccessTokenResponseBody
	GetAccessToken() *string
	SetCode(v string) *GetAccessTokenResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAccessTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAccessTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccessTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAccessTokenResponseBody
	GetSuccess() *bool
}

type GetAccessTokenResponseBody struct {
	// example:
	//
	// null
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// "eyJhbGc****.eyJ********.****TCk"
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// "EA12****-****-****-****-****E5C"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAccessTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetAccessTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccessTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAccessTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAccessTokenResponseBody) SetAccessDeniedDetail(v string) *GetAccessTokenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetAccessToken(v string) *GetAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetCode(v string) *GetAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetHttpStatusCode(v int32) *GetAccessTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetMessage(v string) *GetAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetRequestId(v string) *GetAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetSuccess(v bool) *GetAccessTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
