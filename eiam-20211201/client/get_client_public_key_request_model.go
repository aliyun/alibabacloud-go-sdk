// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetClientPublicKeyRequest
	GetApplicationId() *string
	SetClientPublicKeyId(v string) *GetClientPublicKeyRequest
	GetClientPublicKeyId() *string
	SetInstanceId(v string) *GetClientPublicKeyRequest
	GetInstanceId() *string
}

type GetClientPublicKeyRequest struct {
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
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetClientPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetClientPublicKeyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetClientPublicKeyRequest) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *GetClientPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetClientPublicKeyRequest) SetApplicationId(v string) *GetClientPublicKeyRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetClientPublicKeyRequest) SetClientPublicKeyId(v string) *GetClientPublicKeyRequest {
	s.ClientPublicKeyId = &v
	return s
}

func (s *GetClientPublicKeyRequest) SetInstanceId(v string) *GetClientPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetClientPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
