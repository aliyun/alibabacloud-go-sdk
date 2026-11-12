// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *ImportKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *ImportKeyPairResponseBody
	GetRequestId() *string
}

type ImportKeyPairResponseBody struct {
	// The name of the key pair. The name must be 2 to 64 characters in length. It must start with a letter or a Chinese character. It cannot start with http\\:// or https\\://. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKeyPairResponseBody) SetKeyPairName(v string) *ImportKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
