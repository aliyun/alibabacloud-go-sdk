// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationClientSecretRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DeleteApplicationClientSecretRequest
	GetInstanceId() *string
	SetSecretId(v string) *DeleteApplicationClientSecretRequest
	GetSecretId() *string
}

type DeleteApplicationClientSecretRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client secret ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DeleteApplicationClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationClientSecretRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApplicationClientSecretRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *DeleteApplicationClientSecretRequest) SetApplicationId(v string) *DeleteApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationClientSecretRequest) SetInstanceId(v string) *DeleteApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApplicationClientSecretRequest) SetSecretId(v string) *DeleteApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

func (s *DeleteApplicationClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
