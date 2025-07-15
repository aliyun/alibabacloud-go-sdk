// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateKeyPairResponseBodyData) *CreateKeyPairResponseBody
	GetData() *CreateKeyPairResponseBodyData
	SetRequestId(v string) *CreateKeyPairResponseBody
	GetRequestId() *string
}

type CreateKeyPairResponseBody struct {
	// The objects that are returned.
	Data *CreateKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) GetData() *CreateKeyPairResponseBodyData {
	return s.Data
}

func (s *CreateKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKeyPairResponseBody) SetData(v *CreateKeyPairResponseBodyData) *CreateKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateKeyPairResponseBodyData struct {
	// The time when the key pair was created.
	//
	// example:
	//
	// 2024-06-30 08:45:09.0
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3*****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The PEM-encoded private key that is in PKCS#8 format and adheres to the ADB connection specification.
	//
	// example:
	//
	// MIIEpAIBAAKCAQEAtReyMzLIcBH78EV2zj****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
}

func (s CreateKeyPairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateKeyPairResponseBodyData) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *CreateKeyPairResponseBodyData) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairResponseBodyData) GetPrivateKeyBody() *string {
	return s.PrivateKeyBody
}

func (s *CreateKeyPairResponseBodyData) SetGmtCreated(v string) *CreateKeyPairResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetKeyPairId(v string) *CreateKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetKeyPairName(v string) *CreateKeyPairResponseBodyData {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) SetPrivateKeyBody(v string) *CreateKeyPairResponseBodyData {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponseBodyData) Validate() error {
	return dara.Validate(s)
}
