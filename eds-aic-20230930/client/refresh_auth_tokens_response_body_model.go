// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAuthTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RefreshAuthTokensResponseBodyData) *RefreshAuthTokensResponseBody
	GetData() *RefreshAuthTokensResponseBodyData
	SetRequestId(v string) *RefreshAuthTokensResponseBody
	GetRequestId() *string
}

type RefreshAuthTokensResponseBody struct {
	// The token data.
	Data *RefreshAuthTokensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAuthTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAuthTokensResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAuthTokensResponseBody) GetData() *RefreshAuthTokensResponseBodyData {
	return s.Data
}

func (s *RefreshAuthTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAuthTokensResponseBody) SetData(v *RefreshAuthTokensResponseBodyData) *RefreshAuthTokensResponseBody {
	s.Data = v
	return s
}

func (s *RefreshAuthTokensResponseBody) SetRequestId(v string) *RefreshAuthTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAuthTokensResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshAuthTokensResponseBodyData struct {
	// The model gateway access URL.
	//
	// example:
	//
	// https://ai-gateway.example.com/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The list of tokens.
	Tokens []*RefreshAuthTokensResponseBodyDataTokens `json:"Tokens,omitempty" xml:"Tokens,omitempty" type:"Repeated"`
}

func (s RefreshAuthTokensResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefreshAuthTokensResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAuthTokensResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *RefreshAuthTokensResponseBodyData) GetTokens() []*RefreshAuthTokensResponseBodyDataTokens {
	return s.Tokens
}

func (s *RefreshAuthTokensResponseBodyData) SetBaseUrl(v string) *RefreshAuthTokensResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyData) SetTokens(v []*RefreshAuthTokensResponseBodyDataTokens) *RefreshAuthTokensResponseBodyData {
	s.Tokens = v
	return s
}

func (s *RefreshAuthTokensResponseBodyData) Validate() error {
	if s.Tokens != nil {
		for _, item := range s.Tokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefreshAuthTokensResponseBodyDataTokens struct {
	// The authorization token value.
	//
	// example:
	//
	// cp******lp
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The expiration timestamp.
	//
	// example:
	//
	// 1719648600
	ExpireAt *int64 `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// The validity period in seconds.
	//
	// example:
	//
	// 600
	ExpireSeconds *int64 `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// acp-2zef0gov2nh2l3xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuance timestamp.
	//
	// example:
	//
	// 1719648000
	IssuedAt *int64 `json:"IssuedAt,omitempty" xml:"IssuedAt,omitempty"`
	// The license key.
	//
	// example:
	//
	// lk-abcdef1234567890
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
}

func (s RefreshAuthTokensResponseBodyDataTokens) String() string {
	return dara.Prettify(s)
}

func (s RefreshAuthTokensResponseBodyDataTokens) GoString() string {
	return s.String()
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetAuthToken() *string {
	return s.AuthToken
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetExpireAt() *int64 {
	return s.ExpireAt
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetExpireSeconds() *int64 {
	return s.ExpireSeconds
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetIssuedAt() *int64 {
	return s.IssuedAt
}

func (s *RefreshAuthTokensResponseBodyDataTokens) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetAuthToken(v string) *RefreshAuthTokensResponseBodyDataTokens {
	s.AuthToken = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetExpireAt(v int64) *RefreshAuthTokensResponseBodyDataTokens {
	s.ExpireAt = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetExpireSeconds(v int64) *RefreshAuthTokensResponseBodyDataTokens {
	s.ExpireSeconds = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetInstanceId(v string) *RefreshAuthTokensResponseBodyDataTokens {
	s.InstanceId = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetIssuedAt(v int64) *RefreshAuthTokensResponseBodyDataTokens {
	s.IssuedAt = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) SetLicenseKey(v string) *RefreshAuthTokensResponseBodyDataTokens {
	s.LicenseKey = &v
	return s
}

func (s *RefreshAuthTokensResponseBodyDataTokens) Validate() error {
	return dara.Validate(s)
}
