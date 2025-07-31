// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsKeys(v []*DescribeKmsKeysResponseBodyKmsKeys) *DescribeKmsKeysResponseBody
	GetKmsKeys() []*DescribeKmsKeysResponseBodyKmsKeys
	SetRequestId(v string) *DescribeKmsKeysResponseBody
	GetRequestId() *string
}

type DescribeKmsKeysResponseBody struct {
	// The KMS keys.
	KmsKeys []*DescribeKmsKeysResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 60EEBD77-227C-5B39-86EA-D89163C5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKmsKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBody) GetKmsKeys() []*DescribeKmsKeysResponseBodyKmsKeys {
	return s.KmsKeys
}

func (s *DescribeKmsKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKmsKeysResponseBody) SetKmsKeys(v []*DescribeKmsKeysResponseBodyKmsKeys) *DescribeKmsKeysResponseBody {
	s.KmsKeys = v
	return s
}

func (s *DescribeKmsKeysResponseBody) SetRequestId(v string) *DescribeKmsKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKmsKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKmsKeysResponseBodyKmsKeys struct {
	// The alias of the key.
	//
	// example:
	//
	// key-shh656820f4mh9qxxxxx     alias/test1
	KeyAlias *string `json:"KeyAlias,omitempty" xml:"KeyAlias,omitempty"`
	// The key ID.
	//
	// example:
	//
	// 37291352-xxxx-xxxx-adbf-fd0630a95583
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeKmsKeysResponseBodyKmsKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBodyKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) GetKeyAlias() *string {
	return s.KeyAlias
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) SetKeyAlias(v string) *DescribeKmsKeysResponseBodyKmsKeys {
	s.KeyAlias = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) SetKeyId(v string) *DescribeKmsKeysResponseBodyKmsKeys {
	s.KeyId = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) Validate() error {
	return dara.Validate(s)
}
