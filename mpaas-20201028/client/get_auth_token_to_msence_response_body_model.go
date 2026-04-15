// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenToMsenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasSystemOauthTokenResponse(v *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) *GetAuthTokenToMsenceResponseBody
	GetMpaasSystemOauthTokenResponse() *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse
	SetRequestId(v string) *GetAuthTokenToMsenceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetAuthTokenToMsenceResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *GetAuthTokenToMsenceResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *GetAuthTokenToMsenceResponseBody
	GetSuccess() *bool
}

type GetAuthTokenToMsenceResponseBody struct {
	MpaasSystemOauthTokenResponse *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse `json:"MpaasSystemOauthTokenResponse,omitempty" xml:"MpaasSystemOauthTokenResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuthTokenToMsenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenToMsenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenToMsenceResponseBody) GetMpaasSystemOauthTokenResponse() *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	return s.MpaasSystemOauthTokenResponse
}

func (s *GetAuthTokenToMsenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthTokenToMsenceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetAuthTokenToMsenceResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetAuthTokenToMsenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuthTokenToMsenceResponseBody) SetMpaasSystemOauthTokenResponse(v *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) *GetAuthTokenToMsenceResponseBody {
	s.MpaasSystemOauthTokenResponse = v
	return s
}

func (s *GetAuthTokenToMsenceResponseBody) SetRequestId(v string) *GetAuthTokenToMsenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBody) SetResultCode(v string) *GetAuthTokenToMsenceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBody) SetResultMsg(v string) *GetAuthTokenToMsenceResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBody) SetSuccess(v bool) *GetAuthTokenToMsenceResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBody) Validate() error {
	if s.MpaasSystemOauthTokenResponse != nil {
		if err := s.MpaasSystemOauthTokenResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse struct {
	// example:
	//
	// NPHTGKNR
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// example:
	//
	// 300
	ExpiresIn *string `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// example:
	//
	// 123456
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	// example:
	//
	// -
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// userid001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GetExpiresIn() *string {
	return s.ExpiresIn
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GetOpenId() *string {
	return s.OpenId
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GetPlatform() *string {
	return s.Platform
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) GetUserId() *string {
	return s.UserId
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) SetAuthToken(v string) *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	s.AuthToken = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) SetExpiresIn(v string) *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	s.ExpiresIn = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) SetOpenId(v string) *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	s.OpenId = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) SetPlatform(v string) *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	s.Platform = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) SetUserId(v string) *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse {
	s.UserId = &v
	return s
}

func (s *GetAuthTokenToMsenceResponseBodyMpaasSystemOauthTokenResponse) Validate() error {
	return dara.Validate(s)
}
