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
	SetExpirationTime(v int64) *CreateApplicationClientSecretRequest
	GetExpirationTime() *int64
	SetInstanceId(v string) *CreateApplicationClientSecretRequest
	GetInstanceId() *string
}

type CreateApplicationClientSecretRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The expiration time in UNIX timestamp format, in milliseconds. If this parameter is left empty, the client secret is permanently valid. The minimum validity period that you can set is 1 day, and the maximum validity period is 3 years.
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

func (s *CreateApplicationClientSecretRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *CreateApplicationClientSecretRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationClientSecretRequest) SetApplicationId(v string) *CreateApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) SetExpirationTime(v int64) *CreateApplicationClientSecretRequest {
	s.ExpirationTime = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) SetInstanceId(v string) *CreateApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
