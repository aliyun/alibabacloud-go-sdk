// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCredentialsResponseBody
	GetRequestId() *string
	SetResult(v *CreateCredentialsResponseBodyResult) *CreateCredentialsResponseBody
	GetResult() *CreateCredentialsResponseBodyResult
}

type CreateCredentialsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2E37A447-C010-5A49-9F31-DE12E97710A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response result.
	Result *CreateCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialsResponseBody) GetResult() *CreateCredentialsResponseBodyResult {
	return s.Result
}

func (s *CreateCredentialsResponseBody) SetRequestId(v string) *CreateCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialsResponseBody) SetResult(v *CreateCredentialsResponseBodyResult) *CreateCredentialsResponseBody {
	s.Result = v
	return s
}

func (s *CreateCredentialsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialsResponseBodyResult struct {
	// The workspace ID.
	//
	// example:
	//
	// 12323
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// Indicates whether the credential is enabled.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The generated credential token.
	//
	// example:
	//
	// OS-********
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The credential type.
	//
	// - api-token
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCredentialsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCredentialsResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *CreateCredentialsResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCredentialsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *CreateCredentialsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateCredentialsResponseBodyResult) SetAppGroupId(v int64) *CreateCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *CreateCredentialsResponseBodyResult) SetEnabled(v bool) *CreateCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *CreateCredentialsResponseBodyResult) SetToken(v string) *CreateCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *CreateCredentialsResponseBodyResult) SetType(v string) *CreateCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateCredentialsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
