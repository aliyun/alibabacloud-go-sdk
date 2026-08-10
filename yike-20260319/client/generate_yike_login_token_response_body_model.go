// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateYikeLoginTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresAt(v string) *GenerateYikeLoginTokenResponseBody
	GetExpiresAt() *string
	SetRequestId(v string) *GenerateYikeLoginTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GenerateYikeLoginTokenResponseBody
	GetToken() *string
	SetUserId(v string) *GenerateYikeLoginTokenResponseBody
	GetUserId() *string
}

type GenerateYikeLoginTokenResponseBody struct {
	// The session expiration time (UNIX timestamp in milliseconds).
	//
	// example:
	//
	// 1782008128000
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The logon token.
	//
	// example:
	//
	// ******d6931ff7e89b5eb19484*****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The generated user ID.
	//
	// example:
	//
	// 23253**
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateYikeLoginTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateYikeLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateYikeLoginTokenResponseBody) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *GenerateYikeLoginTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateYikeLoginTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GenerateYikeLoginTokenResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GenerateYikeLoginTokenResponseBody) SetExpiresAt(v string) *GenerateYikeLoginTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateYikeLoginTokenResponseBody) SetRequestId(v string) *GenerateYikeLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateYikeLoginTokenResponseBody) SetToken(v string) *GenerateYikeLoginTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GenerateYikeLoginTokenResponseBody) SetUserId(v string) *GenerateYikeLoginTokenResponseBody {
	s.UserId = &v
	return s
}

func (s *GenerateYikeLoginTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
