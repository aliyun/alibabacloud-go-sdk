// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateCredentialRequestBody) *UpdateCredentialRequest
	GetBody() *UpdateCredentialRequestBody
	SetClientToken(v string) *UpdateCredentialRequest
	GetClientToken() *string
}

type UpdateCredentialRequest struct {
	Body *UpdateCredentialRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequest) GetBody() *UpdateCredentialRequestBody {
	return s.Body
}

func (s *UpdateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialRequest) SetBody(v *UpdateCredentialRequestBody) *UpdateCredentialRequest {
	s.Body = v
	return s
}

func (s *UpdateCredentialRequest) SetClientToken(v string) *UpdateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialRequestBody struct {
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// example:
	//
	// 线上环境调用模型服务使用的 API Key
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateCredentialRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestBody) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *UpdateCredentialRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialRequestBody) SetCredentialMetadata(v string) *UpdateCredentialRequestBody {
	s.CredentialMetadata = &v
	return s
}

func (s *UpdateCredentialRequestBody) SetDescription(v string) *UpdateCredentialRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateCredentialRequestBody) Validate() error {
	return dara.Validate(s)
}
