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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the ClientPublicKey for the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// KEYEqDnDJhztiEAwSin7MZoxGcihzCAuxxxx
	ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
	// A client-generated token that you can use to ensure the idempotence of the request. Make sure that the token is unique for each request. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
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
