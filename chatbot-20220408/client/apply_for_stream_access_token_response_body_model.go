// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForStreamAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ApplyForStreamAccessTokenResponseBody
	GetAccessToken() *string
	SetChannelId(v string) *ApplyForStreamAccessTokenResponseBody
	GetChannelId() *string
	SetRequestId(v string) *ApplyForStreamAccessTokenResponseBody
	GetRequestId() *string
	SetStreamSecret(v string) *ApplyForStreamAccessTokenResponseBody
	GetStreamSecret() *string
}

type ApplyForStreamAccessTokenResponseBody struct {
	// example:
	//
	// 63ba97b4f18a4a04f715c81e8e643938
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// cc9e88c0-4f41-4f1d-a1a9-91a72d2aa27d
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 736994BD-AA35-4742-88C9-E64BE4BAA14B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cc9e88c0-4f41-4f1d-a1a9-91a72d2aa27d
	StreamSecret *string `json:"StreamSecret,omitempty" xml:"StreamSecret,omitempty"`
}

func (s ApplyForStreamAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyForStreamAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ApplyForStreamAccessTokenResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *ApplyForStreamAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyForStreamAccessTokenResponseBody) GetStreamSecret() *string {
	return s.StreamSecret
}

func (s *ApplyForStreamAccessTokenResponseBody) SetAccessToken(v string) *ApplyForStreamAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetChannelId(v string) *ApplyForStreamAccessTokenResponseBody {
	s.ChannelId = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetRequestId(v string) *ApplyForStreamAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) SetStreamSecret(v string) *ApplyForStreamAccessTokenResponseBody {
	s.StreamSecret = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
