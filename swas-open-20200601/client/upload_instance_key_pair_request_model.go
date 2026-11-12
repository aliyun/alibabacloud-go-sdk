// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadInstanceKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UploadInstanceKeyPairRequest
	GetClientToken() *string
	SetInstanceId(v string) *UploadInstanceKeyPairRequest
	GetInstanceId() *string
	SetKeyPairName(v string) *UploadInstanceKeyPairRequest
	GetKeyPairName() *string
	SetPublicKey(v string) *UploadInstanceKeyPairRequest
	GetPublicKey() *string
	SetRegionId(v string) *UploadInstanceKeyPairRequest
	GetRegionId() *string
}

type UploadInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test_gin
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key.
	//
	// example:
	//
	// ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAID5aQ5bM0Am3mWe+upjSXqisUT4DLR6ExwvA0****	- **@**.com
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UploadInstanceKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UploadInstanceKeyPairRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UploadInstanceKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *UploadInstanceKeyPairRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *UploadInstanceKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadInstanceKeyPairRequest) SetClientToken(v string) *UploadInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetInstanceId(v string) *UploadInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetKeyPairName(v string) *UploadInstanceKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetPublicKey(v string) *UploadInstanceKeyPairRequest {
	s.PublicKey = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetRegionId(v string) *UploadInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
