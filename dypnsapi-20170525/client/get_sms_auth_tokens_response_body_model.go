// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsAuthTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmsAuthTokensResponseBody
	GetCode() *string
	SetData(v *GetSmsAuthTokensResponseBodyData) *GetSmsAuthTokensResponseBody
	GetData() *GetSmsAuthTokensResponseBodyData
	SetMessage(v string) *GetSmsAuthTokensResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmsAuthTokensResponseBody
	GetRequestId() *string
}

type GetSmsAuthTokensResponseBody struct {
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
	// The response parameters.
	Data *GetSmsAuthTokensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
}

func (s GetSmsAuthTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsAuthTokensResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsAuthTokensResponseBody) GetData() *GetSmsAuthTokensResponseBodyData {
	return s.Data
}

func (s *GetSmsAuthTokensResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsAuthTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsAuthTokensResponseBody) SetCode(v string) *GetSmsAuthTokensResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetData(v *GetSmsAuthTokensResponseBodyData) *GetSmsAuthTokensResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetMessage(v string) *GetSmsAuthTokensResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetRequestId(v string) *GetSmsAuthTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSmsAuthTokensResponseBodyData struct {
	// The business token.
	//
	// example:
	//
	// FZSMeRbLCiapIBo65NXNHvGbkGDmhs23XWTZDOtZN0g5n/kqSc2FU27Gc9YhGb6dNn9/L9ZXSYiIB6C6LMLQJjyXjRzt5v6pzZXqnjO4cSuPWYUxJvdc8l8OpucEYe8Mx17HxsHDzj0VC4D5+atcrTpJE6jQ7e2QVNjZIPMwsfxELjQS7c****
	BizToken *string `json:"BizToken,omitempty" xml:"BizToken,omitempty"`
	// The time when the token expired. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1631526326000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NSqC****
	StsAccessKeyId *string `json:"StsAccessKeyId,omitempty" xml:"StsAccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// VboZ4xbZ****
	StsAccessKeySecret *string `json:"StsAccessKeySecret,omitempty" xml:"StsAccessKeySecret,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAISiQJ1q6Ft5B2yfSjIr5DEDP/BurtW9PemMEfBrEpsOr5K17XjuDz2IHtLfXFsBusYt/U2nWpX5v4clrxIToR7SFbFY9pb6ZhazBisebDGv8HtR3TcFEjiSwapEBfe8JL4QYeQFaHwGJqEb1TDiVUAo9/TfimjWFqIKICAjYUdAP0cQgi/a0gtZr4UXHwAzvUXLnzML/2gHwf3i27LdipStxF7lHl05NbUoKTeyGKH3AGqlLVF9tite8f9NpczBvolDYfpht4RX7HazStd5yJN8KpLl6Fe8V/FxIrGXAAJv0rdbbOFq4Q1c18hOLJHAKtfsvXmlPNpsevfmpnsx****
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s GetSmsAuthTokensResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmsAuthTokensResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBodyData) GetBizToken() *string {
	return s.BizToken
}

func (s *GetSmsAuthTokensResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetSmsAuthTokensResponseBodyData) GetStsAccessKeyId() *string {
	return s.StsAccessKeyId
}

func (s *GetSmsAuthTokensResponseBodyData) GetStsAccessKeySecret() *string {
	return s.StsAccessKeySecret
}

func (s *GetSmsAuthTokensResponseBodyData) GetStsToken() *string {
	return s.StsToken
}

func (s *GetSmsAuthTokensResponseBodyData) SetBizToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.BizToken = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetExpireTime(v int64) *GetSmsAuthTokensResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeyId(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeyId = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeySecret(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeySecret = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsToken = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) Validate() error {
	return dara.Validate(s)
}
