// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialName(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialName() *string
	SetApplicationFederatedCredentialType(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialType() *string
	SetApplicationId(v string) *CreateApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetAttributeMappings(v []*CreateApplicationFederatedCredentialRequestAttributeMappings) *CreateApplicationFederatedCredentialRequest
	GetAttributeMappings() []*CreateApplicationFederatedCredentialRequestAttributeMappings
	SetDescription(v string) *CreateApplicationFederatedCredentialRequest
	GetDescription() *string
	SetFederatedCredentialProviderId(v string) *CreateApplicationFederatedCredentialRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *CreateApplicationFederatedCredentialRequest
	GetInstanceId() *string
	SetVerificationCondition(v string) *CreateApplicationFederatedCredentialRequest
	GetVerificationCondition() *string
}

type CreateApplicationFederatedCredentialRequest struct {
	// The name of the application federated credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// example_name
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// The type of the application federated credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The attribute mappings.
	AttributeMappings []*CreateApplicationFederatedCredentialRequestAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// description_text
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the federated credential provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// fcp_adasd
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The verification condition.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *CreateApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationFederatedCredentialRequest) GetAttributeMappings() []*CreateApplicationFederatedCredentialRequestAttributeMappings {
	return s.AttributeMappings
}

func (s *CreateApplicationFederatedCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationFederatedCredentialRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *CreateApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationFederatedCredentialRequest) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialName(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialType(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetApplicationId(v string) *CreateApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetAttributeMappings(v []*CreateApplicationFederatedCredentialRequestAttributeMappings) *CreateApplicationFederatedCredentialRequest {
	s.AttributeMappings = v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetDescription(v string) *CreateApplicationFederatedCredentialRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetFederatedCredentialProviderId(v string) *CreateApplicationFederatedCredentialRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetInstanceId(v string) *CreateApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) SetVerificationCondition(v string) *CreateApplicationFederatedCredentialRequest {
	s.VerificationCondition = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequest) Validate() error {
	if s.AttributeMappings != nil {
		for _, item := range s.AttributeMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApplicationFederatedCredentialRequestAttributeMappings struct {
	// The source value expression.
	//
	// example:
	//
	// Append(client.applicationFederatedCredentialId, ":", cert.subject.CN, ":", cert.serialNumber)
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// The target field.
	//
	// example:
	//
	// client.activeSubjectUrn
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
}

func (s CreateApplicationFederatedCredentialRequestAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialRequestAttributeMappings) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) SetSourceValueExpression(v string) *CreateApplicationFederatedCredentialRequestAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) SetTargetField(v string) *CreateApplicationFederatedCredentialRequestAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *CreateApplicationFederatedCredentialRequestAttributeMappings) Validate() error {
	return dara.Validate(s)
}
