// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPrimaryClientPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetPrimaryClientPublicKeyRequest
	GetApplicationId() *string
	SetClientPublicKeyId(v string) *SetPrimaryClientPublicKeyRequest
	GetClientPublicKeyId() *string
	SetClientToken(v string) *SetPrimaryClientPublicKeyRequest
	GetClientToken() *string
	SetInstanceId(v string) *SetPrimaryClientPublicKeyRequest
	GetInstanceId() *string
}

type SetPrimaryClientPublicKeyRequest struct {
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

func (s SetPrimaryClientPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPrimaryClientPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *SetPrimaryClientPublicKeyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetPrimaryClientPublicKeyRequest) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *SetPrimaryClientPublicKeyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetPrimaryClientPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPrimaryClientPublicKeyRequest) SetApplicationId(v string) *SetPrimaryClientPublicKeyRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetPrimaryClientPublicKeyRequest) SetClientPublicKeyId(v string) *SetPrimaryClientPublicKeyRequest {
	s.ClientPublicKeyId = &v
	return s
}

func (s *SetPrimaryClientPublicKeyRequest) SetClientToken(v string) *SetPrimaryClientPublicKeyRequest {
	s.ClientToken = &v
	return s
}

func (s *SetPrimaryClientPublicKeyRequest) SetInstanceId(v string) *SetPrimaryClientPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPrimaryClientPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
