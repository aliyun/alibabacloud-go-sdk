// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppGroupCredentialsResponseBody
	GetRequestId() *string
	SetResult(v *CreateAppGroupCredentialsResponseBodyResult) *CreateAppGroupCredentialsResponseBody
	GetResult() *CreateAppGroupCredentialsResponseBodyResult
}

type CreateAppGroupCredentialsResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// 1-2-3-4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the operation result.
	Result *CreateAppGroupCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppGroupCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppGroupCredentialsResponseBody) GetResult() *CreateAppGroupCredentialsResponseBodyResult {
	return s.Result
}

func (s *CreateAppGroupCredentialsResponseBody) SetRequestId(v string) *CreateAppGroupCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBody) SetResult(v *CreateAppGroupCredentialsResponseBodyResult) *CreateAppGroupCredentialsResponseBody {
	s.Result = v
	return s
}

func (s *CreateAppGroupCredentialsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppGroupCredentialsResponseBodyResult struct {
	// The ID of the application group.
	//
	// example:
	//
	// app_group_123
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// Indicates whether the credential is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The generated credential token.
	//
	// example:
	//
	// generated_token_string
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The credential type.
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppGroupCredentialsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *CreateAppGroupCredentialsResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAppGroupCredentialsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *CreateAppGroupCredentialsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetAppGroupId(v int64) *CreateAppGroupCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetEnabled(v bool) *CreateAppGroupCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetToken(v string) *CreateAppGroupCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetType(v string) *CreateAppGroupCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
