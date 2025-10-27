// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsKeys(v *DescribeKmsKeysResponseBodyKmsKeys) *DescribeKmsKeysResponseBody
	GetKmsKeys() *DescribeKmsKeysResponseBodyKmsKeys
	SetRequestId(v string) *DescribeKmsKeysResponseBody
	GetRequestId() *string
}

type DescribeKmsKeysResponseBody struct {
	// The queried KMS keys.
	KmsKeys *DescribeKmsKeysResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8F91F25F-8BCF-59E3-AF67-3806DB41FD09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKmsKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBody) GetKmsKeys() *DescribeKmsKeysResponseBodyKmsKeys {
	return s.KmsKeys
}

func (s *DescribeKmsKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKmsKeysResponseBody) SetKmsKeys(v *DescribeKmsKeysResponseBodyKmsKeys) *DescribeKmsKeysResponseBody {
	s.KmsKeys = v
	return s
}

func (s *DescribeKmsKeysResponseBody) SetRequestId(v string) *DescribeKmsKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKmsKeysResponseBody) Validate() error {
	if s.KmsKeys != nil {
		if err := s.KmsKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKmsKeysResponseBodyKmsKeys struct {
	KmsKey []*DescribeKmsKeysResponseBodyKmsKeysKmsKey `json:"KmsKey,omitempty" xml:"KmsKey,omitempty" type:"Repeated"`
}

func (s DescribeKmsKeysResponseBodyKmsKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBodyKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) GetKmsKey() []*DescribeKmsKeysResponseBodyKmsKeysKmsKey {
	return s.KmsKey
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) SetKmsKey(v []*DescribeKmsKeysResponseBodyKmsKeysKmsKey) *DescribeKmsKeysResponseBodyKmsKeys {
	s.KmsKey = v
	return s
}

func (s *DescribeKmsKeysResponseBodyKmsKeys) Validate() error {
	if s.KmsKey != nil {
		for _, item := range s.KmsKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKmsKeysResponseBodyKmsKeysKmsKey struct {
	// The alias of the key.
	//
	// example:
	//
	// mykey
	KeyAlias *string `json:"KeyAlias,omitempty" xml:"KeyAlias,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 2e81355b-f8e7-4090-8082-a8f8124a621c
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeKmsKeysResponseBodyKmsKeysKmsKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBodyKmsKeysKmsKey) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBodyKmsKeysKmsKey) GetKeyAlias() *string {
	return s.KeyAlias
}

func (s *DescribeKmsKeysResponseBodyKmsKeysKmsKey) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKmsKeysResponseBodyKmsKeysKmsKey) SetKeyAlias(v string) *DescribeKmsKeysResponseBodyKmsKeysKmsKey {
	s.KeyAlias = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKmsKeysKmsKey) SetKeyId(v string) *DescribeKmsKeysResponseBodyKmsKeysKmsKey {
	s.KeyId = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKmsKeysKmsKey) Validate() error {
	return dara.Validate(s)
}
