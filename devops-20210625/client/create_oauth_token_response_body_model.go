// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateOAuthTokenResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateOAuthTokenResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateOAuthTokenResponseBody
	GetRequestId() *string
	SetResult(v *CreateOAuthTokenResponseBodyResult) *CreateOAuthTokenResponseBody
	GetResult() *CreateOAuthTokenResponseBodyResult
	SetSuccess(v string) *CreateOAuthTokenResponseBody
	GetSuccess() *string
}

type CreateOAuthTokenResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// C2F153F6-BB43-50C4-9F4F-40593203E19A
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateOAuthTokenResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateOAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOAuthTokenResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateOAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOAuthTokenResponseBody) GetResult() *CreateOAuthTokenResponseBodyResult {
	return s.Result
}

func (s *CreateOAuthTokenResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateOAuthTokenResponseBody) SetErrorCode(v string) *CreateOAuthTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetErrorMessage(v string) *CreateOAuthTokenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetRequestId(v string) *CreateOAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetResult(v *CreateOAuthTokenResponseBodyResult) *CreateOAuthTokenResponseBody {
	s.Result = v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetSuccess(v string) *CreateOAuthTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOAuthTokenResponseBodyResult struct {
	// example:
	//
	// 2aeb4cd012af879e54f0d37dfa526f51
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// id
	//
	// example:
	//
	// 30815
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// read:repo
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// code
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s CreateOAuthTokenResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuthTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponseBodyResult) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateOAuthTokenResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateOAuthTokenResponseBodyResult) GetScope() *string {
	return s.Scope
}

func (s *CreateOAuthTokenResponseBodyResult) GetTokenType() *string {
	return s.TokenType
}

func (s *CreateOAuthTokenResponseBodyResult) SetAccessToken(v string) *CreateOAuthTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetId(v string) *CreateOAuthTokenResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetScope(v string) *CreateOAuthTokenResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetTokenType(v string) *CreateOAuthTokenResponseBodyResult {
	s.TokenType = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
