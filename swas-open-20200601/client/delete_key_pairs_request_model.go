// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteKeyPairsRequest
	GetClientToken() *string
	SetKeyPairNames(v []*string) *DeleteKeyPairsRequest
	GetKeyPairNames() []*string
	SetRegionId(v string) *DeleteKeyPairsRequest
	GetRegionId() *string
}

type DeleteKeyPairsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The names of the SSH key pairs. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http\\:// or https\\://. You can specify the names of a maximum of 50 SSH key pairs.
	//
	// This parameter is required.
	KeyPairNames []*string `json:"KeyPairNames,omitempty" xml:"KeyPairNames,omitempty" type:"Repeated"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteKeyPairsRequest) GetKeyPairNames() []*string {
	return s.KeyPairNames
}

func (s *DeleteKeyPairsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKeyPairsRequest) SetClientToken(v string) *DeleteKeyPairsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetKeyPairNames(v []*string) *DeleteKeyPairsRequest {
	s.KeyPairNames = v
	return s
}

func (s *DeleteKeyPairsRequest) SetRegionId(v string) *DeleteKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
