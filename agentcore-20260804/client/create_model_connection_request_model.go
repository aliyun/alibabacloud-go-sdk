// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateModelConnectionRequestBody) *CreateModelConnectionRequest
	GetBody() *CreateModelConnectionRequestBody
	SetClientToken(v string) *CreateModelConnectionRequest
	GetClientToken() *string
}

type CreateModelConnectionRequest struct {
	Body *CreateModelConnectionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateModelConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionRequest) GetBody() *CreateModelConnectionRequestBody {
	return s.Body
}

func (s *CreateModelConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelConnectionRequest) SetBody(v *CreateModelConnectionRequestBody) *CreateModelConnectionRequest {
	s.Body = v
	return s
}

func (s *CreateModelConnectionRequest) SetClientToken(v string) *CreateModelConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelConnectionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelConnectionRequestBody struct {
	// This parameter is required.
	ApiKeys []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OpenAI/v1
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
}

func (s CreateModelConnectionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionRequestBody) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionRequestBody) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *CreateModelConnectionRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateModelConnectionRequestBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateModelConnectionRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateModelConnectionRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateModelConnectionRequestBody) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateModelConnectionRequestBody) SetApiKeys(v []*string) *CreateModelConnectionRequestBody {
	s.ApiKeys = v
	return s
}

func (s *CreateModelConnectionRequestBody) SetDescription(v string) *CreateModelConnectionRequestBody {
	s.Description = &v
	return s
}

func (s *CreateModelConnectionRequestBody) SetEndpoint(v string) *CreateModelConnectionRequestBody {
	s.Endpoint = &v
	return s
}

func (s *CreateModelConnectionRequestBody) SetName(v string) *CreateModelConnectionRequestBody {
	s.Name = &v
	return s
}

func (s *CreateModelConnectionRequestBody) SetProtocol(v string) *CreateModelConnectionRequestBody {
	s.Protocol = &v
	return s
}

func (s *CreateModelConnectionRequestBody) SetProviderType(v string) *CreateModelConnectionRequestBody {
	s.ProviderType = &v
	return s
}

func (s *CreateModelConnectionRequestBody) Validate() error {
	return dara.Validate(s)
}
