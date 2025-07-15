// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMessageTokenResponseBody
	GetRequestId() *string
	SetResult(v *GetMessageTokenResponseBodyResult) *GetMessageTokenResponseBody
	GetResult() *GetMessageTokenResponseBodyResult
}

type GetMessageTokenResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *GetMessageTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMessageTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageTokenResponseBody) GetResult() *GetMessageTokenResponseBodyResult {
	return s.Result
}

func (s *GetMessageTokenResponseBody) SetRequestId(v string) *GetMessageTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageTokenResponseBody) SetResult(v *GetMessageTokenResponseBodyResult) *GetMessageTokenResponseBody {
	s.Result = v
	return s
}

func (s *GetMessageTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMessageTokenResponseBodyResult struct {
	// The token used to establish a persistent connection.
	//
	// example:
	//
	// oauth_cloud_key:***-b0YY5Gy6Q
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// Indicates how long until the token expires. Unit: milliseconds.
	//
	// example:
	//
	// 86400000
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// The updated token. If a token expires, you can call RefreshToken to obtain a new token.
	//
	// example:
	//
	// oauth_cloud_key:****-Q62xggOTdgk3gw=
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s GetMessageTokenResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMessageTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMessageTokenResponseBodyResult) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetMessageTokenResponseBodyResult) GetAccessTokenExpiredTime() *int64 {
	return s.AccessTokenExpiredTime
}

func (s *GetMessageTokenResponseBodyResult) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GetMessageTokenResponseBodyResult) SetAccessToken(v string) *GetMessageTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *GetMessageTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetMessageTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetMessageTokenResponseBodyResult) SetRefreshToken(v string) *GetMessageTokenResponseBodyResult {
	s.RefreshToken = &v
	return s
}

func (s *GetMessageTokenResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
