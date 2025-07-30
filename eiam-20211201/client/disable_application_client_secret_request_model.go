// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationClientSecretRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationClientSecretRequest
	GetInstanceId() *string
	SetSecretId(v string) *DisableApplicationClientSecretRequest
	GetSecretId() *string
}

type DisableApplicationClientSecretRequest struct {
	// The ID of the application for which you want to disable a client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client key ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DisableApplicationClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationClientSecretRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationClientSecretRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationClientSecretRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *DisableApplicationClientSecretRequest) SetApplicationId(v string) *DisableApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationClientSecretRequest) SetInstanceId(v string) *DisableApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationClientSecretRequest) SetSecretId(v string) *DisableApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

func (s *DisableApplicationClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
