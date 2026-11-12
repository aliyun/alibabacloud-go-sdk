// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ImportKeyPairRequest
	GetClientToken() *string
	SetKeyPairName(v string) *ImportKeyPairRequest
	GetKeyPairName() *string
	SetPublicKeyBody(v string) *ImportKeyPairRequest
	GetPublicKeyBody() *string
	SetRegionId(v string) *ImportKeyPairRequest
	GetRegionId() *string
}

type ImportKeyPairRequest struct {
	// A client token used to ensure the idempotence of the request. Generate a unique value for this parameter from your client. **ClientToken*	- can contain only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the key pair. The name must be 2 to 64 characters in length. It must start with a letter or a Chinese character. It cannot start with http\\:// or https\\://. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The content of the public key of the key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbO5Govwhb0iHzoMYKkIQxjlHyHH8nxFsW6KF5saxgYhOwdeIpWngpi+/NDWQKvuOnXFFDh/o3eJJkh3rqP+RlMggt4HLQWOd9TS0f4/cgbAzud1caW9PnankCr****
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
	// The ID of the region where the Simple Application Server instance is located.
	//
	// Call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to view a list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ImportKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairRequest) GetPublicKeyBody() *string {
	return s.PublicKeyBody
}

func (s *ImportKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportKeyPairRequest) SetClientToken(v string) *ImportKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

func (s *ImportKeyPairRequest) SetRegionId(v string) *ImportKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *ImportKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
