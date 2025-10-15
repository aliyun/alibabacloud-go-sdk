// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *DisableApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *DisableApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationFederatedCredentialRequest
	GetInstanceId() *string
}

type DisableApplicationFederatedCredentialRequest struct {
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
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *DisableApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *DisableApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *DisableApplicationFederatedCredentialRequest) SetApplicationId(v string) *DisableApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationFederatedCredentialRequest) SetInstanceId(v string) *DisableApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationFederatedCredentialRequest) Validate() error {
	return dara.Validate(s)
}
