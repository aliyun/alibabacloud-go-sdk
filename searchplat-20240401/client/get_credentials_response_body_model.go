// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCredentialsResponseBody
	GetRequestId() *string
	SetResult(v *GetCredentialsResponseBodyResult) *GetCredentialsResponseBody
	GetResult() *GetCredentialsResponseBodyResult
}

type GetCredentialsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B7C901ED-2BC1-5CFB-BE23-242DE5E3BA5C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The query result.
	Result *GetCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialsResponseBody) GetResult() *GetCredentialsResponseBodyResult {
	return s.Result
}

func (s *GetCredentialsResponseBody) SetRequestId(v string) *GetCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialsResponseBody) SetResult(v *GetCredentialsResponseBodyResult) *GetCredentialsResponseBody {
	s.Result = v
	return s
}

func (s *GetCredentialsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialsResponseBodyResult struct {
	// The workspace ID.
	//
	// example:
	//
	// 111111
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// Indicates whether the credential is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The access credential token.
	//
	// example:
	//
	// OS-********
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The credential type.
	//
	// - api-token.
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCredentialsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCredentialsResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *GetCredentialsResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCredentialsResponseBodyResult) GetToken() *string {
	return s.Token
}

func (s *GetCredentialsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetCredentialsResponseBodyResult) SetAppGroupId(v int64) *GetCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *GetCredentialsResponseBodyResult) SetEnabled(v bool) *GetCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *GetCredentialsResponseBodyResult) SetToken(v string) *GetCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *GetCredentialsResponseBodyResult) SetType(v string) *GetCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCredentialsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
