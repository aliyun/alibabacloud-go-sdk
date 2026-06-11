// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLivyComputeTokenResponseBody
	GetCode() *string
	SetData(v *ListLivyComputeTokenResponseBodyData) *ListLivyComputeTokenResponseBody
	GetData() *ListLivyComputeTokenResponseBodyData
	SetMessage(v string) *ListLivyComputeTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLivyComputeTokenResponseBody
	GetRequestId() *string
}

type ListLivyComputeTokenResponseBody struct {
	// The response code. A value of 1000000 indicates that the request was successful. Other values indicate that the request failed. See the message parameter for failure details.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLivyComputeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLivyComputeTokenResponseBody) GetData() *ListLivyComputeTokenResponseBodyData {
	return s.Data
}

func (s *ListLivyComputeTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLivyComputeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivyComputeTokenResponseBody) SetCode(v string) *ListLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetData(v *ListLivyComputeTokenResponseBodyData) *ListLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetMessage(v string) *ListLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *ListLivyComputeTokenResponseBody) SetRequestId(v string) *ListLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivyComputeTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLivyComputeTokenResponseBodyData struct {
	// The list of tokens.
	Tokens []*ListLivyComputeTokenResponseBodyDataTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
}

func (s ListLivyComputeTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBodyData) GetTokens() []*ListLivyComputeTokenResponseBodyDataTokens {
	return s.Tokens
}

func (s *ListLivyComputeTokenResponseBodyData) SetTokens(v []*ListLivyComputeTokenResponseBodyDataTokens) *ListLivyComputeTokenResponseBodyData {
	s.Tokens = v
	return s
}

func (s *ListLivyComputeTokenResponseBodyData) Validate() error {
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

type ListLivyComputeTokenResponseBodyDataTokens struct {
	// The time when the token was created.
	//
	// example:
	//
	// 1749456094000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The user who created the token.
	//
	// example:
	//
	// alice
	Createdby *string `json:"createdby,omitempty" xml:"createdby,omitempty"`
	// The time when the token expires.
	//
	// example:
	//
	// 1749456994000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The time when the token was last used.
	//
	// example:
	//
	// 1749456098000
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// The token name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The token content.
	//
	// example:
	//
	// 5d37843fb6f1e8
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The token ID.
	//
	// example:
	//
	// lctk-xxxxxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s ListLivyComputeTokenResponseBodyDataTokens) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeTokenResponseBodyDataTokens) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetCreatedby() *string {
	return s.Createdby
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetName() *string {
	return s.Name
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetToken() *string {
	return s.Token
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) GetTokenId() *string {
	return s.TokenId
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetCreateTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.CreateTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetCreatedby(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Createdby = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetExpireTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.ExpireTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetLastUsedTime(v int64) *ListLivyComputeTokenResponseBodyDataTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetName(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Name = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetToken(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.Token = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) SetTokenId(v string) *ListLivyComputeTokenResponseBodyDataTokens {
	s.TokenId = &v
	return s
}

func (s *ListLivyComputeTokenResponseBodyDataTokens) Validate() error {
	return dara.Validate(s)
}
