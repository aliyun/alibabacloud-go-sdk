// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKeyPairRequest
	GetClientToken() *string
	SetKeyPairName(v string) *CreateKeyPairRequest
	GetKeyPairName() *string
	SetRegionId(v string) *CreateKeyPairRequest
	GetRegionId() *string
}

type CreateKeyPairRequest struct {
	// Ensures the idempotence of the request. You can generate the value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the key pair. The name must be 2 to 64 characters in length, and can contain English letters and Chinese characters. It must start with a letter or a Chinese character and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the specified simple application server instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKeyPairRequest) SetClientToken(v string) *CreateKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairRequest) SetRegionId(v string) *CreateKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
