// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialProviderDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderId(v string) *UpdateCredentialProviderDescriptionRequest
	GetCredentialProviderId() *string
	SetDescription(v string) *UpdateCredentialProviderDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCredentialProviderDescriptionRequest
	GetInstanceId() *string
}

type UpdateCredentialProviderDescriptionRequest struct {
	// The ID of the credential provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// A description of the credential provider.
	//
	// > The description can be up to 128 characters long.
	//
	// example:
	//
	// This is an example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCredentialProviderDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderDescriptionRequest) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *UpdateCredentialProviderDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialProviderDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCredentialProviderDescriptionRequest) SetCredentialProviderId(v string) *UpdateCredentialProviderDescriptionRequest {
	s.CredentialProviderId = &v
	return s
}

func (s *UpdateCredentialProviderDescriptionRequest) SetDescription(v string) *UpdateCredentialProviderDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateCredentialProviderDescriptionRequest) SetInstanceId(v string) *UpdateCredentialProviderDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCredentialProviderDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
