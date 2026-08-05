// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCredentialsResponseBody
	GetRequestId() *string
	SetResult(v *UpdateCredentialsResponseBodyResult) *UpdateCredentialsResponseBody
	GetResult() *UpdateCredentialsResponseBodyResult
}

type UpdateCredentialsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *UpdateCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialsResponseBody) GetResult() *UpdateCredentialsResponseBodyResult {
	return s.Result
}

func (s *UpdateCredentialsResponseBody) SetRequestId(v string) *UpdateCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialsResponseBody) SetResult(v *UpdateCredentialsResponseBodyResult) *UpdateCredentialsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateCredentialsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialsResponseBodyResult struct {
	// The workspace ID.
	//
	// example:
	//
	// 12321321
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// Specifies whether the credential is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access credential token.
	//
	// example:
	//
	// OS-****
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The credential type. Valid values:
	//
	// - api-token.
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCredentialsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *UpdateCredentialsResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCredentialsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *UpdateCredentialsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateCredentialsResponseBodyResult) SetAppGroupId(v int64) *UpdateCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *UpdateCredentialsResponseBodyResult) SetEnabled(v bool) *UpdateCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *UpdateCredentialsResponseBodyResult) SetToken(v string) *UpdateCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *UpdateCredentialsResponseBodyResult) SetType(v string) *UpdateCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateCredentialsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
