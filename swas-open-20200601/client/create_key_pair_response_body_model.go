// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *CreateKeyPairResponseBody
	GetKeyPairName() *string
	SetPrivateKeyBody(v string) *CreateKeyPairResponseBody
	GetPrivateKeyBody() *string
	SetRequestId(v string) *CreateKeyPairResponseBody
	GetRequestId() *string
}

type CreateKeyPairResponseBody struct {
	// The name of the key pair. The name must be 2 to 64 characters in length, and can contain English letters and Chinese characters. It must start with a letter or a Chinese character and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The private key is in the PEM-encoded PKCS#8 format.
	//
	// example:
	//
	// MIIEpAIBAAKCAQEAtReyMzLIcBH78EV2zj****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairResponseBody) GetPrivateKeyBody() *string {
	return s.PrivateKeyBody
}

func (s *CreateKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKeyPairResponseBody) SetKeyPairName(v string) *CreateKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetPrivateKeyBody(v string) *CreateKeyPairResponseBody {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
