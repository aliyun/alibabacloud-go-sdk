// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseBodyKmsKeys) *DescribeUserEncryptionKeyListResponseBody
	GetKmsKeys() []*DescribeUserEncryptionKeyListResponseBodyKmsKeys
	SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetRequestId() *string
}

type DescribeUserEncryptionKeyListResponseBody struct {
	// Details about the KMS keys.
	KmsKeys []*DescribeUserEncryptionKeyListResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetKmsKeys() []*DescribeUserEncryptionKeyListResponseBodyKmsKeys {
	return s.KmsKeys
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseBodyKmsKeys) *DescribeUserEncryptionKeyListResponseBody {
	s.KmsKeys = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) Validate() error {
	if s.KmsKeys != nil {
		for _, item := range s.KmsKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserEncryptionKeyListResponseBodyKmsKeys struct {
	// The ID of the KMS key.
	//
	// example:
	//
	// 0b8b1825-fd99-418f-875e-e4dec1dd8715
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBodyKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBodyKmsKeys) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeUserEncryptionKeyListResponseBodyKmsKeys) SetKeyId(v string) *DescribeUserEncryptionKeyListResponseBodyKmsKeys {
	s.KeyId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBodyKmsKeys) Validate() error {
	return dara.Validate(s)
}
