// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateUserAccessTokenRequest
	GetAgentKey() *string
	SetEmail(v string) *GenerateUserAccessTokenRequest
	GetEmail() *string
	SetExpireTime(v int32) *GenerateUserAccessTokenRequest
	GetExpireTime() *int32
	SetExtraInfo(v string) *GenerateUserAccessTokenRequest
	GetExtraInfo() *string
	SetForeignId(v string) *GenerateUserAccessTokenRequest
	GetForeignId() *string
	SetNick(v string) *GenerateUserAccessTokenRequest
	GetNick() *string
	SetTelephone(v string) *GenerateUserAccessTokenRequest
	GetTelephone() *string
}

type GenerateUserAccessTokenRequest struct {
	// The key of the business space. If you do not specify this parameter, the system accesses the default business space. You can obtain the key from the Business Management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The user\\"s email address.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is reserved. You do not need to specify it.
	//
	// example:
	//
	// 1000
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Additional information formatted as a JSON string.
	//
	// example:
	//
	// {
	//
	//     "xx": 123,
	//
	//     "yy": 22332,
	//
	//     "customerNameCard": [
	//
	//         {
	//
	//             "key": "姓名",
	//
	//             "value": "张三"
	//
	//         },
	//
	//         {
	//
	//             "key": "等级",
	//
	//             "value": "v2"
	//
	//         },
	//
	//         {
	//
	//             "key": "联系方式",
	//
	//             "value": "123"
	//
	//         }
	//
	//     ]
	//
	// }
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The external user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8882022040000000171
	ForeignId *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	// The visitor\\"s nickname.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// The user\\"s telephone number.
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
}

func (s GenerateUserAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateUserAccessTokenRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateUserAccessTokenRequest) GetEmail() *string {
	return s.Email
}

func (s *GenerateUserAccessTokenRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GenerateUserAccessTokenRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GenerateUserAccessTokenRequest) GetForeignId() *string {
	return s.ForeignId
}

func (s *GenerateUserAccessTokenRequest) GetNick() *string {
	return s.Nick
}

func (s *GenerateUserAccessTokenRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *GenerateUserAccessTokenRequest) SetAgentKey(v string) *GenerateUserAccessTokenRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetEmail(v string) *GenerateUserAccessTokenRequest {
	s.Email = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetExpireTime(v int32) *GenerateUserAccessTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetExtraInfo(v string) *GenerateUserAccessTokenRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetForeignId(v string) *GenerateUserAccessTokenRequest {
	s.ForeignId = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetNick(v string) *GenerateUserAccessTokenRequest {
	s.Nick = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) SetTelephone(v string) *GenerateUserAccessTokenRequest {
	s.Telephone = &v
	return s
}

func (s *GenerateUserAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
