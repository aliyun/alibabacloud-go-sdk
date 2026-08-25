// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCredentialRequestBody) *CreateCredentialRequest
	GetBody() *CreateCredentialRequestBody
	SetClientToken(v string) *CreateCredentialRequest
	GetClientToken() *string
}

type CreateCredentialRequest struct {
	Body *CreateCredentialRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) GetBody() *CreateCredentialRequestBody {
	return s.Body
}

func (s *CreateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialRequest) SetBody(v *CreateCredentialRequestBody) *CreateCredentialRequest {
	s.Body = v
	return s
}

func (s *CreateCredentialRequest) SetClientToken(v string) *CreateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// example:
	//
	// 线上环境调用模型服务使用的 API Key
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCredentialRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestBody) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *CreateCredentialRequestBody) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateCredentialRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateCredentialRequestBody) SetCredentialMetadata(v string) *CreateCredentialRequestBody {
	s.CredentialMetadata = &v
	return s
}

func (s *CreateCredentialRequestBody) SetCredentialType(v string) *CreateCredentialRequestBody {
	s.CredentialType = &v
	return s
}

func (s *CreateCredentialRequestBody) SetDescription(v string) *CreateCredentialRequestBody {
	s.Description = &v
	return s
}

func (s *CreateCredentialRequestBody) SetName(v string) *CreateCredentialRequestBody {
	s.Name = &v
	return s
}

func (s *CreateCredentialRequestBody) Validate() error {
	return dara.Validate(s)
}
