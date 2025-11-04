// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMessageChatTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GenerateMessageChatTokenResponseBody
	GetAppId() *string
	SetAppSign(v string) *GenerateMessageChatTokenResponseBody
	GetAppSign() *string
	SetNonce(v string) *GenerateMessageChatTokenResponseBody
	GetNonce() *string
	SetRequestId(v string) *GenerateMessageChatTokenResponseBody
	GetRequestId() *string
	SetRole(v string) *GenerateMessageChatTokenResponseBody
	GetRole() *string
	SetTimeStamp(v int64) *GenerateMessageChatTokenResponseBody
	GetTimeStamp() *int64
	SetToken(v string) *GenerateMessageChatTokenResponseBody
	GetToken() *string
	SetUserId(v string) *GenerateMessageChatTokenResponseBody
	GetUserId() *string
}

type GenerateMessageChatTokenResponseBody struct {
	// The AppID of the user.
	//
	// example:
	//
	// ***********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application signature.
	//
	// example:
	//
	// H4sIAAAAAAAE******************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// The nonce used to generate the token.
	//
	// example:
	//
	// AK-***********
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// The request ID.
	//
	// example:
	//
	// req_1234567890abcdef
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role used to generate the token.
	//
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The expiration time. Unit: seconds. Expiration time = Current time + Validity period.
	//
	// example:
	//
	// 1700000000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The generated token.
	//
	// example:
	//
	// acet**********
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the user for joining the channel.
	//
	// example:
	//
	// YOURUSERID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateMessageChatTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateMessageChatTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMessageChatTokenResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GenerateMessageChatTokenResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *GenerateMessageChatTokenResponseBody) GetNonce() *string {
	return s.Nonce
}

func (s *GenerateMessageChatTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateMessageChatTokenResponseBody) GetRole() *string {
	return s.Role
}

func (s *GenerateMessageChatTokenResponseBody) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *GenerateMessageChatTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GenerateMessageChatTokenResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GenerateMessageChatTokenResponseBody) SetAppId(v string) *GenerateMessageChatTokenResponseBody {
	s.AppId = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetAppSign(v string) *GenerateMessageChatTokenResponseBody {
	s.AppSign = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetNonce(v string) *GenerateMessageChatTokenResponseBody {
	s.Nonce = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetRequestId(v string) *GenerateMessageChatTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetRole(v string) *GenerateMessageChatTokenResponseBody {
	s.Role = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetTimeStamp(v int64) *GenerateMessageChatTokenResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetToken(v string) *GenerateMessageChatTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) SetUserId(v string) *GenerateMessageChatTokenResponseBody {
	s.UserId = &v
	return s
}

func (s *GenerateMessageChatTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
