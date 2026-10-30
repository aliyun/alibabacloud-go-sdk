// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *CreateKeyPairRequest
	GetKeyPairName() *string
}

type CreateKeyPairRequest struct {
	// The name of the key pair. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s CreateKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
