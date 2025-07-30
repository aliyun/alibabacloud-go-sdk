// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationClientSecretRequest
	GetApplicationId() *string
	SetInstanceId(v string) *CreateApplicationClientSecretRequest
	GetInstanceId() *string
}

type CreateApplicationClientSecretRequest struct {
	// The ID of the application for which you want to create a client key.
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
}

func (s CreateApplicationClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationClientSecretRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationClientSecretRequest) SetApplicationId(v string) *CreateApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) SetInstanceId(v string) *CreateApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
