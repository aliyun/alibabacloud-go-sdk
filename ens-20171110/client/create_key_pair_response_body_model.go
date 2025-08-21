// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairFingerPrint(v string) *CreateKeyPairResponseBody
	GetKeyPairFingerPrint() *string
	SetKeyPairId(v string) *CreateKeyPairResponseBody
	GetKeyPairId() *string
	SetKeyPairName(v string) *CreateKeyPairResponseBody
	GetKeyPairName() *string
	SetPrivateKeyBody(v string) *CreateKeyPairResponseBody
	GetPrivateKeyBody() *string
	SetRequestId(v string) *CreateKeyPairResponseBody
	GetRequestId() *string
}

type CreateKeyPairResponseBody struct {
	// The fingerprint of the key pair. The message-digest algorithm 5 (MD5) is used based on the public key fingerprint format defined in RFC 4716. For more information, see [RFC 4716](https://tools.ietf.org/html/rfc4716).
	//
	// example:
	//
	// 7880c1ad4687fdbf7a6da2131****
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The ID of the SSH key pair.
	//
	// example:
	//
	// ssh-5lywanlkih1zo9yl8eg****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The private key is encoded with PEM in the PKCS#8 format.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----\\nMIIEogIBAAKCAQE****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *CreateKeyPairResponseBody) GetKeyPairId() *string {
	return s.KeyPairId
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

func (s *CreateKeyPairResponseBody) SetKeyPairFingerPrint(v string) *CreateKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetKeyPairId(v string) *CreateKeyPairResponseBody {
	s.KeyPairId = &v
	return s
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
