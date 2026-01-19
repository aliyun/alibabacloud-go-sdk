// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableClientPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableClientPublicKeyRequest
	GetApplicationId() *string
	SetClientPublicKeyId(v string) *DisableClientPublicKeyRequest
	GetClientPublicKeyId() *string
	SetClientToken(v string) *DisableClientPublicKeyRequest
	GetClientToken() *string
	SetInstanceId(v string) *DisableClientPublicKeyRequest
	GetInstanceId() *string
}

type DisableClientPublicKeyRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用ClientPublicKey的ID
	//
	// This parameter is required.
	//
	// example:
	//
	// KEYEqDnDJhztiEAwSin7MZoxGcihzCAuxxxx
	ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableClientPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableClientPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DisableClientPublicKeyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableClientPublicKeyRequest) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *DisableClientPublicKeyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableClientPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableClientPublicKeyRequest) SetApplicationId(v string) *DisableClientPublicKeyRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableClientPublicKeyRequest) SetClientPublicKeyId(v string) *DisableClientPublicKeyRequest {
	s.ClientPublicKeyId = &v
	return s
}

func (s *DisableClientPublicKeyRequest) SetClientToken(v string) *DisableClientPublicKeyRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableClientPublicKeyRequest) SetInstanceId(v string) *DisableClientPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableClientPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
