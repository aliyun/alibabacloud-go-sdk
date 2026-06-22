// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebofficeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GenerateWebofficeTokenResponseBody
	GetAccessToken() *string
	SetAccessTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody
	GetAccessTokenExpiredTime() *string
	SetRefreshToken(v string) *GenerateWebofficeTokenResponseBody
	GetRefreshToken() *string
	SetRefreshTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody
	GetRefreshTokenExpiredTime() *string
	SetRequestId(v string) *GenerateWebofficeTokenResponseBody
	GetRequestId() *string
	SetWebofficeURL(v string) *GenerateWebofficeTokenResponseBody
	GetWebofficeURL() *string
}

type GenerateWebofficeTokenResponseBody struct {
	// The Weboffice access credential.
	//
	// example:
	//
	// 2d73dd5d87524c5e8a194c3eb5********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The expiration time of the access credential. The credential expires in 30 minutes. Format: YYYY-MM-DDTHH:mm:ss.
	//
	// example:
	//
	// 2021-08-30T13:13:11.347146982Z
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// The Weboffice refresh credential.
	//
	// example:
	//
	// e374995ec532432bb678074d36********
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// The expiration time of the refresh credential. The credential expires in 1 day. Format: YYYY-MM-DDTHH:mm:ss.
	//
	// example:
	//
	// 2021-08-31T12:43:11.347146982Z
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1759315A-CB33-0A75-A72B-62D7********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Weboffice entry URL for previewing or editing documents online.
	//
	// > This URL cannot be opened directly in a browser. You must use it together with the Weboffice JS-SDK and the access credential (AccessToken) to preview or edit documents. For more information, see [Getting Started](https://help.aliyun.com/document_detail/468066.html).
	//
	// example:
	//
	// https://office-cn-shanghai.imm.aliyuncs.com/office/s/dd221b2cdb44fb66e9070d1d70a8b9bbb6d6fff7?_w_tokentype=1
	WebofficeURL *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
}

func (s GenerateWebofficeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GenerateWebofficeTokenResponseBody) GetAccessTokenExpiredTime() *string {
	return s.AccessTokenExpiredTime
}

func (s *GenerateWebofficeTokenResponseBody) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GenerateWebofficeTokenResponseBody) GetRefreshTokenExpiredTime() *string {
	return s.RefreshTokenExpiredTime
}

func (s *GenerateWebofficeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateWebofficeTokenResponseBody) GetWebofficeURL() *string {
	return s.WebofficeURL
}

func (s *GenerateWebofficeTokenResponseBody) SetAccessToken(v string) *GenerateWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRefreshToken(v string) *GenerateWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRequestId(v string) *GenerateWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetWebofficeURL(v string) *GenerateWebofficeTokenResponseBody {
	s.WebofficeURL = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
