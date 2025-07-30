// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ObtainApplicationClientSecretRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ObtainApplicationClientSecretRequest
	GetInstanceId() *string
	SetSecretId(v string) *ObtainApplicationClientSecretRequest
	GetSecretId() *string
}

type ObtainApplicationClientSecretRequest struct {
	// The ID of the application whose client key you want to query.
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

func (s ObtainApplicationClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ObtainApplicationClientSecretRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainApplicationClientSecretRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *ObtainApplicationClientSecretRequest) SetApplicationId(v string) *ObtainApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationClientSecretRequest) SetInstanceId(v string) *ObtainApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationClientSecretRequest) SetSecretId(v string) *ObtainApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

func (s *ObtainApplicationClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
