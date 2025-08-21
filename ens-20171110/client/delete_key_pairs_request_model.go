// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairId(v string) *DeleteKeyPairsRequest
	GetKeyPairId() *string
	SetKeyPairName(v string) *DeleteKeyPairsRequest
	GetKeyPairName() *string
}

type DeleteKeyPairsRequest struct {
	// The ID of the SSH key pair.
	//
	// example:
	//
	// ssh-5lywanlkih1zo9yl8eg****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair. The name must conform to the following naming conventions:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// Before you delete a key pair, you can call the DescribeKeyPairs operation to query existing key pairs.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DeleteKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DeleteKeyPairsRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DeleteKeyPairsRequest) SetKeyPairId(v string) *DeleteKeyPairsRequest {
	s.KeyPairId = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetKeyPairName(v string) *DeleteKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DeleteKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
