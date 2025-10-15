// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *UpdateApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetAttributeMappings(v []*UpdateApplicationFederatedCredentialRequestAttributeMappings) *UpdateApplicationFederatedCredentialRequest
	GetAttributeMappings() []*UpdateApplicationFederatedCredentialRequestAttributeMappings
	SetInstanceId(v string) *UpdateApplicationFederatedCredentialRequest
	GetInstanceId() *string
	SetVerificationCondition(v string) *UpdateApplicationFederatedCredentialRequest
	GetVerificationCondition() *string
}

type UpdateApplicationFederatedCredentialRequest struct {
	// 应用联邦凭证Id
	//
	// This parameter is required.
	//
	// example:
	//
	// afc_aaaaa1111
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 属性映射
	AttributeMappings []*UpdateApplicationFederatedCredentialRequestAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 校验条件
	//
	// example:
	//
	// IsNullOrEmpty("")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetAttributeMappings() []*UpdateApplicationFederatedCredentialRequestAttributeMappings {
	return s.AttributeMappings
}

func (s *UpdateApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationFederatedCredentialRequest) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *UpdateApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetApplicationId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetAttributeMappings(v []*UpdateApplicationFederatedCredentialRequestAttributeMappings) *UpdateApplicationFederatedCredentialRequest {
	s.AttributeMappings = v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetInstanceId(v string) *UpdateApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) SetVerificationCondition(v string) *UpdateApplicationFederatedCredentialRequest {
	s.VerificationCondition = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequest) Validate() error {
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

type UpdateApplicationFederatedCredentialRequestAttributeMappings struct {
	// 源值表达式
	//
	// example:
	//
	// Append(client.applicationFederatedCredentialId, ":", cert.subject.CN, ":", cert.serialNumber)
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// 目标字段
	//
	// example:
	//
	// client.activeSubjectUrn
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
}

func (s UpdateApplicationFederatedCredentialRequestAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialRequestAttributeMappings) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) SetSourceValueExpression(v string) *UpdateApplicationFederatedCredentialRequestAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) SetTargetField(v string) *UpdateApplicationFederatedCredentialRequestAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialRequestAttributeMappings) Validate() error {
	return dara.Validate(s)
}
