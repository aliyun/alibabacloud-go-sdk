// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAuthorizationUrlResponseBody
	GetCode() *string
	SetData(v *GetAuthorizationUrlResponseBodyData) *GetAuthorizationUrlResponseBody
	GetData() *GetAuthorizationUrlResponseBodyData
	SetMessage(v string) *GetAuthorizationUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuthorizationUrlResponseBody
	GetRequestId() *string
}

type GetAuthorizationUrlResponseBody struct {
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
	Data *GetAuthorizationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAuthorizationUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuthorizationUrlResponseBody) GetData() *GetAuthorizationUrlResponseBodyData {
	return s.Data
}

func (s *GetAuthorizationUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuthorizationUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorizationUrlResponseBody) SetCode(v string) *GetAuthorizationUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetData(v *GetAuthorizationUrlResponseBodyData) *GetAuthorizationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetMessage(v string) *GetAuthorizationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetRequestId(v string) *GetAuthorizationUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuthorizationUrlResponseBodyData struct {
	// The authorization URL.
	//
	// example:
	//
	// https://render.****.com/p/s/web-call-minapp/auth-bao?page=commauth/index&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJiaXpUeXBlIjoiQ29tbUF1dGgiLCJiaXpJZCI6IjVmNWZjNjAzZDQzMTQ0MWZiYTZiNjYzM2QyMjIyNzU0IiwiZXhwIjoxNjA4MTkxODQxfQ.5IvBj2nKgr60APtotaIB13vtPVrdsPQ6avIfWxte1pA&_env=prod
	AuthorizationUrl *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
}

func (s GetAuthorizationUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBodyData) GetAuthorizationUrl() *string {
	return s.AuthorizationUrl
}

func (s *GetAuthorizationUrlResponseBodyData) SetAuthorizationUrl(v string) *GetAuthorizationUrlResponseBodyData {
	s.AuthorizationUrl = &v
	return s
}

func (s *GetAuthorizationUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
