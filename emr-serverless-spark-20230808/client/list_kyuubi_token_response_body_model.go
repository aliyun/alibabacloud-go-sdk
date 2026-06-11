// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListKyuubiTokenResponseBodyData) *ListKyuubiTokenResponseBody
	GetData() *ListKyuubiTokenResponseBodyData
	SetRequestId(v string) *ListKyuubiTokenResponseBody
	GetRequestId() *string
}

type ListKyuubiTokenResponseBody struct {
	// The data returned.
	Data *ListKyuubiTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListKyuubiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBody) GetData() *ListKyuubiTokenResponseBodyData {
	return s.Data
}

func (s *ListKyuubiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKyuubiTokenResponseBody) SetData(v *ListKyuubiTokenResponseBodyData) *ListKyuubiTokenResponseBody {
	s.Data = v
	return s
}

func (s *ListKyuubiTokenResponseBody) SetRequestId(v string) *ListKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKyuubiTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKyuubiTokenResponseBodyData struct {
	// A list of tokens.
	Tokens []*ListKyuubiTokenResponseBodyDataTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
}

func (s ListKyuubiTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBodyData) GetTokens() []*ListKyuubiTokenResponseBodyDataTokens {
	return s.Tokens
}

func (s *ListKyuubiTokenResponseBodyData) SetTokens(v []*ListKyuubiTokenResponseBodyDataTokens) *ListKyuubiTokenResponseBodyData {
	s.Tokens = v
	return s
}

func (s *ListKyuubiTokenResponseBodyData) Validate() error {
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

type ListKyuubiTokenResponseBodyDataTokens struct {
	// The names of authorized RAM users.
	AccountNames []*string `json:"accountNames,omitempty" xml:"accountNames,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 1755496031000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The creator\\"s name.
	//
	// example:
	//
	// admin
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1787033891000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The last used time.
	//
	// example:
	//
	// 1760366232121
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// The ARNs of authorized users.
	MemberArns []*string `json:"memberArns,omitempty" xml:"memberArns,omitempty" type:"Repeated"`
	// The token name.
	//
	// example:
	//
	// dev_serverless_spark
	Name      *string   `json:"name,omitempty" xml:"name,omitempty"`
	SparkRole []*string `json:"sparkRole,omitempty" xml:"sparkRole,omitempty" type:"Repeated"`
	// The masked token.
	//
	// example:
	//
	// dxj**********wfg
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The token ID.
	//
	// example:
	//
	// tk-zpi0*****hdv4y
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s ListKyuubiTokenResponseBodyDataTokens) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiTokenResponseBodyDataTokens) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetAccountNames() []*string {
	return s.AccountNames
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetMemberArns() []*string {
	return s.MemberArns
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetName() *string {
	return s.Name
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetSparkRole() []*string {
	return s.SparkRole
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetToken() *string {
	return s.Token
}

func (s *ListKyuubiTokenResponseBodyDataTokens) GetTokenId() *string {
	return s.TokenId
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetAccountNames(v []*string) *ListKyuubiTokenResponseBodyDataTokens {
	s.AccountNames = v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetCreateTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.CreateTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetCreatedBy(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.CreatedBy = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetExpireTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.ExpireTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetLastUsedTime(v int64) *ListKyuubiTokenResponseBodyDataTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetMemberArns(v []*string) *ListKyuubiTokenResponseBodyDataTokens {
	s.MemberArns = v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetName(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.Name = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetSparkRole(v []*string) *ListKyuubiTokenResponseBodyDataTokens {
	s.SparkRole = v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetToken(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.Token = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) SetTokenId(v string) *ListKyuubiTokenResponseBodyDataTokens {
	s.TokenId = &v
	return s
}

func (s *ListKyuubiTokenResponseBodyDataTokens) Validate() error {
	return dara.Validate(s)
}
