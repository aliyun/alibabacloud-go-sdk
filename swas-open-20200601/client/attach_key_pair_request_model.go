// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachKeyPairRequest
	GetClientToken() *string
	SetInstanceIds(v []*string) *AttachKeyPairRequest
	GetInstanceIds() []*string
	SetKeyPairName(v string) *AttachKeyPairRequest
	GetKeyPairName() *string
	SetRegionId(v string) *AttachKeyPairRequest
	GetRegionId() *string
}

type AttachKeyPairRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value from your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. You can specify up to 50 instance IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the key pair. The name must be 2 to 64 characters in length. It must start with a letter or a Chinese character and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the specified lightweight database instance.
	//
	// You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachKeyPairRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachKeyPairRequest) SetClientToken(v string) *AttachKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachKeyPairRequest) SetInstanceIds(v []*string) *AttachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairName(v string) *AttachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairRequest) SetRegionId(v string) *AttachKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *AttachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
