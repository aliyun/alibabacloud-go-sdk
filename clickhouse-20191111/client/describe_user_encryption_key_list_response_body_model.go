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
	KmsKeys []*DescribeUserEncryptionKeyListResponseBodyKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" type:"Repeated"`
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
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
	// example:
	//
	// 0275bd3f-fdbb-4d8c-846b-71b211******
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
