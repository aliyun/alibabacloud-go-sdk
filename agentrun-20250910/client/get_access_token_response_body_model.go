// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAccessTokenResponseBody
	GetCode() *string
	SetData(v *GetAccessTokenResponseBodyData) *GetAccessTokenResponseBody
	GetData() *GetAccessTokenResponseBodyData
	SetRequestId(v string) *GetAccessTokenResponseBody
	GetRequestId() *string
}

type GetAccessTokenResponseBody struct {
	// The status code of the response. A value of `SUCCESS` indicates that the request was successful. For failures, this parameter returns an error code, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetAccessTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The unique request ID. This ID is used for troubleshooting purposes.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccessTokenResponseBody) GetData() *GetAccessTokenResponseBodyData {
	return s.Data
}

func (s *GetAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessTokenResponseBody) SetCode(v string) *GetAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessTokenResponseBody) SetData(v *GetAccessTokenResponseBodyData) *GetAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetAccessTokenResponseBody) SetRequestId(v string) *GetAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessTokenResponseBodyData struct {
	// The personal access token. You can use this token to authenticate subsequent API calls. This authentication method is an alternative to using an Alibaba Cloud AccessKey pair or temporary credentials provided by the Security Token Service (STS).
	//
	// example:
	//
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
}

func (s GetAccessTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccessTokenResponseBodyData) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetAccessTokenResponseBodyData) SetAccessToken(v string) *GetAccessTokenResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *GetAccessTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
