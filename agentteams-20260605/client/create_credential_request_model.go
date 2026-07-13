// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateCredentialRequest
	GetApiKey() *string
	SetClientToken(v string) *CreateCredentialRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCredentialRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCredentialRequest
	GetInstanceId() *string
	SetName(v string) *CreateCredentialRequest
	GetName() *string
}

type CreateCredentialRequest struct {
	// This parameter is required.
	ApiKey      *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCredentialRequest) GetName() *string {
	return s.Name
}

func (s *CreateCredentialRequest) SetApiKey(v string) *CreateCredentialRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateCredentialRequest) SetClientToken(v string) *CreateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialRequest) SetDescription(v string) *CreateCredentialRequest {
	s.Description = &v
	return s
}

func (s *CreateCredentialRequest) SetInstanceId(v string) *CreateCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCredentialRequest) SetName(v string) *CreateCredentialRequest {
	s.Name = &v
	return s
}

func (s *CreateCredentialRequest) Validate() error {
	return dara.Validate(s)
}
