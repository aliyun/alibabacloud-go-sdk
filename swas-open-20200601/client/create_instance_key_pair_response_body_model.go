// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFingerprint(v string) *CreateInstanceKeyPairResponseBody
	GetFingerprint() *string
	SetKeyPairName(v string) *CreateInstanceKeyPairResponseBody
	GetKeyPairName() *string
	SetPrivateKey(v string) *CreateInstanceKeyPairResponseBody
	GetPrivateKey() *string
	SetRequestId(v string) *CreateInstanceKeyPairResponseBody
	GetRequestId() *string
}

type CreateInstanceKeyPairResponseBody struct {
	// The fingerprint of the key pair.
	//
	// example:
	//
	// If2K1ItazA4GlKkWCEhdRj8Wd6czAvK9*****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// ceshi
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key.
	//
	// example:
	//
	// ***
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceKeyPairResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *CreateInstanceKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateInstanceKeyPairResponseBody) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateInstanceKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceKeyPairResponseBody) SetFingerprint(v string) *CreateInstanceKeyPairResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetKeyPairName(v string) *CreateInstanceKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetPrivateKey(v string) *CreateInstanceKeyPairResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetRequestId(v string) *CreateInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
