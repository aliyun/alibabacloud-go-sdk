// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationFederatedCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *DeleteApplicationFederatedCredentialRequest
	GetApplicationFederatedCredentialId() *string
	SetApplicationId(v string) *DeleteApplicationFederatedCredentialRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DeleteApplicationFederatedCredentialRequest
	GetInstanceId() *string
}

type DeleteApplicationFederatedCredentialRequest struct {
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

func (s DeleteApplicationFederatedCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationFederatedCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationFederatedCredentialRequest) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *DeleteApplicationFederatedCredentialRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationFederatedCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApplicationFederatedCredentialRequest) SetApplicationFederatedCredentialId(v string) *DeleteApplicationFederatedCredentialRequest {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *DeleteApplicationFederatedCredentialRequest) SetApplicationId(v string) *DeleteApplicationFederatedCredentialRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationFederatedCredentialRequest) SetInstanceId(v string) *DeleteApplicationFederatedCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApplicationFederatedCredentialRequest) Validate() error {
	return dara.Validate(s)
}
