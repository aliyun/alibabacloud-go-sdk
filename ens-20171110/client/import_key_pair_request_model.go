// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *ImportKeyPairRequest
	GetKeyPairName() *string
	SetPublicKeyBody(v string) *ImportKeyPairRequest
	GetPublicKeyBody() *string
}

type ImportKeyPairRequest struct {
	// The name of the key pair. The name must conform to the following naming conventions:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// You can specify the name of only one key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key of the key pair. You can specify only one public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssh-rsa AAAAB****
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairRequest) GetPublicKeyBody() *string {
	return s.PublicKeyBody
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

func (s *ImportKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
