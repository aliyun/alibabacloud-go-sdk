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
	// The name of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABC1234567*****
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
