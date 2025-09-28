// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAuthTokenResponseBody
	GetCode() *string
	SetMessage(v string) *GetAuthTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuthTokenResponseBody
	GetRequestId() *string
	SetTokenInfo(v *GetAuthTokenResponseBodyTokenInfo) *GetAuthTokenResponseBody
	GetTokenInfo() *GetAuthTokenResponseBodyTokenInfo
}

type GetAuthTokenResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	TokenInfo *GetAuthTokenResponseBodyTokenInfo `json:"TokenInfo,omitempty" xml:"TokenInfo,omitempty" type:"Struct"`
}

func (s GetAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuthTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthTokenResponseBody) GetTokenInfo() *GetAuthTokenResponseBodyTokenInfo {
	return s.TokenInfo
}

func (s *GetAuthTokenResponseBody) SetCode(v string) *GetAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetMessage(v string) *GetAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetTokenInfo(v *GetAuthTokenResponseBodyTokenInfo) *GetAuthTokenResponseBody {
	s.TokenInfo = v
	return s
}

func (s *GetAuthTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuthTokenResponseBodyTokenInfo struct {
	// The business authentication token.
	//
	// >  AccessToken is valid for 10 minutes and can be used repeatedly within its validity period.
	//
	// example:
	//
	// agag****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The API authentication token.
	//
	// >  JwtToken is valid for 1 hour and can be used repeatedly within its validity period.
	//
	// example:
	//
	// aweghd****
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetAuthTokenResponseBodyTokenInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenResponseBodyTokenInfo) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyTokenInfo) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetAuthTokenResponseBodyTokenInfo) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetAccessToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.AccessToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetJwtToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.JwtToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyTokenInfo) Validate() error {
	return dara.Validate(s)
}
