// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCredentialDescriptionRequest
	GetClientToken() *string
	SetCredentialId(v string) *UpdateCredentialDescriptionRequest
	GetCredentialId() *string
	SetDescription(v string) *UpdateCredentialDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCredentialDescriptionRequest
	GetInstanceId() *string
}

type UpdateCredentialDescriptionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// The description of the credential.
	//
	// example:
	//
	// credential_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCredentialDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialDescriptionRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateCredentialDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialDescriptionRequest) SetClientToken(v string) *UpdateCredentialDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialDescriptionRequest) SetCredentialId(v string) *UpdateCredentialDescriptionRequest {
	s.CredentialId = &v
	return s
}

func (s *UpdateCredentialDescriptionRequest) SetDescription(v string) *UpdateCredentialDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateCredentialDescriptionRequest) SetInstanceId(v string) *UpdateCredentialDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
