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
	Body        *UpdateModelConnectionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	ClientToken *string                           `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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
	ApiKeys      []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	Description  *string   `json:"description,omitempty" xml:"description,omitempty"`
	Endpoint     *string   `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Name         *string   `json:"name,omitempty" xml:"name,omitempty"`
	Protocol     *string   `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ProviderType *string   `json:"providerType,omitempty" xml:"providerType,omitempty"`
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
