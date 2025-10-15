// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest
	GetApplicationId() *string
	SetDescription(v string) *UpdateApplicationFederatedCredentialDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest
	GetInstanceId() *string
}

type UpdateApplicationFederatedCredentialDescriptionRequest struct {
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
	// 联邦凭证描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationFederatedCredentialDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) SetApplicationFederatedCredentialId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) SetApplicationId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) SetDescription(v string) *UpdateApplicationFederatedCredentialDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) SetInstanceId(v string) *UpdateApplicationFederatedCredentialDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
