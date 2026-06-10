// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnthropicHost(v string) *GetUserResponseBody
	GetAnthropicHost() *string
	SetApiKeys(v []*GetUserResponseBodyApiKeys) *GetUserResponseBody
	GetApiKeys() []*GetUserResponseBodyApiKeys
	SetAppId(v string) *GetUserResponseBody
	GetAppId() *string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetHost(v string) *GetUserResponseBody
	GetHost() *string
	SetInnerToken(v string) *GetUserResponseBody
	GetInnerToken() *string
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetToken(v string) *GetUserResponseBody
	GetToken() *string
}

type GetUserResponseBody struct {
	AnthropicHost *string                       `json:"AnthropicHost,omitempty" xml:"AnthropicHost,omitempty"`
	ApiKeys       []*GetUserResponseBodyApiKeys `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	AppId         *string                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code          *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Host          *string                       `json:"Host,omitempty" xml:"Host,omitempty"`
	InnerToken    *string                       `json:"InnerToken,omitempty" xml:"InnerToken,omitempty"`
	Message       *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token         *string                       `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetAnthropicHost() *string {
	return s.AnthropicHost
}

func (s *GetUserResponseBody) GetApiKeys() []*GetUserResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *GetUserResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetUserResponseBody) GetInnerToken() *string {
	return s.InnerToken
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetUserResponseBody) SetAnthropicHost(v string) *GetUserResponseBody {
	s.AnthropicHost = &v
	return s
}

func (s *GetUserResponseBody) SetApiKeys(v []*GetUserResponseBodyApiKeys) *GetUserResponseBody {
	s.ApiKeys = v
	return s
}

func (s *GetUserResponseBody) SetAppId(v string) *GetUserResponseBody {
	s.AppId = &v
	return s
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetHost(v string) *GetUserResponseBody {
	s.Host = &v
	return s
}

func (s *GetUserResponseBody) SetInnerToken(v string) *GetUserResponseBody {
	s.InnerToken = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetToken(v string) *GetUserResponseBody {
	s.Token = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.ApiKeys != nil {
		for _, item := range s.ApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserResponseBodyApiKeys struct {
	ApiKey      *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	InnerApiKey *string `json:"InnerApiKey,omitempty" xml:"InnerApiKey,omitempty"`
}

func (s GetUserResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyApiKeys) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetUserResponseBodyApiKeys) GetInnerApiKey() *string {
	return s.InnerApiKey
}

func (s *GetUserResponseBodyApiKeys) SetApiKey(v string) *GetUserResponseBodyApiKeys {
	s.ApiKey = &v
	return s
}

func (s *GetUserResponseBodyApiKeys) SetInnerApiKey(v string) *GetUserResponseBodyApiKeys {
	s.InnerApiKey = &v
	return s
}

func (s *GetUserResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}
