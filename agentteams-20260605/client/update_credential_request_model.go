// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *UpdateCredentialRequest
	GetApiKey() *string
	SetClientToken(v string) *UpdateCredentialRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateCredentialRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCredentialRequest
	GetInstanceId() *string
	SetName(v string) *UpdateCredentialRequest
	GetName() *string
}

type UpdateCredentialRequest struct {
	// This parameter is required.
	ApiKey      *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialRequest) SetApiKey(v string) *UpdateCredentialRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdateCredentialRequest) SetClientToken(v string) *UpdateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialRequest) SetDescription(v string) *UpdateCredentialRequest {
	s.Description = &v
	return s
}

func (s *UpdateCredentialRequest) SetInstanceId(v string) *UpdateCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialRequest) SetName(v string) *UpdateCredentialRequest {
	s.Name = &v
	return s
}

func (s *UpdateCredentialRequest) Validate() error {
	return dara.Validate(s)
}
