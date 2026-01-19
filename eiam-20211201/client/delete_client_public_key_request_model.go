// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteClientPublicKeyRequest
	GetApplicationId() *string
	SetClientPublicKeyId(v string) *DeleteClientPublicKeyRequest
	GetClientPublicKeyId() *string
	SetInstanceId(v string) *DeleteClientPublicKeyRequest
	GetInstanceId() *string
}

type DeleteClientPublicKeyRequest struct {
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

func (s DeleteClientPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientPublicKeyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteClientPublicKeyRequest) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *DeleteClientPublicKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteClientPublicKeyRequest) SetApplicationId(v string) *DeleteClientPublicKeyRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteClientPublicKeyRequest) SetClientPublicKeyId(v string) *DeleteClientPublicKeyRequest {
	s.ClientPublicKeyId = &v
	return s
}

func (s *DeleteClientPublicKeyRequest) SetInstanceId(v string) *DeleteClientPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteClientPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
