// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateModelConnectionRequestBody) *UpdateModelConnectionRequest
	GetBody() *UpdateModelConnectionRequestBody
	SetClientToken(v string) *UpdateModelConnectionRequest
	GetClientToken() *string
}

type UpdateModelConnectionRequest struct {
	// The request body.
	Body *UpdateModelConnectionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token used for idempotence. Not currently supported.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateModelConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionRequest) GetBody() *UpdateModelConnectionRequestBody {
	return s.Body
}

func (s *UpdateModelConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelConnectionRequest) SetBody(v *UpdateModelConnectionRequestBody) *UpdateModelConnectionRequest {
	s.Body = v
	return s
}

func (s *UpdateModelConnectionRequest) SetClientToken(v string) *UpdateModelConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelConnectionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelConnectionRequestBody struct {
	// The list of API keys used to access the upstream model service. The list must contain at least one non-empty value.
	ApiKeys []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// The description of the model connection. The description can be up to 255 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The absolute HTTP or HTTPS address of the upstream model service. The address can be up to 1024 characters in length.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The model connection name. The name must be 1 to 128 non-whitespace characters in length.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The model invocation protocol. Currently, only OpenAI/v1 is supported. If not specified in Settings when the model connection is created, this default value is used.
	//
	// example:
	//
	// OpenAI/v1
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The model provider type.
	//
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
}

func (s UpdateModelConnectionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionRequestBody) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *UpdateModelConnectionRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelConnectionRequestBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateModelConnectionRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateModelConnectionRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateModelConnectionRequestBody) GetProviderType() *string {
	return s.ProviderType
}

func (s *UpdateModelConnectionRequestBody) SetApiKeys(v []*string) *UpdateModelConnectionRequestBody {
	s.ApiKeys = v
	return s
}

func (s *UpdateModelConnectionRequestBody) SetDescription(v string) *UpdateModelConnectionRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateModelConnectionRequestBody) SetEndpoint(v string) *UpdateModelConnectionRequestBody {
	s.Endpoint = &v
	return s
}

func (s *UpdateModelConnectionRequestBody) SetName(v string) *UpdateModelConnectionRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateModelConnectionRequestBody) SetProtocol(v string) *UpdateModelConnectionRequestBody {
	s.Protocol = &v
	return s
}

func (s *UpdateModelConnectionRequestBody) SetProviderType(v string) *UpdateModelConnectionRequestBody {
	s.ProviderType = &v
	return s
}

func (s *UpdateModelConnectionRequestBody) Validate() error {
	return dara.Validate(s)
}
