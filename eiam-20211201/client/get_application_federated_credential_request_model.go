// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *GetApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *GetApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationFederatedCredentialRequest
	GetInstanceId() *string
}

type GetApplicationFederatedCredentialRequest struct {
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

func (s GetApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *GetApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *GetApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *GetApplicationFederatedCredentialRequest) SetApplicationId(v string) *GetApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationFederatedCredentialRequest) SetInstanceId(v string) *GetApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationFederatedCredentialRequest) Validate() error {
	return dara.Validate(s)
}
