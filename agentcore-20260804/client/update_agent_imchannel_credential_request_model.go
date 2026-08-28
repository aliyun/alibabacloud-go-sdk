// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAgentIMChannelCredentialRequestBody) *UpdateAgentIMChannelCredentialRequest
	GetBody() *UpdateAgentIMChannelCredentialRequestBody
	SetClientToken(v string) *UpdateAgentIMChannelCredentialRequest
	GetClientToken() *string
}

type UpdateAgentIMChannelCredentialRequest struct {
	// The request body.
	Body *UpdateAgentIMChannelCredentialRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// A reserved idempotency token. The backend does not provide persistent idempotence guarantee in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAgentIMChannelCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialRequest) GetBody() *UpdateAgentIMChannelCredentialRequestBody {
	return s.Body
}

func (s *UpdateAgentIMChannelCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentIMChannelCredentialRequest) SetBody(v *UpdateAgentIMChannelCredentialRequestBody) *UpdateAgentIMChannelCredentialRequest {
	s.Body = v
	return s
}

func (s *UpdateAgentIMChannelCredentialRequest) SetClientToken(v string) *UpdateAgentIMChannelCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelCredentialRequestBody struct {
	// The channel credential. All fields must be provided and field values must be non-empty strings. DingTalk uses clientID and clientSecret. Lark uses appId and appSecret. WeCom uses botId and secret.
	//
	// This parameter is required.
	Credential map[string]*string `json:"credential,omitempty" xml:"credential,omitempty"`
}

func (s UpdateAgentIMChannelCredentialRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialRequestBody) GetCredential() map[string]*string {
	return s.Credential
}

func (s *UpdateAgentIMChannelCredentialRequestBody) SetCredential(v map[string]*string) *UpdateAgentIMChannelCredentialRequestBody {
	s.Credential = v
	return s
}

func (s *UpdateAgentIMChannelCredentialRequestBody) Validate() error {
	return dara.Validate(s)
}
