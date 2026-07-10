// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationClientSecretExpirationTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationClientSecretExpirationTimeRequest
	GetApplicationId() *string
	SetExpirationTime(v int64) *UpdateApplicationClientSecretExpirationTimeRequest
	GetExpirationTime() *int64
	SetInstanceId(v string) *UpdateApplicationClientSecretExpirationTimeRequest
	GetInstanceId() *string
	SetSecretId(v string) *UpdateApplicationClientSecretExpirationTimeRequest
	GetSecretId() *string
}

type UpdateApplicationClientSecretExpirationTimeRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The expiration time of the client secret. The value is a UNIX timestamp in milliseconds. The minimum validity period that can be set is 1 day, and the maximum validity period is 3 years.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735530123762
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// sct_11111
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s UpdateApplicationClientSecretExpirationTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationClientSecretExpirationTimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) SetApplicationId(v string) *UpdateApplicationClientSecretExpirationTimeRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) SetExpirationTime(v int64) *UpdateApplicationClientSecretExpirationTimeRequest {
	s.ExpirationTime = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) SetInstanceId(v string) *UpdateApplicationClientSecretExpirationTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) SetSecretId(v string) *UpdateApplicationClientSecretExpirationTimeRequest {
	s.SecretId = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeRequest) Validate() error {
	return dara.Validate(s)
}
