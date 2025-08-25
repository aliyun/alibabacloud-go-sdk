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
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// client secret的有效期时间，Unix时间戳格式，单位为毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735530123762
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// OIDC 场景下用于客户端身份验证的客户端密钥
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
