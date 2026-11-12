// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachKeyPairRequest
	GetClientToken() *string
	SetInstanceIds(v []*string) *DetachKeyPairRequest
	GetInstanceIds() []*string
	SetKeyPairName(v string) *DetachKeyPairRequest
	GetKeyPairName() *string
	SetRegionId(v string) *DetachKeyPairRequest
	GetRegionId() *string
}

type DetachKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers from which you want to unbind SSH key pairs. You can specify a maximum of 50 IDs of simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the key pair. The name must be globally unique. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with http\\:// or https\\://.
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachKeyPairRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DetachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachKeyPairRequest) SetClientToken(v string) *DetachKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachKeyPairRequest) SetInstanceIds(v []*string) *DetachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairName(v string) *DetachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairRequest) SetRegionId(v string) *DetachKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *DetachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
